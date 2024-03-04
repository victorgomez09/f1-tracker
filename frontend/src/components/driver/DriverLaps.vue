<script setup lang="ts">
import { PropType } from "vue";

import { TimeStats } from "../../models/driver.model";
import { getTimeColor } from "../../utils/time.util";

const props = defineProps({
  best: Object as PropType<TimeStats>,
  last: Object as PropType<TimeStats>,
});
</script>

<template>
  <div class="flex flex-col items-center">
    <!-- Last -->
    <div class="flex items-center gap-2">
      <span class="font-semibold">Last</span>
      <p
        class="text-lg font-semibold leading-none"
        :class="getTimeColor(props.last!.fastest, props.last!.pb), { 'text-gray-500': !props.last?.value }"
      >
        {{ !!props.last?.value ? props.last?.value : "-- -- ---" }}
      </p>
    </div>

    <!-- Best -->
    <div class="flex items-center gap-2">
      <span class="font-semibold">Best</span>
      <p
        class="text-lg font-semibold leading-none"
        :class="getTimeColor(props.best!.fastest, props.best!.pb), { 'text-gray-500': !props.best?.value }"
      >
        {{ !!props.best?.value ? props.best?.value : "-- -- ---" }}
      </p>
    </div>
  </div>
</template>
../../models/driver.model../../utils/time.util
