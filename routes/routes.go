package routes

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "database/sql"
  _ "github.com/lib/pq"
  "github.com/sirupsen/logrus"
  "github.com/toorop/gin-logrus"
)

const (
  host = "db"
  port = 5432
  user = "pqgotest"
  password = "secret"
  dbname = "pqgotest"
  sslmode = "disable"
)

type RssRoutes struct {
}

func (r *RssRoutes) getDb() *sql.DB {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
    host, port, user, password, dbname, sslmode)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  err = db.Ping()
  if err != nil {
    panic(err)
  }

  return db
}

func (r *RssRoutes) Run(log *logrus.Logger) error {
  db := r.getDb()

  feedRoutes := &FeedRoutes{db: db}

  myGin := gin.Default()
  myGin.Use(ginlogrus.Logger(log), gin.Recovery())

  feedRoutes.Setup(myGin)

  myGin.Run()
  return nil
}
