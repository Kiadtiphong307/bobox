// composables/useAuthForm.ts
export function useAuthForm() {
    const username = ref('')
    const email = ref('')
    const password = ref('')
  
    const router = useRouter()
    const { login, register } = useAuth()
  
    // คือฟังก์ชันสำหรับการสมัครสมาชิก
    async function onRegister() {
      try {
        await register({
          username: username.value,
          email: email.value,
          password: password.value,
        })
        router.push('/profile')
      } catch (e) {
        alert('Register Failed')
      }
    }
  
    // คือฟังก์ชันสำหรับการเข้าสู่ระบบ
    async function onLogin() {
      try {
        await login({ email: email.value, password: password.value })
        router.push('/profile')
      } catch (e) {
        alert('Login Failed')
      }
    }
  
    return {
      username,
      email,
      password,
      onRegister,
      onLogin,
    }
  }
  