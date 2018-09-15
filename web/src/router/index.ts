import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home.vue'
import Notebook from '@/components/notebook/Notebook.vue'
import Dashboard from '@/components/dashboard/Dashboard.vue'
import Search from '@/components/search/Search.vue'
import loginGuard from './loginGuard'

import { sync } from 'vuex-router-sync'
import store from '../store'

Vue.use(Router)

const router = new Router({
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home,
      meta: { requiresAuth: true },
      children: [
        {
          path: 'notebook/:titleFilter?',
          name: 'Notebook',
          component: Notebook,
          meta: { requiresAuth: true },
          props: (route) => {
            return route.query
          }
        },
        {
          path: 'search/',
          name: 'Search',
          component: Search,
          meta: { requiresAuth: true },
          props: (route) => {
              return {query: route.query.q}
          }
        },
        {
          path: '',
          name: 'Dashboard',
          component: Dashboard,
          meta: {requiresAuth: true }
        }
      ]
    }
  ],
  mode: 'history',
  linkExactActiveClass: 'active',
  scrollBehavior: function (to, from, savedPosition) {
    return savedPosition || { x: 0, y: 0 }
  }
})

sync (store, router)

router.beforeEach(loginGuard)

export default router