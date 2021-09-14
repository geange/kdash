import service from '@/utils/request'

export const createCACertificate = (data) => {
  return service({
    url: '/kong/ca_certificate/create',
    method: 'post',
    data
  })
}

export const updateCACertificate = (data) => {
  return service({
    url: '/kong/ca_certificate/update',
    method: 'put',
    data
  })
}

export const deleteCACertificate = (data) => {
  return service({
    url: '/kong/ca_certificate/delete',
    method: 'delete',
    data
  })
}

export const getCACertificate = (params) => {
  return service({
    url: '/kong/ca_certificate/get',
    method: 'get',
    params
  })
}

export const listCACertificate = (params) => {
  return service({
    url: '/kong/ca_certificate/list',
    method: 'get',
    params
  })
}

export const listAllCACertificate = (params) => {
  return service({
    url: '/kong/ca_certificate/all',
    method: 'get',
    params
  })
}
