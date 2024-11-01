package iredis

import (
	"context"
	"fmt"
	"github.com/bsm/redislock"
	"strconv"
	"testing"
	"time"
)

func init() {
	config := Config{
		Mode:    singleMode,
		Address: "127.0.0.1:6379",
		DB:      0,
	}

	Init(&config)
}

func TestSetAndGet(t *testing.T) {
	ctx := context.Background()

	_client := Client()
	fmt.Println(_client.Set(ctx, "key", "value", 0).Result())
	fmt.Println(_client.Get(ctx, "key").Result())

	fmt.Println(_client.SetEx(ctx, "ttlkey", "ttl value", 10*time.Second).Result())
	fmt.Println(_client.Get(ctx, "ttlkey").Result())
}

func getLock() (*redislock.Lock, error) {
	myLock, err := GetLock("lock", 10*time.Second)
	if err != nil {
		return nil, err
	}
	return myLock, nil
}

// 获取锁，锁自动失效
func TestLock(t *testing.T) {
	ctx := context.Background()

	var theLock *redislock.Lock

	fmt.Println("================== 一轮抢锁 =======================")
	// 一轮抢锁
	for i := 0; i < 10; i++ {
		go func() {
			_lock, err := getLock()
			if err != nil {
				fmt.Println("#" + strconv.Itoa(i) + " Get Lock Failed")
			} else {
				fmt.Println("#" + strconv.Itoa(i) + " Get The Lock")
				theLock = _lock
			}
		}()
	}
	time.Sleep(2 * time.Second)
	// 释放锁
	theLock.Release(ctx)
	fmt.Println("Lock Released")

	fmt.Println("================== 二轮抢锁 =======================")
	// 二轮抢锁
	for i := 0; i < 10; i++ {
		go func() {
			_lock, err := getLock()
			if err != nil {
				fmt.Println("#" + strconv.Itoa(i) + " Get Lock Failed")
			} else {
				fmt.Println("#" + strconv.Itoa(i) + " Get The Lock")
				theLock = _lock
			}
		}()
	}
	time.Sleep(2 * time.Second)
	// 释放锁
	theLock.Release(ctx)
	fmt.Println("Lock Released")

	fmt.Println("================== 三轮抢锁 =======================")
	// 三轮抢锁
	for i := 0; i < 10; i++ {
		go func() {
			_lock, err := getLock()
			if err != nil {
				fmt.Println("#" + strconv.Itoa(i) + " Get Lock Failed")
			} else {
				fmt.Println("#" + strconv.Itoa(i) + " Get The Lock")
				theLock = _lock
			}
		}()
	}
	// 释放锁
	time.Sleep(2 * time.Second)
	// 不释放锁
	//theLock.Release(ctx)

	fmt.Println("================== 四轮抢锁 =======================")
	// 四轮抢锁
	for i := 0; i < 10; i++ {
		go func() {
			_lock, err := getLock()
			if err != nil {
				fmt.Println("#" + strconv.Itoa(i) + " Get Lock Failed")
			} else {
				fmt.Println("#" + strconv.Itoa(i) + " Get The Lock")
				theLock = _lock
			}
		}()
	}
	// 谁都抢不到锁
	time.Sleep(11 * time.Second)

	// 锁失效

	fmt.Println("================== 五轮抢锁 =======================")
	// 锁超时，五轮抢锁
	for i := 0; i < 10; i++ {
		go func() {
			_lock, err := getLock()
			if err != nil {
				fmt.Println("#" + strconv.Itoa(i) + " Get Lock Failed")
			} else {
				fmt.Println("#" + strconv.Itoa(i) + " Get The Lock")
				theLock = _lock
			}
		}()
	}
	time.Sleep(2 * time.Second)
}
