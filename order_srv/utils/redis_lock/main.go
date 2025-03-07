package main

import (
	"fmt"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"sync"
	"time"
)

func main() {
	// Create a pool with go-redis (or redigo) which is the pool redisync will
	// use while communicating with Redis. This can also be any pool that
	// implements the `redis.Pool` interface.
	client := goredislib.NewClient(&goredislib.Options{
		Addr: "192.168.15.21:6379",
	})
	pool := goredis.NewPool(client) // or, pool := redigo.NewPool(...)

	// Create an instance of redisync to be used to obtain a mutual exclusion
	// lock.
	rs := redsync.New(pool)

	// Obtain a new mutex by using the same name for all instances wanting the
	// same lock.
	gNum := 2
	mutexname := "421-goods-stock-lock"

	var wg sync.WaitGroup
	wg.Add(gNum)
	for i := 0; i < gNum; i++ {
		go func(i int) {
			defer wg.Done() // wg.Add和wg.Done需要成对出现，否则报错

			mutex := rs.NewMutex(mutexname)

			fmt.Println("开始获取锁", i)

			if err := mutex.Lock(); err != nil {
				panic(err)
			}

			fmt.Println("获取锁成功", i)

			time.Sleep(time.Second * 2)

			fmt.Println("开始释放锁", i)
			if ok, err := mutex.Unlock(); !ok || err != nil {
				panic("unlock failed")
			}

			fmt.Println("释放锁成功", i)
		}(i)
	}
	wg.Wait()
}
