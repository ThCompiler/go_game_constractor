package main

import (
	"flag"
	"fmt"
	"gameconstractor/scg/expr"
	"gameconstractor/scg/generator"
	"os"
	"strings"
)

func main() {
	var (
		out    = flag.String("output", "", "")
		script = flag.String("script", "", "")
	)
	{
		flag.Parse()
		if *out == "" {
			fail("missing output flag")
		}

		if *script == "" {
			fail("missing script flag")
		}
	}

	si, err := expr.LoadScriptInfo(*script)
	if err != nil {
		fail(err.Error())
	}

	outputs, err := generator.Generate(*out, *si)
	if err != nil {
		fail(err.Error())
	}

	fmt.Println(strings.Join(outputs, "\n"))
}

func fail(msg string, vals ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, vals...)
	os.Exit(1)
}
