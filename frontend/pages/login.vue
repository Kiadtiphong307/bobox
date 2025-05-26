<!-- pages/login.vue -->
<template>
    <div class="max-w-md mx-auto p-4">
      <h2 class="text-2xl font-bold mb-4">เข้าสู่ระบบ</h2>
      <form @submit.prevent="onLogin">
        <input v-model="email" placeholder="Email" class="input mb-2 w-full" />
        <input v-model="password" type="password" placeholder="Password" class="input mb-4 w-full" />
        <button type="submit" class="btn btn-primary w-full">Login</button>
      </form>
    </div>
  </template>
  
  <script setup lang="ts">
  const email = ref('')
  const password = ref('')
  const router = useRouter()
  const { login } = useAuth()
  
  async function onLogin() {
    try {
      await login({ email: email.value, password: password.value })
      router.push('/profile')
    } catch (e) {
      alert('Login Failed')
    }
  }
  </script>
  