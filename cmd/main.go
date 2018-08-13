package main

import (
	"github.com/dancing-koala/iridium-go/pkg/repl"
)

func main() {
	interpreter := repl.New()
	interpreter.Run()
}
