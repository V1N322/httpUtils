package util

import (
	"net/http"
	"sync"
)

var (
	clientInstance *http.Client
	once           sync.Once
)

func GetHTTPClient() *http.Client {
	once.Do(func() {
		clientInstance = &http.Client{}
	})
	return clientInstance
}
