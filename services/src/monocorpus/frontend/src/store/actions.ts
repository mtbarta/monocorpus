export default {
    setUser: ({commit}, user) => {
      if (window.localStorage) {
        window.localStorage.setItem('user', JSON.stringify(user))
      }
      commit('SET_USER', user)
    },
    setToken: ({commit}, {token}) => {
      commit ('SET_TOKEN', token)
    }
}

