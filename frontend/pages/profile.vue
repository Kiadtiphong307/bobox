<!-- pages/profile.vue -->
<template>
    <div class="max-w-md mx-auto p-4">
      <h2 class="text-2xl font-bold mb-4">โปรไฟล์ผู้ใช้</h2>
      <div v-if="user">
        <p><strong>Username:</strong> {{ user.username }}</p>
        <p><strong>Email:</strong> {{ user.email }}</p>
        <button class="btn btn-danger mt-4" @click="onLogout">ออกจากระบบ</button>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  definePageMeta({
    middleware: 'auth',
  })
  
  const { user, fetchProfile, logout } = useAuth()
  
  onMounted(() => {
    fetchProfile()
  })
  
  function onLogout() {
    logout()
    navigateTo('/login')
  }
  </script>
  