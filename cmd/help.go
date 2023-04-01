package cmd

import (
	"fmt"
	"os"
)

func Help() {
	fmt.Println("bbg-cli: BBG implementation on the command line.")
	os.Exit(1)
}
