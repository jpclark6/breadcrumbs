package geo

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // used by gorm
)

// Message struct to hold message info and location
type Message struct {
	gorm.Model
	Text     string  `form:"text"`
	Lat      float64 `form:"lat"`
	Long     float64 `form:"long"`
	Distance float64 `form:"distance"`
	Private  bool    `form:"private"`
	Password string  `form:"password"`
}