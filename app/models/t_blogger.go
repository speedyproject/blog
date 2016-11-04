package models

import "time"
import "blog/app/support"
import "encoding/json"

// Blogger model.
type Blogger struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Title      string    `xorm:"not null VARCHAR(50)"`
	Context    string    `xorm:"not null TEXT"`
	LabelId    int       `xorm:"default 0 INT(11)"`
	Passwd     string    `xorm:"VARCHAR(64)"`
	CreateTime time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	CreateBy   int       `xorm:"not null INT(11)"`
	ReadCount  int64     `xorm:"default 0 BIGINT(20)"`
	LeaveCount int64     `xorm:"default 0 BIGINT(20)"`
}

// Get blogger list.
func (b *Blogger) FindList() ([]Blogger, error) {
	// get list data from cache
	list := make([]Blogger, 0)
	res, _ := support.Cache.Get(support.SPY_BLOGGER_LIST).Result()

	if res != "" {
		err := json.Unmarshal([]byte(res), &list)
		if err == nil {
			return list, err
		}
	}
	// if list data is null in cache,get list data in db
	err := support.Xorm.Find(&list)

	if err == nil {
		res, e1 := json.Marshal(&list)
		if e1 != nil {
			support.Cache.Set(support.SPY_BLOGGER_LIST, string(res), 0)
		}
	}

	return list, err
}
