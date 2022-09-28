package main

import (
	"bufio"
	"fmt"
	"os"
)

type AppContext struct {
	// TODO: replace following with appropriate properties.
	X      string
	result string
}

type IAppContext interface {
	process()
	printResult()
}

func (p *AppContext) read(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func (p *AppContext) process() {

	// TODO: replace following with appropriate implementation for processing.
	p.result = fmt.Sprintf("%s", p.X)
}

func (p *AppContext) printResult() {

	// TODO: replace following with appropriate implementation for printing.
	fmt.Println(p.result)
}

func NewAppContext(sc *bufio.Scanner) *AppContext {
	context := &AppContext{}
	// TODO: replace following with appropriate input reading.
	context.X = context.read(sc)
	return context
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var context IAppContext = NewAppContext(sc)
	context.process()
	context.printResult()
}
