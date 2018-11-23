<template>
  <v-container class="note-wrapper">

    <markdown-note :note="note" v-if="type=='markdown'" 
      :updateNote="updateNote" :readOnly="readOnly" :deleteNote="deleteNote"/>
    <tex-note :note="note" v-if="type=='tex'" 
      :updateNote="updateNote" :readOnly="readOnly" :deleteNote="deleteNote"/>
    <arxiv-note :note="note" v-if="type=='arxiv'" 
      :updateNote="updateNote" :readOnly="readOnly" :deleteNote="deleteNote"/>
    <image-note :note="note" v-if="type=='image'" 
      :updateNote="updateNote" :readOnly="readOnly" :deleteNote="deleteNote"/>
  </v-container>
</template>

<script lang="ts">
import debounce from 'lodash.debounce'
import MarkdownNote from './MarkdownNote.vue'
import TexNote from './TexNote.vue'
import ArxivNote from './ArxivNote.vue'
import ImageNote from './ImageNote.vue'
import { mapActions } from 'vuex'
import DeleteMixin from '@/components/notebook/mixins/deletable'
import UpdateMixin from '@/components/notebook/mixins/updatable'

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
  mixins: [
    DeleteMixin,
    UpdateMixin
  ]
}
</script>

<style scoped>
ul {
  padding-left: .5rem;
}
li {
  padding-left: 1rem;
}

</style>
