package logger

import (
	"bls/config"
	"os"
	"time"

	"github.com/elisiei/zlog"
)

var Log *zlog.Logger

func init() {
	Log = zlog.New()
	Log.SetOutput(os.Stdout)
	Log.SetTimeFormat(time.Kitchen)
	Log.EnableColors(true)
	Log.ShowCaller(true)
	if config.Mode == "debug" {
		Log.SetLevel(zlog.LevelDebug)
	} else {
		Log.SetLevel(zlog.LevelInfo)
	}
}
