package log

import "go.uber.org/zap"

var Sugared *zap.SugaredLogger

// ConfigureLogger initialize sugar sugar as global scope variable
func ConfigureLogger() {
	productionLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	Sugared = productionLogger.Sugar()
}
