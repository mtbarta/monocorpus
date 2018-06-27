<template>
  <!-- <v-container fluid grid-list> -->
        <v-layout row>
          <v-flex class="">
            <!-- add drop down select for type of note -->
            <!-- add note button -->
            <v-btn-toggle v-model="noteType">
              <v-btn flat value="markdown">
                Markdown
              </v-btn>
               <v-btn flat value="arxiv">
                Arxiv
              </v-btn>
               <v-btn flat value="tex">
                Latex
              </v-btn>
              <v-btn flat value="image">
                Image
              </v-btn>
            </v-btn-toggle>
            <v-btn color="primary" v-on:click="addNote(pseudoNote.prototype())">
                <v-icon>add</v-icon>
                Note
              </v-btn>
          </v-flex>
          
        </v-layout>
      <!-- </v-container> -->
</template>

<script lang='ts'>
import Note from '@/components/notes/note'
import AddNoteMixin from '../mixins/addable'

export default {
  props: {
   supportedNotes: Array,
   defaultTitle: {
     type: String,
     default: 'Untitled'
   }
  },
  mixins: [
    AddNoteMixin
  ],
  watch : {
      defaultTitle () {
        this.pseudoNote.title = this.defaultTitle
      },
      noteType () {
        this.pseudoNote.type = this.noteType
      }
    },
    // TODO (MB): add in author/team to new notes.
  data () {
    return {
      noteType: 'markdown',
      pseudoNote: new Note({
        title: this.defaultTitle || 'Untitled',
        type: 'markdown'
      })
    }
  }
}
</script>

<style>

</style>
