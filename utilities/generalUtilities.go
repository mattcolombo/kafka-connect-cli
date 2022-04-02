package utilities

import (
	"fmt"
	"os"
)

func CheckMutuallyExclusive(flag1, flag2 bool, message string) {
	if flag1 && flag2 {
		fmt.Println(message)
		os.Exit(1)
	}
}
