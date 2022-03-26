package utils

import (
	"fmt"
	"os"
	"time"
)

func CloseServ() {

	os.Exit(2)
}

func MsWith2Decimal(d time.Duration) string {
	return fmt.Sprintf("%d.%d%d ms", d.Milliseconds(), (d.Microseconds()-(d.Milliseconds()*1000))/100, (d.Microseconds()-(d.Milliseconds()*1000)-((d.Microseconds()-(d.Milliseconds()*1000))/100)*100)/10)
}
