import store from '../store'

export default () => {
  var keycloakAuth = store.getters.SECURITY_AUTH
  return { 'Authorization': 'Bearer ' + keycloakAuth.token }
}