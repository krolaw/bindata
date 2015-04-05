# bindata - Converts binary files to Go const

## Examples
	bindata -o output.go binfile1.jpg binfile2.jpg

Or using go generate:

	//go:generate bindata -o jpegs.go pic1.jpg pic2.jpg pic3.jpg

## Params
	-o Output filename, default is assets.go
	-p Package name, default is name of output filename's parent directory. 
	-w Width of lines, inbetween speechmarks - default 72

## Installation
	go get github.com/krolaw/bindata
	go build -o $GOPATH/bin/bindata $GOPATH/src/github/krolaw/bindata/bindata.go

## Gotchas
Due to [issue #9035](https://github.com/golang/go/issues/9035), you may need to
increase the line width to reduce the number of lines bindata generates.

## Alternatives
* https://github.com/jteeuwen/go-bindata
