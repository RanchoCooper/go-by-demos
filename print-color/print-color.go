package print_color

import (
	"fmt"
	"time"
)

func Main() {
	var (
		green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
		magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
		reset   = string([]byte{27, 91, 48, 109})
	)

	fmt.Printf("%s[message] %v%s \n",
		magenta, time.Now().Format("2006/01/02 - 15:04:05"), reset)
	fmt.Printf("%s[message] %v%s \n",
		green, time.Now().Format("2006/01/02 - 15:04:05"), reset)
}
