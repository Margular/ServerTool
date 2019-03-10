package main

import "flag"

var host = flag.String("l", "0.0.0.0", "监听的IP")
var port = flag.Int("p", 80, "监听的端口")

func init() {
	flag.Parse()
}