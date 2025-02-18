package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("secrets.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		err = errors.Join(err, f.Close())
	}()
}
