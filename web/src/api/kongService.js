import service from '@/utils/request'

export const createService = (data) => {
  return service({
    url: '/kong/service/create',
    method: 'post',
    data
  })
}

export const updateService = (data) => {
  return service({
    url: '/kong/service/update',
    method: 'put',
    data
  })
}

export const deleteService = (data) => {
  return service({
    url: '/kong/service/delete',
    method: 'delete',
    data
  })
}

export const getService = (params) => {
  return service({
    url: '/kong/service/get',
    method: 'get',
    params
  })
}

export const listService = (params) => {
  return service({
    url: '/kong/service/list',
    method: 'get',
    params
  })
}

export const listAllService = (params) => {
  return service({
    url: '/kong/service/all',
    method: 'get',
    params
  })
}
