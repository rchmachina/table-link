package models

type User struct {
	ID         string
	RoleID     string
	Name       string
	Email      string
	LastAccess string
	RoleName   string
}
type CreateNewUser struct {
	ID       string
	RoleID   string
	Name     string
	Email    string
	Password string
}

type Role struct {
	ID   string
	Name string
}

type RoleRight struct {
	RoleID  string
	Section string
	Route   string
	RCreate bool
	RRead   bool
	RUpdate bool
	RDelete bool
}
