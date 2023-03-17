package main

import "github.com/olegshishkin/financier/cmd/web/di"

func main() {
	server := di.Wire()
	server.Start()
}
