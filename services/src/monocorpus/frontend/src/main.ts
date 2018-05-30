// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import Vuetify from 'vuetify'
import VueApollo from 'vue-apollo'
import AsyncComputed from 'vue-async-computed'

import App from './components/App.vue'
import router from './router'
import store from './store'
import config from '../config'

import GraphClient from '@/graphql/graphClient'
import loginGuard from './router/loginGuard'

import 'vuetify/dist/vuetify.min.css'

Vue.use(Vuetify)
Vue.use(AsyncComputed)
Vue.use(VueApollo)

const apolloProvider = new VueApollo({
  defaultClient: GraphClient
})

Vue.config.productionTip = false

new Vue({
  el: '#app',
  router,
  provide: apolloProvider.provide(),
  store: store,
  components: { App },
  template: '<App/>'
})
