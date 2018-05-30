import Keycloak, { KeycloakInstance } from 'keycloak-js'
import config from '../../config'

export const keycloakAuth: KeycloakInstance = Keycloak(config.keycloak.options)

export function getEmail(keycloakAuth) {
    return keycloakAuth.idTokenParsed.email
}

export function getTeam(keycloakAuth) {
    return keycloakAuth.realm
}