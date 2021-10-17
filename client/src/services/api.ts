import axios, { AxiosInstance } from 'axios'
import { User } from '@/types'

const API_URL = process.env.VUE_APP_API_URL

const instance: AxiosInstance = axios.create({
  baseURL: API_URL || 'http://localhost:3000/api/v1/pf'
})

export async function getUser(): Promise<User> {
  return request('/github/user')
}

export async function getProjects(): Promise<User> {
  return request('github/projects')
}

async function request(resource: string): Promise<User> {
  try {
    const response = await instance.get(resource)
    const data = await response.data

    return data.user
  } catch (error) {
    console.log(error)
    return Promise.reject(error)
  }
}
