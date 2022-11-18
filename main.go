package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func main() {
	e := echo.New()
	e.GET("/numbers", numbers)
	e.Logger.Fatal(e.Start(":8080"))
}

// Numbers is an api endpoint that returns numbers based on query
func numbers(c echo.Context) error {
	query := c.QueryParam("size")
	if query == "" {
		return c.String(http.StatusBadRequest, "Please supply a positive integer value such as 0 at least")
	}
	size, err := strconv.Atoi(query)
	resp := response{
		series: []int{},
	}
	var series []int
	if err != nil {
		return c.String(http.StatusBadRequest, "Please supply an integer value")
	}
	if size <= 0 {
		return c.JSON(http.StatusOK, resp)
	}
	for i := 0; i < size; i++ {
		series = append(resp.series, i)
	}
	return c.JSON(http.StatusOK, series)
}

type response struct {
	series []int
}
