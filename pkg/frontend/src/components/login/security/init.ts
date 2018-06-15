import { keycloakAuth } from '@/keycloak'
import GraphClient from '@/graphql/graphClient'
import authHeader from './header'
import store from '@/store'

export default (next, roles) => {
  keycloakAuth.init({ onLoad: 'login-required' })
    .success(() => {
      
      keycloakAuth.updateToken(10)
        .success(() => {
          store.dispatch('authLogin', keycloakAuth)
          // mutation(state, keycloakAuth)
          if (roles) {
            if (keycloakAuth.hasRealmRole(roles[0])) {
              next()
            } else {
              next({ name: 'Unauthorized' })
            }
          } else {
            next()
          }
        })
    })
    .error(() => {
      console.log('failed to login')
    })
  }