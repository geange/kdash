import service from '@/utils/request'

export const createRoute = (data) => {
  return service({
    url: '/kong/route/create',
    method: 'post',
    data
  })
}

export const createRouteInService = (data) => {
  return service({
    url: '/kong/route/service',
    method: 'post',
    data
  })
}

export const updateRoute = (data) => {
  return service({
    url: '/kong/route/update',
    method: 'put',
    data
  })
}

export const deleteRoute = (data) => {
  return service({
    url: '/kong/route/delete',
    method: 'delete',
    data
  })
}

export const getRoute = (params) => {
  return service({
    url: '/kong/route/get',
    method: 'get',
    params
  })
}

export const listRoute = (params) => {
  return service({
    url: '/kong/route/list',
    method: 'get',
    params
  })
}

export const listAllRoute = (params) => {
  return service({
    url: '/kong/route/all',
    method: 'get',
    params
  })
}

export const listAllRouteInService = (params) => {
  return service({
    url: '/kong/route/service',
    method: 'get',
    params
  })
}
