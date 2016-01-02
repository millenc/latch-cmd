#latch-cmd

**latch-cmd** is an unofficial command line tool that lets you interact with the Latch API ([https://latch.elevenpaths.com/](https://latch.elevenpaths.com/ "https://latch.elevenpaths.com/")). With this tool, you can call all the functions of the API directly from the shell in a really easy way. *latch-cmd* is written in the [Go](https://golang.org/) programming language and it's available as a single executable file with no dependencies for all the major operating systems (thanks to Go's cross-compile capabilities).

##Installation & Building

The easiest way to install *latch-cmd* on your system is to download one of the precompiled binaries that match your operating system and architecture:

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
That's it! Go will compile the source code into one single executable file named *latch-cmd* (or `latch-cmd.exe` on Windows).

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

##Configuration

You don't need any configuration to start using *latch-cmd*, but there is some information that you have to provide to every command through flags like the Application ID (`--app`) and the Secret key (`--secret`) for the `app` subcommands for example. 

Typing this flags over and over can be very cumbersome. That's the reason why *latch-cmd* supports the concept of "12 factor" configuration, that lets you provide these values through flags, config files or environment values (or a combination of all of them). The values that can be configured are the following:

| Parameter  	  | Flag 			| Configuration		| Environment variable		|
| --------------- | --------------- | ----------------- | ------------------------- |
| Application ID  | --app (-a)		| app				| LATCH_APP					|
| Secret key	  | --secret (-s)   | secret			| LATCH_SECRET				|
| Proxy			  | --proxy (-p)    | proxy				| LATCH_PROXY				|
| User ID		  | --user (-u)		| user				| LATCH_USER				|
| User Secret key | --secret (-s)	| user_secret		| LATCH_USER_SECRET			|

When looking for these values, *latch-cmd* uses the following priority order (from highest to lowest):

* Flags: passed to the command like `--app` or `--secret`.
* Environment variables: always in uppercase and prefixed by the LATCH keyword like `LATCH_APP` or `LATCH_SECRET`.
* Configuration files: *latch-cmd* will look for configuration files named `latch-cmd` (plus extension) in the following folders (also from highest to lowest priority):
	* Current working directory (`.`).
	* `.latch-cmd` folder inside user's home directory (`$HOME/.latch-cmd`).
	* `/etc/latch-cmd` only in UNIX-like systems.

This means that, for example, if you have provided the Application ID through an environment variable and you also have it defined in a configuration file, the value from the environment variable will be used. When using configuration files, keep in mind that the first configuration file found will be used (following the priority order previously defined).

*latch-cmd* supports reading configuration files in the TOML (recommended), JSON, YAML and HCL formats. You must name the config file `latch-cmd` with the proper extension (`latch-cmd.toml` for TOML, `latch-cmd.json` for JSON and so on). Here's an example of one such file using the TOML format:

```toml
###################################################
#                                                 #
# latch-cmd.toml                                  #
# latch-cmd configuration (TOML)                  #
# you can use the `#` character to write comments #
#                                                 #
###################################################

#Application ID
app="2Wv8UqaT6iZRQEbyG9Kv"

#Secret key
secret="aDYA2qVAv8wLgawGBWxhkv3EuBUgw6RBCy3nRmgv"

#User ID
user="fWZWTpA4Hg2TsMmLMjb2"

#User secret key
user_secret="veRY4LZ7qKVwMgZWriFDgxPvWg7mrPskvYQWA7xm"

#Proxy
proxy="http://8.8.8.8:8080"
```

##Usage

###Return values

Except where noted, *latch-cmd* will return a `0` exit code if the command completes successfully or `-1` if there's any error.

###Application API (`app`)

You can issue the command:

```bash
$ latch-cmd app
```

to get a list of all the available subcommands to interact with Latch's main application API:

```bash
Set of commands to interact with the main application API.

Usage:
  latch-cmd app [flags]
  latch-cmd app [command]

Available Commands:
  pair        Pairs an account with the provided pairing token (--token).
  unpair      Unpairs an account using it's account ID (--account).
  status      Gets the current status of an account using it's account ID (--account).
  lock        Locks an account using it's account ID (--account).
  unlock      Unlocks an account using it's account ID (--account).
  operation   Manages Latch operations
  history     Gets history information about an account. You can filter events between the --from and --to dates.

Flags:
  -a, --app string      Application's ID
  -h, --help            help for app
  -w, --no-shadow       Don't hide secret keys
  -p, --proxy string    Proxy URL
  -s, --secret string   Secret key
  -v, --verbose         Display additional information about what's going on on each call

Use "latch-cmd app [command] --help" for more information about a command.
```

####Global flags

You can pass these flags to every subcommand of the `app` command (and some of them can be provided via configuration, see the appropriate section of this documentation):

* `--app` (`-a`): Application ID (mandatory).
* `--secret` (`-s`): Secret key (mandatory).
* `--proxy` (`-p`): URL of the proxy.
* `--verbose` (`-v`): Prints additional information about what's going on on each call.
* `--no-shadow` (`-w`): Don't hide secret keys.

####Pair (`app pair`)

Command to pair an account with a given token (`--token`, `-t`):

```bash
$ latch-cmd app pair --app=YourAppID --secret=YourSecretKey --token=YourPairingToken
```

If the pairing goes well, *latch-cmd* will print the newly created Account ID. You can use the `--bare` (`-b`) flag to get *latch-cmd* to print only this account ID, thus making it easier to get that value from shell scripts or other programs.

NOTE: From now on we will omit the `--app` and `--secret` flags from the examples for the sake of clarity (it's recommended to create a configuration file/use environment variables instead of passing these flags over and over).

####Unpair (`app unpair`)

Command to unpair an account, given it's Account ID (`--account`, `-i`):

```bash
$ latch-cmd app unpair --account=AccountIDYouWantToUnpair
```

####Status (`app status`)

Command to get the status of an account, given it's Account ID (`--account`, `-i`):

```bash
$ latch-cmd app status --account=YourAccountID
```
This command will return the following exit codes:

* `0`: Account is ON.
* `1`: Account is OFF.
* `-1`: Error.

You can also pass the following optional flags:

* `--nootp` (`-n`): No OTP, don't include the 1-time password in the response.
* `--silent` (`-l`): Silent, don't alert the user of the access.
* `--bare` (`-b`): Bare, prints bare information. With this flag, *latch-cmd* will print `on` if the account is ON or `off` if the account is OFF. If you don't pass the `--no-otp` flag it will also print the 1-time password and the generated time (unix epoch in ms) separated by the `:` character in the following format `[status]:[1-time password]:[generated time]`. This makes it really easy to parse these values in calling shell scripts or programs.

####Lock (`app lock`)

Command to lock an account using it's account ID (`--account`, `-i`). Requires a a GOLD or PLATINUM subscription in order to work:

```bash
$ latch-cmd app lock --account=YourAccountID
```

####Unlock (`app unlock`)

Command to unlock an account using it's account ID (`--account`, `-i`). Requires a a GOLD or PLATINUM subscription in order to work:

```bash
$ latch-cmd app unlock --account=YourAccountID
```

####Operation (`app operation`)

Set of command to manage operations. Prints usage information about the available subcommands (described in the following sections).

####Operation status (`app operation status`)

Command to get the status of an operation, given an Account ID (`--account`, `-i`) and an Operation ID (`--operation`, `-o`):

```bash
$ latch-cmd app operation status --account=YourAccountID --operation=YourOperationID
```

This command accepts the same flags and displays information in the same way as the `app status` command.

####Lock Operation (`app operation lock`)

Command to lock an operation, given an Account ID (`--account`, `-i`) and an Operation ID (`--operation`, `-o`):

```bash
$ latch-cmd app operation lock --account=YourAccountID --operation=YourOperationID
```

####Unlock Operation (`app operation unlock`)

Command to unlock an operation, given an Account ID (`--account`, `-i`) and an Operation ID (`--operation`, `-o`):

```bash
$ latch-cmd app operation unlock --account=YourAccountID --operation=YourOperationID
```

####Add Operation (`app operation add`)

Command to add a new operation. You must provide the parent's application or operation ID (`--parent`, `-i`) and a name for the operation (`--name`, `-n`):

```bash
$ latch-cmd app operation add --parent=ParentID --name=MyNewOperationName
```

If everything goes well, *latch-cmd* will print the ID of the newly created operation. You can also pass these optional flags:

* `--two-factor`, (`-t`): Configures the two-factor authentication setting. Possible values are MANDATORY, OPT_IN and DISABLED (default).
* `--lock-on-request`, (`-l`): Configures the lock on request setting. Possible values are MANDATORY, OPT_IN and DISABLED (default).
* `--bare`, (`-b`): Bare output, displays only the newly created operation's ID.

####Update Operation (`app operation update`)

Command to modify an existing operation. You must provide the operation's ID (`--operation`, `-o`):

```bash
$ latch-cmd app operation update --operation=YourOperationID
```

The values you can modify are:

* `--name`, (`-n`): Name of the operation.
* `--two-factor`, (`-t`): Configures the two-factor authentication setting. Possible values are MANDATORY, OPT_IN and DISABLED.
* `--lock-on-request`, (`-l`): Configures the lock on request setting. Possible values are MANDATORY, OPT_IN and DISABLED.

If you don't provide any of these flags they will keep their original values.

####Delete Operation (`app operation delete`)

Command to delete an existing operation, given it's Operation ID (`--operation`, `-o`):

```bash
$ latch-cmd app operation delete --operation=YourOperationID
```

####Show Operation (`app operation show`)

Command to get information about an operation, given it's Operation ID (`--operation`, `-o`):

```bash
$ latch-cmd app operation show --operation=YourOperationID
```
####History (`app history`)

Command to get history information about an account (`--account`, `-i`). You can use the flags `--from` (`-f`) and `--to` (`-t`) to filter events between those dates.For example, to get all the events that happened in the year 2015 you could do:

```bash
$ latch-cmd app history --account=YourAccountID --from="01-01-2015 00:00:00" --to="31-12-2015 23:59:59"
```

*latch-cmd* will show a table with the events. For example:

```
success  Last seen: 02-01-2016 13:36:58, Client version: [Android - 1.5.1,Android - 1.5.1], History count: 15

+---------------------+------------------+------------+-----+-------+------------------+-------------+-----------+
|        TIME         |      ACTION      |    WHAT    | WAS | VALUE |       NAME       | USER AGENT  |    IP     |
+---------------------+------------------+------------+-----+-------+------------------+-------------+-----------+
| 28-12-2015 12:18:56 | USER_UPDATE      | two_factor |     | on    | Test Application |             | 127.0.0.1 |
| 28-12-2015 12:19:41 | get              | status     | -   | on    | Test Application | Golatch 1.0 | 127.0.0.1 |
| 28-12-2015 12:20:41 | get              | status     | -   | on    | Test Application | Golatch 1.0 | 127.0.0.1 |
| 28-12-2015 12:27:38 | get              | status     | -   | on    | Test Application | Golatch 1.0 | 127.0.0.1 |
| 28-12-2015 12:27:56 | get              | status     | -   | on    | Test Application | Golatch 1.0 | 127.0.0.1 |
| 28-12-2015 12:35:19 | USER_UPDATE      | two_factor |     | on    | Test Application |             | 127.0.0.1 |
| 28-12-2015 12:36:14 | get              | status     | -   | on    | Test Application | Golatch 1.0 | 127.0.0.1 |
| 28-12-2015 12:36:33 | get              | status     | -   | on    | Test operation 1 | Golatch 1.0 | 127.0.0.1 |
| 28-12-2015 12:36:33 | DEVELOPER_UPDATE | status     |     | off   | Test operation 1 | Golatch 1.0 | 127.0.0.1 |
| 28-12-2015 12:37:21 | get              | status     | -   | off   | Test operation 1 | Golatch 1.0 | 127.0.0.1 |
| 28-12-2015 12:37:30 | USER_UPDATE      | status     | off | on    | Test operation 1 |             | 127.0.0.1 |
| 28-12-2015 12:37:39 | get              | status     | -   | on    | Test operation 1 | Golatch 1.0 | 127.0.0.1 |
| 28-12-2015 12:37:39 | DEVELOPER_UPDATE | status     | on  | off   | Test operation 1 | Golatch 1.0 | 127.0.0.1 |
| 28-12-2015 12:38:04 | get              | status     | -   | off   | Test operation 1 | Golatch 1.0 | 127.0.0.1 |
| 28-12-2015 12:40:29 | get              | status     | -   | off   | Test operation 1 | Golatch 1.0 | 127.0.0.1 |
+---------------------+------------------+------------+-----+-------+------------------+-------------+-----------+
```

###User API (`user`)

You can issue the command:

```bash
$ latch-cmd user
```

to get a list of all the available subcommands to interact with Latch's User API:

```bash
Set of commands to interact with the user API (manage applications and subscription information).

Usage:
  latch-cmd user [flags]
  latch-cmd user [command]

Available Commands:
  subscription Gets information about your current subscription.
  application  Manages Latch applications

Flags:
  -h, --help            help for user
  -w, --no-shadow       Don't hide secret keys
  -p, --proxy string    Proxy URL
  -s, --secret string   User secret key
  -u, --user string     User ID
  -v, --verbose         Display additional information about what's going on on each call

Use "latch-cmd user [command] --help" for more information about a command.
```

Please keep in mind that all these commands require a GOLD or PLATINUM subscription in order to work.

####Global flags

You can pass these flags to every subcommand of the `user` command (and some of them can be provided via configuration, see the appropriate section of this documentation):

* `--user` (`-u`): User ID (mandatory).
* `--secret` (`-s`): User Secret key (mandatory).
* `--proxy` (`-p`): URL of the proxy.
* `--verbose` (`-v`): Prints additional information about what's going on on each call.
* `--no-shadow` (`-w`): Don't hide secret keys.

####Subscription (`user subscription`)

Command to get information about your current subscription:

```batch
$ latch-cmd user subscription
```

####Applications (`user application`)

Set of commands to manage (create, update, delete,show) applications. Prints usage information about the available subcommands (described in the following sections):

```bash
$ latch-cmd user application
```

```bash
Manages Latch applications

Usage:
  latch-cmd user application [flags]
  latch-cmd user application [command]

Available Commands:
  add         Adds a new application
  update      Updates an existing application
  delete      Deletes an existing application
  show        Shows information about your applications

Flags:
  -h, --help   help for application

Global Flags:
  -w, --no-shadow       Don't hide secret keys
  -p, --proxy string    Proxy URL
  -s, --secret string   User secret key
  -u, --user string     User ID
  -v, --verbose         Display additional information about what's going on on each call

Use "latch-cmd user application [command] --help" for more information about a command.
```

####Add Application (`user application add`)

Command to create a new application. You must provide the name of the application through the `--name` (`-n`) flag:

```bash
$ latch-cmd user application add --name=MyNewApplication
```

If everything goes well, *latch-cmd* will create the new application and display the Application's ID and Secret Key (hidden by default, use the `--no-shadow` flag to see it). You can also use the `--bare` (`-b`) flag to display only the Application ID and Secret key separated by the `:` character (`[Application ID]:[Secret key]). This can be useful to process this values in calling shell scripts or programs.

You can also provide additional information using these optional flags:

* `--email` (`-e`): Contact email.
* `--phone` (`-c`): Contact phone.
* `--two-factor`, (`-t`): Configures the two-factor authentication setting. Possible values are MANDATORY, OPT_IN and DISABLED.
* `--lock-on-request`, (`-l`): Configures the lock on request setting. Possible values are MANDATORY, OPT_IN and DISABLED.

####Update Application (`user application update`)

Command to update an existing application. You must provide the Application's ID (`--app`, `-a`):

```bash
$ latch-cmd user application update --app=MyApplicationID
```

You can modify the following values:

* `--name` (`-n`): Name of the application
* `--email` (`-e`): Contact email.
* `--phone` (`-c`): Contact phone.
* `--two-factor`, (`-t`): Configures the two-factor authentication setting. Possible values are MANDATORY, OPT_IN and DISABLED.
* `--lock-on-request`, (`-l`): Configures the lock on request setting. Possible values are MANDATORY, OPT_IN and DISABLED.

Please note that if you don't pass any of these flags Latch will keep their original values.

####Delete Application (`user application delete`)

Command to delete an existing application, given it's Application's ID (`--app`, `-a`):

```bash
$ latch-cmd user application delete --app=MyApplicationID
```

####Show Applications (`user application show`)

Command to show information about the existing applications:

```bash
$ latch-cmd user application show
```