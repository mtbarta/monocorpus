import noteFilter from '@/components/notes/noteFilter.ts'

describe('NoteFilter.test.js', () => {
  it('creates a copy on get', () => {
    const filter = new noteFilter({
      team: 'test'
    })
    const result = filter.getNotebookQuery()

    filter.team = 'not_test'

    expect(result.team).toBe('test')
  })

  it('does not set a to date', () => {
    const filter = new noteFilter({
      team: 'test'
    })
    const result = filter.getNotebookQuery()

    expect(result.to).not.toBeDefined()
  })

  it('creates a copy', () => {
    const filter = new noteFilter({
      team: 'test'
    })
    const result = filter.copy()

    filter.team = 'not_test'

    expect(result.team).toBe('test')
  })

  it('fetches older queries', () => {
    const filter = new noteFilter({
      team: 'test'
    })
    
    const olderFilter = filter.fetchOlderNotesQuery(1, 'day')

    expect(filter.from).toBe(olderFilter.to)
    expect(olderFilter.from).toBeLessThan(olderFilter.to)
  })
})