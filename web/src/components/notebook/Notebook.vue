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
    <v-layout row justify-center v-if="!$apollo.queries.notes.loading && notes.length == 0 && tryFetchingNotes == true">
      <v-flex class="text-xs-center pt-3">
        <v-divider class="mb-3"/>
        <v-btn @click="isTriggerFirstLoad = true" v-if="!isTriggerFirstLoad"
         large>Load More</v-btn>
        <infinite-loading v-else @infinite="fetchOlderNotes" >
          <span slot="no-more">
            No more notes found.
          </span>
        </infinite-loading>
      </v-flex>
    </v-layout>
    
    
    <!-- <div v-else>
      No more notes found.
    </div> -->
    <!-- end load more -->
  </v-container>
</template>

<script lang='ts'>
import { Component, Emit, Inject, Model, Prop, Provide, Vue, Watch } from 'vue-property-decorator'
import NoteFilter from '@/components/notes/noteFilter'
import notesQuery from './graphql/getNotes.graphql'
import GroupedCollection from '@/components/lists/groupedCollection'
import config from '../../../config'
import { mapActions, mapState, mapGetters } from 'vuex'
import * as moment from 'moment'
import { normalizeDate } from '@/util/dateHelper'

import EditMenu from './menu/EditMenu.vue'
import NoteList from '@/components/lists/List.vue'
import * as InfiniteLoading from 'vue-infinite-loading'
import Gettable from './mixins/gettable'

export default  {
  components: {
    EditMenu,
    NoteList,
    InfiniteLoading
  },
  props: {
    titleFilter: {
      type: String,
      default: ''
    }
  },
  mixins: [
    Gettable
  ],
  data () {
    return {
      supportedNotes: config.notebook.supportedNotes,
      notes: [],
      skipUpdates: false,
      error: null,
      isTriggerFirstLoad: false,
      tryFetchingNotes: true
    }
  },
  watch: {
    titleFilter (val, oldVal) {
      this.updateNoteFilterTitle(val)
    }
  },
  methods: {
    ...mapActions('notebook', [
      'updateNoteFilter',
      'updateNoteFilterTitle',
      'updateNoteFilterAuthors',
    ]),
    sortingFunc: (a: any, b: any) => {
      return moment.utc(b).unix() - moment.utc(a).unix()
    },
    groupingFunc: ({dateCreated}) => {
      return normalizeDate(dateCreated)
    },
  },
  computed: {
    ...mapGetters('notebook', [
      'noteFilter'
    ]),
    noteCollection() {
      return new GroupedCollection(this.notes, this.groupingFunc, this.sortingFunc)
    }
  }
}

</script>

<style>

</style>
