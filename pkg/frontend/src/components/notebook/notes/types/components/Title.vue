<<template>
  <v-toolbar flat dense class="titleElement">
    
    <v-toolbar-title>
      <span class="header title" 
        v-if="!editing || readOnly"
        @click="enableEditing">
          {{this.tempValue}}
      </span>
      <v-text-field
          v-if="editing && !readOnly" 
          v-model="tempValue"
          label="title"
          @keyup.enter="saveEdit"
          @blur="saveEdit">
      </v-text-field>
      <router-link v-if="!editing" 
        tag="a" 
        class="pageLink" 
        :to="{path: '/notebook', query: {titleFilter: title}}"
        v-on:click.native="$store.dispatch('notebook/updateNotebookTitleFilter', title)">
          <a>
            <v-icon>filter_list</v-icon>
          </a>
        </router-link>

    </v-toolbar-title>
    <v-spacer />
    <v-toolbar-items>
      <v-btn flat light right disabled >
        <v-icon>device_access_time</v-icon> 
          {{ this.formatDate}}
      </v-btn>
      <!-- <v-btn flat>
          <v-icon color="grey lighten-1">delete</v-icon>
      </v-btn> -->
      <delete-button 
        :deleteNote="deleteNote" :id="id" />
    </v-toolbar-items>
  </v-toolbar>
</template>

<script lang='ts'>
import { normalizeDate } from '@/util/dateHelper'
import DeleteButton from './DeleteButton.vue'

export default {
  components: {
    DeleteButton
  },
  props: {
    title: String,
    updateTitle: Function,
    date: Number,
    readOnly: Boolean,
    deleteNote: Function,
    id: String
  },
  data () {
    return {
      tempValue: this.title,
      editing: false
    }
  },
  methods:{
    enableEditing: function(){
      // this.tempValue = this.title;
      this.editing = true;
    },
    disableEditing: function(){
      // this.tempValue = null;
      this.editing = false;
    },
    saveEdit: function(){
      // However we want to save it to the database
      // this.title = this.tempValue;
      this.updateTitle(this.tempValue);
      this.disableEditing();
    }
  },
  computed: {
    formatDate () {
      return normalizeDate(this.date, true)
    }
  }
}
</script>

<style scoped>
.note-info {
  font-size: .75rem;
  vertical-align: center;
  /* display:inline-block;
  align:right;
  float:right; */
}
</style>
