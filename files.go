package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const inputFilenameBase = "file"

// DownloadImage downloads an image into the images directory
func DownloadImage(url string, dest string, filenum int) (string, error)  {
	filename := getFilename(url, dest, filenum)
	out, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func getFilename(url string, dest string, filenum int) string {
	filenameWithoutExtension := fmt.Sprintf("%s%d", inputFilenameBase, filenum)
	filename := fmt.Sprintf("%s%s", filenameWithoutExtension, filepath.Ext(url))
	
	return filepath.Join(dest, filename)
}