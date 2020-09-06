<h1 align="center">no more 5k</h1>

<p align="center">
  <img width="300" height="150" src="https://user-images.githubusercontent.com/15611134/92327311-ebc21000-f082-11ea-82a6-07a3b9f1516c.png" alt="The doctl mascot." />
</p>

<p align="center">
  <a href="https://github.com/mhughdo/nm5/workflows/AutoRelease/badge.svg">
    <img src="https://github.com/mhughdo/nm5/workflows/AutoRelease/badge.svg" alt="Build Status" />
  </a>
  <a href="https://godoc.org/github.com/mhughdo/nm5">
    <img src="https://godoc.org/github.com/mhughdo/nm5?status.svg" alt="GoDoc" />
  </a>
  <a href="https://goreportcard.com/report/github.com/mhughdo/nm5">
    <img src="https://goreportcard.com/badge/github.com/mhughdo/nm5" alt="Go Report Card" />
  </a>
</p>

```
nm5 is a command line interface (CLI) that ensure you won't lose 5000VND (anymore).

Usage:
  nm5 [command]
  nm5 [flags]

Available Commands:
  config          Show all configs
  cron            Run a cron job that automatically send message at 16:46
  help            Get help about any command
  run             Send message
  set-cookie      Set cookie in config file to send message
  set-token       Set token in config file (automatically set token using cookie if no arguments are passed) to send message
  set-room        Set room id

Flags:
  -h              Get helo about any command (work with all commands)

Use "nm5 [command] -h" for more information about a command.
```

- [Installing `nm5`](#installing-nm5)
  - [Downloading a Release from GitHub](#downloading-a-release-from-github)
  - [Building the Development Version from Source](#building-the-development-version-from-source)
- [Configuring Default Values](#configuring-default-values)
- [Uninstalling `nm5`](#uninstalling-nm5)
- [Examples](#examples)

## Installing `nm5`

### Downloading a Release from GitHub

Visit the [Releases
page](https://github.com/mhughdo/nm5/releases) for the
[`nm5` GitHub project](https://github.com/mhughdo/nm5), and find the
appropriate archive for your operating system and architecture.
Download the archive from from your browser or copy its URL and
retrieve it to your home directory with `wget` or `curl`.

For example, with `wget`:

```
cd ~
wget https://github.com/mhughdo/nm5/releases/download/v<version>/nm5_<version>_linux_amd64.tar.gz
```

Or with `curl`:

```
cd ~
curl -OL https://github.com/mhughdo/nm5/releases/download/v<version>/nm5_<version>_linux_amd64.tar.gz
```

Extract the binary:

```
tar xf ~/nm5_<version>_linux_amd64.tar.gz
```

Or download and extract with this oneliner:
```
curl -sL https://github.com/mhughdo/nm5/releases/download/v<version>/nm5_<version>_linux_amd64.tar.gz | tar -xzv
```

where `<version>` is the full semantic version, e.g., `1.1.1`.

Move the `doctl` binary to somewhere in your path. For example, on GNU/Linux and OS X systems:

```
sudo mv ~/nm5 /usr/local/bin
sudo chmod +x /usr/local/bin/nm5
```

Window version is not supported yet.


### Building the Development Version from Source

If you have a Go environment configured, you can install the development version of `nm5` from
the command line.

```
go get github.com/mhughdo/nm5
```

While the development version is a good way to take a peek at
`nm5`'s latest features before they get released, be aware that it
may have bugs. Officially released versions will generally be more
stable.


## Configuring Default Values

The `nm5` configuration file is stored at $HOME/.no-more-5k.yaml

You can directly change the config file but make sure file format stays the same.

Save and close the file. The next time you use `nm5`, the new default values you set will be in effect.

## Uninstalling `nm5`

Remove binary file (if you already move it to /usr/local/bin bin folder)

```
rm -rf /usr/local/bin/nm5
```

Remove config file
```
rm -rf $HOME/.no-more-5k.yaml
```


## Examples

Below are a few common usage examples.

* Send message:
```
nm5 run
```
* Set cron job:
```
nm5 cron
```

* Get help about set cookie command:
```
nm5 sc -h
```

* Set cookie:
```
nm5 sc 3tg2ku3a4vs2jmofu6paptrcne
```
* Set token:
```
nm5 st or nm5 st b229775d5036c15990287c7979d84a511e4096235f51f6779dca7
```
* Set room id (default: 195722902):
```
nm5 sr 195722902
```
