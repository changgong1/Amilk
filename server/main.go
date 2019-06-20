package main

import (
	"./gosrc"
	"github.com/wonderivan/logger"
)

func main() {
	logger.SetLogger(`{"Console": {"level": "DEBG","color": true}}`)
	gosrc.Calculate(1)
	logger.Debug(`{"Console":"DEBG"}`)
}
