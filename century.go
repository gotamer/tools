// Century Time
// Is the time duration in float64 since 2000-01-01 UTC  
package century

import (
	"fmt"
	"time"
)

var Date2000 = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)

func Now() float64 {
	return time.Since(Date2000).Seconds()
}

func At(t time.Time) float64 {
	return t.Sub(Date2000).Seconds()
}