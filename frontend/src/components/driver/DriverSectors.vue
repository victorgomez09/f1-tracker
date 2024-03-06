<script setup lang="ts">
import { PropType, defineProps } from "vue";

import { Sector } from "../../models/driver.model";
import { getTimeColor } from "../../utils/time.util";

const props = defineProps({
  sectors: Object as PropType<Sector[]>,
  driverDisplayName: String,
});
</script>

<template>
  <div className="flex gap-2">
    <div class="flex flex-col gap-[0.2rem]" v-for="sector in props.sectors!">
      <div class="flex h-[10px] flex-row gap-1">
        <div
          v-for="status in sector.segments"
          class="badge badge-primary badge-xs"
          :class="[
            { 'bg-warning': status === 2048 || status === 2052 }, // TODO unsure
            { 'bg-success': status === 2049 },
            { 'bg-violet-600': status === 2051 },
            { 'bg-info': status === 2064 },
            { 'bg-base-content': status === 0 },
          ]"
        />
      </div>

      <p
        class="text-lg font-semibold leading-none"
        :class="
          (getTimeColor(sector.current.fastest, sector.current.pb),
          { 'text-gray-500': !sector.current.value })
        "
      >
        {{
          !!sector.current.value
            ? sector.current.value
            : !!sector.last.value
            ? sector.last.value
            : "-- ---"
        }}
      </p>
    </div>
  </div>
</template>
