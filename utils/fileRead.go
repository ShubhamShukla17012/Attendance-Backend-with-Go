package utils

import (
	"fmt"
	"os"
	"sync"
)

const (
	EmployeeFile string = "employees.json"
)

var (
	fileLock sync.RWMutex = sync.RWMutex{}
)

func FileReadOpration() ([]byte, error) {
	fileLock.Lock()
	defer fileLock.Unlock()
	rawContents, err := os.ReadFile(EmployeeFile)
	if err != nil {
		return []byte{}, fmt.Errorf("error while reading employee file, err : %s", err.Error())
	}
	return rawContents, nil
}
