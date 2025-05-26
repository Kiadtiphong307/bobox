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
    import { useAuthForm } from '~/Constants/auth/useAuthForm'
    
  const { username, email, password, onRegister } = useAuthForm()
  </script>
  
  
  