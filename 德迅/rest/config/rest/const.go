package rest

import "time"

const (
	PackageSize = 30

	Jwt_Context_Key_User_Obj = "jwt_author"
	Jwt_Bearer               = "Bearer "
	Jwt_Bearer_Length        = len(Jwt_Bearer)
	Jwt_Token_Expire         = time.Hour * 24 * 365

	Jwt_Salt_Len    = 32
	Jwt_Default_Pwd = "888888"

)
