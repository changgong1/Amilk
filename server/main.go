package main

import (
	"./gosrc"
	"github.com/wonderivan/logger"
)

func main() {
	gosrc.Calculate(1)
	logger.Debug(`{"Console":"DEBG"}`)
}
