// Package zip handles file compression and decompression.
package zip

import (
	"compress/gzip"
	"io"

	"github.com/knightjdr/genemap/pkg/fs"
)

// Gunzip a file.
func Gunzip(source, target string) error {
	reader, err := fs.Instance.Open(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	archive, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	defer archive.Close()

	writer, err := fs.Instance.Create(target)
	if err != nil {
		return err
	}
	defer writer.Close()

	_, err = io.Copy(writer, archive)
	return err
}
