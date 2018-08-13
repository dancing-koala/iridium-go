package iridium

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	CMD_QUIT    = ".quit"
	CMD_HISTORY = ".history"
)

type REPL struct {
	commandBuffer []string
	vm            *VM
}

func NewREPL() *REPL {
	return &REPL{
		commandBuffer: make([]string, 0, 8),
		vm:            New(),
	}
}

func (repl *REPL) Run() {
	fmt.Println("Welcome to Iridium! Let's be productive!")

	reader := bufio.NewReader(os.Stdin)
	delimiter := byte('\n')
	linePrefix := ">>> "

	for {
		fmt.Print(linePrefix)

		input, err := reader.ReadString(delimiter)

		if err != nil {
			fmt.Println(err)
			return
		}

		// convert CRLF to LF
		input = strings.Replace(input, "\n", "", -1)
		input = strings.TrimSpace(input)

		repl.commandBuffer = append(repl.commandBuffer, input)

		switch input {
		case CMD_QUIT:
			fmt.Println("Farewell! Have a great day!")
			return

		case CMD_HISTORY:
			for _, val := range repl.commandBuffer {
				fmt.Println(val)
			}

		default:
			fmt.Println("Invalid input!")
		}
	}
}
