package common

import (
	"fmt"
)

func Log(content string, err error) {
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(content)
}
