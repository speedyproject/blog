package models

import "time"
import "blog/app/support"
import "encoding/json"
import "strconv"

// Blogger model.
type Blogger struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Title      string    `xorm:"not null VARCHAR(50)"`
	Context    string    `xorm:"not null TEXT"`
	TagId      string    `xorm:"VARCHAR(20)"`
	LabelId    string    `xorm:"VARCHAR(20)"`
	Passwd     string    `xorm:"VARCHAR(64)"`
	CreateTime time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	CreateBy   int       `xorm:"not null INT(11)"`
	ReadCount  int64     `xorm:"default 0 BIGINT(20)"`
	LeaveCount int64     `xorm:"default 0 BIGINT(20)"`
}

// Get blogger list.
func (b *Blogger) FindList() ([]Blogger, error) {
	// get list data from cache.
	list := make([]Blogger, 0)
	res, _ := support.Cache.Get(support.SPY_BLOGGER_LIST).Result()

	if res != "" {
		err := json.Unmarshal([]byte(res), &list)
		if err == nil {
			return list, err
		}
	}
	// if list data is null in cache,get list data in db.
	err := support.Xorm.Find(&list)

	if err == nil {
		res, e1 := json.Marshal(&list)
		if e1 != nil {
			support.Cache.Set(support.SPY_BLOGGER_LIST, string(res), 0)
		}
	}

	return list, err
}

//Add new blogger.
func (b *Blogger) New() (bool, error) {

	blog := new(Blogger)

	blog.Title = b.Title
	blog.Context = b.Context
	blog.CreateBy = b.CreateBy
	blog.CreateTime = time.Now()
	blog.Passwd = b.Passwd
	blog.LabelId = b.LabelId
	blog.TagId = b.TagId

	has, err := support.Xorm.InsertOne(&blog)

	// refurbish cache.
	if err == nil {
		list := make([]Blogger, 0)
		err := support.Xorm.Find(&list)
		if err == nil {
			res, e1 := json.Marshal(&list)
			if e1 != nil {
				support.Cache.Set(support.SPY_BLOGGER_LIST, string(res), 0)
			}
		}
	}

	return has > 0, err
}

// find blogger by id.
func (b *Blogger) FindById() (*Blogger, error) {

	blog := new(Blogger)
	// Get single blogger from cache.
	res, e1 := support.Cache.Get(support.SPY_BLOGGER_SINGLE + strconv.Itoa(b.Id)).Result()

	if e1 == nil {
		e2 := json.Unmarshal([]byte(res), &blog)
		if e2 == nil {
			return blog, nil
		}
	}
	// if cache not blogger data, find in db.
	_, err := support.Xorm.Id(b.Id).Get(&blog)

	if err != nil {
		return blog, err
	}

	return blog, err
}

// Update blogger.
func (b *Blogger) Update() (bool, error) {

	has, err := support.Xorm.Id(b.Id).Update(&b)

	if err == nil {
		// refurbish cache.
		res, e1 := json.Marshal(&b)
		if e1 == nil {
			support.Cache.Del(support.SPY_BLOGGER_SINGLE + strconv.Itoa(b.Id))
			support.Cache.Set(support.SPY_BLOGGER_SINGLE+strconv.Itoa(b.Id), string(res), 0)
		}
	}

	return has > 0, err
}

// Delete blogger.
func (b *Blogger) Del() (bool, error) {

	has, err := support.Xorm.Id(b.Id).Delete(&b)

	if err == nil {
		// Delete cache.
		support.Cache.Del(support.SPY_BLOGGER_SINGLE + strconv.Itoa(b.Id))
	}

	return has > 0, err
}
