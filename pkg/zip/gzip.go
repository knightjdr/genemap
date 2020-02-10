// Package zip handles file compression and decompression.
package zip

import (
	"compress/gzip"

	"github.com/knightjdr/genemap/pkg/fs"
)

// Gzip a string to a file.
func Gzip(sourceData, target string) error {
	writer, err := fs.Instance.Create(target)
	if err != nil {
		return err
	}
	defer writer.Close()

	zipper := gzip.NewWriter(writer)
	_, err = zipper.Write([]byte(sourceData))
	if err != nil {
		return err
	}

	if err := zipper.Close(); err != nil {
		return err
	}

	return err
}
