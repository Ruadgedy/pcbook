package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ruadgedy/pcbook/pb"
	"github.com/jinzhu/copier"
	"log"
	"sync"
)

var ErrAlreadyExists = fmt.Errorf("the laptop ID is already exists")

type LaptopStore interface {
	Save(laptop *pb.Laptop) error
	Find(id string) (*pb.Laptop, error)
	Search(ctx context.Context,filter *pb.Filter,found func(laptop *pb.Laptop)error) error
}

type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	// map的key是电脑id，value是电脑对象
	data map[string]*pb.Laptop
}

func (store *InMemoryLaptopStore) Search(
	ctx context.Context,
	filter *pb.Filter,
	found func(laptop *pb.Laptop) error,
) error {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	for _, laptop := range store.data {
		// heavy processing
		//time.Sleep(time.Second)

		// 判断客户端会话是否取消
		if ctx.Err() == context.Canceled || ctx.Err() == context.DeadlineExceeded {
			log.Print("context is cancelled")
			return errors.New("context is canceled")
		}

		if isQualified(filter, laptop) {
			// deep copy
			other,err := deepCopy(laptop)
			if err != nil {
				return err
			}

			err = found(other)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func isQualified(filter *pb.Filter, laptop *pb.Laptop) bool {
	if laptop.PriceUsd > filter.MaxPriceUsd {
		return false
	}
	if laptop.Cpu.NumberCores < filter.MinCpuCores {
		return false
	}
	if laptop.Cpu.MinGhz < filter.MinCpuGhz {
		return false
	}
	if toBit(laptop.Ram) < toBit(filter.MinRam) {
		return false
	}
	return true
}

func toBit(memory *pb.Memory) uint64 {
	value := memory.Value
	switch memory.Unit {
	case pb.Memory_BIT:
		return value
	case pb.Memory_BYTE:
		return value<<3
	case pb.Memory_KILOBYTE:
		return value<<13
	case pb.Memory_MEGABYTE:
		return value<<23
	case pb.Memory_GIGABYTE:
		return value<<33
	case pb.Memory_TERABYTE:
		return value<<43
	default:
		return 0
	}
}

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

func deepCopy(laptop *pb.Laptop) (*pb.Laptop, error) {
	// deep copy
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil,fmt.Errorf("cannot copy laptop data: %v", err)
	}
	return other,nil
}

func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	// deep copy
	other,err := deepCopy(laptop)
	if err != nil {
		return err
	}

	store.data[other.Id] = other
	return nil
}

func (store *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}

	return deepCopy(laptop)
}

type DBLaptopStore struct {
}

func (store *DBLaptopStore) Save(laptop *pb.Laptop) error {
	panic("implement me")
}
