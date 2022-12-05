package debug

import (
	"bufio"
	"fmt"
	"os"
)

func Debug() bool {
	return os.Getenv("DEBUG") == "true"
}

var LogIO = bufio.NewWriter(os.Stdout)

func LogIt(s string) {
	fmt.Fprint(LogIO, fmt.Sprintf("%s ", s))
}
