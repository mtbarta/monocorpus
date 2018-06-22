import NoteFilter from '../notes/noteFilter'
import graphClient from '@/graphql/graphClient'
import Note from '../notes/note'
import ADD_NOTE_MUTATION from '../graphql/addNote.graphql'
import getNotes from '../graphql/getNotes.graphql'
import updateNote from '../graphql/updateNote.graphql'
import deleteNote from '../graphql/deleteNote.graphql'

export default {
  UPDATE_NOTE_FILTER (state, noteFilter: NoteFilter) {
    state.noteFilter = Object.assign(state.noteFilter, noteFilter)
  },
  UPDATE_NOTE_FILTER_TITLE (state, args: any) {
    state.noteFilter.title = args.title
  },
  UPDATE_NOTE_FILTER_AUTHORS(state: any, emails: string[]) {
    state.noteFilter.authors = emails
  },
  UPDATE_NOTE_FILTER_TEAM(state:any, team: string) {
    state.noteFilter.team = team
  },
  STORE_INITIAL_FILTER(state:any, query:any) {
    state.startingFilter = query
  },
  UPDATE_NOTEBOOK_TITLE_FILTER(state:any, title: string) {
    state.startingFilter.title = title
  },
  ADD_NOTE(state: any, note:Note) {
    graphClient.mutate({
      mutation: ADD_NOTE_MUTATION,
      variables: note,
      update: (store, {data: { createNote }}) => {
        const data:any = store.readQuery({query: getNotes, variables: state.startingFilter.getNotebookQuery()})

        data.notes.unshift(createNote)
        store.writeQuery({query: getNotes, variables: state.startingFilter.getNotebookQuery(), data})
      },
      optimisticResponse: {
          __typename: 'Mutation',
          createNote: {
            __typename: 'Note',
            ...note
          }
      }
    })
  },
  UPDATE_NOTE(state: any, note:Note) {
    graphClient.mutate({
      mutation: updateNote,
      variables: note,
      update: (store, {data: { updateNote }}) => {
        const updateNoteInData = (data, note) => {
          const index = data.notes.findIndex(i => i.id === note.id)
          if (index !== -1) {
            data.notes[index] = Object.assign(data.notes[index], note)
          }
        }

        const data:any = store.readQuery({query: getNotes, variables: state.startingFilter.getNotebookQuery()})

        updateNoteInData(data, updateNote)
        store.writeQuery({query: getNotes, variables: state.startingFilter.getNotebookQuery(), data})
      },
      optimisticResponse: {
          __typename: 'Mutation',
          updateNote: {
            __typename: 'Note',
            ...note
          }
      }
    })
  },
  DELETE_NOTE(state: any, id:string) {
    graphClient.mutate({
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

        const data:any = store.readQuery({query: getNotes, variables: state.startingFilter.getNotebookQuery()})

        deleteNoteInData(data, updateNote)
        store.writeQuery({query: getNotes, variables: state.startingFilter.getNotebookQuery(), data})
      },
      // optimisticResponse: {
      //     __typename: 'Mutation',
      //     updateNote: {
      //       __typename: 'Note',
      //       ...note
      //     }
      // }
    })
  }
}