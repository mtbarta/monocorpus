import Vue from 'vue'

import { shallowMount, mount, createLocalVue } from '@vue/test-utils'
import Note from '@/components/notebook/notes/note.ts'
import MarkdownNote from '@/components/notebook/notes/types/MarkdownNote.vue'
import * as moment from 'moment'
import Editor from '@/components/notebook/notes/codemirror/editor.vue'

const localVue = createLocalVue();

describe('Notebook.test.js', () => {
  let cmp
  let note

  beforeEach(() => {
    cmp = shallowMount(MarkdownNote, { // Create a shallow instance of the component
      propsData: {
        note: new Note({
          dateCreated: moment.utc('2018-01-01'),
          body: 'hello'
        }),
        updateNote: jest.fn(),
        deleteNote: jest.fn()
      },
      sync: false,
      attachToDocument: true,
      localVue
    })
  })

  it('onCodeChange creates a copy', () => {
    // Within cmp.vm, we can access all Vue instance methods
    const oldNote = cmp.vm.note
    cmp.vm.onCodeChange('test')

    let fnNote = cmp.vm.updateNote.mock.calls[0][0]

    expect(fnNote).not.toBe(oldNote)
  })

  it('updates the note body', () => {
    // Within cmp.vm, we can access all Vue instance methods
    cmp.vm.onCodeChange('test')

    let fnNote = cmp.vm.updateNote.mock.calls[0][0]

    expect(fnNote.body).toBe('test')
  })

  it('updateTitle creates a copy', () => {
    cmp.vm.updateTitle('test')

    let fnNote = cmp.vm.updateNote.mock.calls[0][0]

    expect(fnNote).not.toBe(note)
  })

  it('updates the note title', () => {
    cmp.vm.updateTitle('test')

    let fnNote = cmp.vm.updateNote.mock.calls[0][0]

    expect(fnNote.title).toBe('test')
  })

  // it('renders the editor', () => {
  //   cmp.setData({
  //     isEditing: true,
  //   })
  //   cmp.setProps({
  //     readOnly: false
  //   })
  //   cmp.find('div.text-space').trigger('click')
  //   expect(cmp.html()).toContain('editor')
  // })

  it('renders the html div', () => {
    cmp.setData({
      isEditing: false,
    })
    cmp.setProps({
      readOnly: false
    })
    
    expect(cmp.contains(Editor)).toBe(false)
    expect(cmp.html()).toContain('renderedNote')
  })

  it('respects readOnly notes', () => {
    cmp.setData({
      isEditing: true, //should not show editor if readOnly
    })
    cmp.setProps({
      readOnly: true
    })

    expect(cmp.html()).toContain('<div class="renderedNote"')  })
})
