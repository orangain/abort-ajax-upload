package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.File("/", "index.html")
	e.POST("/upload", func(c echo.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, file.Filename)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
