import Note from '../note'

export default class NamedCollection {
  title: string|number
  notes: Note[]

  constructor(title: string|number, notes: Note[]) {
    this.title = title
    this.notes = notes
  }
}