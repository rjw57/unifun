package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"uctricks"
)

func main() {
	fontname := flag.String("f", "blackletter", "Specify which 'font' to use. Use -l to see all.")
	list := flag.Bool("l", false, "List fonts and exit.")
	noNewline := flag.Bool("n", false, "Don't print trailing newline.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: unifun [options] text\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	// print a list of fonts if asked
	if *list {
		for _, name := range uctricks.FontList() {
			fmt.Println(name)
		}
		return
	}

	// concatentate all arguments
	s := strings.Join(flag.Args(), " ")

	// output the result
	font, err := uctricks.FontNamed(*fontname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "No such font: %v", *fontname)
		os.Exit(1)
	}
	fmt.Print(font.Apply(s))

	// print a neline if required
	if !*noNewline {
		fmt.Println("")
	}
}
