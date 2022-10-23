package main

import (
	"github.com/tianshengdiwangK/simProject/router"
)

func main() {
	router.NewRouter().Run(":9090")

}
