package bootstrap

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/jiny3/go-app/services"
)

// Start start the application.
func Start() {
	app := APP("config/app.toml")
	r := gin.Default()
	r.GET("/hello", services.Hello)

	if err := r.Run(strings.Join([]string{":", strconv.Itoa(app.Port)}, "")); err != nil {
		logrus.WithError(err).Fatal("failed to start server")
	}
}
