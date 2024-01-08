// 5 在main.js中导入并挂载路由模块
import { createApp } from 'vue'
import Home from './Home.vue'
 
// 导入路由模块
import router from './router'
 
const app = createApp(Home)
 
// 挂载路由模块
app.use(router)
 
app.mount('#app')