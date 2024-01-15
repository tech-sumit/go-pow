package pkg

import "go.uber.org/zap"

func SetupLogger() (*zap.SugaredLogger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	defer logger.Sync() // Flush the logger before exiting
	return logger.Sugar(), nil
}
