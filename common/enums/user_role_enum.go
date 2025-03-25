package enums

type userRoleIDEnum struct {
	Admin    int
	Manager  int
	Employee int
}

var UserRoleID = userRoleIDEnum{
	Admin:    4,
	Manager:  5,
	Employee: 6,
}

