package time

import (
	"fmt"
	"time"
)

// CurrentTime just return the current time nothing fancy
func CurrentTime() string {
	var x = time.Now().Format(time.UnixDate)
	return x
}

func main() {
	fmt.Println(CurrentTime())
}
