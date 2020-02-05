// Package download has functions for downloading files.
package download

import (
	"io"
	"net/http"

	"github.com/knightjdr/genemap/pkg/fs"
)

// HTTP will download an http url to a local file.
func HTTP(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := fs.Instance.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return err
}
