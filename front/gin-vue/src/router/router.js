import { createRouter, createWebHistory } from 'vue-router'

// 导入你需要的页面组件
import HelloWorld from '@/components/HelloWorld.vue'
// import Register from '@/pages/RegisterPage.vue'
// import Login from '@/pages/LoginPage.vue'

// 创建路由配置
const routes = [
  { path: '/', name: 'home', component: HelloWorld },
  { path: '/register', name: 'register', component: () => import('@/pages/RegisterPage.vue') },
  { path: '/login', name: 'login', component: () => import('@/pages/LoginPage.vue') }
  // { path: '/register', component: Register },
  // { path: '/login', component: Login }
  // { path: '/about', component: About }

]

// 创建 Vue Router 实例
const router = createRouter({
  history: createWebHistory(), // 使用 HTML5 历史模式
  routes // 将路由配置传递给 router 实例
})


// 导出 router 实例
export default router