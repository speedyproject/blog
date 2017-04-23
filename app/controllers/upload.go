package controllers

import (
	"blog/app/service"
	"bytes"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"

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
	Path        string
}

func (c *Upload) Before() revel.Result {
	// Rendering useful info here.
	c.ViewArgs["action"] = c.Controller.Action

	return nil
}

func (c *Upload) HandleUpload() revel.Result {
	fileCount := len(c.Params.Files["file"])
	filesInfo := make([]FileInfo, fileCount)

	for kk, vv := range c.Params.Files["file"] { //此处的file须和view中一致。
		// Create buffer
		buf := new(bytes.Buffer)
		// create a tmpfile and assemble your multipart from there (not tested)
		w := multipart.NewWriter(buf)
		// Create file field
		fsrc, _ := vv.Open()
		filepath, fsize, err := service.StoreFile(vv.Filename, fsrc)
		if err != nil {
			revel.ERROR.Println("store file error: ", err)
		}
		fsrc.Close()
		w.Close()
		revel.TRACE.Println("index ", filepath)

		filesInfo[kk] = FileInfo{
			ContentType: vv.Header.Get("Content-Type"),
			Filename:    vv.Filename,
			Size:        int(fsize / 1024),
			Path:        filepath,
		}
	}

	return c.RenderJSON(map[string]interface{}{
		"Count":  fileCount,
		"Files":  filesInfo,
		"Status": "Successfully uploaded",
	})
}
