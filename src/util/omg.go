package util

import (
	"fmt"
)

func OMG(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	//os.Exit(1)
	panic(err)
}
