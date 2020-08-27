package counter

import (
	"bufio"
	"strings"
	"testing"
)


func TestCountGo(t *testing.T) {
	FileUrlName := "https://golang.org\n"
	expect := 20

	s1 := strings.NewReader(FileUrlName)
	r := bufio.NewReader(s1)
	_, count := CountGo(r, 5)

	if count != expect {
		t.Errorf("Wrong count for %s: expected %d got %d", FileUrlName, expect, count)
	}

	FileUrlName = "https://golang.org\nhttps://golang.org\n/etc/passwd\n"
	expect = 40

	s1 = strings.NewReader(FileUrlName)
	r = bufio.NewReader(s1)
	_, count = CountGo(r, 5)
	if count != expect {
		t.Errorf("Wrong count for %s: expected %d got %d", FileUrlName, expect, count)
	}

}
