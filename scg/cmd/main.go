package main

import (
	"flag"
	"fmt"
	"github.com/ThCompiler/go_game_constractor"
	"github.com/ThCompiler/go_game_constractor/scg/expr"
	"github.com/ThCompiler/go_game_constractor/scg/generator"
	"os"
	"strings"
)

func main() {
	var (
		out     = flag.String("output", "", "")
		script  = flag.String("script", "", "")
		version = flag.Bool("version", true, "")
	)
	{
		flag.Parse()
		if *version {
			fmt.Printf(go_game_constractor.Version)
			return
		}

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
