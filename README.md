# hellogo-cli

An example of a simple Golang CLI for helping others get Go'ing

This is meant as a learning tool and nothing else. Do NOT depend on this for anything!

## Root Command

```shell
$ ./hellogo-cli

This CLI used to show others a working model of the various
features that Cobra offers. It's also a great example of just
how easy it is to build a portable CLI with Go + Cobra

Usage:
  hellogo-cli [command]

Available Commands:
  command1    Command 1 does a thing
  command2    Command 2 has flags
  command3    Command 3 has sub-commands
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/.hellogo-cli.yaml)
  -h, --help            help for hellogo-cli
  -t, --toggle          Help message for toggle

Use "hellogo-cli [command] --help" for more information about a command.
```

## Command1

Command 1 will list files in your current working directory

```shell
$ ./hellogo-cli command1

command1 called
.git
LICENSE
README.md
cmd
go.mod
go.sum
hellogo-cli
main.go
pkg
```

## Command 2

Command 2 has flags that can be used to influence the execution of code

```none
Usage:
  hellogo-cli command2 [flags]

Flags:
  -h, --help          help for command2
  -n, --name string   A name to reference for a custom greeting (default "Stranger")
  -t, --toggle-name   Toggle name specific greeting
```

Example 1:

```shell
$ ./hellogo-cli command2

command2 called

Hello Stranger, I hope you are having a great day!
```

Example 2:

```shell
$ ./hellogo-cli command2 --name Joe

command2 called

Hello Joe, I hope you are having a great day!
```

Example 3:

```shell
$ ./hellogo-cli command2 --toggle-name

command2 called

I hope you are having a great day!
```

## Command 3

```none
Usage:
  hellogo-cli command3 [flags]
  hellogo-cli command3 [command]

Available Commands:
  subcommand1 This is the first sub-command for Command 3
  subcommand2 This is the second sub-command for Command 3
```

Example 1:

```shell
$ ./hellogo-cli command3 subcommand1

subcommand1 called
```

## Building CLI Binary

There are many ways to build a Go binary, but one of the simplest methods is this:

```shell
$ go build
```

## Bootstrapping CLI files

[Cobra](https://github.com/spf13/cobra) makes it really easy to bootstrap/scaffold a CLI project. Here are the steps I used for this example:

### Start with fresh Git repo cloned locally**

```shell
$ git clone git@github.com:phenixblue/hellogo-cli.git
```

### Initialize your Go project

```shell
$ go mod init github.com/phenixblue/hellogo-cli
```

### Intialize your Cobra application

```shell
$ cobra init --pkg-name github.com/phenixblue/hellogo-cli -a "The WebRoot, Inc."
```

### Add Commands

```shell
$ cobra add command1
$ cobra add command2
$ cobra add command3
```

### Add Sub-Commands

This adds sub-commands by adding the `-p` flag and specifying `command3` as the parent command.

```shell
$ cobra add subcommand1 -p command3Cmd
$ cobra add subcommand2 -p command3Cmd
```
