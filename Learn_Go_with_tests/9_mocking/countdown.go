package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countdownStart = 3
const finalWord = "Go!"

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprint(out, i, "\n")
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
