<template>
    <vc-date-picker
        mode='range'
        v-model='selectedDate'
        :max-date='new Date()'
        show-caps>
    </vc-date-picker>
</template>

<script lang='ts'>
import Vue from 'vue';
import { Component, Prop, Watch } from 'vue-property-decorator'
import { mapActions } from 'vuex'
import * as moment from 'moment'
import VCalendar from 'v-calendar'
import 'v-calendar/lib/v-calendar.min.css'

Vue.use(VCalendar, {
  firstDayOfWeek: 2,  // Monday
  componentPrefix: 'vc'
});

@Component({
    methods: mapActions('notebook', [
        'updateNoteFilterStart',
        'updateNoteFilterEnd'
    ])
})
export default class NotePicker extends Vue {
    updateNoteFilterStart: Function
    updateNoteFilterEnd: Function

    @Prop({default: moment().unix(), type: Number})
    start: number

    @Prop({default: moment().unix(), type: Number})
    end: number

    @Watch('selectedDate.start')
    onStartChanged(value: Date, oldvalue: Date) {
        const date = moment(value).utc().unix()
        this.updateNoteFilterStart(date)
    }

    @Watch('selectedDate.end')
    onEndChanged(value: Date, oldvalue: Date) {
        const date = moment(value).utc().unix()
        this.updateNoteFilterEnd(date)
    }

    data() {
        return {
            selectedDate: {
                start: moment.unix(this.start).local().toDate(),
                end: moment.unix(this.end).local().toDate()
            }
        }
    }

}
</script>

<style>

</style>
