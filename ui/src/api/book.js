import request from '@/utils/request'

export function bookList() {
  return request({
    url: '/books',
    method: 'get',
  })
}

export function bookDelete(data) {
  return request({
    url: '/books',
    method: 'delete',
    params: query
  })
}

// export function bookAdd(data) {
//   return request({
//     url: '/book/book_add',
//     method: 'post',
//     data
//   })
// }

// export function bookUpdate(data) {
//   return request({
//     url: '/book/book_update',
//     method: 'post',
//     data
//   })
// }
