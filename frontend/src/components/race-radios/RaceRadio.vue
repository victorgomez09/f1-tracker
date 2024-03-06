<script setup lang="ts">
import { PropType } from "vue";

import RadioMessage from "./RadioMessage.vue";
import { TeamRadioType } from "../../models/radio.model";

defineProps({
  teamRadios: Array as PropType<TeamRadioType[]>
})
</script>

<template>
  <ul class="flex flex-col gap-2 p-2">
    <RadioMessage v-for="radio in teamRadios" :radio="radio" />
  </ul>
</template>
