import NoteFilter from '@/components/notes/noteFilter'

export default {
  UPDATE_NOTE_FILTER (state, noteFilter: NoteFilter) {
    state.noteFilter = Object.assign(state.noteFilter, noteFilter)
  },
  UPDATE_NOTE_FILTER_TITLE (state, title: string) {
    state.noteFilter.title = title
  },
  UPDATE_NOTE_FILTER_AUTHORS(state: any, emails: string[]) {
    state.noteFilter.authors = emails
  },
  UPDATE_NOTE_FILTER_TEAM(state:any, team: string) {
    state.noteFilter.team = team
  }
}