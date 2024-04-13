package ui

import (
	"github.com/Frank-Mayer/vied/project"
	"github.com/Frank-Mayer/vied/ui/style"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Wizard() *project.Project {
	var p *project.Project = nil

	rl.InitWindow(400, 300, "ViEd Wizard")
	defer rl.CloseWindow()

	rl.SetTargetFPS(targetFPS)
	rl.SetExitKey(rl.KeyEscape)
	rl.SetConfigFlags(windowConfigFlags)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(style.Background)
		drawOutline(p)
		rl.DrawFPS(0, 0)
		rl.EndDrawing()
	}

	return p
}
