import { createApp } from 'vue'
import App from './App.vue'


// 导入 BootstrapVueNext
import { BootstrapVueNext } from 'bootstrap-vue-next' 

// 引入 Bootstrap 5 的 CSS
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue-next/dist/bootstrap-vue-next.css'

import router from './router'  // 导入 router 配置





const app = createApp(App)

// 使用 Bootstrap-Vue-Next
app.use(BootstrapVueNext)

// 使用 Vue Router
app.use(router)

app.mount('#app')