/**
* @description :
* @author : Jarick
* @Date : 2022-07-06
* @Url : http://CloudWebOps
 */
package util

import "time"

func TimeFuncDuration(Num int64, TimeType string) (timeDuration time.Duration) {
	switch {
	case TimeType == "Millisecond":
		timeDuration = time.Duration(Num * int64(time.Millisecond))
	case TimeType == "Microsecond":
		timeDuration = time.Duration(Num * int64(time.Microsecond))
	case TimeType == "Nanosecond":
		timeDuration = time.Duration(Num * int64(time.Nanosecond))
	case TimeType == "Second":
		timeDuration = time.Duration(Num * int64(time.Second))
	case TimeType == "Minute":
		timeDuration = time.Duration(Num * int64(time.Minute))
	}
	return
}
