import service from '@/utils/request'

export const createUpstream = (data) => {
  return service({
    url: '/kong/target/create',
    method: 'post',
    data
  })
}

export const deleteService = (data) => {
  return service({
    url: '/kong/target/delete',
    method: 'delete',
    data
  })
}

export const listUpstream = (params) => {
  return service({
    url: '/kong/target/list',
    method: 'get',
    params
  })
}

export const listAllUpstream = (params) => {
  return service({
    url: '/kong/target/all',
    method: 'get',
    params
  })
}

export const markHealthy = (data) => {
  return service({
    url: '/kong/target/healthy',
    method: 'post',
    data
  })
}

export const markUnhealthy = (data) => {
  return service({
    url: '/kong/target/unhealthy',
    method: 'post',
    data
  })
}
