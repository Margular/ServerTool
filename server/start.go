package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"strconv"
)

func Run() {
	if err := GlobalOptions.validate(); err != nil {
		log.Fatal(err)
	}

	if !GlobalOptions.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.NoRoute(notFound)

	r.GET("/ip", ip)
	r.GET("/download/:filename", download)
	log.Fatal(r.Run(net.JoinHostPort(GlobalOptions.Host, strconv.Itoa(int(GlobalOptions.Port)))))
}
