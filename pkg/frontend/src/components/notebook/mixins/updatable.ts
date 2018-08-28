import Vue from 'vue'
import ComponentOptions from 'vue-apollo'
import updateNote from '@/graphql/updateNote.graphql'
import notesQuery from '@/graphql/getNotes.graphql'
import debounce from 'lodash.debounce'
import { mapGetters } from 'vuex'


//requires this.noteFilter
export default Vue.extend({
    name: 'updatable',

    data () {
      return {
        debounceTime: 500
      }
    },

    created () {
      this.updateNote = debounce(function(note) {
        this._updateNote(note)
      }, this.debounceTime).bind(this)
    },

    methods: {
      _updateNote(note) {
        this.$apollo.mutate({
          mutation: updateNote,
          variables: note,
          update: (store, {data: { updateNote }}) => {
            const updateNoteInData = (data, note) => {
              const index = data.notes.findIndex(i => i.id === note.id)
              if (index !== -1) {
                data.notes[index] = Object.assign(data.notes[index], note)
              }
            }
    
            const data:any = store.readQuery({query: notesQuery, variables: this.noteFilter.getNotebookQuery()})
    
            updateNoteInData(data, updateNote)
            store.writeQuery({query: notesQuery, variables: this.noteFilter.getNotebookQuery(), data})
          },
          optimisticResponse: {
            __typename: 'Mutation',
            updateNote: {
              __typename: 'Note',
              ...note
            }
          }
        })
      }
    },

    computed: {
      ...mapGetters('notebook', [
        'noteFilter'
      ])
    }

})