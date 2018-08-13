package main

import (
	"github.com/dancing-koala/iridium-go/pkg/iridium"
)

func main() {
	repl := iridium.NewREPL()
	repl.Run()
}
