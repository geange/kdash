import service from '@/utils/request'

export const createCertificate = (data) => {
  return service({
    url: '/kong/certificate/create',
    method: 'post',
    data
  })
}

export const updateCertificate = (data) => {
  return service({
    url: '/kong/certificate/update',
    method: 'put',
    data
  })
}

export const deleteCertificate = (data) => {
  return service({
    url: '/kong/certificate/delete',
    method: 'delete',
    data
  })
}

export const getCertificate = (params) => {
  return service({
    url: '/kong/certificate/get',
    method: 'get',
    params
  })
}

export const listCertificate = (params) => {
  return service({
    url: '/kong/certificate/list',
    method: 'get',
    params
  })
}

export const listAllCertificate = (params) => {
  return service({
    url: '/kong/certificate/all',
    method: 'get',
    params
  })
}
