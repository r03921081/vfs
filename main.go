package main

import (
	"bufio"
	"fmt"
	"os"
	"r03921081/vfs/constant"
	"r03921081/vfs/controller"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("# ")
		scanner.Scan()
		command := scanner.Text()

		if strings.ToLower(command) == constant.CommandExit.String() {
			os.Exit(0)
		}

		controller.CommandController.Handle(command)
	}
}
