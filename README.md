tldr
====

A [tldr](https://github.com/tldr-pages/tldr) client in Go, focusing on speed by storing pages in a database.
The choice of Go also allows us to support all platforms which Go compiles to.

## Installation
You can use [Go](https://golang.org/)'s tooling

```
go get github.com/elecprog/tldr
go install github.com/elecprog/tldr
```

or download a binary for Linux or Windows from the [release page](https://github.com/elecprog/tldr/releases/latest/).

## Usage
You can print information for one or more commands by using:

```
tldr command1 [command2 ...]
```

This client downloads all tldr pages on the first run (resulting in a database of about 2&nbsp;MB) which should only take a couple of seconds. To redownload the pages and rebuild the database you can use:

```
tldr -u
```

The database is then stored in the cache directory of your platform.