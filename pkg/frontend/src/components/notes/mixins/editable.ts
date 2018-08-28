import Vue from 'vue'

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