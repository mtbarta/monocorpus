<template>
  <v-layout>
    <v-container>
      <v-progress-linear v-if="loading==true" v-bind:indeterminate="true"></v-progress-linear>

      <note-list
        :title="query"
        :notes="searchResults"
        :readOnly="true"
      />
      
      <v-card flat v-if="ranQuery && searchResults.length == 0">
        <v-card-text>
          <h2>Sorry.</h2> 
          
          There were no search results.
        </v-card-text>
      </v-card>
      
    </v-container>
  </v-layout>
</template>

<script lang='ts'>
import NoteList from '@/components/notebook/notes/lists/List.vue'
import { Component, Emit, Inject, Model, Prop, Provide, Vue, Watch } from 'vue-property-decorator'
import { searchQuery } from '@/graphql/noteQueries'


export default {
  components: {
    NoteList
  },
  props: {
    query: {
      type: String
    }
  },
  data () {
    return {
      loading: false,
      ranQuery: false,
      searchResults: []
    }
  },
  apollo: {
    $loadingKey: 'loading',
    searchResults: {
      query: searchQuery,
      variables () {
        return {
          query: this.query
        }
      },
      manual: true,
      result (result) {
        this.ranQuery = true
        this.searchResults = result.data.search || []
      }
    }
  }
}
</script>

<style>

</style>
