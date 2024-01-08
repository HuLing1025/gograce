// 1 从vue-router按需导入两个方法
// createRouter方法用于创建路由实例对象
// createWebHashHistory方法用于指定路由的工作模式（hash模式）
import { createRouter, createWebHashHistory } from 'vue-router'
// 2 导入需要使用路由控制的组件
import Home from './Home.vue'

// 3 创建路由对象
const router = createRouter({
  // 3.1 通过 history 属性，指定路由的工作模式
  history: createWebHashHistory(),
  // 3.2 通过 routes 数组，指定路由规则
  // path 是 hash 地址，component 是要展示的组件
  routes: [
    { path: '/', component: Home },
  ],
})

    // 4、向外共享路由对象
export default router