import Vue from 'vue'
import VueRouter from 'vue-router'
import LoginCom from '../views/LoginCom.vue'
import AdminCom from '../views/AdminCom.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'login',
    component: LoginCom
  },
  {
    path: '/admin',
    name: 'admin',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: AdminCom
  }
]

const router = new VueRouter({
  routes
})

export default router
