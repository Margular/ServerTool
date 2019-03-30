package server

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

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

	// 文件合法性校验
	if !regexp.MustCompile(`^[a-zA-Z0-9._]+$`).MatchString(fileName) {
		notFound(c)
		return
	}

	targetPath := filepath.Join(GlobalOptions.DownloadBasedir, fileName)

	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		notFound(c)
		return
	}

	//Seems this headers needed for some browsers (for example without this headers Chrome will download files as txt)
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/octet-stream")
	c.File(targetPath)
}
