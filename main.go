package main

import (
	"os"
	"wolfpack/app"
	"wolfpack/config"

	"github.com/fatih/color"
	flags "github.com/jessevdk/go-flags"
)

// options are command-line options that are provided by the user.
type options struct {
	Verbose  bool   `short:"v" long:"verbose" description:"Enable verbose output"`
	Quiet    bool   `short:"q" long:"quiet" description:"Silence output"`
	Parallel bool   `short:"p" long:"parallel" description:"Execute tasks across multiple hosts in parallel"`
	DryRun   bool   `short:"d" long:"dryrun" description:"Print tasks to be executed without actually executing any tasks"`
	Url      string `long:"url" description:"Ask for a Wordpress URL"`
	Key      string `long:"key" description:"Specify an key generate by woocommerce"`
	Path     string `long:"path" description:"Ask for a Path where the csv file is to be updated"`
	Secret   string `long:"secret" description:"Ask for a Secret generate by woocommerce to use for woocommerce authentication"`
}

// main runs the command line parsing and validations. This function will also start the application logic execution.
func main() {
	// Parse command line arguments
	var opts options
	_, err := flags.ParseArgs(&opts, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}

	// Convert to internal config
	cfg := config.New()
	cfg.Url = opts.Url
	cfg.Key = opts.Key
	cfg.Path = opts.Path
	cfg.Secret = opts.Secret
	cfg.Parallel = opts.Parallel
	cfg.DryRun = opts.DryRun

	// Run the App
	err = app.Run(cfg)
	if err != nil {
		if !cfg.Quiet {
			color.Red("Error executing: %s", err)
		}
		os.Exit(1)
	}
	if !cfg.Quiet {
		color.Green("Execution completed successfully")
	}
}
