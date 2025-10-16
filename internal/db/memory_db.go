package db

import (
	"sync"
)

type MemoryDb struct {
	Movies sync.Map
	Characters sync.Map
}

func New() *MemoryDb {
	return &MemoryDb{ }
}