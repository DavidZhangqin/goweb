package main

import log "github.com/cihub/seelog"

func main() {
	defer log.Flush()
	log.Info("Hello from Seelog!")
	consoleWriter, _ := log.NewConsoleWriter()
	formatter, _ := log.NewFormatter("%Level %Msg %File%n")
	root, _ := log.NewSplitDispatcher(formatter, []interface{}{consoleWriter})
	constraints, _ := log.NewMinMaxConstraints(log.TraceLvl, log.CriticalLvl)
	specificConstraints, _ := log.NewListConstraints([]log.LogLevel{log.InfoLvl, log.ErrorLvl})
	ex, _ := log.NewLogLevelException("*", "*main.go", specificConstraints)
	exceptions := []*log.LogLevelException{ex}
	logger := log.NewAsyncLoopLogger(log.NewLoggerConfig(constraints, exceptions, root))
	log.ReplaceLogger(logger)
	log.Trace("This should not be seen")
	log.Debug("This should not be seen")
	log.Info("Test")
	log.Error("Test2")
}
