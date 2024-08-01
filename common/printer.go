package common

import (
	"fmt"
	"os"
	"r03921081/vfs/constant"
)

type IPrinter interface {
	PrintSuccess(message string)
	PrintWarning(message string)
	PrintError(message string)
}

var Printer IPrinter = &printerImpl{}

type printerImpl struct{}

func (p *printerImpl) PrintSuccess(message string) {
	fmt.Fprintln(os.Stdout, message)
}

func (p *printerImpl) PrintWarning(message string) {
	fmt.Fprintln(os.Stdout, constant.PrefixWaring+message)
}

func (p *printerImpl) PrintError(message string) {
	fmt.Fprintln(os.Stderr, constant.PrefixError+message)
}
