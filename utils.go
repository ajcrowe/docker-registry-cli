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

func statusOK(w *tabwriter.Writer, status int) bool {
	if status == 200 {
		return true
	} else if status == 404 {
		fmt.Fprintf(w, "Error: 404 - Not found")
	} else if status == 400 {
		fmt.Fprintf(w, "Error: 400 - Invalid data")
	} else if status == 401 {
		fmt.Fprintf(w, "Error: 401 - Requires authorisation")
	} else {
		fmt.Fprintf(w, "Error: %d - Unknown error", status)
	}
	return false
}
