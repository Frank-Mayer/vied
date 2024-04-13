package ui

import (
	"github.com/charmbracelet/log"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func SetupLogger() {
	rl.SetTraceLogLevel(rl.LogAll)
	rl.SetTraceLogCallback(func(level int, msg string) {
		switch level {
		case int(rl.LogDebug), int(rl.LogTrace):
			log.Debug(msg)
		case int(rl.LogInfo):
			log.Info(msg)
		case int(rl.LogWarning):
			log.Warn(msg)
		case int(rl.LogError):
			log.Error(msg)
		case int(rl.LogFatal):
			log.Fatal(msg)
		}
	})

}
