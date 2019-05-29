package main

import (
	"fmt"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 9}))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"message": "welcome :)",
		})
	})

	e.GET("/macros/exec/:macro", func(c echo.Context) error {
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

	e.Start(*flagHTTPAddr)
}
