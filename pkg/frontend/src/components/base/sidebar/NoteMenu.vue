<template>
  <v-list
    v-on-clickaway="resetActive"
  >
    <v-subheader> Notes </v-subheader>
        <v-list-tile v-for="item in noteTitles" 
            :key="item.index" 
            dense 
            ripple
            @click="navclick(item)"
            :value="item.index == activeIndex"
        >
          <v-list-tile-action>
            <v-icon> note </v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
              {{item.title}}
          </v-list-tile-content>
        </v-list-tile>
  </v-list>
</template>

<script lang='ts'>
import { mapState, mapGetters, mapActions } from 'vuex'
import titleQuery from '@/graphql/noteTitles.graphql'
import { mixin as clickaway } from 'vue-clickaway';

export default {
    name: 'note-menu',
    props: {
      sidebarClosing: Boolean
    },
    mixins: [
      clickaway
    ],
    data () {
      return {
        noteTitles: [],
        menuOpen: false,
        active: -1
      }
    },
    methods: {
      navclick(item) {
        this.active = item.index
        this.$router.push({name: 'Notebook', query: {titleFilter: item.title}})
      },
      resetActive() {
        this.active = -1
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
        return this.noteFilter.getTitleQuery()
      },
      activeIndex() {
        return this.active
      }
    },
    apollo: {
      titles: {
        query: titleQuery,
        variables () {
          return this.titleFilter
        },
        manual: true,
        result (result) {
          if (result.data.notes) {
            this.noteTitles = Array.from(
              new Set(result.data.notes.map((note, index) => {
                return {
                  index: index,
                  title: note.title,
                  active: false
                }
              })))
          }
        }
      }
    }
}
</script>

<style>

</style>
