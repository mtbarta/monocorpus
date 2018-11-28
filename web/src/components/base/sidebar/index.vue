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
    
    <v-list-tile dense ripple 
      v-for="item in pages"
      :key="item.name"
      v-model="item.active"
      :to="item.link">
      <v-list-tile-action>
        <v-icon>{{item.icon}} </v-icon>
      </v-list-tile-action>
      <v-list-tile-content>
        <span>{{item.name}}</span>
      </v-list-tile-content>
  </v-list-tile>

    <note-menu />

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
import LogoutLink from './LogoutLink.vue'

export default Vue.extend({
  // type inference enabled
  components: {
    NoteMenu,
    UserMenu,
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
  },

  data() {
    return {
      pages: [
                {
                    name: 'Dashboard',
                    link: '/',
                    icon: 'collections_bookmark'
                },
                {
                    name: 'Notebook',
                    link: '/notebook',
                    icon: 'collections_bookmark'
                }
            ]
    }
  }
})
</script>

<style scoped>
#nav-bar {
  z-index: 2;
}
</style>