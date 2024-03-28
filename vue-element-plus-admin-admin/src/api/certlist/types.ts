export type TableData = {
  id: string
  author: string
  title: string
  content: string
  importance: number
  display_time: string
  pageviews: number
}

export type PageInfo = {
  pageIndex: bigint
  pageSize: bigint
  // offset: number
  // limit: number
}

export type CertStatus = {
  id: bigint
  Status: bigint
}

export interface PersonalCert {
  id: bigint
  user_id: bigint
  real_name: string
  card_id: string
  sex: string
  birthday: string
  city: string
  card_front: string
  card_back: string
  status: number
}
