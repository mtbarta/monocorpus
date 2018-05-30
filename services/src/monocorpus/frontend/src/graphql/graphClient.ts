import { ApolloClient } from 'apollo-client'
import { HttpLink } from 'apollo-link-http'
import { InMemoryCache } from 'apollo-cache-inmemory'
import { ApolloLink, concat } from 'apollo-link'
import { setContext } from 'apollo-link-context'
import config from '../../config'
import { keycloakAuth } from '@/keycloak'
import authStore from '@/components/login/store'

const authMiddleware = setContext((_, { headers }) => {
  // const access_token = keycloakAuth.token || null;
  const access_token = authStore.state.auth.token
  
  return {
    headers: {
      ...headers,
      authorization: access_token ? `Bearer ${access_token}` : null
    }
  }
})

const createNewClient = (uri: string, dev: boolean) => {
  const httpLink = new HttpLink({
    uri: config.api.host
  });
  return new ApolloClient({
    link: concat(authMiddleware, httpLink),
    cache: new InMemoryCache(),
    connectToDevTools: dev
  })
}

const isDev = config.NODE_ENV === '"development"'
export default createNewClient(config.api.host, isDev)

// export const getClient = () => {
//   const access_token = localStorage.getItem('token');
//   const headers = {
//     authorization: access_token ? `Bearer ${access_token}` : null
//   };
//   const httpLink = new HttpLink({
//     uri: config.api.host,
//     headers: headers
//   });
//   return new ApolloClient({
//     link: httpLink,
//     cache: new InMemoryCache().restore((<any>window).__APOLLO_STATE__)
//   });
// };