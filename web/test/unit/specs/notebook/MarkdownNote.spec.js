import { shallowMount, mount } from '@vue/test-utils'
import Note from '@/components/notes/note.ts'
import MarkdownNote from '@/components/notes/MarkdownNote.vue'
import * as moment from 'moment'
import Editor from '@/components/notes/components/codemirror/editor.vue'
import TitleBox from '@/components/notes/components/Title.vue'
import Vuetify from 'vuetify'
import { createLocalVue } from '@vue/test-utils'

describe('MarkdownNote.test.js', () => {
  let cmp
  let note
  let localVue

  beforeEach(() => {
    localVue = createLocalVue()
    localVue.use(Vuetify)

    cmp = mount(MarkdownNote, { // Create a shallow instance of the component
      propsData: {
        note: new Note({
          dateCreated: moment.utc('2018-01-01'),
          body: 'hello',
          title: 'test'
        }),
        updateNote: jest.fn(),
        deleteNote: jest.fn()
      },
      stubs: {
        Editor: '<div class="editor" />',
        TitleBox: '<div class="title" />'
      },
      sync: false,
      attachToDocument: true,
      localVue: localVue
    })
  })

  it('renders the editor', (done) => {
    cmp.setData({
      isEditing: true,
    })
    cmp.setProps({
      readOnly: false
    })
    cmp.find('div.text-space').trigger('click')
    expect(cmp.vm.isEditing).toBe(true)
    cmp.find('div.text-space').trigger('focus')
    cmp.vm.$nextTick(() => {
      expect(cmp.html()).toContain('editor')
      done()
    })
  })

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

    expect(cmp.html()).toContain('<div class="renderedNote"')  
  })

})
