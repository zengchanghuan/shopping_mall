package models

type Manager struct {
	Id       int
	Username string
	Password string
	Mobile   string
	Email    string
	Status   int
	RoleId   int
	AddTime  int
	IsSuper  int
	//配置关联关系
	Role Role `gorm:"foreignKey:RoleId;references:Id"`
}

func (Manager) TableName() string {
	return "manager"
}
