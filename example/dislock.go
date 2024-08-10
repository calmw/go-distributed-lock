package main

import (
	dislock "github.com/calmw/go-distributed-lock"
	"log"
	"os"
	"time"
)

const (
	DefaultServerAddr = "127.0.0.1:6000"
)

var Lk *dislock.Lock

func main() {
	serverAddr := os.Getenv("SERVER_ADDR")
	if len(serverAddr) == 0 {
		serverAddr = DefaultServerAddr
	}

	// retryMillisecondDur 加锁/释放锁不成功, 距下次重试的时间间隔(单位：毫秒)
	// retryMaxTimes 设置为0时，加锁/释放锁不成功就一直等待; 不为0时，尝试到该次数后强制执行加锁/释放锁
	Lk = dislock.NewLock(serverAddr, 200, 3000)

	lock("order_lock", "clientA")
	log.Println("订单操作")
	time.Sleep(time.Second * 10)
	unlock("order_lock", "clientA")
}

func lock(lockName, clientId string) {
	t := time.Now()
	log.Printf("%s %s, 加锁... \n", lockName, clientId)
	ok, retryTimes, err := Lk.Lock(lockName, clientId)

	log.Printf("%s %s %d (%v %v %v), 执行加锁后的业务... \n", lockName, clientId, time.Since(t), ok, retryTimes, err)
}

func unlock(lockName, clientId string) {
	log.Printf("%s %s, 释放锁... \n", lockName, clientId)
	t := time.Now()
	ok, retryTimes, err := Lk.UnLock(lockName, clientId)

	log.Printf("%s %s %d (%v %v %v),执行释放锁后的业务... \n", lockName, clientId, time.Since(t), ok, retryTimes, err)
}
