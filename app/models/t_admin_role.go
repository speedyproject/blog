package models

//AdminRole model
type AdminRole struct {
	Id       int    `xorm:"not null pk autoincr INT(11)"`
	RoleDes  string `xorm:"VARCHAR(20)"`
	RoleType int    `xorm:"INT(11)"`
}
