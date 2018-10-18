import { keycloakAuth } from '../keycloak'
/**
 * before fetching a page through graphql,
 * check whether the token is expired and 
 * optionally update it.
 * 
 * options should be a RequestInit, i think.
 */
export default async (uri?: string | Request, options?: any): Promise<any> => {
  let isExpired = keycloakAuth.isTokenExpired(30)

  if (isExpired) {
    await keycloakAuth.updateToken(30)

    let newAccessToken = keycloakAuth.token
    options.headers.authorization = `Bearer ${newAccessToken}`
  }
  
  return fetch(uri, options)
}