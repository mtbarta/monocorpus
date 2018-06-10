import moment from 'moment'

export class NoteOptions {
  title?: string
  dateCreated?: number
  dateModified?: number
  body?: string
  link?: string
  team?: string
  type?: string
  author?: string
  id?: string

  image?: string

  searchScore?: number
}

export default class Note {
  title: string
  dateCreated: number
  dateModified: number
  body: string
  link: string
  team: string
  type: string
  author: string
  id: string
  image: string
  
  searchScore?: number

  constructor (opts: NoteOptions) {
    this.title = opts.title ? opts.title.slice(0) : ''
    this.dateCreated = moment.unix(opts.dateCreated).unix() || moment.utc().unix()
    this.dateModified = moment.unix(opts.dateModified).unix() || moment.utc().unix()

    this.body = opts.body ? opts.body.slice(0) : ''
    this.link = opts.link || ''
    this.team = opts.team || ''
    this.type = opts.type || ''
    this.author = opts.author || ''
    this.id = opts.id || ''

    this.image = opts.image || ''

    this.searchScore = opts.searchScore
  }

  prototype (): Note {
    return new Note({
      title: String(this.title),
      body: String(this.body),
      type: this.type,
      link: this.link,
      team: String(this.team),
      author: String(this.author),
      image: String(this.image)
    })
  }
}

