import service from '@/utils/request'

export const createConsumer = (data) => {
  return service({
    url: '/kong/consumer/create',
    method: 'post',
    data
  })
}

export const updateConsumer = (data) => {
  return service({
    url: '/kong/consumer/update',
    method: 'put',
    data
  })
}

export const deleteConsumer = (data) => {
  return service({
    url: '/kong/consumer/delete',
    method: 'delete',
    data
  })
}

export const getConsumer = (data) => {
  return service({
    url: '/kong/target/get',
    method: 'get',
    data
  })
}

export const getConsumerByCustomID = (data) => {
  return service({
    url: '/kong/consumer/custom_id',
    method: 'get',
    data
  })
}

export const listConsumer = (params) => {
  return service({
    url: '/kong/consumer/list',
    method: 'get',
    params
  })
}

export const listAllConsumer = (params) => {
  return service({
    url: '/kong/consumer/all',
    method: 'get',
    params
  })
}
