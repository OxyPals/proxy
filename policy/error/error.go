package error

import (
	"fmt"
	"os"
)

func HandleAndExit(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}
