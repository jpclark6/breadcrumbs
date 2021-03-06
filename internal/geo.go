// Package geo handles the main logic and web endpoints
package geo

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // used by gorm
)

var db *gorm.DB
var err error

// SetupRouter creates endpoints and calls additional logic
func SetupRouter() {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("$PORT must be set")
	} else {
		fmt.Println("Port", port)
	}

	router := SetupRouterSettings()

	router.Run(":" + port)
}

func SetupRouterSettings() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("templates/*")

	geo := router.Group("/")
	{
		geo.GET("", rootEndpoint)
		geo.POST("/submitbreadcrumb", submitBreadcrumb)
		geo.GET("/getbreadcrumbs", getBreadcrumbs)
	}
	router.Static("/web", "./web")
	router.NoRoute(endpointNotFound)
	return router
}

// SetupDatabase sets up the database connection
func SetupDatabase() {
	dbSSL := os.Getenv("DB_SSL")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	databaseURL := fmt.Sprintf("sslmode=%v user=%v password=%v host=%v port=%v dbname=%v",
		dbSSL,
		username,
		password,
		host,
		5432,
		name,
	)
	db, err = gorm.Open("postgres", databaseURL)
	if err != nil {
		fmt.Println("could not connect to database. err:", err)
	}
	db.AutoMigrate(&Message{})
}

func endpointNotFound(c *gin.Context) {
	c.Writer.WriteString("404: not found")
}

func rootEndpoint(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

func submitBreadcrumb(c *gin.Context) {
	var message Message
	if c.ShouldBind(&message) == nil {
		log.Print(message.Text)
	}
	fmt.Println("text:", message.Text)
	if message.Text != "" {
		go writeBreadcrumbToDB(message)

		c.HTML(http.StatusOK, "submitted.tmpl", gin.H{
			"title": "Main website",
		})
	} else {
		c.HTML(http.StatusOK, "submittedNoText.tmpl", gin.H{
			"title": "Main website",
		})
	}
}

func writeBreadcrumbToDB(message Message) {
	message.Lat = message.Lat + 0.005*(1-rand.Float64())
	message.Long = message.Long + 0.005*(1-rand.Float64())
	db.Create(&message)
}

func getBreadcrumbs(c *gin.Context) {
	lat, long, err := parseLatLong(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"invalid request": err})
	}
	messages := findNearbyMessages(lat, long)
	messages = FindDistances(messages, lat, long)
	messages = RoundMessageValues(messages)
	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

func parseLatLong(c *gin.Context) (lat float64, long float64, err error) {
	var latLong Message
	err = c.ShouldBindQuery(&latLong)
	lat = latLong.Lat
	long = latLong.Long
	return lat, long, err
}

func findNearbyMessages(lat, long float64) []Message {
	var messages []Message
	db.Limit(5).Where("lat >= ? AND lat <= ? AND long >= ? and long <= ?",
		lat-0.015, lat+0.015, long-0.015, long+0.015).Find(&messages)
	return messages
}

func FindDistances(messages []Message, lat float64, long float64) []Message {
	for i := 0; i < len(messages); i++ {
		message := messages[i]
		deltaXDeg := math.Abs(long - message.Long)
		deltaYDeg := math.Abs(lat - message.Lat)

		deltaXMiles := deltaXDeg / .0140 * 1.5 // appx 1.5 mi/.0140 deg for east/west in US
		deltaYMiles := deltaYDeg / .0144 * 1.0 // appx 1.0 mi/.0144 deg for north/south in US

		distance := math.Sqrt(math.Pow(deltaXMiles, 2) + math.Pow(deltaYMiles, 2))

		messages[i].Distance = distance
	}

	sort.Slice(messages, func(i, j int) bool {
		return messages[i].Distance < messages[j].Distance
	})

	return messages
}

func RoundMessageValues(messages []Message) []Message {
	for i := 0; i < len(messages); i++ {
		message := messages[i]
		message.Distance = math.Floor(message.Distance*1000) / 1000
		message.Lat = math.Floor(message.Lat*1000) / 1000
		message.Long = math.Floor(message.Long*1000) / 1000
		messages[i] = message
	}
	return messages
}
