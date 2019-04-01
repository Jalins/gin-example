package main

import "gin-example/routers"

func main() {
	r := routers.Router()
	r.Run()
}
