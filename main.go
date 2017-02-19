package main

import (
	"fmt"
	"github.com/jinzhu/now"
	"github.com/labstack/echo"
	"github.com/urfave/cli"
	"net/url"
	"os"
	"strconv"
	"time"
)

type H map[string]interface{}

func mainHandler(c echo.Context) error {
	return c.String(200, "Timestamp Microservice")
}

func dataHandler(c echo.Context) error {
	data := c.Param("data")
	unix, err := strconv.Atoi(data)
	if err != nil {
		dateString, err := url.QueryUnescape(data)
		if err != nil {
			return c.JSON(200, H{"unix": nil, "natural": nil})
		}
		t, err := now.Parse(dateString)
		if err != nil {
			return c.JSON(200, H{"unix": nil, "natural": nil})
		} else {
			return c.JSON(200, H{"unix": t.Unix(), "natural": t.Format("January 02, 2006")})
		}
	} else {
		return c.JSON(200, H{"unix": unix, "natural": time.Unix(int64(unix), 0).Format("January 02, 2006")})
	}
}

func start(c *cli.Context) error {
	now.TimeFormats = append(now.TimeFormats, "January 02, 2006")
	port := c.Int("port")
	e := echo.New()
	e.GET("/", mainHandler)
	e.GET("/:data", dataHandler)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
	return nil
}

func main() {
	app := cli.NewApp()
	app.Author = "Alain Gilbert"
	app.Email = "alain.gilbert.15@gmail.com"
	app.Name = "Timestamp Microservice"
	app.Usage = "Timestamp Microservice"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   "port",
			Value:  3001,
			Usage:  "Webserver port",
			EnvVar: "PORT",
		},
	}
	app.Action = start
	app.Run(os.Args)
}
