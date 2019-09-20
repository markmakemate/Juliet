package common

import "sync"

func init()  {

	DomainEntity = &Domain{
		Mutex: sync.Mutex{},
		Id:    0,
	}
}

