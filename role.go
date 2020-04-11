package thousandeyes

// AccountGroupRole - an account group role
type AccountGroupRole struct {
	RoleName                 string       `json:"roleName,omitempty"`
	RoleID                   int          `json:"roleId,omitempty"`
	HasManagementPermissions int          `json:"hasManagementPermissions,omitempty"`
	Builtin                  int          `json:"builtin,omitempty"`
	Permissions              []Permission `json:"permissions,omitempty"`
}

// Permission - permission attached to roles
type Permission struct {
	IsManagementPermission int    `json:"isManagementPermission"`
	Label                  string `json:"label"`
	PermissionID           int    `json:"permissionId"`
}
