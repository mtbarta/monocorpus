<<template>
<!-- <div class="note"> -->  
  <v-card>
      <!-- <v-card-title>
      
      </v-card-title> -->
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
    <!-- <div class="footer">
        <a v-for="btn in note.buttons" v-bind:class="'btn btn-' + btn.type + ' btn-xs'" v-bind:href="btn.href" v-bind:target="btn.target">{{btn.message}}</a>
    </div><div class="footer">
        <a v-for="btn in note.buttons" v-bind:class="'btn btn-' + btn.type + ' btn-xs'" v-bind:href="btn.href" v-bind:target="btn.target">{{btn.message}}</a>
    </div> -->

</v-card>
</template>

<script>
require("codemirror/mode/markdown/markdown")

import { mixin as clickaway } from 'vue-clickaway';
import Editor from '@/components/notebook/notes/codemirror/editor.vue'
import TitleBox from './components/Title'
import marked from 'marked';
import {normalizeDate} from '@/util/dateHelper'
import Note from '@/components/notebook/notes/note'
import sanitize from 'sanitize-html'

marked.setOptions({
  gfm: true,
  smartLists: true,
  breaks: true
})
/**
 * line will be a note object.
 * 
 * title, body, icon, color, time, buttons
 */
export default {
    name: 'MarkdownNote',
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
            name: "markdown",
            highlightFormatting: true
          },
          lineNumbers: false,
          line: true,
          lineWrapping: true,
          viewportMargin: 25
        }
      }
    },
    watch: {
      // note () {
      //   // this.title = this.note.title,
      //   this.code = this.note.body
      // }
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
      editingNote() {
        this.isEditing = true
      },
      renderingNote() {
        this.isEditing = false
      },
      renderCode(string) {
        return marked(string)
      },
      formatDate(date) {
        return normalizeDate(date)
      }
      
    }
  }
</script>

<style scoped>

</style>
