package main

import (
	"github.com/boyter/scc/processor"
	"github.com/spf13/cobra"
	"os"
)

//go:generate go run scripts/include.go
func main() {
	const version = "2.6.0"

	// include in text output in some formats
	processor.Version = version

	//f, _ := os.Create("scc.pprof")
	//pprof.StartCPUProfile(f)
	//defer pprof.StopCPUProfile()

	rootCmd := &cobra.Command{
		Use:     "scc",
		Short:   "scc [FILE or DIRECTORY]",
		Long:    "Sloc, Cloc and Code. Count lines of code in a directory with complexity estimation.\nBen Boyter <ben@boyter.org> + Contributors",
		Version: version,
		Run: func(cmd *cobra.Command, args []string) {
			processor.DirFilePaths = args
			processor.ConfigureGc()
			processor.ConfigureLazy(true)
			processor.Process()
		},
	}

	flags := rootCmd.PersistentFlags()

	flags.Int64Var(
		&processor.AverageWage,
		"avg-wage",
		56286,
		"average wage value used for basic COCOMO calculation",
	)
	flags.BoolVar(
		&processor.DisableCheckBinary,
		"binary",
		false,
		"disable binary file detection",
	)
	flags.BoolVar(
		&processor.Files,
		"by-file",
		false,
		"display output for every file",
	)
	flags.BoolVar(
		&processor.Cocomo,
		"cocomo",
		false,
		"remove COCOMO calculation output",
	)
	flags.BoolVar(
		&processor.Ignore,
		"no-ignore",
		false,
		"disables .ignore file logic",
	)
	flags.BoolVar(
		&processor.GitIgnore,
		"no-gitignore",
		false,
		"disables .gitignore file logic",
	)
	flags.BoolVar(
		&processor.Debug,
		"debug",
		false,
		"enable debug output",
	)
	flags.StringSliceVar(
		&processor.PathBlacklist,
		"exclude-dir",
		[]string{".git", ".hg", ".svn"},
		"directories to exclude",
	)
	flags.IntVar(
		&processor.GcFileCount,
		"file-gc-count",
		10000,
		"number of files to parse before turning the GC on",
	)
	flags.StringVarP(
		&processor.Format,
		"format",
		"f",
		"tabular",
		"set output format [tabular, wide, json, csv, cloc-yml]",
	)
	flags.StringSliceVarP(
		&processor.WhiteListExtensions,
		"include-ext",
		"i",
		[]string{},
		"limit to file extensions [comma separated list: e.g. go,java,js]",
	)
	flags.BoolVarP(
		&processor.Languages,
		"languages",
		"l",
		false,
		"print supported languages and extensions",
	)
	flags.BoolVarP(
		&processor.Complexity,
		"no-complexity",
		"c",
		false,
		"skip calculation of code complexity",
	)
	flags.BoolVarP(
		&processor.Duplicates,
		"no-duplicates",
		"d",
		false,
		"remove duplicate files from stats and output",
	)
	flags.StringArrayVarP(
		&processor.Exclude,
		"not-match",
		"M",
		[]string{},
		"ignore files and directories matching regular expression",
	)
	flags.StringVarP(
		&processor.FileOutput,
		"output",
		"o",
		"",
		"output filename (default stdout)",
	)
	flags.StringVarP(
		&processor.SortBy,
		"sort",
		"s",
		"files",
		"column to sort by [files, name, lines, blanks, code, comments, complexity]",
	)
	flags.BoolVarP(
		&processor.Trace,
		"trace",
		"t",
		false,
		"enable trace output. Not recommended when processing multiple files",
	)
	flags.BoolVarP(
		&processor.Verbose,
		"verbose",
		"v",
		false,
		"verbose output",
	)
	flags.BoolVarP(
		&processor.More,
		"wide",
		"w",
		false,
		"wider output with additional statistics (implies --complexity)",
	)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
