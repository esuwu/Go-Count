package counter

import (
	"bufio"
	"fmt"
	"github.com/esuwu/Counting-Go/workers"
	"log"
	"strings"
	"sync"
)

func CountGo(reader *bufio.Reader, k int) (*strings.Builder, int) {
	var total int
	output := strings.Builder{}
	wg := new(sync.WaitGroup)
	jobs := make(chan workers.ResultT, k)
	queueCh := make(chan struct{}, k)
	readingFromCh := make(chan struct{})

	go func() {
		for job := range jobs {
			<-queueCh
			var result string
			if job.Err != nil {
				result = fmt.Sprintf("Error for %s: %s\n", job.From, job.Err.Error())
			} else {
				result = fmt.Sprintf("Count for %s: %d\n", job.From, job.Count)
				total += job.Count
			}
			output.Write([]byte(result))
		}
		readingFromCh<- struct{}{}
	}()

	for fileUrlName, err := reader.ReadString('\n'); len(fileUrlName) != 0; fileUrlName, err = reader.ReadString('\n') {
		if err != nil {
			log.Print("Error: ", err)
		}
		wg.Add(1)
		go workers.Worker(fileUrlName, jobs, wg, queueCh)
	}

	wg.Wait()
	close(jobs)
	<-readingFromCh
	return &output, total
}