import * as types from './types'
import {KeycloakInstance} from 'keycloak-js'
import userData from '../userData'

// export default class Getters {
//     [types.SECURITY_AUTH]: Function = (state:any): KeycloakInstance => {
//         return state.auth
//     }
// }

export default {
  [types.SECURITY_AUTH](state) {
    return state.auth
  }
}