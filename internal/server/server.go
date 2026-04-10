package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"ppoliv2/internal/repository"
)

func New() (*gin.Engine, error) {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}

	repo := repository.NewHomeRepository(db)
	if err = repo.MigrateAndSeed(); err != nil {
		return nil, err
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		data, dataErr := repo.HomeData()
		if dataErr != nil {
			c.String(500, dataErr.Error())
			return
		}
		c.HTML(200, "base", data)
	})

	return r, nil
}
