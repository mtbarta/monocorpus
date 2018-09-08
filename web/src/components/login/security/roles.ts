import store from '../store'

export default (role) => {
  const keycloakAuth = store.getters.SECURITY_AUTH
  if (keycloakAuth.authenticated) {
    return keycloakAuth.hasRealmRole(role)
  } else {
    return false
  }
}