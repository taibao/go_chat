package main

import "go_chat/router"

func main() {
	r := router.Router()
	err := r.Run(":1234")
	if err != nil {
		return
	}
}
