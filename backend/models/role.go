package models

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
	RoleGuest Role = "guest"
)

var RolePermissions = map[Role][]string{
	RoleAdmin: {"create_listing", "update_listing", "delete_listing", "view_listing"},
	RoleUser:  {"create_listing", "view_listing"},
	RoleGuest: {"view_listing"},
}
