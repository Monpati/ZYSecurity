import request from '@/axios'
import type { PersonalCert, CorpCert } from './types'

export const personalCertApi = (data: PersonalCert): Promise<IResponse<PersonalCert>> => {
  return request.post({ url: 'http://192.168.2.112:8080/v1/cert/person', data })
}

export const personalCertStatusApi = () => {
  return request.post({ url: 'http://192.168.2.112:8080/v1/cert/person/status' })
}

export const corpCertApi = (data: CorpCert): Promise<IResponse<CorpCert>> => {
  return request.post({ url: 'http://192.168.2.112:8080/v1/cert/corp', data })
}

export const corpCertStatusApi = () => {
  return request.post({ url: 'http://192.168.2.112:8080/v1/cert/corp/status' })
}
