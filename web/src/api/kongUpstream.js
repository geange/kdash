import service from '@/utils/request'

export const createUpstream = (data) => {
  return service({
    url: '/kong/upstream/create',
    method: 'post',
    data
  })
}

export const updateUpstream = (data) => {
  return service({
    url: '/kong/upstream/update',
    method: 'put',
    data
  })
}

export const deleteService = (data) => {
  return service({
    url: '/kong/upstream/delete',
    method: 'delete',
    data
  })
}

export const getUpstream = (params) => {
  return service({
    url: '/kong/upstream/get',
    method: 'get',
    params
  })
}

export const listUpstream = (params) => {
  return service({
    url: '/kong/upstream/list',
    method: 'get',
    params
  })
}

export const listAllUpstream = (params) => {
  return service({
    url: '/kong/upstream/all',
    method: 'get',
    params
  })
}
