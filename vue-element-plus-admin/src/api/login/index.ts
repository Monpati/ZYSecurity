import request from '@/axios'
import type { UserType } from './types'

interface RoleParams {
  roleName: string
}

export const codeApi = (): Promise<IResponse> => {
  return request.get({ url: 'http://192.168.2.112:8080/v1/getCode' })
  // return request.get({ url: '/getCode' })
}

export const regApi = (data: UserReg): Promise<IResponse<UserReg>> => {
  return request.post({ url: 'http://192.168.2.112:8080/v1/register', data })
  // return request.post({ url: '/register', data })
}

export const loginApi = (data: UserType): Promise<IResponse<UserType>> => {
  return request.post({ url: 'http://192.168.2.112:8080/v1/login', data })
  // return request.post({ url: 'http://localhost:8080/v1/login', data })
  // return request.post({ url: '/mock/user/login', data })
}

export const loginOutApi = (): Promise<IResponse> => {
  return request.get({ url: 'http://192.168.2.112:8080/v1/logout' })
  // return request.get({ url: '/logout' })
  // return request.get({ url: '/mock/user/loginOut' })
}

export const getUserListApi = ({ params }: AxiosConfig) => {
  return request.get<{
    code: string
    data: {
      list: UserType[]
      total: number
    }
  }>({ url: '/mock/user/list', params })
}

export const getAdminRoleApi = (
  data: RoleParams
): Promise<IResponse<AppCustomRouteRecordRaw[]>> => {
  return request.post({ url: 'http://192.168.2.112:8080/v1/role', data })
  // return request.post({ url: '/role', data })
}

export const getTestRoleApi = (params: RoleParams): Promise<IResponse<string[]>> => {
  return request.get({ url: '/mock/role/list2', params })
}
