package main

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

type Reusable struct {
}

type ObjPool struct {
	//缓冲可重用对象
	bufferChan chan *Reusable
}

func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufferChan = make(chan *Reusable, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objPool.bufferChan <- &Reusable{}
	}
	return &objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*Reusable, error) {
	select {
	case ret := <-p.bufferChan:
		return ret, nil
	//超时控制
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

func (p *ObjPool) ReleaseObj(obj *Reusable) error {
	select {
	case p.bufferChan <- obj:
		return nil
	default:
		return errors.New("overfolw")
	}
}

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}
	fmt.Println("Done...")
}
