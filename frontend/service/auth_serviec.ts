import type { User } from '@/types/user'

const baseURL = '/api/auth'

export async function loginService(payload: { email: string; password: string }): Promise<{ token: string; user: User }> {
  try {
    const res = await $fetch<{ token: string; user: User }>(`${baseURL}/login`, {
      method: 'POST',
      body: payload,
    })
    return res
  } catch (err: any) {
    throw new Error(err?.data?.message || 'Login failed')
  }
}

export async function registerService(payload: { username: string; email: string; password: string }): Promise<{ token: string; user: User }> {
  try {
    const res = await $fetch<{ token: string; user: User }>(`${baseURL}/register`, {
      method: 'POST',
      body: payload,
    })
    return res
  } catch (err: any) {
    throw new Error(err?.data?.message || 'Register failed')
  }
}

export async function fetchProfileService(token: string): Promise<{ user: User }> {
  try {
    const res = await $fetch<{ user: User }>('/api/user/profile', {
      headers: { Authorization: `Bearer ${token}` },
    })
    return res
  } catch (err: any) {
    throw new Error(err?.data?.message || 'Fetch profile failed')
  }
}
