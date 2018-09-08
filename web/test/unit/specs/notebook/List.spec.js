import { mount } from '@vue/test-utils'
import List from '@/components/lists/List.vue'
import * as moment from 'moment'

import GroupedCollection from '@/components/lists/groupedCollection'
import Vuetify from 'vuetify'
import { createLocalVue } from '@vue/test-utils'

describe('Notebook.test.js', () => {
  let cmp
  let localVue

  beforeEach(() => {
    localVue = createLocalVue()
    localVue.use(Vuetify)

    cmp = mount(Notebook, {
      propsData: {
        title: '',
        notes: [],
        readOnly: false
      },
      stubs: {
        NoteWrapper: '<div class="notewrapper" />',
      },
      sync: false,
      attachToDocument: false,
      localVue: localVue,
      mocks: {
        $apollo: {
          mutate: jest.fn()
        }
      }
    })
  })

  it('shows a load more button', () => {
    cmp.setData({
      notes: []
    })

    expect(cmp.html()).toContain('notelist')
  })

})