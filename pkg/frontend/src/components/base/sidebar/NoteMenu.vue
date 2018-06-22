<<template>
  <v-list>
    <v-list-group v-model="menuOpen">
      <v-list-tile slot="activator">
        <v-list-tile-action>
          <v-icon>collections_bookmark </v-icon>
        </v-list-tile-action>
        <v-list-tile-content>
          <span>Notes</span>
        </v-list-tile-content>
        
      </v-list-tile>
        
      <v-list-tile v-for="title in noteTitles" 
        :key="title" 
        dense 
        ripple 
        :to="{name: 'Notebook', query: {titleFilter: title}}"
        v-on:click.native="$store.dispatch('notebook/updateNotebookTitleFilter', title)">
        <v-list-tile-action>
          <v-icon> note </v-icon>
        </v-list-tile-action>
        <v-list-tile-content>
          <!-- <router-link class="pageLink" :to="{path: '/notebook', query: {titleFilter: title}}"> -->
             {{title}}
          <!-- </router-link> -->
        </v-list-tile-content>
      </v-list-tile>
    </v-list-group>
  </v-list>
</template>

<script lang='ts'>
import { getNoteTitles, notesTitleQuery } from '@/graphql/noteQueries'
import TitleFilter from '@/components/notebook/notes/titleFilter'
import gql from 'graphql-tag'
import { mapState, mapGetters, mapActions } from 'vuex'

export default {
    name: 'note-menu',
    props: {
      sidebarClosing: Boolean
    },
    data () {
      return {
        // noteFilter: this.$store.state.notebook.noteFilter,
        noteTitles: [],
        active: false,
        menuOpen: false
      }
    },
    computed: {
      ...mapGetters('notebook', [
        'noteFilter'
      ]),
      ...mapActions('notebook', [
        'updateNotebookTitleFilter'
      ]),
      titleFilter() {
        return TitleFilter.fromNoteFilter(this.noteFilter)
      }
    },
    apollo: {
      titles: {
        query: notesTitleQuery,
        variables () {
          return this.titleFilter
        },
        manual: true,
        result (result) {
          if (result.data.notes) {
            this.noteTitles = Array.from(new Set(result.data.notes.map(note => note.title)))
          }
        }
      }
    }
}
</script>

<style>

</style>
