#latch-cmd

**latch-cmd** is an unofficial command line tool that lets you interact with the Latch API ([https://latch.elevenpaths.com/](https://latch.elevenpaths.com/ "https://latch.elevenpaths.com/")). With this tool, you can call all the functions of the API directly from the shell in a really easy way. *latch-cmd* is written in the [Go](https://golang.org/) programming language and it's available as a single executable file with no dependencies for all the major operating systems (thanks to Go's cross-compile capabilities).

##Installation & Building

The easiest way to install `latch-cmd` on your system is to download one of the precompiled binaries that match your operating system and architecture:

* Windows: 32 bits / 64 bits
* Mac OS X: 32 bits / 64 bits
* GNU/Linux: 32 bits / 64 bits
* FreeBSD: 32 bits / 64 bits

and then place the executable file (`latch-cmd` or `latch-cmd.exe` for Windows) somewhere in your PATH (`/usr/bin` on UNIX-based systems for example). Make sure you have permissions to execute this file. Once installed, you can run this program by opening a terminal and writing:

```bash
$ latch-cmd
```

which will output the following usage (help) information page:

```
Latch-cmd is an unofficial command line tool that lets you interact with the Latch API (https://latch.elevenpaths.com/).

Usage:
  latch-cmd [flags]
  latch-cmd [command]

Available Commands:
  app         Set of commands to interact with the main application API.
  user        Set of commands to interact with the user API (manage applications and subscription information).
  about       Version and authorship information

Flags:
  -h, --help   help for latch-cmd

Use "latch-cmd [command] --help" for more information about a command.
```

You can also build *latch-cmd* from the source code contained in [this github repository](https://github.com/millenc/latch-cmd "this repository"). In order to do so, you must have Go previously installed on your system (please refer to the official documentation on how to install and configure Go if you don't have it already). Then follow the next steps:

1: Get the code

``` bash
$ go get github.com/millenc/latch-cmd 
```
2: Move to the package source folder:

``` bash
$ cd $GOPATH/src/github.com/millenc/latch-cmd
```

2: Build the code

``` bash
$ go build
```
That's it! Go will compile the source code into one single executable file named `latch-cmd` (or `latch-cmd.exe` on Windows).

##Getting help

You can get help about any command using the `--help` (`-h`) flag or the `help` command. For example, to get help about the `app status` (command to get the status of an account) you can use:

```bash
$ latch-cmd app status --help
$ latch-cmd app status -h
$ latch-cmd help app status
```
which will show all the flags available for that command and other usage information:

```
Gets the current status of an account using it's account ID (--account).

Usage:
  latch-cmd app status [flags]

Flags:
  -i, --account string   Account ID
  -b, --bare             Bare output (print only essential information, useful when handling the results in shell scripts for example)
  -n, --nootp            No OTP
  -l, --silent           Silent (requires SILVER, GOLD or PLATINUM subscription)

Global Flags:
  -a, --app string      Application's ID
  -w, --no-shadow       Don't hide secret keys
  -p, --proxy string    Proxy URL
  -s, --secret string   Secret key
  -v, --verbose         Display additional information about what's going on on each call
```


