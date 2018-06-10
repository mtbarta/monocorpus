<<template>
  <v-list>
      <v-list-tile avatar>
        <v-list-tile-avatar color='blue'>
          <v-icon dark class="initials">{{this.userName[0]}}</v-icon>
        </v-list-tile-avatar>
        <v-list-tile-content>
          <v-list-tile-title>{{this.userName}}</v-list-tile-title>
        </v-list-tile-content>
        <v-list-tile-action>
          <v-btn icon ripple
            :href="accountManagementURL">
            <v-icon color="grey lighten-1">settings</v-icon>
          </v-btn>
        </v-list-tile-action>
      </v-list-tile>
  </v-list>
</template>

<script>
// import { getNoteTitles } from '@/graphql/noteQueries'
// import gql from 'graphql-tag'
import {mapState} from 'vuex'
import { keycloakAuth } from '@/keycloak'

export default {
    name: 'user-menu',
    data () {
      return {
        identity: {},
        accountManagementURL: keycloakAuth.createAccountUrl()
      }
    },
    mounted () {
    },
    methods: {
    },
    computed: {
      ...mapState({
        firstName(state) {
          return state.login.firstName
        },
        lastName(state) {
          return state.login.lastName
        },
        userName(state) {
          const first = state.login.firstName
          const last = state.login.lastName

          return [first, last].join(' ')
        }
      })
    }
}
</script>

<style scoped>
.initials {
  font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
}
</style>
