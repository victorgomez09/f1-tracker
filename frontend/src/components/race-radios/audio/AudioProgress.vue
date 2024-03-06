<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps({
    duration: Number,
    progress: Number
})

const percent = computed(() => Math.round((props.progress! / props.duration!) * 100) || 0)
const currentDuration = computed(() => (Math.round((props.progress! + Number.EPSILON) * 100) / 100) || 0)
const totalDuration = computed(() => (Math.round((props.duration! + Number.EPSILON) * 100) / 100) || 0)
</script>

<template>
    <div class="flex flex-col items-center w-20">
        <progress
            class="progress progress-primary z-0"
            :value="percent"
            max="100"
            ></progress>

        <span class="">{{ currentDuration }} - {{ totalDuration }}</span>
    </div>
</template>