# pCloud-cli - pCloud Command Line Interface

Command line interface to pCloud API written in Go.

## Usage

``` shell
pcloud-cli --help
```

Run `pcloud-cli authorize` to authorize pCloud-cli and save settings to `~/.pcloud-cli.json`.

## Build

To build pcloud-cli, download it and run `make build`.

## Functions implemented (so far)
These functions are implemented in the CLI.
For suggestions and/or missing functions, please submit an issue.

* authorize
* folder
  * create
  * delete
  * list
  * rename
* file
  * checksum
  * copy
  * delete
  * get
  * rename
  * upload
* version