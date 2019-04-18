package factory

import (
	"fmt"
	"strings"
)

type DataStore interface {
	Name() string
	FindUserNameById(id int64) (string, error)
}

var errUserNotFound = fmt.Errorf("User not found")

func CreateDatastore(conf map[string]string) (DataStore, error) {
	var engineName string
	var ok bool

	// Query configuration for datastore, by default is "memory"
	if engineName, ok = conf["DATASTORE"]; !ok {
		engineName = "memory"
	}

	engineFactory, ok := datastoreFactories[engineName]
	if !ok {
		// Factory has not been registered
		return nil, fmt.Errorf("Invalid Datastore name. Must be one of: %s", strings.Join(listAllDataStoreFactories(), ", "))
	}

	// Run the factory with the configuration
	return engineFactory.Create(conf)
}

func listAllDataStoreFactories() []string {
	// Make a list of all available datastore factories for logging
	availableDatastores := make([]string, len(datastoreFactories))
	for k := range datastoreFactories {
		availableDatastores = append(availableDatastores, k)
	}
	return availableDatastores
}
