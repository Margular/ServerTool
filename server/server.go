package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

// options of server
type options struct {
	Host            string
	Port            uint
	DownloadBasedir string
	Debug           bool
}

// validate parameters
func (opt *options) validate() error {
	if err := opt.validateHost(); err != nil {
		return err
	}

	if err := opt.validatePort(); err != nil {
		return err
	}

	return nil
}

func (opt *options) validateHost() error {
	ip := net.ParseIP(opt.Host)

	if ip == nil {
		return fmt.Errorf("error param host %s", opt.Host)
	}

	return nil
}

func (opt *options) validatePort() error {
	if opt.Port > 65535 {
		return fmt.Errorf("listen on port %d is impossible", opt.Port)
	}

	return nil
}

// only one server in lifecycle
var s *server

func Server() *server {
	if s == nil {
		s = &server{}
	}

	return s
}

type server struct {
	opt *options
}

func (s *server) Options() *options {
	if s.opt == nil {
		s.opt = &options{}
	}

	return s.opt
}

func (s *server) Run() {
	if err := s.opt.validate(); err != nil {
		log.Fatal(err)
	}

	if !s.opt.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// prevent dir brute forcing
	r.NoRoute(s.notFound)

	r.GET("/ip", s.ip)
	r.GET("/download/:filename", s.download)
	log.Fatal(r.Run(net.JoinHostPort(s.opt.Host, strconv.Itoa(int(s.opt.Port)))))
}

func (s *server) notFound(c *gin.Context) {
	c.String(http.StatusNotFound, "Not Found")
}

func (s *server) ip(c *gin.Context) {
	if host, _, err := net.SplitHostPort(c.Request.RemoteAddr); err == nil {
		c.String(http.StatusOK, host)
	}
}

func (s *server) download(c *gin.Context) {
	fileName := path.Clean(c.Param("filename"))

	// 文件合法性校验
	if strings.ContainsRune(fileName, os.PathSeparator) {
		s.notFound(c)
		if s.opt.Debug {
			log.Printf("detect path separator '%c' in filename", os.PathSeparator)
		}
		return
	}

	targetPath := filepath.Join(s.opt.DownloadBasedir, fileName)

	// 再次校验路径合法性
	if !strings.HasPrefix(targetPath, s.opt.DownloadBasedir) {
		s.notFound(c)
		if s.opt.Debug {
			log.Printf(`detect malicious filename "%s", download basedir "%s", target path %s`,
				fileName, s.opt.DownloadBasedir, targetPath)
		}
		return
	}

	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		s.notFound(c)
		return
	}

	//Seems this headers needed for some browsers (for example without this headers Chrome will download files as txt)
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/octet-stream")
	c.File(targetPath)
}