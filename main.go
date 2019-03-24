package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

import "github.com/gin-gonic/gin"

var downloadPath = "./download"
var validFileName = regexp.MustCompile(`^[a-zA-Z0-9._]+$`).MatchString

var host = flag.String("l", "0.0.0.0", "监听的IP")
var port = flag.String("p", "80", "监听的端口")

func main() {
	flag.Parse()
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.NoRoute(notFound)

	r.GET("/ip", ip)
	r.GET("/download/:filename", download)
	log.Fatal(r.Run(net.JoinHostPort(*host, *port)))
}

func notFound(c *gin.Context) {
	c.String(http.StatusNotFound, "Not Found")
}

func ip(c *gin.Context) {
	if host, _, err := net.SplitHostPort(c.Request.RemoteAddr); err == nil {
		c.String(http.StatusOK, host)
	}
}

func download(c *gin.Context) {
	fileName := c.Param("filename")

	if !validFileName(fileName) {
		notFound(c)
		return
	}

	targetPath := filepath.Join(downloadPath, fileName)

	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		notFound(c)
		return
	}

	//Seems this headers needed for some browsers (for example without this headers Chrome will download files as txt)
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename=" + fileName )
	c.Header("Content-Type", "application/octet-stream")
	c.File(targetPath)
}