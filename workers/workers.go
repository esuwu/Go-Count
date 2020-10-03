package workers

import (
	"bytes"
	"github.com/esuwu/Counting-Go/open"
	"io"
	"strings"
	"sync"
)

type ResultT struct {
	Count int
	From  string
	Err   error
}

func countInReader(r io.Reader, word string) (int, error) {
	var count int
	buf := make([]byte, 4096)
	counter := countInChunk(word)

	for {
		n, err := io.ReadFull(r, buf)
		if err != io.ErrUnexpectedEOF && err != nil {
			return 0, err
		}
		count += counter(buf[:n])
		if err == io.ErrUnexpectedEOF {
			break
		}
	}
	return count, nil
}

func countInChunk(word string) func(data []byte) int {
	length := len(word) - 1
	border := make([]byte, length*2)
	w := []byte(word)

	return func(data []byte) int {
		if len(data) < length {
			return bytes.Count(append(border[length:], data...), w)
		}
		copy(border[length:], data[:length])

		count := bytes.Count(border, w) + bytes.Count(data, w)
		copy(border[:length], data[len(data)-length:])
		return count
	}
}

func Worker(FileUrlName string, jobChan chan ResultT, wg *sync.WaitGroup, queueCh chan struct{}) {
	const word = "Go"
	queueCh<- struct{}{}
	defer wg.Done()
	FileUrlName = strings.TrimSuffix(FileUrlName, "\n")
	reader, err  := open.Open(FileUrlName)
	defer reader.Close()

	if err != nil {
		jobChan<-ResultT{From: FileUrlName, Err: err}
		return
	}
	count, err := countInReader(reader, word)
	if err != nil {
		jobChan<-ResultT{From: FileUrlName, Err: err}
		return
	}
	jobChan<-ResultT{
		Count: count,
		From: FileUrlName,
		Err: nil,
	}
	return
}
