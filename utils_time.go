package glog

import "time"

func RandomSleep(min, max int, duration time.Duration) {
	time.Sleep(time.Duration(GetRandomInt(min, max)) * duration)
}
