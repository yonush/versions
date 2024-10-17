# This how we want to name the binary output
BINARY=goversions.exe

# These are the values we want to pass for VERSION and BUILD
# git tag 1.0.1
# git commit -am "One more change after the tags"
VERSION=`git describe --tags`
BUILD=`date +%FT%T%z`

# Setup the -ldflags option for go build here, interpolate the variable values
# Go linker documentation https://pkg.go.dev/cmd/link
# -w -s Omit symbol table and debug information
# -X Set the value of the string variable in importpath named name to value.
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD}"

# Builds the project
build:
	go build ${LDFLAGS} -o ${BINARY} -v	

# Installs our project: copies binaries
install:
	go install ${LDFLAGS}

# Cleans our project: deletes binaries
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

# Install and run the module that prepares the Windows resource files for version embedding
init:
	go install github.com/tc-hib/go-winres@latest
	go-winres init

.PHONY: clean install init