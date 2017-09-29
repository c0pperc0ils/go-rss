package main

import (
  "github.com/sirupsen/logrus"
  "github.com/c0pperc0ils/go-rss/routes"
)

func main() {
  log := logrus.New()
  log.Info("worked")

  server := routes.RssRoutes{}

  server.Run(log)
}
