package main

import (
	"bufio"
	"fmt"
	"github.com/mskcc/ddp-spec-date-collector/ddp"
	"os"
)

func main() {
	var dmpIds []string
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		dmpIds = append(dmpIds, input.Text())
	}
	fmt.Println("token: %s\n", token)
}
