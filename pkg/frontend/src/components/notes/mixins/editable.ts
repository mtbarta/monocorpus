import Vue from 'vue'
import ComponentOptions from 'vue-apollo'

export default Vue.extend({
    name: 'editable',

    data () {
      return {
        isEditing: false
      }
    },

    methods: {
      editingNote() {
        this.isEditing = true
      },
      renderingNote() {
        this.isEditing = false
      },
    }
})