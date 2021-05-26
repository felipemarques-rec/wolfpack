/*
Package config is the internal configuration.
*/
package config

// Config provides a configuration structure used within the Efs2 application.
type Config struct {
	// Verbose if set to true, will enable verbose output of command execution.
	Verbose bool

	// Quiet if set to true, will silence all output and focus only on exit codes.
	Quiet bool

	// Parallel if set to true, will trigger Efs2 to execute the Efs2file against all hosts in parallel.
	Parallel bool

	// DryRun if set to true, will prevent the execution of tasks against hosts. The end-user will see instructions logged to stdout.
	DryRun bool

	// Url
	Url string

	// Key
	Key string

	// Path
	Path string

	// Secret
	Secret string
}

// New will return Config populated with pre-defined defaults.
func New() Config {
	c := Config{}
	return c
}
