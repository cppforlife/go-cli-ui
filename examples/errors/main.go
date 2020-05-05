package main

import (
	"fmt"

	"github.com/cppforlife/go-cli-ui/errors"
)

func main() {
	fmt.Printf("%s", errors.NewMultiLineError(fmt.Errorf("Some error: Another error: Something else")))
}
