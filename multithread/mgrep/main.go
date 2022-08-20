package main

import (
	"fmt"
	"golangchan/multithread/worker"
	"golangchan/multithread/worklist"
	"os"
	"path/filepath"
	"sync"
)

func discoverDirs(wl *worklist.Worklist, path string) {
	dir, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("error al leer el path: ", err)
		return
	}
	for _, entry := range dir {
		if entry.IsDir() {
			nextPath := filepath.Join(path, entry.Name())
			discoverDirs(wl, nextPath)
		} else {
			wl.Add(worklist.NewJob(filepath.Join(path, entry.Name())))
		}
	}
}

func main() {
	var workersWg sync.WaitGroup
	wl := worklist.New(100)

	resultados := make(chan worker.Resultado, 100)
	numeroWorkers := 10
	workersWg.Add(1)

	go func() {
		defer workersWg.Done()
		discoverDirs(&wl, "")
		wl.Finalize(numeroWorkers)
	}()

	for i := 0; i < numeroWorkers; i++ {
		workersWg.Add(1)
		go func() {
			defer workersWg.Done()
			for {
				workEntry := wl.Next()
				if workEntry.Path != "" {
					workResult := worker.FindInFile(workEntry.Path, "")
					if workResult != nil {
						for _, r := range workResult.Inner {
							resultados <- r
						}
					}
				} else {
					return
				}
			}

		}()
	}

	//utilizo una goroutine para bloquear mientras espero a que se terminen los otros trabajos
	blockWorkersWg := make(chan struct{})
	go func() {
		workersWg.Wait()
		close(blockWorkersWg)
	}()

	var displayWg sync.WaitGroup
	displayWg.Add(1)

	go func() {
		for {
			select {
			case r := <-resultados:
				fmt.Printf("%v[%v]:%v\n", r.Path, r.NumeroLinea, r.Linea)
			case <-blockWorkersWg:
				if len(resultados) == 0 {
					displayWg.Done()
					return
				}
			}
		}
	}()

	displayWg.Wait()

}
