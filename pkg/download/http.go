package download

import (
	"io"
	"net/http"

	"github.com/knightjdr/genemap/pkg/fs"
)

// HTTP will download an http url to a local file.
func HTTP(url string, headers map[string]string, targetFile string) error {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := fs.Instance.Create(targetFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}
