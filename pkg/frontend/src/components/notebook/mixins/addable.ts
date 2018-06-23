import Vue from 'vue'
import ComponentOptions from 'vue-apollo'
import ADD_NOTE_MUTATION from '@/graphql/addNote.graphql'
import notesQuery from '@/graphql/getNotes.graphql'
import { mapGetters, mapState } from 'vuex'

//requires this.noteFilter
export default Vue.extend({
    name: 'addable',

    methods: {
      addNote(note) {
        this.$apollo.mutate({
          mutation: ADD_NOTE_MUTATION,
          variables: {
            ...note,
            author: this.email
          },
          update: (store, {data: { createNote }}) => {
            const data:any = store.readQuery({query: notesQuery, variables: this.noteFilter.getNotebookQuery()})
    
            data.notes.unshift(createNote)
            store.writeQuery({query: notesQuery, variables: this.noteFilter.getNotebookQuery(), data})
          },
          optimisticResponse: {
            __typename: 'Mutation',
            createNote: {
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
      ]),
      ...mapState({
        email(state: any) {
          return state.login.email
        }
      })
    }
})