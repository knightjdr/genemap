// Package download has functions for downloading files.
package download

import (
	"io"
	"net/http"
	"os"
)

// File will download a url to a local file.
func File(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return err
}
