[![Build Status](https://travis-ci.org/yonush/versions.svg?branch=master)](https://travis-ci.org/yonush/versions)

# Embedding version information within a Go project

- [Go linker documentation](https://pkg.go.dev/cmd/link)
- [Versioning binaries in go](https://dev.to/gcdcoder/versioning-binaries-in-go-52al)
- [Using ldflags to set versions](https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications)
- [Make for windows](https://github.com/maweil/MakeForWindows/releases) Binaries in the make4windows folder

A simple Go command line tool ([go-winres](github.com/tc-hib/go-winres)) for embedding usual resources in Windows executables built with Go:

``` sh
go install github.com/tc-hib/go-winres@latest
```