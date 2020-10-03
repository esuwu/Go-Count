package workers

import (
	"github.com/esuwu/Counting-Go/open"
	"testing"
)

func TestCountInReader(t *testing.T) {
	expect := 20
	const word = "Go"
	FileUrlName := "https://golang.org"
	reader, err  := open.Open(FileUrlName)
	defer reader.Close()

	if err != nil {
		t.Error("Error in opening file or url")
	}
	count, err := countInReader(reader, word)
	if err != nil {
		t.Error("Error in counting")
	}

	if count != expect {
		t.Errorf("Wrong count for %s: expected %d got %d", FileUrlName, expect, count)
	}

}

