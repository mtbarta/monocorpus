import store from '@/store'
import { keycloakAuth } from '@/keycloak'

export default () => {
  // var keycloakAuth = store.getters.SECURITY_AUTH
  localStorage.setItem("token", null)
  store.getters.SECURITY_AUTH.logout()
  store.dispatch('authLogout')
  
}