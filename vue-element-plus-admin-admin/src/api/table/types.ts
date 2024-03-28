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
  pageIndex: number
  pageSize: number
  // offset: number
  // limit: number
}
