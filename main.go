package main

import "go_project/gin_project/router"

func main() {
	r := router.SetupRouter()
	r.Run()

}
