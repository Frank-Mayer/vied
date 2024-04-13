package ui

import (
	"github.com/Frank-Mayer/vied/project"
	"github.com/Frank-Mayer/vied/ui/style"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	windowConfigFlags uint32 = rl.FlagVsyncHint | rl.FlagWindowResizable | rl.FlagMsaa4xHint
	targetFPS                = 120
)

func init() {
	SetupLogger()
	rl.SetConfigFlags(windowConfigFlags)
}

func Init(p *project.Project) {
	rl.InitWindow(0, 0, "ViEd")
	defer rl.CloseWindow()

	rl.SetTargetFPS(targetFPS)
	rl.SetExitKey(0)
	rl.SetConfigFlags(windowConfigFlags)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(style.Background)
		drawOutline(p)
		rl.DrawFPS(0, 0)
		rl.EndDrawing()
	}
}
