package main

import (
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
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"message": "welcome :)",
		})
	})

	e.GET("/macros/exec/:macro", func(c echo.Context) error {
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

		return c.JSON(200, map[string]interface{}{
			"result": res,
		})
	})

	e.GET("aggregators/exec/:aggregator", func(c echo.Context) error {
		aggKey := c.Param("aggregator")
		agg := configs.Aggragators[aggKey]
		if nil == agg {
			return c.JSON(404, map[string]interface{}{
				"error": fmt.Sprintf("Aggregator %s cannot be found", aggKey),
			})
		}

		ret := map[string]interface{}{}
		errs := map[string]string{}

		for _, m := range agg {
			macro := configs.Macros[c.Param("macro")]
			if macro == nil || macro.Private {
				continue
			}

			res, _, err := execMacro(m)

			if nil != err {
				errs[m] = err.Error()
			} else {
				ret[m] = res
			}
		}

		if len(errs) > 0 {
			return c.JSON(500, map[string]interface{}{
				"errors": errs,
			})
		}

		return c.JSON(200, map[string]interface{}{
			"result": ret,
		})
	})

	log.Fatal(e.Start(*flagHTTPAddr))
}
