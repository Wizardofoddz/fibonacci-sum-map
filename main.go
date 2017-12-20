package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBytes, err := ioutil.ReadAll(bufio.NewReader(os.Stdin))
	if err != nil {
		log.Panicln("Error reading from stdin:", err.Error())
	}

	input := strings.Trim(string(inputBytes), "\n\t ")
	n, err := strconv.ParseUint(string(input), 10, 64)
	if err != nil {
		log.Panicln("Error parsing stdin as uint64:", err.Error())
	}

	fmt.Println(Fibonacci(n))
}

// Fibonacci computes the nth digit of the Fibonacci Sequence
func Fibonacci(n uint64) uint64 {
	sequence := []uint64{0, 1}

	if n < 2 {
		return sequence[n]
	}

	for i := uint64(2); i <= n; i++ {
		lastValue := sequence[len(sequence)-1]
		secondToLastValue := sequence[len(sequence)-2]

		sequence = append(sequence, lastValue+secondToLastValue)
	}

	return sequence[len(sequence)-1]
}
