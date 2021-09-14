import service from '@/utils/request'

export const createSNI = (data) => {
  return service({
    url: '/kong/sni/create',
    method: 'post',
    data
  })
}

export const updateSNI = (data) => {
  return service({
    url: '/kong/sni/update',
    method: 'put',
    data
  })
}

export const deleteSNI = (data) => {
  return service({
    url: '/kong/sni/delete',
    method: 'delete',
    data
  })
}

export const getSNI = (params) => {
  return service({
    url: '/kong/sni/get',
    method: 'get',
    params
  })
}

export const getSNIForCertificate = (params) => {
  return service({
    url: '/kong/sni/certificate',
    method: 'get',
    params
  })
}

export const listSNI = (params) => {
  return service({
    url: '/kong/sni/list',
    method: 'get',
    params
  })
}

export const listAllSNI = (params) => {
  return service({
    url: '/kong/sni/all',
    method: 'get',
    params
  })
}
