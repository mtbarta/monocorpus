<template>
  <nav-bar>

    <user-menu />
    
    <path-link
      name="Dashboard"
    />
    <path-link
      name="Notebook"
    />

    <note-menu />

    <v-spacer /> 
    
    <v-list>
      <logout-link/>
    </v-list>

  </nav-bar>
</template>

<script lang='ts'>
import Vue from 'vue'
import Navbar from '../Navbar.vue'
import NoteMenu from './NoteMenu.vue'
import UserMenu from './UserMenu.vue'
import PathLink from '../PathLink.vue'
import LogoutLink from '../LogoutLink.vue'

export default Vue.extend({
  // type inference enabled
  components: {
    Navbar,
    NoteMenu,
    UserMenu,
    PathLink,
    LogoutLink
  }
})
</script>

