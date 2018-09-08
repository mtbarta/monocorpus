<template>
  <!-- <v-text-field
          name="search"
          hide-details single-line solo
          append-icon="search"
          :append-icon-cb="search"
          placeholder="search"
          v-model="searchInput"
        >
  </v-text-field> -->
  <v-layout row align-center style="max-width: 450px">
        <v-text-field
          v-model="searchInput"
          placeholder="Search..."
          single-line
          append-icon="search"
          :append-icon-cb="search"
          hide-details
          @keyup.enter="search"
        ></v-text-field>
</v-layout>
</template>

<script lang='ts'>
export default {
  name: 'search-bar',
  data () {
    return {
      searchInput: ''
    }
  },
  methods: {
    search () {
      this.$router.push({
        name: 'Search', 
        query: {
          q: this.searchInput
        }
      })
    }
  }
}
</script>

<style>

</style>
