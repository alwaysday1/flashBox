<template>
  <header>闪念盒子</header>
  <div>
    <el-button type="primary">注册</el-button>
    <el-button type="primary">登录</el-button>
  </div>
  <div>
    <el-input v-model="idea" placeholder="记录此刻的想法"></el-input>
    <el-button type="primary">发送</el-button>
  </div>
</template>

<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import HelloWorld from './components/HelloWorld.vue'
import axios from 'axios'
import { ref,onMounted } from "vue";
const idea = ref('')
const service = axios.create({
  baseURL: '/api',
  timeout: 5000
})

onMounted(() => {
  service.get('/ping').then((res) => {
    console.log(res?.data)
    // el.value = res?.data
  })
  service.get('/').then((res) => {
    console.log(res?.data)
  })
  service.post('/user/add', {
    name: 'test',
    phone: '18'
  }).then((res) => {
    console.log(res?.data)
  })
  service.get('/user/list/test').then((res) => {
    console.log(res?.data)
    // name.value = res?.data.data
  })
})
</script>

<style scoped lang="less"></style>
