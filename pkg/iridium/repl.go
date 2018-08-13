package iridium

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	CMD_QUIT      = ".quit"
	CMD_HISTORY   = ".history"
	CMD_PROGRAM   = ".program"
	CMD_REGISTERS = ".registers"
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

		case CMD_PROGRAM:
			fmt.Println("Listing instructions currently in VM's program slice:")

			for _, val := range repl.vm.program {
				fmt.Println(val)
			}

			fmt.Println("End of program listing")

		case CMD_REGISTERS:
			fmt.Println("Listing registers and all contents:")
			fmt.Println(repl.vm.registers)
			fmt.Println("End of registers listing")

		default:
			values, err := parseHex(input)

			if err != nil {
				fmt.Println("Could not parse input, please enter 4 groups of 2 characters or a command.")
			} else {
				repl.vm.addToProgram(values...)
				repl.vm.runOnce()
			}
		}
	}
}

func parseHex(hex string) ([]uint8, error) {

	parts := strings.Split(hex, " ")
	result := make([]uint8, len(parts))

	for i := range parts {
		value, err := strconv.ParseUint(parts[i], 16, 8)

		if err != nil {
			return nil, err
		}

		result[i] = uint8(value)
	}

	return result, nil
}
