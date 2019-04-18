package factory

import "sync"

type MemoryDataStoreFactory struct{}

// Create is a factory method that return the common interface DataStore
func (mdsf *MemoryDataStoreFactory) Create(conf map[string]string) (DataStore, error) {
	return &MemoryDataStore{
		Users: map[int64]string{
			1: "mnbbrown",
			0: "root",
		},
		RWMutex: sync.RWMutex{},
	}, nil
}
