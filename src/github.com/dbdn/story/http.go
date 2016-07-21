package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

type Response struct {
	ID  int
	Msg string
}

func getResponse(id int, msg string) *Response {
	resp := new(Response)
	resp.ID = id
	resp.Msg = msg
	return resp
}

func createUser(c echo.Context) error {
	return c.JSON(http.StatusCreated, getResponse(0, "ok"))
}

func NewHttpServer() {
	e := echo.New()
	e.POST("/users", createUser)

	e.Run(standard.New(":1323"))
}
