<template>
  <v-list-tile dense ripple :to="{name: name}">
      <v-list-tile-action>
        <v-icon>collections_bookmark </v-icon>
      </v-list-tile-action>
      <v-list-tile-content>
        <span>{{this.name}}</span>
      </v-list-tile-content>
  </v-list-tile>
</template>

<script lang='ts'>
export default {
  props: {
    name: {
      type: String,
      required: true
    }
  }
}
</script>

<style>

</style>
