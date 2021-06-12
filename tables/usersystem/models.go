package usersystem

/*User define en usuario*/
type User struct {
	PersonID     int    `json:"PersonID"`
	UserName     string `json:"UserName"`
	Email        string `json:"Email"`
	PasswordUser string `json:"PasswordUser"`
	ImageURL     string `json:"ImageURL"`
	RoleUserID   int    `json:"RoleUserID"`
	StateID      int    `json:"StateID"`
}

/*Role define el tipo de rol */
type Role struct {
	RoleID int    `json:"RoleID"`
	Name   string `json:"Name"`
}

/*Company define el tipo de rol */
type Company struct {
	CompanyID int    `json:"CompanyID"`
	Code      string `json:"Code"`
}

/*TokenLogin token que será generado al iniciar sesión*/
type TokenLogin struct {
	Token string `json:"Token,omitempty"`
}
