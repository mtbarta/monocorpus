<template>
  <v-navigation-drawer id="nav-bar"
      persistent
      enable-resize-watcher
      fixed
      app
      permanent
      :mini-variant="value"
      :clipped="clipped"
>

    <user-menu />
    
    <v-list>
      <notebook-link/>
      <!-- <v-divider></v-divider> -->
      <note-menu :sidebarClosing="value"/>
    </v-list>

    <v-spacer /> 
    
    <v-list>
      <logout-link/>
    </v-list>

  </v-navigation-drawer>
</template>

<script lang='ts'>
import Vue from 'vue'
import NoteMenu from './NoteMenu.vue'
import UserMenu from './UserMenu.vue'
import NotebookLink from './NotebookLink.vue'
import LogoutLink from './LogoutLink.vue'

export default Vue.extend({
  // type inference enabled
  components: {
    NoteMenu,
    UserMenu,
    NotebookLink,
    LogoutLink
  },
  props: {
    value: {
      type: Boolean
    },
    mini: {
      type: Boolean
    },
    clipped: {
      type: Boolean
    }
  }
})
</script>

<style scoped>
#nav-bar {
  z-index: 2;
}
</style>
