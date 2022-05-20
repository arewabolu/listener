package listener

import (
	"io"
	"net/http"
	"os"
	"path"
)

func reqClient() *http.Client {
	client := &http.Client{}
	return client
}

func fCreator(fName string) *os.File {

	nwFile, err := os.Create(fName)
	if err != nil {
		panic("could not create file")
	}
	return nwFile
}

func GoGetter(url string, c *http.Client) (err error) {
	if url == "" {
		panic("url is empty. Expected non-empty url!")
	}

	resp, err := c.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	fileName := path.Base(url)
	nwFile := fCreator(fileName + ".html")
	defer nwFile.Close()

	_, err = io.Copy(nwFile, resp.Body)
	return
}
