package models

import (
	"blog/app/support"

	"github.com/revel/revel"
)

//Setting model
type Setting struct {
	Key   string `xorm:"not null pk VARCHAR(20)"`
	Value string `xorm:"not null VARCHAR(255)"`
}

//Loaded setting info to cache.
func LoadCache() {
	set := new(Setting)
	res, err := set.FindAll()

	if err != "" {
		revel.ERROR.Printf("Loaded setting info to cache error: %v", err)
	}
	if len(res) > 0 {
		for i := 0; i < len(res); i++ {
			s := res[0]
			support.Cache.Set(s.Key, s.Value, 0)
		}
	}
}

//find all setting info.
func (s *Setting) FindAll() ([]Setting, string) {

	set := make([]Setting, 0)

	err := support.Xorm.Find(&set)

	if err != nil {
		return set, err.Error()
	}

	return set, ""
}
