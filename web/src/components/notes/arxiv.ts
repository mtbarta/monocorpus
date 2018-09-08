import axios from 'axios'
import { parseString } from 'xml2js'
import * as moment from 'moment'

export const arxivSearch = async (id: string): Promise<ArxivArticle> => {
  let url = 'http://export.arxiv.org/api/query'

  let request = await axios.get(url, {
    params: {
      'id_list': id
    }
  })
  let res: ArxivArticle = null
  parseString(request.data, (err, parsed) => {
    if (err != null) {
      return {}
    } else {
      let feed = parsed.feed

      let entry = coerceEntry(feed.entry[0])

      res = entry
    }
  })
  return res
}

interface ArticleOptions {
  id: string
  updated: Date
  published: Date
  title: string
  summary: string
  links: {href: string, title: string}[]
  authors: string[]
  categories: string[]
}

class ArxivArticle {
  id: string
  updated: Date
  published: Date
  title: string
  summary: string
  links: {href: string, title: string}[]
  authors: string[]
  categories: string[]

  constructor(entry: ArticleOptions) {
    Object.assign(this, entry)
  }
}

const coerceEntry = function (entry: any): ArxivArticle {
  return new ArxivArticle({
    id: entry.id[0],
    updated: new Date(entry.updated[0]),
    published: new Date(entry.published[0]),
    title: entry.title[0].trim().replace(/\s+/g, ' '),
    summary: entry.summary[0].trim().replace(/\s+/g, ' '),
    links: entry.link.map(function (link) {
      return {
        href: link['$']['href'],
        title: link['$']['title']
      }
    }),
    authors: entry.author.map(function (author) {
      return author['name'][0]
    }),
    categories: entry.category.map(function (category) {
      return category['$']['term']
    })
  })
}

const unique = function (a, k) {
  var a_, i, j, known, len
  a_ = []
  known = {}
  for (j = 0, len = a.length; j < len; j++) {
    i = a[j]
    if (!known[i[k]]) {
      known[i[k]] = true
      a_.push(i)
    }
  }
  return a_
}