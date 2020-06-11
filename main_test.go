package main

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/jpclark6/breadcrumbs/internal"
    
    "github.com/stretchr/testify/assert"
)

func TestCSSEndpoint(t *testing.T) {
    router := geo.SetupRouterSettings()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/web/css/style.css", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
    assert.Contains(t, w.Body.String(), "background-color")
}

func TestCSSResetEndpoint(t *testing.T) {
    router := geo.SetupRouterSettings()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/web/css/reset.css", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
    assert.Contains(t, w.Body.String(), "background-color")
}

func TestJavascriptEndpoint(t *testing.T) {
    router := geo.SetupRouterSettings()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/web/javascript/script.js", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
    assert.Contains(t, w.Body.String(), "logLocation")
}

func TestRootEndpoint(t *testing.T) {
    router := geo.SetupRouterSettings()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
    assert.Contains(t, w.Body.String(), "Breadcrumbs")
}

func TestFindDistances(t *testing.T) {
    lat := 100.0
    long := 100.0
    message1 := geo.Message{Lat: lat, Long: long}
    message2 := geo.Message{Lat: lat-1, Long: long}
    messages := []geo.Message{message1, message2}
    returnMessages := geo.FindDistances(messages, 100.0, 100.0)
    assert.Equal(t, returnMessages[0].Distance, 0.0)
    assert.Equal(t, returnMessages[1].Distance, 69.44444444444444)
}

func TestRoundMessageValues(t *testing.T) {
    lat := 100.0001
    long := 100.0001
    message := geo.Message{Lat: lat, Long: long}
    messages := []geo.Message{message}
    messages = geo.RoundMessageValues(messages)
    assert.Equal(t, messages[0].Lat, 100.0)
    assert.Equal(t, messages[0].Long, 100.0)
}