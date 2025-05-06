package jewelrymodel

type User struct {
	Id          int    `json:"-" db:"id"`
	Login       string `json:"login" db:"login"`
	Password    string `json:"password" db:"password"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	FatherName  string `json:"father_name" db:"father_name"`
	Age         int    `json:"age" db:"age"`
	Email       string `json:"email" db:"email"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
}

type LoginUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
