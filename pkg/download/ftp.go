package download

import (
	"io"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/knightjdr/genemap/pkg/fs"
)

// FTP will download an ftp url to a local file.
func FTP(url, sourceFile, targetFile string) error {
	connection, err := ftp.Dial(url, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return err
	}

	err = connection.Login("anonymous", "anonymous")
	if err != nil {
		return err
	}

	resp, err := connection.Retr(sourceFile)
	if err != nil {
		return err
	}

	out, err := fs.Instance.Create(targetFile)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp)

	if err := connection.Quit(); err != nil {
		return err
	}

	return err
}
