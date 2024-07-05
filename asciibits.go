package main

import (
	"fmt"
	"os"

	"github.com/teris-io/cli"
)

func main() {

	asciiCmd := cli.NewCommand("ascii", "convert string of decimals to ascii characters.").
		WithShortcut("a").
		WithArg(cli.NewArg("string", "string of decimals")).
		WithOption(cli.NewOption("separator", "Separator used when parsing decimals").WithChar('s').WithType(cli.TypeString)).
		WithAction(ascii)

	decimalCmd := cli.NewCommand("decimal", "convert string of ascii characters to array of decimals.").
		WithShortcut("d").
		WithArg(cli.NewArg("string", "string of ascii characters")).
		WithAction(decimal)

	versionCmd := cli.NewCommand("version", "output current version").
		WithShortcut("v").
		WithAction(version)

	app := cli.New("asciibits - used to convert array of decimals in a string to human readable ascii characters or vice versa").
		WithOption(cli.NewOption("separator", "Separator to use when parsing string").WithChar('s').WithType(cli.TypeString)).
		WithCommand(asciiCmd).
		WithCommand(decimalCmd).
		WithCommand(versionCmd)

	os.Exit(app.Run(os.Args, os.Stdout))
}

func version(args []string, options map[string]string) int {
	fmt.Fprintln(os.Stdout, Version)
	return 0
}

func ascii(args []string, options map[string]string) int {
	if options["separator"] == "" {
		options["separator"] = " "
	}

	t, err := ParseDecimals(args[0], options["separator"])

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed parsing decimals: %s", err)
		return 1
	}

	fmt.Fprintln(os.Stdout, t.DecimalsToASCIIString())

	return 0
}

func decimal(args []string, options map[string]string) int {
	if options["separator"] == "" {
		options["separator"] = " "
	}

	t, err := StringToDecimals(args[0])

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed parsing string: %s", err)
		return 1
	}

	fmt.Fprintln(os.Stdout, t.String(options["separator"]))

	return 0
}
