<template>
  <v-list
    v-on-clickaway="resetActive"
  >
    <v-subheader> Notes </v-subheader>
      <v-subheader class="caption text-md-left">
        <note-picker
            :start="noteFilter.from"
            :end="noteFilter.to" >
          </note-picker>
      </v-subheader>

        <v-list-tile v-for="item in uniqueTitles" 
            :key="item.index" 
            class="note-menu-title"
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
import NotePicker from '@/components/notebook/menu/NotePicker.vue'

export default {
    name: 'note-menu',
    props: {
      sidebarClosing: Boolean
    },
    components: {
      NotePicker
    },
    mixins: [
      clickaway
    ],
    data () {
      return {
        titles: [],
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
      startTime() {
        return this.noteFilter.from
      },
      endTime() {
        return this.noteFilter.to
      },
      activeIndex() {
        return this.active
      },
      uniqueTitles() {
        return this.titles.map((title, index) => {
              return {
                  index: index,
                  title: title,
                  active: false
                }
              }
            )
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
            this.titles = Array.from(
              new Set(result.data.notes.map((note, index) => {
                return note.title
              })))
          }
        }
      }
    }
}
</script>

<style>

</style>
