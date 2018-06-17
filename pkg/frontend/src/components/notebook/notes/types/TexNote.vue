<<template>
  <v-card>

    <title-box :title="title"
                :date="note.dateCreated"
                :updateTitle="updateTitle" 
                :readOnly="readOnly"
                :deleteNote="deleteNote"
                :id="note.id"
                />

    <v-card-text
        @click="editingNote"
        v-on-clickaway="renderingNote">
      <Editor v-if="isEditing && !readOnly"
        :code="note.body"
        :changeThrottle="500"
        :onReady="onMounted"
        :onCodeChange="onCodeChange"
        :options="editorOptions"
        >
      </Editor>
      <div class="renderedNote" 
          v-if="!isEditing || readOnly"
          v-html="renderCode(note.body)">

      </div>
    </v-card-text>
    <div class="footer">
        <a v-for="btn in note.buttons" v-bind:class="'btn btn-' + btn.type + ' btn-xs'" v-bind:href="btn.href" v-bind:target="btn.target">{{btn.message}}</a>
    </div>

  </v-card>
</template>

<script>
require('codemirror/mode/stex/stex')
import { mixin as clickaway } from 'vue-clickaway';
import Editor from '@/components/notebook/notes/codemirror/editor.vue'
import TitleBox from './components/Title'
import { mapActions } from 'vuex'
import katex from 'katex'
import Note from '../note'
// import sanitize from 'sanitize-html'

/**
 * line will be a note object.
 * 
 * title, body, icon, color, time, buttons
 */
export default {
    name: 'TexNote',
    mixins: [ clickaway ],
    components: {
      Editor,
      TitleBox
    },
    data () {
      return {
        code: this.note.body,
        title: this.note.title,
        isEditing: false,
        editorOptions: {
          height: 'auto',
          tabSize: 4,
          mode:  {
            name: 'stex',
            highlightFormatting: true
          },
          lineNumbers: false,
          line: true,
          viewportMargin: Infinity
        }
      }
    },
    watch: {
      note () {
        this.title = this.note.title,
        this.code = this.note.body
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
        let n = new Note(this.note)
        n.body = newCode
        this.updateNote(n)
      },
      updateTitle(title) {
        let n = new Note(this.note)
        n.title = title
        this.updateNote(n)
      },
      editingNote() {
        this.isEditing = true
      },
      renderingNote() {
        this.isEditing = false
      },
      renderCode(s) {
        return katex.renderToString(s, {
          throwOnError: false
        })
      }
      
    }
}
</script>

<style scoped>
/* .note {
  padding: 1px 0 0 5px;
  background-color:white
} */
.note-body {
  padding-top: 5px;
  min-height: 10%;
}
.header {
  margin: 0px;
}
.footer {
  min-height: 1px;
}
</style>
