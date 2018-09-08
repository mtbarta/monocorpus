import Vue from 'vue'
import ComponentOptions from 'vue-apollo'
import deleteNote from '@/graphql/deleteNote.graphql'
import notesQuery from '@/graphql/getNotes.graphql'
import { mapGetters } from 'vuex'


//requires this.noteFilter
export default Vue.extend({
    name: 'deletable',

    methods: {
      deleteNote(id) {
        this.$apollo.mutate({
          mutation: deleteNote,
          variables: {
           id 
          },
          update: (store, {data: { deleteNote }}) => {
            const deleteNoteInData = (data, note) => {
              const index = data.notes.findIndex(i => i.id === note.id)
              if (index !== -1) {
                data.notes.splice(index, 1)
              }
            }
    
            const data:any = store.readQuery({query: notesQuery, variables: this.noteFilter.getNotebookQuery()})
    
            deleteNoteInData(data, deleteNote)
            store.writeQuery({query: notesQuery, variables: this.noteFilter.getNotebookQuery(), data})
          }
        })
      }
    },

    computed: {
      ...mapGetters('notebook', [
        'noteFilter'
      ]),
    }
})