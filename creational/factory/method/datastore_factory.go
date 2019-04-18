package factory

import "log"

type DataStoreFactory interface {
	Create(conf map[string]string) (DataStore, error)
}

// datastoreFactories is the variable for storing the factory methods
var datastoreFactories = make(map[string]DataStoreFactory)

// Register adds a new DataStore implementation into the datastoreFactories variable
func Register(name string, factory DataStoreFactory) {
	if factory == nil {
		log.Panicf("Datastore factory %s does not exist.", name)
	}
	_, registered := datastoreFactories[name]
	if registered {
		log.Printf("Datastore factory %s already registered. Ignoring.", name)
	}
	datastoreFactories[name] = factory
}

func init() {
	Register("postgres", &PostgreSQLDataStoreFactory{})
	Register("memory", &MemoryDataStoreFactory{})
}
