package cli

import (
	"bufio"
	"os"
)

// AwaitCLIcommand is an endless loop that breaks only after user types anything into console or
// after Enter is pressed
func AwaitCLIcommand() string {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		return scanner.Text()
	}
}
