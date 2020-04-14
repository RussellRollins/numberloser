package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if err := inner(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func inner() error {
	var df = flag.Int("distance", 0, "how far to rotate the number")
	var nf = flag.String("number", "", "the phone number to rotate")
	flag.Parse()

	distance := *df
	number := *nf

	// validate distance
	if distance == 0 || distance > 9 || distance < -9 {
		return fmt.Errorf("must rotate the digit between -9 and 9 places, excluding 0")
	}

	// "correct" negative rotations by making them go forward
	if distance < 0 {
		distance = 10 + distance
	}

	if number == "" {
		return fmt.Errorf("phone number to rotate must not be empty")
	}

	var newNumber []rune
	for _, char := range number {
		if digit, err := strconv.Atoi(string(char)); err == nil {
			digit = (digit + distance) % 10

			// TODO: There's no way this is the best solution.
			strDigit := fmt.Sprintf("%d", digit)
			runeDigits := []rune(strDigit)
			newRune := runeDigits[0]

			newNumber = append(newNumber, newRune)
		} else {
			newNumber = append(newNumber, char)
		}
	}

	fmt.Printf("input : %+q\n", number)
	fmt.Printf("result: %+q\n", string(newNumber))

	return nil
}
