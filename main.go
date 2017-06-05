package main

import (
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "./handler"
)

func main() {
  // create echo instance

  e := echo.New()

  // middleware for all request
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // routing
  e.GET("/hello", handler.MainPage())

  // up server
  e.Start(":4000")
}
