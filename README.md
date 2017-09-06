# pCloud-cli - pCloud Command Line Interface

[![Go Report Card](https://goreportcard.com/badge/github.com/storvik/pcloud-cli)](https://goreportcard.com/report/github.com/storvik/pcloud-cli)

Command line interface to pCloud API written in Go.

## Usage

```
pcloud-cli --help
```

Run `pcloud-cli authorize` to authorize pCloud-cli and save authorization data to `~/.pcloud-cli.json`.

## Build

To build pcloud-cli binary, clone and run `make pcloud-cli`.

## Installation

Installing pcloud-cli can be done by cloning git repo and running `make install`.

## Usage

Main command help:

```
Usage:
  pCloud-cli [flags]
  pCloud-cli [command]

Available Commands:
  authorize   Authorize with pCloud.
  file        Actions to manage files.
  folder      Actions to manage folders.
  help        Help about any command
  version     Print the version number of pCloud-cli

Flags:
      --config string   config file (default is $HOME/.pcloud-cli.json)
  -h, --help            help for pCloud-cli
      --token string    bearer token to access API, can be used when not using config file
  -v, --verbose         verbose output for debugging

```

File sub-command help:

```
Usage:
  pCloud-cli file [flags]
  pCloud-cli file [command]

Available Commands:
  checksum    Calculate chacksums of file.
  copy        Copy file to another location.
  delete      Delete file.
  get         Get remote file url and download it.
  rename      Rename / Move source file.
  upload      Upload local file to remote folder.

Flags:
  -h, --help   help for file

Global Flags:
      --config string   config file (default is $HOME/.pcloud-cli.json)
      --token string    bearer token to access API, can be used when not using config file
  -v, --verbose         verbose output for debugging

```

Folder sub-command help:

```
Usage:
  pCloud-cli folder [flags]
  pCloud-cli folder [command]

Available Commands:
  create      Create folder.
  delete      Delete folder.
  list        List folders in pCloud directory
  rename      Rename / Move folder.

Flags:
  -h, --help   help for folder

Global Flags:
      --config string   config file (default is $HOME/.pcloud-cli.json)
      --token string    bearer token to access API, can be used when not using config file
  -v, --verbose         verbose output for debugging

```
