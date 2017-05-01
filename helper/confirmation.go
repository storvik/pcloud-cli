package helper

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// AskConfirmation prints the string str and prompts the user for
// yes/no. Accept all yes, YES, y, Y, no, NO, n, N.
func AskConfirmation(str string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", str)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}
