# Wolfpack

Do you need to access a woocommerce database and list all subscriptions compared to an excel file and update the status column? Meet Wolfpack, A dead simple configuration management tool that is powered by shell scripts.

Wolfpack is an idea to combine the shell scripts philosophy of [fss](https://github.com/brandonhilkert/fucking_shell_scripts) with the simplicity of a `Dockerfile`.

[![PkgGoDev](https://pkg.go.dev/badge/github.com/felipemarques-rec/wolfpack)](https://pkg.go.dev/github.com/felipemarques-rec/wolfpack) [![Go Report Card](https://goreportcard.com/badge/github.com/felipemarques-rec/wolfpack)](https://goreportcard.com/report/github.com/felipemarques-rec/wolfpack) [![Build Status](https://travis-ci.com/felipemarques-rec/wolfpack.svg?branch=main)](https://travis-ci.com/felipemarques-rec/wolfpack) [![Coverage Status](https://coveralls.io/repos/github.com/felipemarques-rec/wolfpack/badge.svg)](https://coveralls.io/github.com/felipemarques-rec/wolfpack)

## Getting Started

Let's take a look at how easy it is to use Wolfpack.

### Installation

Wolfpack is simple to install, with the fastest method being to download one of our [binary releases](https://github.com/felipemarques-rec/wolfpack/releases).

It is also possible to install Wolfpack with Go (requires v1.14+).

```console
go get -u github.com/felipemarques-rec/wolfpack
```

Once installed, we can start defining our steps.

#### Command-Line Options

Wolfpack offers several additional options, such as parallel execution and various authentication methods.

```console
  -v, --verbose   Enable verbose output
  -q, --quiet     Silence output
  -p, --parallel  Execute tasks across multiple hosts in parallel
  -d, --dryrun    Print tasks to be executed without actually executing any tasks
      --url=      Ask for a Wordpress URL
      --key=      Specify an key generate by woocommerce
      --path      Ask for a Path where the csv file is to be updated 
      --secret=   Ask for a Secret generate by woocommerce to use for woocommerce authentication
```

## Call to Action

Wolfpack is a small project to fit the fine line between complex configuration management and simple shell scripts.

For those interested in helping develop Wolfpack. The time, skills, and perspectives you contribute to this project are valued. Please reference our [Issues Page](https://github.com/felipemarques-rec/wolfpack/issues) for open ideas and our [Contributing Guide](CONTRIBUTING.md) for contribution details.

If you like Wolfpack, please tell others about it by sharing this project on the social media site of your choice.
