<template>
  <v-layout row justify-center>
      <v-btn flat
          @click.stop="dialog=true">
          <v-icon color="grey lighten-1">delete</v-icon>
      </v-btn>

      <v-dialog class='dialog' max-width="400" v-model="dialog">
        <v-card>
          <v-card-title>
            <span>Are you sure you want to delete this note?</span>
            <v-spacer></v-spacer>
          </v-card-title>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="secondary" flat @click.stop="dialog=false">Cancel</v-btn>
            <v-btn color="primary" flat @click.stop="del()">Yes</v-btn>

          </v-card-actions>
        </v-card>
      </v-dialog>
  </v-layout>
</template>

<script lang='ts'>
import { Component, Model, Prop, Vue } from 'vue-property-decorator'

@Component
export default class DeleteButton extends Vue {

  dialog: boolean = false

  /**
   * the delete function to call. this should be passed in from
   * the note wrapper.
   */
  @Prop()
  deleteNote: (id: string) => {}

  /**
   * the id of the note. this is a mongo ObjectId
   */
  @Prop()
  id: string

  del(): void {
    this.deleteNote(this.id)
    this.dialog = false
  }
}
</script>

<style scoped>
.dialog {
  max-width: 200px;
}
</style>
