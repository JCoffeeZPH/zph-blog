package main

import (
	"time"
	"zph/logger"
)

//var log = logger.Log

func main() {
	//go handler.SyncBlogDataHandler()
	//r := router.InitRouter()
	//err := r.Run(":8282")
	//if err != nil {
	//	log.Errorf("r.Run error. err: %+v", err)
	//}
	log := logger.Log
	for i := 0; i < 2; i++ {
		log.Debugf("test debug, time: %d", i)
		log.Infof("test info, time: %d", i)
		log.Errorf("test error, time: %d", i)
		time.Sleep(2 * time.Second)
	}

}
