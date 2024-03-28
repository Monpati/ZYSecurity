import request from '@/axios'
import type { TableData } from './types'

export const getTableListApi = (data: PageInfo): Promise<IResponse<PageInfo>> => {
  // return request.get({ url: '/mock/example/list', params })
  return request.post({ url: 'http://192.168.2.112:8080/v1/accounts', data })
}

export const getCardTableListApi = (params: any) => {
  return request.get({ url: '/mock/card/list', params })
}

export const getTreeTableListApi = (params: any) => {
  return request.get({ url: '/mock/example/treeList', params })
}

export const saveTableApi = (data: Partial<TableData>): Promise<IResponse> => {
  return request.post({ url: '/mock/example/save', data })
}

export const getTableDetApi = (id: string): Promise<IResponse<TableData>> => {
  return request.get({ url: '/mock/example/detail', params: { id } })
}

export const delTableListApi = (ids: string[] | number[]): Promise<IResponse> => {
  return request.post({ url: '/mock/example/delete', data: { ids } })
}
