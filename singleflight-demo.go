package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"time"
)

func main() {
	fn := func() (interface{}, error) {
		time.Sleep(time.Second)
		return time.Now(), nil
	}
	sg := new(singleflight.Group)
	for i := 0; i < 200; i++ {
		go func() {
			// 这里 使用 sg.Do 可以让多个goroutine只有一个被使用，并且其它的获取相同的数据
			// 这可以解决 缓存击穿的问题：一个缓存实效，瞬间多个请求把 数据库搞死。
			v, err, shared := sg.Do("jerry", fn)
			fmt.Println(v,err,shared)
		}()
	}
	time.Sleep(time.Second * 5)
}
