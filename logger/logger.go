package logger

import (
	"fmt"
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"github.com/gookit/slog/rotatefile"
	"zph/config"
)

var Log *slog.Logger

//func init() {
//	getLogFileDate()
//	f, err := handler.NewFileHandler("./logs/"+logSuffix+"-slog.log", false)
//	if err != nil {
//		slog.Panicf("init slog fail, err: %+v", err)
//	}
//	f.Configure(func(h *handler.FileHandler) {
//		h.NoBuffer = true
//	})
//
//	Log = slog.NewWithHandlers(f)
//}
//
//func getLogFileDate() {
//	today := time.Now().Format("20060102")
//	if today > logSuffix {
//		logSuffix = today
//	}
//}

var logLevelMap = map[string]uint32{
	"debug": uint32(slog.DebugLevel),
	"info":  uint32(slog.InfoLevel),
	"warn":  uint32(slog.WarnLevel),
	"error": uint32(slog.ErrorLevel),
}

func init() {
	logPath := config.LogConfig().LogPath
	configLogLevel := config.LogConfig().LogLevel
	handlers := make([]slog.Handler, 0)
	if level, ok := logLevelMap[configLogLevel]; ok {
		for k, v := range logLevelMap {
			if v <= level {
				logFile := logPath + k + ".log"
				h := handler.MustRotateFile(logFile, 3*rotatefile.EveryHour, handler.WithLogLevels(slog.Levels{slog.Level(v)}),
					handler.WithUseJSON(false),
					handler.WithBuffSize(0),
				)
				handlers = append(handlers, h)
			}
		}
	} else {
		panic(fmt.Errorf("not allowed level"))
	}
	l := slog.NewWithHandlers(handlers...)
	Log = l
}

type Slog struct {
	slog.Handler
}
