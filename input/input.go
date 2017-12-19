package input

import (
	"os"
	"bufio"
	"fmt"
)

func Ask(task string) string {
	reader := bufio.NewReader(os.Stdin)
    fmt.Print(task + "? ")
    text, _ := reader.ReadString('\n')
    return text[0:len(text)-1]
}