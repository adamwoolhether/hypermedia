package contacts

type Contact struct {
	ID        int    `json:"id"`
	FirstName string `json:"first"`
	LastName  string `json:"last"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}
