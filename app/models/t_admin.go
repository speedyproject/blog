package models

import (
	"blog/app/support"
	"strings"
	"time"
	"github.com/revel/revel"
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
func (a *Admin) SignIn(request *revel.Request) (*Admin, string) {

	admin := new(Admin)

	if a.Name == "" || a.Passwd == "" {
		return admin, "username or passwd can't be null."
	}

	sign_key := support.Cache.Get(support.SPY_CONF_MD5_KEY).String()
	sign := &support.Sign{Src:a.Passwd, Key: sign_key}

	a.Passwd = sign.GetMd5()

	_, err := support.Xorm.Where("name = ? and passwd = ?", a.Name, a.Passwd).Get(admin)

	if err != nil {
		return admin, "login failed."
	}

	if admin.Lock > 0 {
		return admin, "login failed, the account is lock."
	}

	if strings.EqualFold(a.Name, admin.Name) && strings.EqualFold(a.Passwd, admin.Passwd) {
		lastIp := support.GetRequestIP(request)

		ad := new(Admin)
		ad.LastIp = lastIp
		ad.LastLogin = time.Now()
		_, e1 := support.Xorm.Id(admin.Id).Get(ad)

		if e1 != nil {
			revel.ERROR.Println(e1)
		}

		return admin, ""
	}

	return admin, "login failed."
}

//Add new admin user.
func (a *Admin) New() (int64, string) {

	if a.Name == "" || a.Passwd == "" {
		return 0, "username or passwd can't be null."
	}

	md5_key := support.Cache.Get(support.SPY_CONF_MD5_KEY).String()
	sign_key := support.Cache.Get(support.SPY_CONF_SIGN_KEY).String()

	revel.INFO.Printf("MD5_Key: %s, Sign_Key: %s", md5_key, sign_key)

	passwd := &support.Sign{Src: a.Passwd, Key: md5_key}
	sign := &support.Sign{Src: a.Name + a.Passwd, Key: sign_key}

	a.Sign = sign.GetMd5()
	a.Passwd = passwd.GetMd5()

	revel.INFO.Println(a)

	res, err := support.Xorm.InsertOne(a)

	if err != nil {
		revel.ERROR.Println(err)
		return 0, "create new admin user failed."
	}

	return res, ""
}

//Admin change password.
func (a *Admin) ChangePasswd(old_pwd, new_pwd string) (bool, string) {

	if old_pwd == "" || new_pwd == "" {
		return false, "old passwd or new passwd can't be null."
	}

	key := support.Cache.Get(support.SPY_CONF_MD5_KEY).String()

	o := &support.Sign{Src: old_pwd, Key: key}
	n := &support.Sign{Src: new_pwd, Key: key}

	old_pwd = o.GetMd5()
	new_pwd = n.GetMd5()

	admin := new(Admin)
	_, e1 := support.Xorm.Id(a.Id).Get(admin)

	if e1 != nil {
		return false, e1.Error()
	}

	if !strings.EqualFold(old_pwd, admin.Passwd) {
		return false, "change passwd failed, old passwd error."
	}

	admin = new(Admin)
	admin.Passwd = new_pwd
	has, e2 := support.Xorm.Id(a.Id).Update(&admin)

	if e2 != nil {
		return false, e2.Error()
	}

	return has > 0, ""
}
