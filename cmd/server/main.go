package main

import (
	"flag"
	"github.com/Margular/ServerTool/server"
)

func main() {
	flag.StringVar(&server.Server().Options().Host, "h", "0.0.0.0", "监听的IP")
	flag.UintVar(&server.Server().Options().Port, "p", 80, "监听的端口")
	flag.StringVar(&server.Server().Options().DownloadBasedir, "b", "download", "服务器下载目录")
	flag.BoolVar(&server.Server().Options().Debug, "d", false, "调试选项")

	flag.Parse()

	server.Server().Run()
}