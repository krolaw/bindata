# bindata - Converts binary files to Go const

## UNTESTED
No really, I haven't checked the output matches the binaries.  Tomorrow...

## Examples
	bindata -o output.go binfile1.jpg binfile2.jpg

Or using go generate:
	//go:generate bindata -o jpegs.go pic1.jpg pic2.jpg pic3.jpg

## Params
* -o Output filename, default is assets.go
* -p Package name, default is name of output filename's parent directory 

## Installation
	go get github.com/krolaw/bindata
	go build -o $GOPATH/bin/bindata $GOPATH/src/github/krolaw/bindata/bindata.go

## Alternatives
* https://github.com/jteeuwen/go-bindata