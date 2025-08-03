package main

import (
	"github.com/jiny3/gopkg/hookx"
	"github.com/sirupsen/logrus"

	"github.com/jiny3/go-app/bootstrap"
)

func init() {
	// use hookx.Exit(func() { ... }) to register a function to be called on exit
	// wait for the application to exit gracefully
	defer hookx.ExitWait()

	// initialize the application configuration
	bootstrap.Init()
}

func main() {
	// This is the entry point of the application.
	// You can add your application logic here.
	logrus.WithField("author", "jiny3").Info("go-app template")
	bootstrap.Start()
}
