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
		out     string
		script  string
		update  bool
		version bool
	)
	{
		flagStringVar(&out, "output", "o", "", "Path to directory, where you want generate code for script")
		flagStringVar(&script, "script", "s", "", "Path to config file as yaml, json or xml with info about script")
		flagBoolVar(&update, "update", "u", false, "Show program version")
		flagBoolVar(&version, "version", "v", false, "Save user changes in files")
	}

	{
		flag.Parse()
		if version {
			fmt.Printf(go_game_constractor.Version)
			return
		}

		if out == "" {
			fail("missing output flag")
		}

		if script == "" {
			fail("missing script flag")
		}
	}

	si, err := expr.LoadScriptInfo(script)
	if err != nil {
		fail(err.Error())
	}

	outputs, err := generator.Generate(out, *si, update)
	if err != nil {
		fail(err.Error())
	}

	fmt.Println(strings.Join(outputs, "\n"))
}

func flagStringVar(str *string, longName string, shortName string, defaultValue string, description string) {
	flag.StringVar(str, longName, defaultValue, description)
	flag.StringVar(str, shortName, defaultValue, description)
}

func flagBoolVar(str *bool, longName string, shortName string, defaultValue bool, description string) {
	flag.BoolVar(str, longName, defaultValue, description)
	flag.BoolVar(str, shortName, defaultValue, description)
}

func fail(msg string, vals ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, msg, vals...)
	os.Exit(1)
}
