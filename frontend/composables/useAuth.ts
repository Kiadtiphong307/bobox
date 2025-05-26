import { useCookie, useState } from '#app'
import type { User } from '@/types/user'
import { loginService, registerService, fetchProfileService } from '@/service/auth_serviec'

export function useAuth() {
  const token = useCookie<string | null>('token')
  const user = useState<User | null>('user', () => null)

  // คือฟังก์ชันสำหรับการเข้าสู่ระบบ
  async function login(payload: { email: string; password: string }) {
    const res = await loginService(payload)
    token.value = res.token
    user.value = res.user
  }
// คือฟังก์ชันสำหรับการสมัครสมาชิก
  async function register(payload: { username: string; email: string; password: string }) {
    const res = await registerService(payload)
    token.value = res.token
    user.value = res.user
  }
// คือฟังก์ชันสำหรับการดึงข้อมูลของผู้ใช้งาน
  async function fetchProfile() {
    if (!token.value) return
    const res = await fetchProfileService(token.value)
    user.value = res.user
  }
// คือฟังก์ชันสำหรับการออกจากระบบ
  function logout() {
    token.value = null
    user.value = null
  }

  return { user, token, login, register, fetchProfile, logout }
}
