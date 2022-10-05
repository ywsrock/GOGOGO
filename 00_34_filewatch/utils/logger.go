package utils

import (
	"go.uber.org/zap"
	"sync"
)

var logger *zap.Logger
var once sync.Once

func NewLogger() (*zap.Logger, error) {
	var err error
	once.Do(func() {
		logger, err = zap.NewProduction()
		//	cfg := zap.NewProductionConfig()
		//	fmt.Println("")
		//	cfg.OutputPaths = []string{`C:\Users\en.b\Desktop\test\src\run.log`}
		//	logger, err = cfg.Build()
	})

	if err != nil {
		return nil, err
	}
	return logger, nil
}
