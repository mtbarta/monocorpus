import Vue from 'vue'
import ComponentOptions from 'vue-apollo'
import notesQuery from '@/graphql/getNotes.graphql'
import * as moment from 'moment'

export default Vue.extend({
    name: 'gettable',

    data () {
      return {
        hasMore: true,
        page: 1,
        prevAmount: 3,
        prevUnit: 'days'
      }
    },


    apollo: {
        notes: {
          query: notesQuery,
          variables () {
            return this.noteFilter.getNotebookQuery()
          }
        }
    },

    methods: {
      fetchOlderNotes(state) {
        let to = this.noteFilter.from
        let from = moment.utc(this.noteFilter.from)
                         .subtract(this.prevAmount * this.page, this.prevUnit)
                         .unix()

        this.$apollo.queries.notes.fetchMore({
          variables: {
            ...this.noteFilter.getNotebookQuery(),
            to,
            from
          },
          updateQuery: (previousResult, { fetchMoreResult }) => {
            // fetchMoreResult is Object { notes: [...] }
            const newNotes = fetchMoreResult.notes
            if (!newNotes) {
              state.complete()
              return;
            }
            newNotes.forEach((note) => {
              note.__typename = 'Note'
            })

            if (newNotes.length == 0) {
              state.complete()
              this.hasMore = false
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
            this.page++
          }
        })
      }
    }
})