package controllers

import (
	"bytes"
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"mime/multipart"
	"os"

	"github.com/revel/revel"
)

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

type Upload struct {
	*revel.Controller
}

type FileInfo struct {
	ContentType string
	Filename    string
	RealFormat  string `json:",omitempty"`
	Resolution  string `json:",omitempty"`
	Size        int
	Status      string `json:",omitempty"`
}

func (c *Upload) Before() revel.Result {
	// Rendering useful info here.
	c.RenderArgs["action"] = c.Controller.Action

	return nil
}

func (c *Upload) HandleUpload() revel.Result {
	var files [][]byte
	c.Params.Bind(&files, "file")
	filesInfo := make([]FileInfo, len(files))

	for _, vv := range c.Params.Files["file"] { //此处的file须和view中一致。
		// Create buffer
		buf := new(bytes.Buffer)
		// create a tmpfile and assemble your multipart from there (not tested)
		w := multipart.NewWriter(buf)
		// Create file field
		fsrc, _, _ := c.Request.FormFile("file")
		fdst, err := os.OpenFile(revel.BasePath+"/public/file/"+vv.Filename, os.O_CREATE|os.O_WRONLY, 0677)
		if err != nil {
			fmt.Println("err ", err)
			return nil
		}
		defer fdst.Close()
		// Write file field from file to upload
		_, err = io.Copy(fdst, fsrc)
		if err != nil {
			fmt.Println("e")
			return nil
		}
		fsrc.Close()
		w.Close()
	}

	return c.RenderJson(map[string]interface{}{
		"Count":  len(files),
		"Files":  filesInfo,
		"Status": "Successfully uploaded",
	})
}
