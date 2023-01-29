package main

import (
	"github.com/pkg/errors"

	"github.com/olegshishkin/financier/cmd/web/di"
)

func main() {
	router := di.WireStubs()

	if err := router.Run(); err != nil {
		panic(errors.Wrap(err, "Web Server error"))
	}
}
