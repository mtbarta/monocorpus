<template>
  <v-list-tile color="grey" dense ripple href @click="logout">
      <v-list-tile-action>
        <v-icon>exit_to_app </v-icon>
      </v-list-tile-action>
      <v-list-tile-content>
       
        <!-- <span class="body-1 logout">Log out</span> -->
        Log out
        <!-- </v-btn> -->
      </v-list-tile-content>
  </v-list-tile>
</template>

<script lang='ts'>
import security from '@/components/login/security'

export default {
  methods: {
    logout () {
      security.logout()
    }
  }
}
</script>

<style>
.logout {
  text-decoration-color: grey;
}
</style>
