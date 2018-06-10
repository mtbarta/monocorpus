<template>
  <v-app id="app">
    <router-view/>
  </v-app>
</template>

<script lang='ts'>
import Toolbar from './base/header.vue'
import Sidebar from './base/sidebar/index.vue'

export default {
  name: 'App',
  components: {
    Toolbar,
    Sidebar
  },
  data () {
    return {
      clipped: true,
      drawer: false,
      fixed: false,
      miniVariant: false,
    }
  },
  methods: {
      logout () {
        this.$store.commit('SET_USER', null)
        this.$store.commit('SET_TOKEN', null)

        if (window.localStorage) {
          window.localStorage.setItem('user', null)
          window.localStorage.setItem('token', null)
        }

        this.$router.push('/login')
      }
    }
}
</script>

<style>
/* @import '../node_modules/vuetify/dist/vuetify.min.css'; */
/* #app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
} */
</style>
