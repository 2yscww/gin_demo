

const userRoutes = [
    { path: '/register', name: 'register', component: () => import('@/pages/RegisterPage.vue') },
    { path: '/login', name: 'login', component: () => import('@/pages/LoginPage.vue') },
    {
        path: '/user/profile',
        name: 'profile',
        component: () => import('@/pages/profileInfo.vue'),
        meta: { requireAuth: true }
        // 添加路由守卫
    }


];

export default userRoutes;


