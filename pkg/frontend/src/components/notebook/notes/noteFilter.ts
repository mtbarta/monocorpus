import { NoteOptions, default as Note } from './note'
import * as moment from 'moment'

class NoteFilterOptions extends NoteOptions {
  to?: number
  from?: number
  authors?: string[]
  title?: string
  body?: string
  link?: string
  team?: string
  type?: string
  id?: string
}

export default class NoteFilter {
  to?: number
  from?: number
  authors?: string[]
  title?: string
  body?: string
  link?: string
  team?: string
  type?: string
  id?: string

  constructor(opts: NoteFilterOptions) {
    this.to = opts.to || moment.utc().unix()
    this.from = opts.from || moment.utc().startOf('week').subtract(1, 'week').unix()
    this.authors = opts.authors || []

    this.title = opts.title ? opts.title.slice(0) : null
    this.body = opts.body ? opts.body.slice(0) : null
    this.link = opts.link || null
    this.team = opts.team || null
    this.type = opts.type || null
    this.id = opts.id || null
  }

  public copy(): NoteFilter {
    const copy = JSON.parse(JSON.stringify(this))
    return new NoteFilter(copy)
  }

  /**
   * sending a 'to' field for the main notebook query makes it
   * hard to continually update a cache through apollo.
   * 
   * We can make our lives easier by not sending this parameter.
   */
  public getNotebookQuery () {
    const str = JSON.stringify(this);
    const {to,id, ...rest} = JSON.parse(str)

    return rest;
  }

  public fetchOlderNotesQuery(amount:any, duration:string): NoteFilter {
    const newQuery:NoteFilter  = new NoteFilter(this)
    newQuery.to = newQuery.from

    const newFrom = moment.unix(newQuery.from).subtract(amount, duration)
    newQuery.from = newFrom.unix()
    newQuery.id = null

    return newQuery
  }
}