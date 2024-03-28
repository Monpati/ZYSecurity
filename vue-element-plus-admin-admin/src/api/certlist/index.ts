import request from '@/axios'
import type { TableData } from './types'

export const getTableListApi = (data: PageInfo): Promise<IResponse<PageInfo>> => {
  // return request.get({ url: '/mock/example/list', params })
  return request.post({ url: 'http://192.168.2.112:8080/v1/certs/person', data })
}

export const getCorpListApi = (data: PageInfo): Promise<IResponse<PageInfo>> => {
  return request.post({ url: 'http://192.168.2.112:8080/v1/certs/corp', data })
}

export const updatePersonStatusAllow = (id, data: CertStatus) => {
  return request.post({ url: `http://192.168.2.112:8080/v1/cert/person/${id}/status`, data })
}

export const updatePersonStatusReject = (id, data: CertStatus) => {
  return request.post({ url: `http://192.168.2.112:8080/v1/cert/person/${id}/status`, data })
}

export const updateCorpStatusAllow = (id, data: CertStatus) => {
  return request.post({ url: `http://192.168.2.112:8080/v1/cert/corp/${id}/status`, data })
}

export const updateCorpStatusReject = (id, data: CertStatus) => {
  return request.post({ url: `http://192.168.2.112:8080/v1/cert/corp/${id}/status`, data })
}

export const getCertStatusApi = () => {
  return request.post({ url: 'http://192.168.2.112:8080/v1/cert/person/status' })
}
