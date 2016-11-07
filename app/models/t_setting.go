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

//Get setting value for key.
func (s *Setting) Get() (string, error) {
	res, err := support.Cache.Get(s.Key).Result()

	if res == "" {
		set := new(Setting)
		has, err := support.Xorm.Where("key = ?", s.Key).Get(&set)
		if has && err == nil {
			return set.Value, err
		}
	}

	return res, err
}

//Put setting K,V in cache and db.
func (s *Setting) Put() (bool, error) {
	has, err := support.Xorm.InsertOne(s)
	if err == nil && has > 0 {
		support.Cache.Set(s.Key, s.Value, 0)
	}
	return has > 0, err
}

//Update value for key.
func (s *Setting) Update() (bool, error) {

	set := new(Setting)
	set.Value = s.Value
	has, err := support.Xorm.Where("key = ?", s.Key).Update(&set)

	if err == nil && has > 0 {
		support.Cache.Del(s.Key)
		support.Cache.Set(s.Key, s.Value, 0)
	}

	return has > 0, err
}

//SiteInfo model
type SiteInfo struct {
	Title      string
	SubTitle   string
	Url        string
	Seo        string
	Reg        string
	Foot       string
	Statistics string
	Status     string
}

//Query site setting info.
func (s *Setting) GetSiteInfo() (*SiteInfo, string) {

	site := new(SiteInfo)
	data, err := s.FindAll()

	if err != "" {
		return site, err
	}

	if len(data) > 0 {
		for _, tmp := range data {
			switch tmp.Key {
			case "site-foot":
				site.Foot = tmp.Value
			case "site-reg":
				site.Reg = tmp.Value
			case "site-seo":
				site.Seo = tmp.Value
			case "site-status":
				site.Status = tmp.Value
			case "site-subtitle":
				site.SubTitle = tmp.Value
			case "site-title":
				site.Title = tmp.Value
			case "site-url":
				site.Url = tmp.Value
			case "site-statistics":
				site.Statistics = tmp.Value
			}
		}
	}
	return site, err
}

//Add new site info
func (s *Setting) NewSiteInfo(title, subtitle, url, seo, reg, foot,
	statistics, status string) error {

	set := new(Setting)

	if title != "" {
		s.Key = "site-title"
		res, err := s.Get()
		if err == nil && res != "" {
			set.Key = "site-title"
			set.Value = title
			has, err := set.Update()
			revel.INFO.Printf("NewSiteInfo::Put has: %v,error: %v", has, err)
			if err != nil {
				return err
			}
		} else {
			set.Key = "site-title"
			set.Value = title
			has, err := set.Put()
			revel.INFO.Printf("NewSiteInfo::Put has: %v,error: %v", has, err)
			if err != nil {
				return err
			}
		}
	}

	if subtitle != "" {
		s.Key = "site-subtitle"
		res, err := s.Get()
		if err == nil && res != "" {
			set.Key = "site-subtitle"
			set.Value = subtitle
			has, err := set.Update()
			revel.INFO.Printf("NewSiteInfo::Put has: %v,error: %v", has, err)
			if err != nil {
				return err
			}
		} else {
			set.Key = "site-subtitle"
			set.Value = subtitle
			has, err := set.Put()
			revel.INFO.Printf("NewSiteInfo::Put has: %v,error: %v", has, err)
			if err != nil {
				return err
			}
		}
	}

	if url != "" {
		s.Key = "site-url"
		res, err := s.Get()
		if err == nil && res != "" {
			set.Key = "site-url"
			set.Value = url
			has, err := set.Update()
			revel.INFO.Printf("NewSiteInfo::Put has: %v,error: %v", has, err)
			if err != nil {
				return err
			}
		} else {
			set.Key = "site-url"
			set.Value = url
			has, err := set.Put()
			revel.INFO.Printf("NewSiteInfo::Put has: %v,error: %v", has, err)
			if err != nil {
				return err
			}
		}
	}

	if seo != "" {
		s.Key = "site-seo"
		res, err := s.Get()
		if err == nil && res != "" {
			set.Key = "site-seo"
			set.Value = seo
			has, err := set.Update()
			revel.INFO.Printf("NewSiteInfo::Put has: %v,error: %v", has, err)
			if err != nil {
				return err
			}
		} else {
			set.Key = "site-seo"
			set.Value = seo
			has, err := set.Put()
			revel.INFO.Printf("NewSiteInfo::Put has: %v,error: %v", has, err)
			if err != nil {
				return err
			}
		}
	}

	if reg != "" {
		s.Key = "site-reg"
		res, err := s.Get()
		if err == nil && res != "" {
			set.Key = "site-reg"
			set.Value = reg
			has, err := set.Update()
			revel.INFO.Printf("NewSiteInfo::Put has: %v,error: %v", has, err)
			if err != nil {
				return err
			}
		} else {
			set.Key = "site-reg"
			set.Value = reg
			has, err := set.Put()
			revel.INFO.Printf("NewSiteInfo::Put has: %v,error: %v", has, err)
			if err != nil {
				return err
			}
		}
	}

	if foot != "" {
		s.Key = "site-foot"
		res, err := s.Get()
		if err == nil && res != "" {
			set.Key = "site-foot"
			set.Value = foot
			has, err := set.Update()
			revel.INFO.Printf("NewSiteInfo::Put has: %v,error: %v", has, err)
			if err != nil {
				return err
			}
		} else {
			set.Key = "site-foot"
			set.Value = foot
			has, err := set.Put()
			revel.INFO.Printf("NewSiteInfo::Put has: %v,error: %v", has, err)
			if err != nil {
				return err
			}
		}
	}

	if statistics != "" {
		s.Key = "site-statistics"
		res, err := s.Get()
		if err == nil && res != "" {
			set.Key = "site-statistics"
			set.Value = statistics
			has, err := set.Update()
			revel.INFO.Printf("NewSiteInfo::Put has: %v,error: %v", has, err)
			if err != nil {
				return err
			}
		} else {
			set.Key = "site-statistics"
			set.Value = statistics
			has, err := set.Put()
			revel.INFO.Printf("NewSiteInfo::Put has: %v,error: %v", has, err)
			if err != nil {
				return err
			}
		}
	}

	if status != "" {
		s.Key = "site-status"
		res, err := s.Get()
		if err == nil && res != "" {
			set.Key = "site-status"
			set.Value = status
			has, err := set.Update()
			revel.INFO.Printf("NewSiteInfo::Put has: %v,error: %v", has, err)
			if err != nil {
				return err
			}
		}
		set.Key = "site-status"
		set.Value = status
		has, err := set.Put()
		revel.INFO.Printf("NewSiteInfo::Put has: %v,error: %v", has, err)
		if err != nil {
			return err
		}
	}
	return nil
}
