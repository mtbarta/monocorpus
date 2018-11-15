<template>
<!-- <div class="note"> -->  
  <v-card>
    <title-box :title="title"
                 :date="note.dateCreated"
                 :updateTitle="updateTitle" 
                 :readOnly="readOnly"
                 :deleteNote="deleteNote"
                 :id="note.id"
                 />
    <v-divider />
    <v-card-text class="text-space"
        @click="editingNote"
        v-on-clickaway="renderingNote">

      <Editor v-if="isEditing && !readOnly"
        class="editor"
        :code="code"
        :onReady="onMounted"
        :onCodeChange="onCodeChange"
        :onFocus="onFocus"
        :options="editorOptions"
        >
      </Editor>

      <div class="renderedNote" 
          v-if="!isEditing || readOnly"
          v-html="renderCode(code)">

      </div>
    </v-card-text>
  </v-card>
</template>

<script>
require("codemirror/mode/gfm/gfm")

import { mixin as clickaway } from 'vue-clickaway';
import Editor from './components/codemirror/editor.vue'
import TitleBox from './components/Title'
import Note from './note'
import markdownable from './mixins/markdownable'
import editable from './mixins/editable'

/**
 * line will be a note object.
 * 
 * title, body, icon, color, time, buttons
 */
export default {
    name: 'MarkdownNote',
    mixins: [ 
      clickaway,
      markdownable,
      editable
    ],
    components: {
      Editor,
      TitleBox,
    },
    data () {
      return {
        code: this.note.body,
        title: this.note.title,
        editorOptions: {
          height: 'auto',
          tabSize: 4,
          mode:  {
            name: "gfm",
            highlightFormatting: false
          },
          lineNumbers: false,
          line: true,
          lineWrapping: true,
          viewportMargin: 25
        }
      }
    },
    props: {
      note: Object,
      updateNote: Function,
      readOnly: {
        type: Boolean,
        default: false
      },
      deleteNote: Function
    },
    methods: {
      onMounted(editor) {
        this.editor = editor;
      },
      onCodeChange(newCode) {
        this.code = newCode

        let n = new Note(this.note)
        n.body = newCode
        this.updateNote(n)
      },
      onFocus(cm) {
        // console.log("focus")
      },
      updateTitle(title) {
        this.title = title
        let n = new Note(this.note)
        n.title = title
        this.updateNote(n)
      },  
    }
  }
</script>

<style>
/* highlightjs style */
@import '/static/highlightjs/a11y-light.css';
@import '/static/katex.min.css';

  .markdown .inline-katex .katex {
    display: inline;
    text-align: initial;
    line-height: 1.8em;
  }

</style>
