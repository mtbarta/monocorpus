<template>
  <v-container>
    <v-card
      max-width=400
    >
      <v-card-title> <h2> Most Common Titles in the Past Two Weeks </h2> </v-card-title>
      <v-data-table
        :items="titleCounts"
        item-key="title"
        :loading="loading"
        :headers="headers"
        :pagination.sync="pagination"
      >
        <template slot="headers" slot-scope="props">
          <tr>
            <th
              v-for="header in props.headers"
              :key="header.text"
              :class="['column sortable', pagination.descending ? 'desc' : 'asc', header.value === pagination.sortBy ? 'active' : '']"
              @click="changeSort(header.value)"
            >
            >
              <v-icon small>arrow_upward</v-icon>
              {{ header.text }}
            </th>
          </tr>
        </template>

        <template slot="items" slot-scope="props">
        <tr :active="props.selected" @click="props.selected = !props.selected">
          <td>{{ props.item.title }}</td>
          <td class="text-xs-right">{{ props.item.count }}</td>
        </tr>
      </template>
      </v-data-table>
    </v-card>
  </v-container>
</template>

<script lang='ts'>
import Vue from 'vue'
import { mapActions, mapState, mapGetters } from 'vuex'
import titleQuery from '@/graphql/noteTitles.graphql'

export default Vue.extend({
  
  data () {
    return {
      titles: [],
      loading: false,
      headers: [
        {
          text: 'title',
          align: 'right',
          value: 'title'
        },
        {
          text: 'count',
          align: 'left',
          value: 'count'
        }
        ],
      pagination: {
        sortBy: 'count'
      },
    }
  },

  apollo: {
    $loadingKey: 'loading',
    titles: {
      query: titleQuery,
      variables () {
        return this.titleFilter
      },
      manual: true,
      result (result) {
        if (result.data.notes) {
          this.titles = result.data.notes
        }
      }
    }
  },

  methods: {
    count(array) {
      let count = {};
      array.forEach(val => count[val] = (count[val] || 0) + 1);
      return count;
    },
    changeSort (column) {
        if (this.pagination.sortBy === column) {
          this.pagination.descending = !this.pagination.descending
        } else {
          this.pagination.sortBy = column
          this.pagination.descending = false
        }
      }
  },

  computed: {
    ...mapGetters('notebook', [
      'noteFilter'
    ]),
    titleFilter() {
      return this.noteFilter.getTitleQuery()
    },
    titleCounts() {
      let titles = this.titles.map((x) => x.title)
      let counts = this.count(titles)

      return Object.keys(counts).map((key) => {
        return {
          title: key,
          count: counts[key]
        }
      })
    }
  }
})
</script>

<style>

</style>
