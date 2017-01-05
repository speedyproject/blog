package models

import "blog/app/support"
import "github.com/revel/revel"

//AdminRole model
type AdminRole struct {
	Id       int    `xorm:"not null pk autoincr INT(11)"`
	RoleDes  string `xorm:"VARCHAR(20)"`
	RoleType int    `xorm:"INT(11)"`
}

//find role data for type
func (a *AdminRole) FindByType(id int) (*AdminRole, string) {

	ar := new(AdminRole)

	_, err := support.Xorm.Where("role_type = ?", id).Get(&ar)

	if err != nil {
		return ar, err.Error()
	}

	revel.INFO.Println(ar)

	return ar, ""
}
