package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"main/pkg/cursor"
	. "main/pkg/debug"
)

func main() {
	filename := "./input.txt"

	buf, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	text := string(buf)

	if Debug() {
		LogIt(fmt.Sprintf("input_length=%d", len(text)))
	}

	ok, output := cursor.Scan(text)

	if ok {
		if Debug() {
			LogIt("scan=ok")
			LogIt("\n")
			LogIO.Flush()
		}
		fmt.Printf("\noutput=%d\n", output)
	} else {
		if Debug() {
			LogIt(fmt.Sprintf("length=%d scan=not_ok", len(text)))
			LogIt("\n")
			LogIO.Flush()
		}
		panic("wat")
	}
}
