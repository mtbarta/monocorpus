<template>
  <v-toolbar
      app
      :clipped-left="clipped"
  >
    <v-toolbar-side-icon @click.stop="toggleDrawer()"></v-toolbar-side-icon>
    <v-toolbar-title>{{this.title}}</v-toolbar-title>
    <v-spacer></v-spacer>
    <!-- <v-toolbar-items class="hidden-xs-only">
      <v-btn flat>Link One</v-btn>
      <v-btn flat>Link Two</v-btn>
      <v-btn flat>Link Three</v-btn>
      
    </v-toolbar-items> -->
    <searchbar class="hidden-xs-only"/>
  </v-toolbar>
</template>

<script lang='ts'>
import Vue from 'vue'
import Searchbar from './search/Searchbar.vue'
import * as en  from '@/content/en'

export default Vue.extend({
  components: {
    Searchbar
  },
  props: {
    value: {
      type: Boolean
    },
    clipped: {
      type: Boolean
    }
  },
  data () {
    return {
      title: en.title
    }
  },
  methods: {
    toggleDrawer: function () {
      this.$emit('input', !this.value);
    }
  }
})
</script>

<style>

</style>
