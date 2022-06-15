import Vue from 'vue'
import VueRouter from 'vue-router'
import LoginCom from '../views/LoginCom.vue'
import AdminCom from '../views/AdminCom.vue'

// page router components
import IndexCom from '../components/admin/IndexCom.vue'
import AddArt from '../components/article/AddArt.vue'
import ArtList from '../components/article/ArtList.vue'
import CateList from '../components/category/CateList.vue'
import UserList from '../components/user/UserList.vue'

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
    component: AdminCom,
    children: [
      { path: '/index', component: IndexCom },
      { path: '/addart', component: AddArt },
      { path: '/artlist', component: ArtList },
      { path: '/catelist', component: CateList },
      { path: '/userlist', component: UserList }
    ]
  }
]

const router = new VueRouter({
  routes
})

router.beforeEach((to, from, next) => {
  const token = window.sessionStorage.getItem('token')
  if (to.path === '/login') return next()
  if (!token && to.path === '/admin') {
    next('/login')
  } else {
    next()
  }
})

export default router
