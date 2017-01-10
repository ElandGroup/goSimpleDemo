package main

import (
	"Hellogo/test"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

var (
	e = echo.New()
)

func main() {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello api")
	})
	fmt.Println("Hello go")
	test.TestBase()
	test.TestStruct()
	test.TestInterface()
	e.Start(":1003")
}
