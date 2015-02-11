package main

import (
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	outputFilename := flag.String("o", "assets.go", "Name of output go file.")
	packageName := flag.String("p", "", "Package Name - defaults to dir of output go file.")
	width := flag.Int("w", 72, "Width between speechmarks per line.")

	flag.Parse()

	fo, err := os.OpenFile(*outputFilename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0775)
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()

	if *packageName == "" {
		*packageName, err = filepath.Abs(filepath.Dir(*outputFilename))
		if err != nil {
			log.Fatal(err)
		}
		*packageName = filepath.Base(*packageName)
	}

	fo.WriteString("package " + *packageName + "\n\nconst (")

	for _, fileName := range flag.Args() {
		if err := writeFile(fo, fileName, *width); err != nil {
			log.Fatal(err)
		}
	}

	fo.WriteString(")\n")
}

func writeFile(fo io.Writer, fileName string, width int) error {
	fi, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer fi.Close()
	if _, err := fo.Write([]byte("\n\t" + strings.Replace(fileName, ".", "_", -1) + " = \"")); err != nil {
		return err
	}
	if _, err := io.Copy(&HexWriter{&DivideWriter{fo, width - width%4, []byte("\" +\n\t\t\""), 0}}, fi); err != nil {
		return err
	}
	if _, err := fo.Write([]byte("\"\n")); err != nil {
		return err
	}
	return nil
}

type DivideWriter struct {
	io.Writer
	Length   int
	Divider  []byte
	Position int
}

func (l *DivideWriter) Write(data []byte) (n int, err error) {
	var nn int
	for len(data)+l.Position > l.Length {
		if nn, err = l.Writer.Write(data[:l.Length-l.Position]); err != nil {
			return n + nn, err
		}
		n += nn
		if nn, err = l.Writer.Write(l.Divider); err != nil {
			return n, err
		}
		data = data[l.Length-l.Position:]
		l.Position = 0
	}
	l.Writer.Write(data)
	l.Position += len(data)
	n += len(data)
	return
}

const hexChars = "0123456789abcdef"

type HexWriter struct {
	io.Writer
}

func (h *HexWriter) Write(data []byte) (n int, err error) {
	// TODO think about whether garbage generated is a problem
	for i, v := range data {
		if _, err := h.Writer.Write([]byte{'\\', 'x', hexChars[v/16], hexChars[v%16]}); err != nil {
			return i, err
		}
	}
	return len(data), nil
}
