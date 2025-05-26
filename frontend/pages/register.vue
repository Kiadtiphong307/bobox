<!-- pages/register.vue -->
<template>
    <div class="max-w-md mx-auto p-4">
      <h2 class="text-2xl font-bold mb-4">สมัครสมาชิก</h2>
      <form @submit.prevent="onRegister">
        <input v-model="username" placeholder="Username" class="input mb-2 w-full" />
        <input v-model="email" placeholder="Email" class="input mb-2 w-full" />
        <input v-model="password" type="password" placeholder="Password" class="input mb-4 w-full" />
        <button type="submit" class="btn btn-primary w-full">Register</button>
      </form>
    </div>
  </template>
  
  <script setup lang="ts">
  const username = ref('')
  const email = ref('')
  const password = ref('')
  const router = useRouter()
  const { register } = useAuth()
  
  async function onRegister() {
    try {
      await register({ username: username.value, email: email.value, password: password.value })
      router.push('/profile')
    } catch (e) {
      alert('Register Failed')
    }
  }
  </script>
  