package models

import (
	"blog/app/support"
	"strings"
	"time"

	"github.com/revel/revel"
)

//Admin model
type Admin struct {
	Id        int       `xorm:"not null pk autoincr INT(11)"`
	Name      string    `xorm:"not null VARCHAR(15)"`
	Passwd    string    `xorm:"not null VARCHAR(64)"`
	Email     string    `xorm:"VARCHAR(45)"`
	Skey      string    `xorm:"not null VARCHAR(64)"`
	Lock      int       `xorm:"default 0 INT(11)"`
	RoleId    int       `xorm:"default 1001 INT(11)"`
	LastIp    string    `xorm:"default '0.0.0.0' VARCHAR(20)"`
	LastLogin time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

//Admin sign in.
func (a *Admin) SignIn(request *revel.Request) (*Admin, string) {

	admin := new(Admin)

	if a.Name == "" || a.Passwd == "" {
		return admin, "username or passwd can't be null."
	}

	//Get MD5 key in cache
	signKey, _ := support.Cache.Get(support.SPY_CONF_MD5_KEY).Result()
	sign := &support.Sign{Src: a.Passwd, Key: signKey}
	a.Passwd = sign.GetMd5()

	_, err := support.Xorm.Where("name = ? and passwd = ?", a.Name, a.Passwd).Get(admin)

	if err != nil {
		return admin, err.Error()
	}

	revel.INFO.Printf("Admin user info: %v", admin)

	if admin.Lock > 0 {
		return admin, "login failed, the account is lock."
	}

	if strings.EqualFold(a.Name, admin.Name) && strings.EqualFold(a.Passwd, admin.Passwd) {
		lastIP := support.GetRequestIP(request)

		ad := new(Admin)
		ad.LastIp = lastIP
		ad.LastLogin = time.Now()
		_, e1 := support.Xorm.Id(admin.Id).Update(ad)

		if e1 != nil {
			revel.ERROR.Println(e1)
		}

		return admin, ""
	}

	return admin, "login failed."
}

//Add new admin user.
func (a *Admin) New() (int64, string) {

	if a.Name == "" || a.Passwd == "" || a.Email == "" {
		return 0, "username or passwd can't be null."
	}
	//Get MD5/Sign key in cache
	md5Key, _ := support.Cache.Get(support.SPY_CONF_MD5_KEY).Result()
	signKyey, _ := support.Cache.Get(support.SPY_CONF_SIGN_KEY).Result()

	revel.INFO.Printf("MD5_Key: %s, Sign_Key: %s", md5Key, signKyey)

	passwd := &support.Sign{Src: a.Passwd, Key: md5Key}
	sign := &support.Sign{Src: a.Name + a.Passwd, Key: signKyey}

	a.Skey = sign.GetMd5()
	a.Passwd = passwd.GetMd5()
	a.LastLogin = time.Now()

	if a.RoleId <= 0 {
		a.RoleId = 1003
	}

	revel.INFO.Println(a)

	res, err := support.Xorm.InsertOne(a)

	if err != nil {
		revel.ERROR.Println(err)
		return 0, "create new admin user failed."
	}

	return res, ""
}

//Admin change password.
func (a *Admin) ChangePasswd(oldPwd, newPwd string) (bool, string) {

	if oldPwd == "" || newPwd == "" {
		return false, "old passwd or new passwd can't be null."
	}
	//Get MD5 key in cache
	key, _ := support.Cache.Get(support.SPY_CONF_MD5_KEY).Result()

	o := &support.Sign{Src: oldPwd, Key: key}
	n := &support.Sign{Src: newPwd, Key: key}

	oldPwd = o.GetMd5()
	newPwd = n.GetMd5()

	admin := new(Admin)
	_, e1 := support.Xorm.Id(a.Id).Get(admin)

	if e1 != nil {
		return false, e1.Error()
	}

	if !strings.EqualFold(oldPwd, admin.Passwd) {
		return false, "change passwd failed, old passwd error."
	}

	admin = new(Admin)
	admin.Passwd = newPwd
	has, e2 := support.Xorm.Id(a.Id).Update(&admin)

	if e2 != nil {
		return false, e2.Error()
	}

	return has > 0, ""
}

//Group show user's admin Group
func (a *Admin) Group() string {
	if a.RoleId == 1001 {
		return "管理员"
	} else if a.RoleId == 1002 {
		return "发布者"
	} else {
		return "订阅者"
	}
}

// List all user
func (a *Admin) List() ([]Admin, error) {
	return a.listByDB()
}

func (a *Admin) listByDB() ([]Admin, error) {
	users := make([]Admin, 0)
	err := support.Xorm.Find(&users)
	return users, err
}
