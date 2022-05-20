package listener

import (
	"os"
	"path"
	"testing"
)

func TestGoGetter(t *testing.T) {
	url := "https://google.com"

	c := reqClient()
	err := GoGetter(url, c)
	if err != nil {
		t.Error("An error occured during this request")
	}

	fileName := path.Base(url)
	fInfo, err := os.Stat(fileName + ".html")
	if err != nil {
		t.Error("Could not read file info")
	}

	fSize := fInfo.Size()
	if fSize < 1 {
		t.Error("response body is too low")
	}
}
