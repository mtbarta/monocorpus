<template>
  <v-container>
    <edit-menu :supportedNotes="supportedNotes"
                :defaultTitle="titleFilter"/>

    <v-progress-linear v-if="$apollo.queries.notes.loading==true" v-bind:indeterminate="true"></v-progress-linear>

    <note-list v-for="collection in noteCollection.groups"
        :title="collection.title"
        :notes="collection.notes"
        :key="collection.title"
    >
    </note-list>

    <!-- load more -->
    <v-layout row justify-center v-if="!$apollo.queries.notes.loading && notes.length == 0">
      <v-flex class="text-xs-center pt-3">
        <v-divider class="mb-3"/>
        <v-btn large
        @click="fetchOlderNotes" 
        >Load More</v-btn>

      </v-flex>
    </v-layout>
    
    <infinite-loading v-else-if="$apollo.queries.notes.loading != true && numCallsAfterEmpty < 3" @infinite="fetchOlderNotes" >
      <span slot="no-more">
        No more notes found.
      </span>
    </infinite-loading>
    <div v-else>
      No more notes found.
    </div>
    <!-- end load more -->
  </v-container>
</template>

<script lang='ts'>
import { Component, Emit, Inject, Model, Prop, Provide, Vue, Watch } from 'vue-property-decorator'
import NoteFilter from '../notebook/notes/noteFilter'
import notesQuery from './graphql/getNotes.graphql'
import GroupedCollection from './notes/lists/groupedCollection'
import config from '../../../config'
import { today } from '@/util/dateHelper'
import { mapActions, mapState, mapGetters } from 'vuex'
import * as moment from 'moment'
import { normalizeDate } from '@/util/dateHelper'

import EditMenu from './menu/EditMenu.vue'
import NoteList from './notes/lists/List.vue'
import InfiniteLoading from 'vue-infinite-loading'

export default  {
  components: {
    EditMenu,
    NoteList,
    InfiniteLoading
  },
  props: {
    titleFilter: String
  },
  data () {
    return {
      supportedNotes: config.notebook.supportedNotes,
      notes: [],
      skipUpdates: false,
      hasMore: true,
      numCallsAfterEmpty: 0, //infinite loading complete() isn't working.
      error: null
    }
  },
  mounted() {
    this.storeInitialFilter(this.noteFilter.copy())
  },
  apollo: {
      notes: {
        query: notesQuery,
        variables () {
          let q = {
            ...this.startingFilter.getNotebookQuery(),
            title: this.titleFilter
          }
          return q
        }
      }
  },
  methods: {
    ...mapActions('notebook', [
      'updateNoteFilter',
      'updateNoteFilterTitle',
      'updateNoteFilterAuthors',
      'storeInitialFilter'
    ]),
    sortingFunc: (a: any, b: any) => {
      // this func is done on the keys, which are formatted date strings.
      return moment.utc(b).unix() - moment.utc(a).unix()
    },
    groupingFunc: ({dateCreated}) => {
      return normalizeDate(dateCreated)
    },
    /**
     * this is how vue-apollo says to fetch more from a query.
     */
    fetchOlderNotes(state) {
      const newFilter = this.noteFilter.fetchOlderNotesQuery(3, 'days')

      this.$apollo.queries.notes.fetchMore({
        variables: newFilter,
        updateQuery: (previousResult, { fetchMoreResult }) => {
          // fetchMoreResult is Object { notes: [...] }
          const newNotes = fetchMoreResult.notes
          newNotes.forEach((note) => {
            note.__typename = 'Note'
          })

          if (newNotes.length == 0) {
            state.complete()
            this.hasMore = false
            this.numCallsAfterEmpty += 1
          } 
          else {
            state.loaded()
            return {
              notes: [
                ...previousResult.notes, 
                ...newNotes
              ]
            }
          }

          
        }
      })
      // this.noteFilter.from = newFilter.from
      this.updateNoteFilter(newFilter)
      // 
    }
  },
  computed: {
    ...mapGetters('notebook', [
      'noteFilter',
      'startingFilter'
    ]),
    noteCollection() {
      return new GroupedCollection(this.notes, this.groupingFunc, this.sortingFunc)
    }
  }
}

</script>

<style>

</style>
