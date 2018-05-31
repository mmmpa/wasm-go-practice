package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mmmpa/iso-go/data"
	"github.com/labstack/echo/middleware"
)

var store = []data.Message{}

func main() {
	e := echo.New()

	e.GET("api/messages", indexMessages)
	e.POST("api/messages", createMessage)
	e.Static("/", "public")

	e.Use(cors())
	e.Logger.Fatal(e.Start(":1323"))
}

func indexMessages(c echo.Context) error {
	return c.JSON(http.StatusOK, store)
}

func createMessage(c echo.Context) error {
	message := new(data.Message)

	if err := c.Bind(message); err != nil {
		return err
	}

	result := data.Validate(message)

	if result.Valid {
		store = append(store, *message)
		return c.JSON(http.StatusCreated, data.MessageCreationResult{ID: 1})
	} else {
		return c.JSON(http.StatusBadRequest, result)
	}
}

func cors() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		},
	)
}
