
import axios from '@/vue-axios'
import header from '../security/header'
import * as types from './types'

export default {
  [types.SECURITY_AUTH] (state, keycloakAuth) {
    state.auth = keycloakAuth

    const id = keycloakAuth.idTokenParsed

    state.email = id.email
    state.firstName = id.given_name
    state.lastName = id.family_name

    // rootState.notebook.noteFilter.emails = [id.email]

    axios.defaults.headers.common = header()
  }
}