import Vue from 'vue'
import Vuex from 'vuex'
import state from './state'
import actions from './actions'
import mutations from './mutations'
import getters from './getters'

import notebook from '@/components/notebook/store'
import login from '@/components/login/store'

Vue.use(Vuex)

export default new Vuex.Store({
  state,
  actions,
  mutations,
  getters,
  modules: {
    notebook,
    login
  }
})
