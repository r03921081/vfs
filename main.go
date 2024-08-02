package main

import (
	"bufio"
	"fmt"
	"os"
	"r03921081/vfs/common"
	"r03921081/vfs/constant"
	"r03921081/vfs/controller"
	"r03921081/vfs/util"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("# ")
		scanner.Scan()
		command := scanner.Text()

		if !util.IsValidInput(command, util.ValidCommand) {
			common.Printer.PrintError(constant.ErrMsgCommandShouldNotBeLongerThan)
			continue
		}

		controller.CommandController.Handle(command)
	}
}
