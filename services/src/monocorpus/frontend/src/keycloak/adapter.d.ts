import Vue from 'vue'
import { KeyCloakPlugin } from './adapter'

declare module 'vue/types/vue' {
  interface Vue {
    $keycloak: KeyCloakPlugin
  }
}