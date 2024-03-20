package controller

import "Dexun/config"

func GetAdminRoles() []string {
	roles := []string{
		config.OAuth_Type_Admin,
	}
	return roles
}
