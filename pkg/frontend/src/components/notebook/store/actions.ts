export default {
    updateNoteFilter: ({commit}, noteFilter) => {
      commit('UPDATE_NOTE_FILTER', noteFilter)
    },
    updateNoteFilterTitle: ({commit}, args) => {
      commit('UPDATE_NOTE_FILTER_TITLE', args)
    },
    updateNoteFilterAuthors: ({commit}, emails) => {
      commit("UPDATE_NOTE_FILTER_AUTHORS", emails)
    },
    storeInitialFilter: ({commit}, query) => {
      commit("STORE_INITIAL_FILTER", query)
    },
    updateNotebookTitleFilter: ({commit}, title) => {
      //title filter needs to be updated for both filters.
      commit("UPDATE_NOTEBOOK_TITLE_FILTER", title)
      commit("UPDATE_NOTE_FILTER_TITLE", {title})
    },
    addNote: ({commit}, note) => {
      commit("ADD_NOTE", note)
    },
    updateNote: ({commit}, note) => {
      commit("UPDATE_NOTE", note)
    },
    deleteNote: ({commit}, id) => {
      commit("DELETE_NOTE", id)
    }
  }