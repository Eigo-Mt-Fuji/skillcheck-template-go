package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AppContext struct {
	// TODO: replace following with appropriate properties.
	X      int
	Y      int
	Result int
}

type IAppContext interface {
	Process()
	PrintResult()
}

func (p *AppContext) read(sc *bufio.Scanner) (int, int) {
	sc.Scan()
	line := strings.Split(sc.Text(), " ")

	X, err := strconv.Atoi(line[0])
	if err != nil {
		os.Exit(1)
	}
	Y, err := strconv.Atoi(line[1])
	if err != nil {
		os.Exit(1)
	}
	return X, Y
}

func (p *AppContext) Process() {

	// TODO: replace following with appropriate implementation for processing.
	p.Result = p.X + p.Y
}

func (p *AppContext) PrintResult() {

	// TODO: replace following with appropriate implementation for printing.
	fmt.Println(p.Result)
}

func NewAppContext(sc *bufio.Scanner) *AppContext {
	context := &AppContext{}
	// TODO: replace following with appropriate input reading.
	X, Y := context.read(sc)
	context.X = X
	context.Y = Y

	return context
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	// TODO: make sure required(or suitable) initial / maximum buffer size
	buf := make([]byte, 16*1024)
	sc.Buffer(buf, 256*1024)

	var context IAppContext = NewAppContext(sc)
	context.Process()
	context.PrintResult()
}
