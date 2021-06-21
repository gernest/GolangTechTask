package GolangTechTask

import "go.uber.org/zap"

var Logger *zap.Logger

func init() {
	var err error
	c := zap.NewProductionConfig()
	c.DisableStacktrace = true
	c.Level.SetLevel(zap.DebugLevel)
	Logger, err = c.Build(
		zap.WithCaller(false),
	)
	if err != nil {
		panic(err)
	}
}
