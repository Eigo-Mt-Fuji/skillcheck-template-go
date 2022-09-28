package main

import (
	"bufio"
	"fmt"
	"os"
)

type AppContext struct {
	// TODO: replace following with appropriate properties.
	X      string
	Result string
}

type IAppContext interface {
	Process()
	PrintResult()
}

func (p *AppContext) read(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func (p *AppContext) Process() {

	// TODO: replace following with appropriate implementation for processing.
	p.Result = p.X
}

func (p *AppContext) PrintResult() {

	// TODO: replace following with appropriate implementation for printing.
	fmt.Println(p.Result)
}

func NewAppContext(sc *bufio.Scanner) *AppContext {
	context := &AppContext{}
	// TODO: replace following with appropriate input reading.
	context.X = context.read(sc)
	return context
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
        // TODO: make sure required(or suitable) initial / maximum buffer size 
        buf := make([]byte, 16 * 1024)
        sc.Buffer(buf, 256 * 1024)

	var context IAppContext = NewAppContext(sc)
	context.Process()
	context.PrintResult()
}
