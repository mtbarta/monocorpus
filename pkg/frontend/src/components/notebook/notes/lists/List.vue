<<template>
  <!-- <v-container fluid grid-list> -->
    <v-layout column wrap>
        <v-divider />
        <v-subheader v-text="title" />

        <transition-group name="notelist" tag="span">
          <note-wrapper v-for="note in notes"
            :note="note" :type="note.type" :key="note.id" :readOnly="readOnly"/>
        </transition-group>
    </v-layout>
  <!-- </v-container> -->
</template>

<script lang='ts'>
import { Component, Emit, Inject, Model, Prop, Provide, Vue, Watch } from 'vue-property-decorator'

import NoteWrapper from '../NoteWrapper.vue'
import Note from '../note'

@Component({
  components: {
    NoteWrapper
  }
})
export default class List extends Vue {
  @Prop()
  title: string

  @Prop()
  notes: Note[]

  @Prop({default: false})
  readOnly: boolean

}

</script>

<style scoped>
.notelist {
  transition: all 0.5s;

}
.notelist-enter, .notelist-enter-active, .notelist-leave-to .notelist-leave-active {
  opacity: 0;
  transform: scale(0);
  /* transition: all 0.5s; */
  transition: opacity 1s;
}
.notelist-enter-to {
  opacity: 1;
  transform: scale(1);
}

.notelist-leave-active {
}

.notelist-move {
  opacity: 1;
  transition: all 0.5s;
}

/* .notelist-enter-to {
  opacity: 1;
  transform: scale(1);
}

.notelist-enter-active, .notelist-leave-active {
  transition: opacity 0.3s, transform 0.3s;
  transform-origin: left center;
}
.notelist-enter, .notelist-leave-to .list-leave-active {
  opacity: 0;
  transition: all 0.5s;
} */
</style>
