package support

import (
	"crypto/md5"
	"encoding/hex"
)

type Sign struct {
	Src string
	Key string
}

// Create md5 strings.
func (s *Sign) GetMd5() string {

	if s.Src == "" || s.Key == "" {
		return ""
	}

	data := s.Src + s.Key

	h := md5.New()
	h.Write([]byte(data))
	cipherStr := h.Sum(nil)

	return hex.EncodeToString(cipherStr)
}
