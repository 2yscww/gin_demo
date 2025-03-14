import { createApp } from 'vue'
import App from './App.vue'

// 导入 BootstrapVueNext
import { BootstrapVueNext } from 'bootstrap-vue-next' 

// 引入 Pinia
import { createPinia } from 'pinia';  

// 引入 Bootstrap 5 的 CSS
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue-next/dist/bootstrap-vue-next.css'

import router from './router/router.js'  // 导入 router 配置
import request from './utils/request.js' //引入request




const app = createApp(App)

// 创建 Pinia 实例
const pinia = createPinia();  

// 使用 Bootstrap-Vue-Next
app.use(BootstrapVueNext)

// 使用 Vue Router
app.use(router)

// 将 Pinia 注入到 Vue 应用中
app.use(pinia);


app.mount('#app')

// 让 Vue 全局可用 axios（可选）
app.config.globalProperties.$axios = request