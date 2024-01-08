// 5 在main.js中导入并挂载路由模块
import { createApp } from 'vue'
import App from './App.vue'
import Antd  from 'ant-design-vue'
import 'ant-design-vue/dist/reset.css';
import * as AntdIcons from '@ant-design/icons-vue';
 
// 导入路由模块
import router from './router'
 
const app = createApp(App)

for (const key in AntdIcons) {
  app.component(key, AntdIcons[key])
}
 
app.use(Antd)
// 挂载路由模块
app.use(router)
 
app.mount('#app')