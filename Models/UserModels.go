package Models

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username" binding:"required,alphanum"`
	Password   string `json:"password" binding:"required,min=6,alphanum"`
	Status     int    `json:"status"`
	Last_login string `json:"lastlogin"`
}

func (user *User) TableName() string {
	return "tbl_user"
}

type CreateUser struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role"`
}

func (user *CreateUser) TableName() string {
	return "tbl_user"
}

type DeleteUser struct {
	ID       int    `json:"id" binding:"required"`
	Username string `json:"username"`
	Status   int    `json:"status"`
}

func (user *DeleteUser) TableName() string {
	return "tbl_user"
}

type InputDeleteUser struct {
	Status int `json:"status"`
}

type UpdateUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func (user *UpdateUser) TableName() string {
	return "tbl_user"
}

type InputUpdateUser struct {
	Username string `json:"username"`
}

type ValidateUser struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Status     int    `json:"status"`
	Role       string `json:"role"`
	Last_login string `json:"lastlogin"`
}

func (user *ValidateUser) TableName() string {
	return "tbl_user"
}

type InputLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ViewUser struct {
	ID         int    `json:"id"`
	Username   string `json:"username" binding:"required,alphanum"`
	Last_login string `json:"lastlogin"`
}

func (user *ViewUser) TableName() string {
	return "v_user"
}
