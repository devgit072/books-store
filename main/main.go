package main

import "github.com/devgit072/books-store/web_server"

func main() {
	c := &web_server.Controller{}

	c.StartServer()
}
