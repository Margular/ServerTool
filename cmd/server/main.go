package main

import (
	"flag"
	"github.com/Margular/ServerTool/server"
)

func main() {
	flag.StringVar(&server.GlobalOptions.Host, "h", "0.0.0.0", "监听的IP")
	flag.UintVar(&server.GlobalOptions.Port, "p", 80, "监听的端口")
	flag.StringVar(&server.GlobalOptions.DownloadBasedir, "b", "download", "服务器下载目录")
	flag.BoolVar(&server.GlobalOptions.Debug, "d", false, "调试选项")

	flag.Parse()

	server.Run()
}
