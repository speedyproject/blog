package models

import (
	"blog/app/support"
	"github.com/alecthomas/log4go"
	"net/http"
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
func (a *Admin) SignIn(request *http.Request) (Admin, error) {

	if a.Name == "" || a.Passwd == "" {
		return error.Error("username or passwd can't be null.")
	}

	admin := new(Admin)

	sign_key := support.Cache.Get(support.SPY_CONF_MD5_KEY).String()
	sign := &support.Sign{a.Passwd, Key: sign_key}

	a.Passwd = sign.GetMd5()

	_, err := utils.Orm.Where("name = ? and passwd = ?", a.Name, a.Passwd).Get(&admin)

	if err != nil {
		return admin, error.Error("login failed.")
	}

	if admin.Lock > 0 {
		return admin, error.Error("login failed, the account is lock.")
	}

	if strings.EqualFold(a.Name, admin.Name) && strings.EqualFold(a.Passwd, admin.Passwd) {
		lastIp := support.GetRequestIP(request)

		ad := new(Admin)
		ad.LastIp = lastIp
		ad.LastLogin = time.Time.UTC()
		_, e1 := support.Xorm.Id(admin.Id).Get(ad)

		if e1 != nil {
			log4go.Error(e1)
		}

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

	admin := new(Admin)
	_, e1 := support.Xorm.Id(a.Id).Get(admin)

	if e1 != nil {
		return false, e1
	}

	if !strings.EqualFold(old_pwd, admin.Passwd) {
		return false, error.Error("change passwd failed, old passwd error.")
	}

	admin = new(Admin)
	admin.Passwd = new_pwd
	has, e2 := support.Xorm.Id(a.Id).Update(&admin)

	if e2 != nil {
		return false, e2
	}

	return has > 0, nil
}
