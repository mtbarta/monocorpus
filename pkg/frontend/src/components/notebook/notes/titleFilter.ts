import NoteFilter from './noteFilter'

export default class TitleFilter {
  to?: number
  from?: number
  authors?: string[]
  team?: string
  type?: string

  constructor (noteFilter: NoteFilter) {
    // this.to = noteFilter.to
    this.from = noteFilter.from
    this.authors = noteFilter.authors
    this.team = noteFilter.team
    this.type = noteFilter.type
  }

  static fromNoteFilter(noteFilter: NoteFilter) {
    return new TitleFilter(noteFilter)
  }
}