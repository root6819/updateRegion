package main

import (
	"sync"
	"updateRegion/utils"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		utils.DoGetInfo()
	}()
	wg.Wait()
}
