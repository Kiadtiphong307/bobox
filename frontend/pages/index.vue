<template>
  <div class="p-6 flex flex-col items-center justify-center h-screen">
    <h1 class="text-2xl font-bold mb-4">หน้าแรก</h1>
    <NuxtLink to="/products" class="text-blue-500 hover:underline"> ดูสินค้า </NuxtLink>
    <NuxtLink to="/login" class="text-blue-500 hover:underline"> เข้าสู่ระบบ </NuxtLink>
    <NuxtLink to="/register" class="text-blue-500 hover:underline"> สมัครสมาชิก </NuxtLink>
    <NuxtLink to="/profile" class="text-blue-500 hover:underline"> โปรไฟล์ </NuxtLink>
  </div>
</template>