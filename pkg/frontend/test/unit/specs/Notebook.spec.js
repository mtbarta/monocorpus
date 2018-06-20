// import Vue from 'vue'

// import { shallowMount } from '@vue/test-utils'
// import Notebook from '@/components/notebook/Notebook'
// import noteFilter from '@/components/notebook/notes/noteFilter.ts'
// import Note from '@/components/notebook/notes/note.ts'
// import GroupCollection from '@/components/notebook/notes/lists/groupedCollection.ts'
// import namedCollection from '@/components/notebook/notes/lists/namedCollection.ts'
// import * as moment from 'moment'


// describe('Notebook.test.js', () => {
//   let cmp
//   let notes

//   beforeEach(() => {
//     notes = jest.fn()
//               .mockReturnValue([
//                 new Note({
//                   dateCreated: moment.utc('2018-01-01')
//                 }),
//               ])
//     cmp = shallowMount(Notebook, { // Create a shallow instance of the component
//       data: {
//         notes: [],
//         skipUpdates: false,
//         hasMore: true,
//         numCallsAfterEmpty: 0, //infinite loading complete() isn't working.
//         error: null,
//         isTriggerFirstLoad: false,
//         tryFetchingNotes: true
//       },
//       propsData: {
//         titleFilter: ''
//       },
//       mocks: {
//         $apollo: {
//           notes
//         }
//       }
//     })
//   })

//   it('calls notes once', () => {
//     // Within cmp.vm, we can access all Vue instance methods
//     expect(notes).toBeCalled()
//   })

//   it('creates one day', () => {
//     expect(cmp.notes).toHaveProperty('2018-01-01')

//     expect(cmp.notes['2018-01-01']).toHaveLength(1)
//   })
// })
