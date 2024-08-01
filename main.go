package main

import (
	"bufio"
	"fmt"
	"os"
	"r03921081/vfs/controller"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("# ")
		scanner.Scan()
		command := scanner.Text()
		controller.CommandController.Handle(command)
	}
}
