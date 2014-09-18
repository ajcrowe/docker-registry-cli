package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func getTabWriter() (w *tabwriter.Writer) {
	w = new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 0, 3, ' ', 0)
	return w
}

func writeLine(w *tabwriter.Writer, values ...string) {

	line := ""
	for _, v := range values {
		line += v + "\t"
	}
	fmt.Fprintln(w, line)
}

func writeHeader(w *tabwriter.Writer, values ...string) {

	header := ""
	for _, v := range values {
		header += v + "\t"
	}
	fmt.Fprintln(w, header)
}
