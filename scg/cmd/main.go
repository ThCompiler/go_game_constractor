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
		out     = flag.String("output", "", "Path to directory, where you want generate code for script")
		script  = flag.String("script", "", "Path to config file as yaml, json or xml with info about script")
		update  = flag.Bool("update", false, "Return current version")
		version = flag.Bool("version", false, "Return current version")
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

	outputs, err := generator.Generate(*out, *si, *update)
	if err != nil {
		fail(err.Error())
	}

	fmt.Println(strings.Join(outputs, "\n"))
}

func fail(msg string, vals ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, vals...)
	os.Exit(1)
}
