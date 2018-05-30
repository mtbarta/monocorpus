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

    <v-card-text class="text-space">
      <form enctype="multipart/form-data" novalidate v-if="image==null && !readOnly">
        <div>
          <input type="file" :disabled="loading" @change="onFilesChange"
            accept="image/*" class="input-file">
            <!-- <p v-if="isInitial">
              Drag your file(s) here to begin<br> or click to browse
            </p> -->
        </div>
      </form>
      <v-card-media v-else>
        <viewer :options="viewerOpts" class="viewer">
          <img :src="image" class="uploadedImage"/>
        </viewer>
      </v-card-media>
      <!-- {{this.image}} -->
    </v-card-text>
    <!-- <div class="footer">
        <a v-for="btn in note.buttons" v-bind:class="'btn btn-' + btn.type + ' btn-xs'" v-bind:href="btn.href" v-bind:target="btn.target">{{btn.message}}</a>
    </div><div class="footer">
        <a v-for="btn in note.buttons" v-bind:class="'btn btn-' + btn.type + ' btn-xs'" v-bind:href="btn.href" v-bind:target="btn.target">{{btn.message}}</a>
    </div> -->

</v-card>
</template>

<script>

import TitleBox from './components/Title'
import {normalizeDate} from '@/util/dateHelper'
import Note from '@/components/notebook/notes/note'
import 'viewerjs/dist/viewer.css'
import Viewer from "v-viewer/src/component.vue"


/**
 * line will be a note object.
 * 
 * title, body, icon, color, time, buttons
 */
export default {
    name: 'ImageNote',
    components: {
      TitleBox,
      Viewer
    },
    data () {
      return {
        image: this.note.image,
        title: this.note.title,
        loading: false,
        error: null,
        viewerOpts: {
          title: false,
          toolbar: {
            zoomIn: 4,
            zoomOut: 4,
            oneToOne: 4,
            reset: 4,
            prev: false,
            play: {
              show: false,
              size: 'large',
            },
            next: false,
            rotateLeft: 4,
            rotateRight: 4,
            flipHorizontal: 4,
            flipVertical: 4,
          }
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

      onFocus(cm) {
        // console.log("focus")
      },
      updateTitle(title) {
        this.title = title
        let n = new Note(this.note)
        n.title = title
        this.updateNote(n)
      },
      formatDate(date) {
        return normalizeDate(date)
      },
      async onFilesChange(e) {
        const n = new Note(this.note)
        const files = e.target.files || e.dataTransfer.files

        if (files.length !== 1) {
          return
        }

        if (files[0].size > 16000000) {
          this.error = "image is greater than 16MB"
          return
        }

        const reader = new FileReader()

        reader.onload = (event) => {
          n.image = event.target.result
          this.image = event.target.result
        }

        this.loading = true
        reader.readAsDataURL(files[0])
    
        this.updateNote(n)

        this.loading = false

      }
    }
  }
</script>

<style scoped>
/* .text-space {
  border-top: 1px solid;
} */
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
.renderedNote {
  padding-left: 1em;
}

.uploadedImage {
  height: 400px;
  width: auto;
}

.viewer {
  cursor: pointer;
}
</style>
