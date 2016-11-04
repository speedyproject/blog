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
		return
	}
	if len(res) > 0 {
		//select db 1
		support.Cache.Pipeline().Select(1)
		for i := 0; i < len(res); i++ {
			s := res[0]
			support.Cache.Set(s.Key, s.Value, 0)
		}
		//select db 0
		support.Cache.Pipeline().Select(0)
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

//Get setting value for key.
func (s *Setting) Get() (string, error) {
	//select db 1
	support.Cache.Pipeline().Select(1)
	res, err := support.Cache.Get(s.Key).Result()
	//select db 0
	support.Cache.Pipeline().Select(0)
	return res, err
}

//Put setting K,V in cache and db.
func (s *Setting) Put() (bool, error) {
	//select db 1
	support.Cache.Pipeline().Select(1)

	has, err := support.Xorm.InsertOne(s)

	if err == nil && has > 0 {
		support.Cache.Set(s.Key, s.Value, 0)
	}
	//select db 0
	support.Cache.Pipeline().Select(0)
	return has > 0, err
}

//Update value for key.
func (s *Setting) Update() (bool, error) {
	//select db 1
	support.Cache.Pipeline().Select(1)

	set := new(Setting)
	set.Value = s.Value
	has, err := support.Xorm.Where("key = ?", s.Key).Update(&set)

	if err == nil && has > 0 {
		support.Cache.Del(s.Key)
		support.Cache.Set(s.Key, s.Value, 0)
	}

	//select db 0
	support.Cache.Pipeline().Select(0)
	return has > 0, err
}
