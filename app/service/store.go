package service

import (
	"fmt"
	"io"
	"os"

	"mime/multipart"

	"github.com/revel/revel"
)

// StoreFile to store a file to local
// return filepaht, file size and error
func StoreFile(filename string, fsrc multipart.File) (string, int64, error) {
	path := revel.BasePath + "/public/file/" + filename
	fdst, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0677)
	if err != nil {
		fmt.Println("err ", err)
		return "", -1, nil
	}
	defer fdst.Close()
	// Write file field from file to upload
	_, err = io.Copy(fdst, fsrc)
	if err != nil {
		fmt.Println("e")
		return "", -1, nil
	}
	finfo, _ := fdst.Stat()
	return "public/file/" + filename, finfo.Size(), nil
}
