package main

import (
	"log"
	"log/syslog"
	"net/http"

	"github.com/goproxy/goproxy"
	"github.com/goproxy/goproxy/cacher"
)

func main() {
	log.Println("Starting Go modules proxy server ...")

	logger, err := syslog.NewLogger(syslog.LOG_NOTICE, log.LstdFlags)
	if err != nil {
		log.Fatalln(err)
	}

	// create and start go proxy instance
	gp := goproxy.New()
	gp.Cacher = &cacher.Disk{Root: "/var/lib/goproxy"}
	gp.ErrorLogger = logger

	http.ListenAndServe("0.0.0.0:8080", gp)
}
