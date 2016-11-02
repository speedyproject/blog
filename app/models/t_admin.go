package models

import (
	"blog/app/support"
	"encoding/json"
	"github.com/alecthomas/log4go"
	"strconv"
	"strings"
	"time"
	"wechat/utils"
)

type Admin struct {
	Id        int       `xorm:"not null pk autoincr INT(11)"`
	Name      string    `xorm:"not null VARCHAR(15)"`
	Passwd    string    `xorm:"not null VARCHAR(64)"`
	Email     string    `xorm:"VARCHAR(45)"`
	Sign      string    `xorm:"not null VARCHAR(64)"`
	Lock      int       `xorm:"default 0 INT(11)"`
	LastIp    string    `xorm:"default '0.0.0.0' VARCHAR(20)"`
	LastLogin time.Time `xorm:"TIMESTAMP"`
}

//Admin sign in.
func (a *Admin) SignIn() (Admin, error) {

	if a.Name == "" || a.Passwd == "" {
		return error.Error("username or passwd can't be null.")
	}

	//first try login for cache use sign.
	sign_key := support.Cache.Get(support.SPY_CONF_SIGN_KEY).String()

	s := &support.Sign{Src: a.Name + a.Passwd, Key: sign_key}

	sign := s.GetMd5()

	res := support.Cache.Get(sign).String()

	admin := new(Admin)
	e := json.Unmarshal([]byte(res), admin)

	if e == nil {
		if strings.EqualFold(a.Name, admin.Name) && strings.EqualFold(a.Passwd, admin.Passwd) {
			return admin, nil
		}
	}

	//if cache login failed, handle db login.
	s.Src = a.Passwd
	a.Passwd = s.GetMd5()

	_, err := utils.Orm.Where("name = ? and passwd = ?", a.Name, a.Passwd).Get(&admin)

	if err != nil {
		return admin, error.Error("login failed.")
	}

	if strings.EqualFold(a.Name, admin.Name) && strings.EqualFold(a.Passwd, admin.Passwd) {
		return admin, nil
	}

	return admin, error.Error("login failed.")
}

//Add new admin user.
func (a *Admin) New() (int64, error) {

	if a.Name == "" || a.Passwd == "" {
		return false, error.Error("username or passwd can't be null.")
	}

	md5_key := support.Cache.Get(support.SPY_CONF_MD5_KEY).String()
	sign_key := support.Cache.Get(support.SPY_CONF_SIGN_KEY).String()

	log4go.Debug("MD5_Key: %s, Sign_Key: %s", md5_key, sign_key)

	passwd := &support.Sign{Src: a.Passwd, Key: md5_key}.GetMd5()
	sign := &support.Sign{Src: a.Name + a.Passwd, Key: sign_key}.GetMd5()

	a.Sign = sign
	a.Passwd = passwd

	log4go.Debug(a)

	res, err := support.Xorm.InsertOne(a)

	if err != nil {
		log4go.Debug(err)
		return false, error.Error("create new admin user failed.")
	}

	json, _ := json.Marshal(&a)
	support.Cache.Set(support.USER_LOGIN_BY_SIGN, string(json), 0)

	return res, nil
}

//Admin change password.
func (a *Admin) ChangePasswd(old_pwd, new_pwd string) (bool, error) {

	if old_pwd == "" || new_pwd == "" {
		return false, error.Error("old passwd or new passwd can't be null.")
	}

	key := support.Cache.Get(support.SPY_CONF_MD5_KEY).String()
	old_pwd = &support.Sign{Src: old_pwd, Key: key}.GetMd5()
	new_pwd = &support.Sign{Src: new_pwd, Key: key}.GetMd5()

	res := support.Cache.Get(support.USER_DATA_BY_ID + strconv.Itoa(a.Id)).String()
	admin := new(Admin)
	e1 := json.Unmarshal([]byte(res), admin)

	if e1 != nil {
		return false, error.Error("change fiald.")
	}

	if !strings.EqualFold(old_pwd, admin.Passwd) {
		return false, error.Error("change fiald, old passwd error.")
	}

	ad := new(Admin)
	ad.Passwd = new_pwd
	has, e2 := support.Xorm.Id(a.Id).Update(ad)

	if e2 != nil {
		return false, e2
	}

	support.Cache.Del(support.USER_DATA_BY_ID + strconv.Itoa(a.Id))
	admin.Passwd = new_pwd
	tmp, _ := json.Marshal(&admin)
	support.Cache.Set(support.USER_DATA_BY_ID+strconv.Itoa(a.Id), string(tmp), 0)

	return has > 0, nil
}
