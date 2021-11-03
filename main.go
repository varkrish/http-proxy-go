package main

import (
	"net/http"
	"os"
	"io/ioutil"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	 "github.com/labstack/gommon/log"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.SetLevel(log.DEBUG)
		e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hey There!")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})
	e.GET("/proxy", func(c echo.Context) error {

		tr := &http.Transport{}
		client := &http.Client{Transport: tr}
		url := c.FormValue("url")
		e.Logger.Debug("Context is " + c.FormValue("url"))
		resp, err := client.Get(
			url,
		  )
		  if err != nil {
			return err
		  }
		  defer resp.Body.Close()
		  body, err := ioutil.ReadAll(resp.Body)
		  // Unmarshal the response into a ExampleResponse struct	  
		  return c.HTML(http.StatusOK, string(body))	  
		//return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
