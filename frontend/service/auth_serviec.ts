import type { User } from '@/types/user'
import { BASE_API_URL  } from '~/Constants/api/ApiAuth'

// คือฟังก์ชันสำหรับการเข้าสู่ระบบ
export async function loginService(payload: { email: string; password: string }): Promise<{ token: string; user: User }> {
  try {
    const res = await $fetch<{ token: string; user: User }>(`${BASE_API_URL }/login`, {
      method: 'POST',
      body: payload,
    })
    return res
  } catch (err: any) {
    throw new Error(err?.data?.message || 'Login failed')
  }
}

// คือฟังก์ชันสำหรับการสมัครสมาชิก
export async function registerService(payload: 
    {   username: string; 
        email: string;
        password: string 
    }): Promise<{ token: string; user: User }> {
  try {
    const res = await $fetch<{ token: string; user: User }>(`${BASE_API_URL }/register`, {
      method: 'POST',
      body: payload,
    })
    return res
  } catch (err: any) {
    throw new Error(err?.data?.message || 'Register failed')
  }
}

// คือฟังก์ชันสำหรับการดึงข้อมูลของผู้ใช้งาน
export async function fetchProfileService(token: string): Promise<{ user: User }> {
  try {
    const res = await $fetch<{ user: User }>(`${BASE_API_URL }/profile`, {
      headers: { Authorization: `Bearer ${token}` },
    })
    return res
  } catch (err: any) {
    throw new Error(err?.data?.message || 'Fetch profile failed')
  }
}
