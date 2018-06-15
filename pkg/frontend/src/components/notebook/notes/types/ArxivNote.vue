<<template>
<!-- <div class="note"> -->
  <v-card color="" class="">
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
          <v-text-field
              name="arxiv-id"
              label="arxiv ID"
              id="arxiv-id"
              v-model="link"
            ></v-text-field>

      <v-progress-linear v-if="waiting==true" v-bind:indeterminate="true"></v-progress-linear>
      
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
          v-html="renderCode(this.code)">

      </div>
    </v-card-text>
    <div class="footer">
        <a v-for="btn in note.buttons" v-bind:class="'btn btn-' + btn.type + ' btn-xs'" v-bind:href="btn.href" v-bind:target="btn.target">{{btn.message}}</a>
    </div>

<!-- </div> -->
</v-card>
</template>

<script lang='ts'>
import { mixin as clickaway } from 'vue-clickaway';
import Editor from '@/components/notebook/notes/codemirror/editor.vue'
import TitleBox from './components/Title.vue'
import { arxivSearch } from './util/arxiv'
import marked from 'marked'
import util from 'util'
import {normalizeDate} from '@/util/dateHelper'
import Note from '../note'
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
    name: 'ArxivNote',
    mixins: [ clickaway ],
    components: {
      Editor,
      TitleBox
    },
    data () {
      return {
        code: this.note.body,
        title: this.note.title,
        link: this.note.link,
        waiting: false,
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
          viewportMargin: 20,
        }
      }
    },
    watch: {
      note () {
        this.title = this.note.title
        this.code = this.note.body
      }
    },
    props: ['note', 'updateNote', 'deleteNote', 'readOnly'],
    asyncComputed: {
      // async arxivArticle () {
      //   const article = await getArxivInfo();

      //   return formatArxivArticle(article);
      // },
      async retrieveArxivArticle() {
        if (this.link == null || this.link == '') {
          return null
        }

        if (this.note.body.length > 0) {
          return this.note.body
        }

        this.waiting = true
        const newlink = this.link;

        let linkParts = this.link.split("/")
        const id = linkParts[linkParts.length-1]

        const article = await arxivSearch(id)
        
        let n = new Note(this.note)
        n.body = this.formatArxivArticle(article)
        this.code = n.body
        n.link = newlink

        this.updateNote(n)
        this.waiting = false

        // return res;
      },
    },
    methods: {
      onMounted(editor) {
        this.editor = editor;
      },
      onCodeChange(newCode) {
        let n = new Note(this.note)
        n.body = newCode

        n.link = this.link
        this.updateNote(n)
      },
      onFocus(cm) {
        // console.log("focus")
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
      formatDate(date) {
        return normalizeDate(date)
      },
      renderCode(code) {
        return sanitize(marked(code))
      },
      formatArxivArticle(article) {
        const str = util.format("### %s\n%s\n**%s**\n***\n%s\n",
          article.title, article.id, article.authors.join(", "), article.summary
        )       

        return str
      }
      
    }
  }
</script>

<style scoped>
.note {
  padding: 1px 0 0 5px;
  background-color:white
}
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
