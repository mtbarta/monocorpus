import { shallowMount } from '@vue/test-utils'
import NoteMenu from '@/components/base/sidebar/NoteMenu.vue'

import Vuetify from 'vuetify'
import { createLocalVue } from '@vue/test-utils'

describe('NoteMenu.test.js', () => {
  let cmp
  let localVue

  beforeEach(() => {
    localVue = createLocalVue()
    localVue.use(Vuetify)

    cmp = shallowMount(NoteMenu, {
      propsData: {
        titles: [],
        readOnly: false
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

  it('shows only unique titles', () => {
    cmp.setData({
      titles: ['one', 'one']
    })

    expect(cmp.uniqueTitles).toHaveLength(1)
  })

})