import * as types from './types'
import { getEmail, getTeam } from '@/keycloak'

export default {
  authLogin ({ state, commit, rootState }, keycloakAuth) {
    commit(types.SECURITY_AUTH, keycloakAuth)
    const email = getEmail(keycloakAuth)
    const team = getTeam(keycloakAuth)

    localStorage.setItem("token", keycloakAuth.token)
    // console.log(JSON.stringify(keycloakAuth.idTokenParsed))

    commit("notebook/UPDATE_NOTE_FILTER_AUTHORS", [email])
    commit("notebook/UPDATE_NOTE_FILTER_TEAM", team)
  },
  authLogout ({ commit }) {
    commit(types.SECURITY_AUTH)
  }
}