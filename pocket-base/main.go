package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/api/init-check", func(c echo.Context) error {
			var total int

			err := app.DB().
				Select("count(*)").
				From("_admins").
				Row(&total)
			if err != nil {
				return err
			}
			if total > 0 {
				return c.JSON(http.StatusOK, map[string]interface{}{
					"isSetup": true,
					"message": "Initial setup is complete",
				})
			}
			return c.JSON(http.StatusOK, map[string]interface{}{
				"isSetup": false,
				"message": "Initial setup needed",
			})
		})
		return nil
	})
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
