import _Vue from 'vue'
import {default as keycloakAdapter, KeycloakInstance} from 'keycloak-js'

interface KeyCloakOptions {
    keycloakOptions: Object
    keycloakInitOptions: Object
    refreshTime: number
}

let installed = false

// export default class KeyCloak {
    
//     install 
// }

export function KeyCloakPlugin<KeyCloakOptions>(Vue: typeof _Vue, options?: any): void {
    if (installed) {
        return
    }
    installed = true

    const keycloak: KeycloakInstance = keycloakAdapter(options.keycloakOptions)

    const watch: any = new Vue({
        data () {
            return {
                user: '',
                adapter: keycloak
            }
        }
    })
    keycloak.init(options.keycloakInitOptions)
            .success((isAuthenticated: boolean) => {

                if (isAuthenticated) {
                    watch.user = keycloak.idTokenParsed

                    setTimeout(() => {
                      keycloak.updateToken(options.refreshTime + 2)
                    }, options.refreshTime * 1000)
                }
            })
    

    Vue.prototype.$keycloak = watch
    // Vue.prototype.$user = watch.user
}

export default KeyCloakPlugin