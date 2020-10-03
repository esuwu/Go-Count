package open

import "testing"

func TestOpenError(t *testing.T) {
	FileUrlName := "kek111"
	reader, err  := Open(FileUrlName)
	defer reader.Close()
	if err == nil {
		t.Error("There is a problem with opening files")
	}
}
