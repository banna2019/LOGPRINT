// 最好是小写字母,并且最好不要带下划线,最好见名知义
package logprint

import (
	"fmt"
	"time"
)

func Debug(msg interface{}) {
	t := time.Now()
	fmt.Printf("[debug] %s: %s \n", t.Format("2006-01-02 15:04:05.000"), msg)
}
