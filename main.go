package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 9}))
	// e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"message": "welcome :)",
			"macros":  len(configs.Macros),
		})
	})

	e.GET("/:macro", func(c echo.Context) error {
		m := configs.Macros[c.Param("macro")]
		if nil == m {
			return c.JSON(404, map[string]string{
				"error": "not found",
			})
		}

		if m.Private {
			return c.JSON(403, map[string]string{
				"error": "you don't have permission to access this resource",
			})
		}

		res, cached, err := execMacro(c.Param("macro"))
		if err != nil {
			return c.JSON(500, map[string]string{
				"error": err.Error(),
			})
		}

		if cached {
			c.Response().Header().Set("X-Cached-Version", "true")
		}

		if _, err := json.Marshal(res); err != nil {
			return c.JSON(500, map[string]interface{}{
				"error":   err.Error(),
				"payload": fmt.Sprintf("%v", res),
			})
		}

		return c.JSON(200, map[string]interface{}{
			"result": res,
		})
	})

	log.Fatal(e.Start(*flagHTTPAddr))
}
