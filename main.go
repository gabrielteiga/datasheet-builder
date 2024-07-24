package main

import (
	"os"
	"sync"

	"github.com/gabrielteiga/datasheet-builder/src/main/domain/services"
)

var fileService services.FileService

func main() {
	ROOT, _ := os.Getwd()
	PATH := ROOT + "/src/resources/"

	FILES := fileService.GetFiles(PATH)

	var wg sync.WaitGroup

	for _, file := range FILES {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fileService.FormatFiles(PATH + file)
		}()
	}

	wg.Wait()
}
