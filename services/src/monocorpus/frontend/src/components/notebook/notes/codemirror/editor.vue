<template>
  <codemirror ref="codeEditor"
              :value="code" 
              :options="options"
              @ready="onReady"
              @focus="onFocus"
              @input="onCodeChange">
  </codemirror>
</template>


<script lang='ts'>
import { codemirror } from 'vue-codemirror'

// require styles
require('codemirror/lib/codemirror.css')
// require('./codemirror.css')

// require more codemirror resource...

// component
export default {
  components: {
    codemirror
  },
  props: {
    'code': {
      type: String
    },
    'options': {
      type: Object
    },
    'onReady': {
      type: Function,
      default: () => {}
    },
    'onFocus': {
      type: Function,
      default: () => {}
    },
    'onCodeChange': {
      type: Function,
      default: () => {}
    }
  }
}
</script>

<style scope>
/* @import 'codemirror/lib/codemirror.css'; */
/* @import './codemirror.css'; */

</style>
