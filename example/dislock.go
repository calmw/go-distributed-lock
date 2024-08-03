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
	Lk = dislock.NewLock(serverAddr, 200, 3000)

	TaskLock("order_lock", "userA")
	log.Println("订单操作")
	time.Sleep(time.Second * 10)
	TaskUnLock("order_lock", "userA")
}

func TaskLock(lockName, clientId string) {
	t := time.Now()
	log.Printf("%s %s, 加锁...", lockName, clientId)
	ok, retryTimes, err := Lk.Lock(lockName, clientId)

	log.Printf("%s %s %d (%v %v %v), 执行加锁后的业务...", lockName, clientId, time.Since(t), ok, retryTimes, err)
}

func TaskUnLock(lockName, clientId string) {
	log.Printf("%s %s, 释放锁...", lockName, clientId)
	t := time.Now()
	ok, retryTimes, err := Lk.UnLock(lockName, clientId)

	log.Printf("%s %s %d (%v %v %v),执行释放锁后的业务...", lockName, clientId, time.Since(t), ok, retryTimes, err)
}
