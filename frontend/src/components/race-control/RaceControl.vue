<script setup lang="ts">
import { PropType, computed, defineProps } from "vue";
import moment from "moment";

import { F1RaceControlMessage } from "../../models/race-control.model";
import { sortUtc } from "../../utils/time.util";

const props = defineProps({
  messages: Object as PropType<F1RaceControlMessage[]>,
});

console.log("messages", props.messages);

const messagesSorted = computed(() => props.messages?.sort(sortUtc));
</script>

<template>
  <ul class="flex flex-col gap-2 p-2">
    <li class="flex items-center text-nowrap" v-for="msg in messagesSorted">
      <div class="mr-2">
        <time :dateTime="moment.utc(msg.utc).local().format('HH:mm:ss')">{{
          moment.utc(msg.utc).local().format("HH:mm:ss")
        }}</time>
        {{ "·" }}
        <time
          class="text-gray-600"
          :dateTime="moment.utc(msg.utc).format('HH:mm')"
        >
          {{ moment.utc(msg.trackTime).format("HH:mm") }}
        </time>
      </div>

      <div class="flex items-center col-span-11 gap-1 text-left">
        <div v-if="msg.flag">
          <div
            class="badge badge-outline"
            :class="[
              {
                'badge-warning':
                  msg.flag === 'YELLOW' || msg.flag === 'DOUBLE YELLOW',
              },
              { 'badge-error': msg.flag === 'RED' },
              { 'badge-success': msg.flag === 'GREEN' },
              { 'badge-info': msg.flag === 'BLUE' },
              { 'badge-ghost': msg.flag === 'BLACK AND WHITE' },
              { 'badge-success': msg.flag === 'CLEAR' },
            ]"
          >
            FLAG
          </div>
        </div>

        <p class="text-sm">{{ msg.message }}</p>
      </div>
    </li>
  </ul>
</template>
