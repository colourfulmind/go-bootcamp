package calculations

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	Min = -100000
	Max = 100000
)

// Flags is a struct to mark values needed to be calculated.
type Flags struct {
	Mean, Median, Mode, Sd bool
}

// Input outputs the usage message.
func (f Flags) Input() {
	fmt.Println("Enter the sequence of numbers. Separate numbers by newline. Press 'q' + 'Enter' to calculate:")
	f.InitNumbers()
}

// InitNumbers initiates user input.
func (f Flags) InitNumbers() {
	sc := bufio.NewScanner(os.Stdin)
	var numbers []int
	for sc.Scan() {
		text := sc.Text()
		if text == "q" {
			break
		}
		i, err := strconv.Atoi(text)
		if err != nil || i < Min || i > Max {
			err := errors.New("input must be an integer strictly between -100000 and 100000, try again")
			fmt.Println(err)
			continue
		}
		numbers = append(numbers, i)
	}
	s := Statistics{}
	s.Output(f, numbers)
}
