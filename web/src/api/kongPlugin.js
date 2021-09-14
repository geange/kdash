import service from '@/utils/request'

export const createPlugin = (data) => {
  return service({
    url: '/kong/plugin/create',
    method: 'post',
    data
  })
}

export const updatePlugin = (data) => {
  return service({
    url: '/kong/plugin/update',
    method: 'put',
    data
  })
}

export const deletePlugin = (data) => {
  return service({
    url: '/kong/plugin/delete',
    method: 'delete',
    data
  })
}

export const getUpstream = (params) => {
  return service({
    url: '/kong/plugin/get',
    method: 'get',
    params
  })
}

export const listPlugin = (params) => {
  return service({
    url: '/kong/plugin/list',
    method: 'get',
    params
  })
}

export const listAllPlugin = (params) => {
  return service({
    url: '/kong/plugin/all',
    method: 'get',
    params
  })
}

export const listAllPluginForService = (params) => {
  return service({
    url: '/kong/plugin/service',
    method: 'get',
    params
  })
}

export const listAllPluginForConsumer = (params) => {
  return service({
    url: '/kong/plugin/consumer',
    method: 'get',
    params
  })
}

export const listAllPluginForRoute = (params) => {
  return service({
    url: '/kong/plugin/route',
    method: 'get',
    params
  })
}

export const pluginValidate = (params) => {
  return service({
    url: '/kong/plugin/validate',
    method: 'post',
    params
  })
}

export const pluginSchema = (params) => {
  return service({
    url: '/kong/plugin/schema',
    method: 'get',
    params
  })
}

