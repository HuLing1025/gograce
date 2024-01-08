import Axios from 'axios'

const axios = Axios.create({
  baseURL: "", // you can set baseURL here
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json'
  }
})

axios.interceptors.request.use(config => {
  return config
}, error => {
  return Promise.reject(error)
})

axios.interceptors.response.use(response => {
  return response
}, error => {
  return Promise.reject(error)
})

const get = (url, params) => {
  return axios.get(url, { params: params })
}

const post = (url, data) => {
  return axios.post(url, data)
}

export {
  get, post
}