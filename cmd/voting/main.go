package main

import (
	"os"

	"github.com/buffup/GolangTechTask"
	"go.uber.org/zap"
)

func main() {
	if err := GolangTechTask.App().Run(os.Args); err != nil {
		GolangTechTask.Logger.Error("Exited program", zap.String("error", err.Error()))
		os.Exit(1)
	}
}
