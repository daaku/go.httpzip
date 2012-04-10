package httpzip

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Read(res *http.Response) (*zip.Reader, error) {
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response body with error %s", err)
	}
	r := bytes.NewReader(b)
	return zip.NewReader(r, int64(r.Len()))
}

func ReadURL(url string) (*zip.Reader, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf(
			"Fetching zip URL %s failed with error %s.", url, err)
	}
	defer res.Body.Close()
	return Read(res)
}
