package main

import (
	"fmt"
	"strconv"
	"io"
	"os"
	"flag"
	"github.com/johnryan-wf/csvmaker/generator"
)

func main() {
	var rows *string = flag.String("rows", "", "number of rows")
	var cols *string = flag.String("cols", "", "number of columns")
	flag.Parse()

	var rowsInt int64
	var colsInt int64
	var err error
	rowsInt, err = strconv.ParseInt(*rows, 0, 0)
	if err != nil {
		fmt.Println("no cols specified")
		flag.Usage()
		return
	}

	colsInt, err = strconv.ParseInt(*cols, 0, 0)
	if err != nil {
		fmt.Println("no rows specified")
		flag.Usage()
		return
	}

	g := generator.New(int(rowsInt), int(colsInt))
	io.Copy(os.Stdout, g)
}

