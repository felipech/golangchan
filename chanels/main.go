package main

import (
	"fmt"
	"golang.org/x/exp/rand"
	"time"
)

type Job int

func calculoLargo(i Job) int {
	duracion := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(duracion)
	fmt.Printf("Job %d completado en %v\n", i, duracion)
	return int(i) * 30
}

func generarTrabajos() []Job {
	//make el tipo, largo, capacidad
	jobs := make([]Job, 0, 100)
	for i := 0; i < 100; i++ {
		jobs = append(jobs, Job(rand.Intn(10000)))
	}
	return jobs
}

// hacemos el canal para procesar los trabajos
func runJob(resultChan chan int, i Job) {
	resultChan <- calculoLargo(i)

}

func main() {
	rand.Seed(uint64(time.Now().UnixNano()))
	jobs := generarTrabajos()
	resultChan := make(chan int, 10)
	for i := 0; i < len(jobs); i++ {
		go runJob(resultChan, jobs[i])
	}

	resultCount := 0
	sum := 0
	for {
		result := <-resultChan
		sum += result
		resultCount += 1
		if resultCount == len(jobs) {
			break
		}
	}
	fmt.Println("sum is: ", sum)

}

//16026150
