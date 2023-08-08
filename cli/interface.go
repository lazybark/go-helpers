package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// AwaitStringInput awaits for any text from user in os.Stdin
func AwaitStringInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}

// AwaitStringInputNotEmpty awaits for text with len>0 from user in os.Stdin
func AwaitStringInputNotEmpty() string {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		if scanner.Text() != "" {
			return scanner.Text()
		}
	}
}

// AwaitNumberInput awaits for user input and then tries to parse into number
func AwaitNumberInput() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	d, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, fmt.Errorf("[AwaitNumberInput] %w", err)
	}

	return d, nil
}
