import gql from 'graphql-tag'

import { normalizeDate } from '../util/dateHelper'
import GroupedCollection from '../components/notebook/notes/lists/groupedCollection'
import Note from '../components/notebook/notes/note'
import * as moment from 'moment'
import store from '../store'

// let lastGetVars = store.state.notebook.startingFilter
let lastGetVars = store.state.notebook.startingFilter

export const notesQuery = gql`query GetNotes($to: Float, $from:Float, $authors: [String], $title: String, $ids: [ID]) {
  notes(todate: $to, fromdate: $from, authors: $authors, title: $title, ids:$ids) {
    id,
    title,
    body,
    dateCreated,
    author,
    team,
    link,
    type
  }
}`

export const notesTitleQuery = gql`query GetNotes($to: Float, $from:Float, $authors: [String], $title: String, $ids: [ID]) {
  notes(todate: $to, fromdate: $from, authors: $authors, title: $title, ids:$ids) {
    title
  }
}`

export const searchQuery = gql`query GetNotes($query: String, $to: Float, $from:Float, $authors: [String], $team: String) {
  search(query: $query, todate: $to, fromdate: $from, authors: $authors, team: $team) {
    id,
    title,
    body,
    dateCreated,
    author,
    team,
    type
  }
}`

/**
 * retrieves notes. this is the foundation of the notebook view.
 */
// export const getNotesQuery = () => {
//   return {
//     query: notesQuery,
//     variables () {
//       return this.noteFilter
//     },
//     manual: true,
//     result (result) {
//       lastGetVars = Object.assign({}, this.noteFilter)

//       const sortingFunc: (a: any, b: any) => number = (a: any, b: any) => {
//         // this func is done on the keys, which are formatted date strings.
//         return moment.utc(b).unix() - moment.utc(a).unix()
//       }

//       const groupingFunc: Function = ({dateCreated}) => {
//         return normalizeDate(dateCreated)
//       }

//       this.notes = new GroupedCollection(result.data.notes, groupingFunc, sortingFunc)
//     }
//   }
// }

export const getNoteTitles = function () {
  return {
    query: notesTitleQuery,
    variables () {
      return this.noteFilter
    },
    manual: true,
    result (result) {
      this.noteTitles = Array.from(new Set(result.data.notes.map(note => note.title)))
    }
  }
}

export const updateNote = (note) => {
  return {
    mutation: gql`mutation UpdateNotes($id: ID, $title:String, $body:String, $author:String, $team:String, $dateCreated:String, $dateModified:String, $type:String, $link:String){
      updateNote(id: $id, title: $title, body:$body, author:$author, team:$team, dateCreated:$dateCreated, dateModified:$dateModified, type:$type, link:$link) {
        id,
        title,
        body,
        author,
        team,
        type,
        link,
        dateCreated,
        dateModified
      }
    }`,
    variables: note,
    // update: (store, {data: updateNote}) => {
    //   const data = store.readQuery({query: notesQuery, variables: lastGetVars.getNotebookQuery()})
    //   updateNoteInData(data, updateNote)
    //   store.writeQuery({query: notesQuery, variables: lastGetVars.getNotebookQuery(), data})
    // },
    refetchQueries: ['GetNotes'],
    optimisticResponse: {
      __typename: 'Mutation',
      updateNote: {
        __typename: 'Note',
        ...note
      }
    }
  }
}

const updateNoteInData = (data, note) => {
  const index = data.notes.findIndex(i => i.id === note.id)
  if (index !== -1) {
    data.notes[index] = Object.assign(data.notes[index], note)
  }
}

export const addNote = (note: Note) => {
  return {
    mutation: gql`mutation AddNote($id: ID, $title:String, $body:String, $author:String, $team:String, $dateCreated:String, $dateModified:String, $type:String, $link:String){
      createNote(id: $id, title: $title, body:$body, author:$author, team:$team, dateCreated:$dateCreated, dateModified:$dateModified, type:$type, link:$link) {
        id,
        title,
        body,
        author,
        team,
        dateCreated,
        dateModified,
        type,
        link
      }
    }`,
    variables: note,
    update: (store, {data: { createNote }}) => {
      console.log(note)
      console.log(store)
      console.log(JSON.stringify(lastGetVars))
      // const variables = lastGetVars.getNotebookQuery()
      // variables.authors = [note.author]
      const query = Object.keys(store.data.data.ROOT_QUERY)[0]
      console.log(query)
      const data = store.readQuery({query: notesQuery, variables: lastGetVars.getNotebookQuery()})

      addNoteInData(data, createNote)
      store.writeQuery({query: notesQuery, variables: lastGetVars.getNotebookQuery(), data})
    },
    // refetchQueries: ['GetNotes'],
    optimisticResponse: {
      __typename: 'Mutation',
      createNote: {
        __typename: 'Note',
        ...new Note(note)
      }
    }
  }
}

const addNoteInData = (data, note) => {
  data.notes.unshift(note)
}

/**
 * delete note
 */

export const deleteNote = (note) => {
  return {
    mutation: gql`mutation DeleteNote($id: ID, $title:String, $body:String, $author:String, $team:String, $dateCreated:String, $dateModified:String, $type:String, $link:String){
      deleteNote(id: $id, title: $title, body:$body, author:$author, team:$team, dateCreated:$dateCreated, dateModified:$dateModified, type:$type, link:$link) {
        id,
        title,
        body,
        author,
        team,
        dateCreated,
        dateModified,
        type,
        link
      }
    }`,
    variables () {
      return {
        id: note.id || null,
        title: note.title || null,
        body: note.body || null,
        author: note.author || null,
        team: note.team || null,
        dateCreated: note.dateCreated || null,
        dateModified: note.dateModified || null,
        type: note.type || null,
        link: note.link || null
      }
    },
    optimisticResponse: {
      __typename: 'Mutation',
      deleteNote: {
        __typename: 'Note',
        ...note
      }
    }
  }
}
