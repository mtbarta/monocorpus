<template>
  <v-container class="note-wrapper">

    <markdown-note :note="note" v-if="type=='markdown'" 
      :updateNote="update" :readOnly="readOnly" :deleteNote="del"/>
    <tex-note :note="note" v-if="type=='tex'" 
      :updateNote="update" :readOnly="readOnly" :deleteNote="del"/>
    <arxiv-note :note="note" v-if="type=='arxiv'" 
      :updateNote="update" :readOnly="readOnly" :deleteNote="del"/>
    <image-note :note="note" v-if="type=='image'" 
      :updateNote="update" :readOnly="readOnly" :deleteNote="del"/>
  </v-container>
</template>

<script lang="ts">
import debounce from 'lodash.debounce'
import MarkdownNote from './types/MarkdownNote.vue'
import TexNote from './types/TexNote.vue'
import ArxivNote from './types/ArxivNote.vue'
import ImageNote from './types/ImageNote.vue'
// import {
//   updateNote,
//   addNote,
//   deleteNote
// } from '@/graphql/noteQueries'
import { mapActions } from 'vuex'

const THROTTLE_MS = 500

export default {
  name: 'NoteWrapper',
  components: {
    MarkdownNote,
    TexNote,
    ArxivNote,
    ImageNote
  },
  props: {
    note: Object,
    type: String,
    readOnly: {
      type: Boolean,
      default: false
    }
  },
  data () {
    return {
      THROTTLE_MS: THROTTLE_MS
    }
  },
  methods: {
    ...mapActions('notebook', [
      'updateNote',
      'deleteNote'
    ]),
    del (id) {
      this.deleteNote(id);
    },
    update: debounce(function(note) {
      this.updateNote(note)
    }, THROTTLE_MS)
  },
}
</script>

<style>
ul {
  padding-left: .5rem;
}
li {
  padding-left: 1rem;
}
p {
  margin-bottom: 0px !important;
}
</style>
