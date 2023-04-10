package main

import (
	"learn-gin/routers"
)

func main() {
	PORT := ":8000"
	routers.StartServer().Run(PORT)
}