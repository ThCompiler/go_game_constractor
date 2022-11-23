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
		help    bool
		server  bool
	)
	{
		flagStringVar(&out, "output", "o", "", "Path to directory, where you want generate code for script")
		flagStringVar(&script, "script", "s", "", "Path to config file as yaml, json or xml with info about script")
		flagBoolVar(&update, "update", "u", false, "Show program version")
		flagBoolVar(&version, "version", "v", false, "Save user changes in files")
		flagBoolVar(&help, "help", "h", false, "Save user changes in files")
		flagBoolVar(&server, "http-server", "", false, "Save user changes in files")
	}

	{
		flag.Parse()
		if help {
			printHelp()

			return
		}

		if version {
			_, _ = fmt.Printf("%s", go_game_constractor.Version) //nolint:forbidigo //golangci not support forbidigo default patterns

			return
		}

		if out == "" {
			failFlag("Missing output flag")
		}

		if script == "" {
			failFlag("Missing script flag")
		}
	}

	si, err := expr.LoadScriptInfo(script)
	if err != nil {
		fail(err.Error())
	}

	outputs, err := generator.Generate(out, *si, update, server)
	if err != nil {
		fail(err.Error())
	}

	_, _ = fmt.Printf("%s", strings.Join(outputs, "\n")) //nolint:forbidigo //golangci not support forbidigo default patterns
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

func failFlag(msg string, vals ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, "Incorrect option: "+msg+"\n", vals...)

	printHelp()

	os.Exit(1)
}

func printHelp() {
	_, _ = fmt.Printf("%s", //nolint:forbidigo //golangci not support forbidigo default patterns
		`SCG - script generator. Generate script structs, functions for store texts of script in redis from yml, or json, or xml file.

Usage:
    scg ( (-o | --output=<file>) (-s | --script=<file>) | [options] | (-v | --version) | (-h | --help) )

Options:
    -u --update         save user changes in files
    --http-server       generates a basic http server 

Other args:
    -o --output=<file>  path to dir where need generate files
    -s --script=<file>  path to config file
    -v --version        show program version
    -h --help           help info

Note:
    With the --update flag, user changes are saved unchanged. Comments are embedded in the code with the code 
    that was generated based on the new initializing file. These comments are limited to the lines // >>>>>>> Generated.
    The decision to apply the changes remains with you, as well as the decision to remove unnecessary functionality.`,
	)
}
