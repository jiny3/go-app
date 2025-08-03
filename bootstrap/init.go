package bootstrap

import (
	"net/http"

	"github.com/jiny3/gopkg/hookx"
	"github.com/jiny3/gopkg/logx"
	"github.com/sirupsen/logrus"
)

// Init you can initialize your application here.
func Init() {
	withPprof()
	withLogrus()
	withExitTip()
}

func withPprof() {
	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			panic(err)
		}
	}()
}

func withLogrus() {
	logx.InitLogrus(logx.WithLevel(logrus.DebugLevel), logx.WithAllText("stderr.log"), logx.WithOpsJSON("ops.log"))
}

func withExitTip() {
	hookx.Exit(func() {
		logrus.Debug("Application is exiting. Please wait...")
	})
}
