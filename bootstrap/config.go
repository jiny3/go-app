// Description: This file contains the configuration for the application
package bootstrap

import (
	"github.com/jiny3/gopkg/configx"
	"github.com/sirupsen/logrus"
)

type app struct {
	Port int `toml:"port"`
}

func APP(path string) app {
	app := app{}
	err := configx.Read(path, &app)
	if err != nil {
		logrus.WithField("path", path).WithError(err).Fatal("read config failed")
	}
	return app
}
