import { createRouter, createWebHistory } from 'vue-router'
import userRoutes from './userRouter'
import { useUserStore } from '@/stores/module/user';



// 导入你需要的页面组件
// import HelloWorld from '@/components/HelloWorld.vue'
// import Register from '@/pages/RegisterPage.vue'
// import Login from '@/pages/LoginPage.vue'

// 创建路由配置
const baseRoutes = [
  { path: '/', name: 'home', component: () => import('@/components/HelloWorld.vue') },

  // { path: '/register', component: Register },
  // { path: '/login', component: Login }
  // { path: '/about', component: About }

];

// 合并路由
const routes = [...baseRoutes, ...userRoutes];

// 创建 Vue Router 实例
const router = createRouter({
  history: createWebHistory(), // 使用 HTML5 历史模式
  routes // 将路由配置传递给 router 实例
})



// 添加路由全局前置守卫

router.beforeEach((to, from, next) => {

  const userStore = useUserStore();

  const token = userStore.token;

  if (to.meta.requireAuth) {
    // 判断是否需要登录

    if(token){
      
      // TODO 还需要判断token是否有效
      next();
      
    } else {
      next('/login');
    }
  } else {
    next();
  }

});


// 导出 router 实例
export default router;