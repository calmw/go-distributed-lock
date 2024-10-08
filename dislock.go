package dislock

import (
	"context"
	"errors"
	"github.com/calmw/go-distributed-lock/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"time"
)

type Lock struct {
	RetryMillisecondDur int // 尝试加锁的时间间隔 单位毫秒
	RetryMaxTimes       int // 最大加锁失败次数,超过后将强制释放锁，然后再枷锁。因为某个节点可能在加锁未释放的情况下异常退出. 如果设置为0，加不到锁就一直阻塞
	Lk                  service.LockServiceClient
}

func NewLock(ServerAddr string, retryMillisecondDur, retryMaxTimes int) *Lock {
	return &Lock{
		RetryMillisecondDur: retryMillisecondDur,
		RetryMaxTimes:       retryMaxTimes,
		Lk:                  NewLockServiceClient(ServerAddr),
	}
}

// Lock
/*
lockName 必选参数，锁的名称，可根据业务使用多种锁
clientId 必选参数，加锁的节点ID，分布式的每个节点ID必须唯一
*/
func (l *Lock) Lock(lockName, clientId string) (bool, int, error) {
	var result bool
	var err error
	var retryTimes int
	if l.RetryMaxTimes == 0 {
		for {
			retryTimes++
			ok, e := lock(l.Lk, lockName, clientId)
			if ok {
				result = true // 加锁成功
				err = nil
				break
			}
			err = e
			time.Sleep(time.Millisecond * time.Duration(l.RetryMillisecondDur))
		}
	} else {
		for i := 1; i <= l.RetryMaxTimes; i++ {
			retryTimes = i
			if i == l.RetryMaxTimes {
				r, er := forceLock(l.Lk, lockName, clientId)
				return r, retryTimes, er
			}
			ok, e := lock(l.Lk, lockName, clientId)
			if ok {
				result = true // 加锁成功
				err = nil
				break
			}
			err = e
			time.Sleep(time.Millisecond * time.Duration(l.RetryMillisecondDur))
		}
	}

	return result, retryTimes, err
}

// UnLock
/*
lockName 必选参数，锁的名称，可根据业务使用多种锁
clientId 必选参数，加锁的节点ID，分布式的每个节点ID必须唯一
*/
func (l *Lock) UnLock(lockName, clientId string) (bool, int, error) {
	var result bool
	var err error
	var retryTimes int
	if l.RetryMaxTimes == 0 {
		for {
			retryTimes++
			ok, e := unLock(l.Lk, lockName, clientId)
			if ok {
				err = nil
				result = true // 释放锁成功
				break
			}
			err = e
			if e != nil && strings.Contains(e.Error(), "does not exist") {
				return result, retryTimes, err
			}
			time.Sleep(time.Millisecond * 200)
		}
	} else {
		for i := 1; i <= l.RetryMaxTimes; i++ {
			retryTimes = i
			if i == l.RetryMaxTimes {
				r, er := forceUnLock(l.Lk, lockName)
				return r, retryTimes, er
			}
			ok, e := unLock(l.Lk, lockName, clientId)
			if ok {
				err = nil
				result = true // 释放锁成功
				break
			}
			err = e
			if e != nil && strings.Contains(e.Error(), "does not exist") {
				return result, retryTimes, err
			}
			time.Sleep(time.Millisecond * 200)
		}
	}

	return result, retryTimes, err
}

// ForceUnLock
/*
lockName 必选参数，锁的名称，可根据业务使用多种锁
*/
func (l *Lock) ForceUnLock(lockName string) (bool, error) {
	res, err := forceUnLock(l.Lk, lockName)
	return res, err
}

// ForceLock
/*
lockName 必选参数，锁的名称，可根据业务使用多种锁
clientId 必选参数，加锁的节点ID，分布式的每个节点ID必须唯一
*/
func (l *Lock) ForceLock(lockName, clientId string) (bool, error) {
	res, err := forceLock(l.Lk, lockName, clientId)
	return res, err
}

func lock(l service.LockServiceClient, lockName, clientId string) (bool, error) {
	result, err := l.Lock(context.Background(), &service.LockRequest{
		ClientId: clientId,
		LockName: lockName,
	})

	s := status.Convert(err) // status.Convert函数分别访问错误代码和错误消息
	if s.Code() != codes.OK || err != nil {
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
		return false, err
	}
	return result.Result, errors.New(result.Msg)
}

func forceLock(l service.LockServiceClient, lockName, clientId string) (bool, error) {
	result, err := l.ForceLock(context.Background(), &service.ForceLockRequest{
		LockName: lockName,
		ClientId: clientId,
	})

	s := status.Convert(err) // status.Convert函数分别访问错误代码和错误消息
	if s.Code() != codes.OK || err != nil {
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
		return false, err
	}
	return result.Result, errors.New(result.Msg)
}
