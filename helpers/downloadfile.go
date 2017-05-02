package helpers

import (
	"errors"
	"io"
	"net/http"
	"os"
)

// Download file from url and save it as filepath
func DownloadFile(url, filepath string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return errors.New("Download file error: " + err.Error())
	}
	defer out.Close()

	if err != nil {
		return errors.New("Download file error: " + err.Error())
	}

	request, err := http.NewRequest("GET", url, nil)
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return errors.New("Download file error: " + err.Error())
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return errors.New("Download file error: " + err.Error())
	}

	return nil
}
