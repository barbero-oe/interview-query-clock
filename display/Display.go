package display

import (
	"fmt"
	"time"
)

func SimpleConsolePrinter(t time.Time) {
	fmt.Printf("%02d:%02d:%02d\n", t.Hour(), t.Minute(), t.Second())
}
