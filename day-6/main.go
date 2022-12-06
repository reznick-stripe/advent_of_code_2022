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
	packetLength := 4
	messageLength := 14

	buf, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	text := string(buf)

	if Debug() {
		LogIt(fmt.Sprintf("input_length=%d", len(text)))
	}

	output, err := cursor.Scan(text, packetLength)

	if err != nil {
		if Debug() {
			LogIt(fmt.Sprintf("length=%d scan=not_ok err=%s", len(text), err))
			LogIt("\n")
			LogIO.Flush()
		}
		log.Fatal(err)
	} else {
		fmt.Printf("\npacket_output=%d\n", output)
		if Debug() {
			LogIt("scan=ok")
			LogIt("\n")
		}
	}

	output, err = cursor.Scan(text, messageLength)

	if err != nil {
		if Debug() {
			LogIt(fmt.Sprintf("length=%d scan=not_ok err=%v", len(text), err))
			LogIt("\n")
			LogIO.Flush()
		}
		log.Fatal(err)
	} else {
		if Debug() {
			LogIt("scan=ok")
			LogIt("\n")
			LogIO.Flush()
		}
		fmt.Printf("\nmessage_output=%d\n", output)
	}
}
