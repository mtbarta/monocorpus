import store from '../store'
import security from '@/components/login/security'

export default function loginGuard (to, from, next) {
    if (to.meta.requiresAuth) {
      const auth = store.state.login.auth
      if (!auth.authenticated) {
        security.init(next, to.meta.roles)
      } else {
        if (to.meta.roles) {
          if (security.roles(to.meta.roles[0])) {
            next()
          } else {
            next({ name: 'Unauthorized' })
          }
        } else {
          next()
        }
      }
    } else {
      next()
    }
  }