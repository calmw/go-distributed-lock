package dislock

import (
	"context"
	"errors"
	"github.com/calmw/go-distributed-lock/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

type Lock struct {
	MillisecondDur int // 尝试加锁的时间间隔 单位毫秒
	MaxRetryTimes  int // 最大加锁失败次数,超过后将强制释放锁，然后再枷锁。因为某个节点可能在加锁未释放的情况下异常退出
	Lk             service.LockServiceClient
}

func NewLock(ServerAddr string, millisecondDur, maxRetryTimes int) *Lock {
	return &Lock{
		MillisecondDur: millisecondDur,
		MaxRetryTimes:  maxRetryTimes,
		Lk:             NewLockServiceClient(ServerAddr),
	}
}

// Lock
/*
lockName 必选参数，锁的名称，可根据业务使用多种锁
clientId 必选参数，加锁的节点ID，分布式的每个节点ID必须唯一
*/
func (l *Lock) Lock(lockName, clientId string) (bool, error) {
	var result bool
	var err error
	for i := 1; i <= l.MaxRetryTimes; i++ {
		if i == l.MaxRetryTimes-1 {
			r, er := forceUnLock(l.Lk, lockName)
			log.Printf("%s-%s, 加锁%d次失败, 强制释放锁 %v %v", lockName, clientId, i, r, er)
		}
		ok, e := lock(l.Lk, lockName, clientId)
		if ok {
			result = true // 加锁成功
			err = nil
			break
		}
		err = e
		time.Sleep(time.Millisecond * time.Duration(l.MillisecondDur))
	}
	log.Printf("%s-%s 加锁 result %d %v", lockName, clientId, result, err)

	return result, err
}

// UnLock
/*
lockName 必选参数，锁的名称，可根据业务使用多种锁
clientId 必选参数，加锁的节点ID，分布式的每个节点ID必须唯一
*/
func (l *Lock) UnLock(lockName, clientId string) (bool, error) {
	var result bool
	var err error
	for i := 1; i <= l.MaxRetryTimes; i++ {
		if i == l.MaxRetryTimes {
			r, er := forceUnLock(l.Lk, lockName)
			log.Printf("%s-%s, 释放锁%d次失败, 强制释放锁 %v %v", lockName, clientId, i, r, er)
			return r, er
		}
		ok, e := unLock(l.Lk, lockName, clientId)
		if ok {
			err = nil
			result = true // 释放锁成功
			break
		}
		err = e
		time.Sleep(time.Millisecond * 200)
	}
	log.Printf("%s-%s 释放锁 result %d %v", lockName, clientId, result, err)

	return result, err
}

// ForceUnLock
/*
lockName 必选参数，锁的名称，可根据业务使用多种锁
*/
func (l *Lock) ForceUnLock(lockName string) (bool, error) {
	res, err := forceUnLock(l.Lk, lockName)
	log.Printf("%s, 强制释放锁 %v %v", lockName, res, err)
	return res, err
}

func lock(l service.LockServiceClient, lockName, clientId string) (bool, error) {
	result, err := l.Lock(context.Background(), &service.LockRequest{
		ClientId: clientId,
		LockName: lockName,
	})

	s := status.Convert(err) // status.Convert函数分别访问错误代码和错误消息
	if s.Code() != codes.OK || err != nil {
		//log.Fatalf("Request failed: %v-%v\n", s.Code(), s.Message())
		return false, err
	}
	return result.Result, errors.New(result.Msg)
}

func unLock(l service.LockServiceClient, lockName, clientId string) (bool, error) {
	result, err := l.UnLock(context.Background(), &service.UnLockRequest{
		ClientId: clientId,
		LockName: lockName,
	})

	s := status.Convert(err) // status.Convert函数分别访问错误代码和错误消息
	if s.Code() != codes.OK || err != nil {
		log.Fatalf("Request failed: %v-%v\n", s.Code(), s.Message())
		return false, err
	}
	return result.Result, errors.New(result.Msg)
}

func forceUnLock(l service.LockServiceClient, lockName string) (bool, error) {
	result, err := l.ForceUnLock(context.Background(), &service.ForceUnLockRequest{
		LockName: lockName,
	})

	s := status.Convert(err) // status.Convert函数分别访问错误代码和错误消息
	if s.Code() != codes.OK || err != nil {
		//log.Fatalf("Request failed: %v-%v\n", s.Code(), s.Message())
		return false, err
	}
	return result.Result, errors.New(result.Msg)
}
