export interface UserLoginType {
  username: string
  password: string
}

export interface UserType {
  username: string
  password: string
  token: string
  // role: string
  // roleId: string
}

export interface UserReg {
  username: string
  password: string
  email: string
  tel_num: string
  code: string
}
