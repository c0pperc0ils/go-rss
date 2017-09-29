package routes

import (
  "github.com/gin-gonic/gin"
  "database/sql"
)

type FeedRoutes struct {
  db *sql.DB
}

func (f *FeedRoutes) Setup(g *gin.Engine) error {
  g.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{"message": "pong"})
  })

  return nil
}
