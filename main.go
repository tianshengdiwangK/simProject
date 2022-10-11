package main

import log "github.com/tianshengdiwangK/simProject/log"

func init() {
	log.InitLogger() //初始化log
}
func main() {
	log.CwLog().Warn("this is a test!") //测试log
}
