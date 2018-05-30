import axios from 'axios'
import config from '../../config'

const API_URL = config.api.host

export default axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json'
  },
  withCredentials: true
})