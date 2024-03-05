<script setup lang="ts">
import { PropType, VNodeRef, ref } from "vue";
import moment from "moment";

import AudioControls from "./AudioControls.vue";
import DriverTag from "../driver/DriverTag.vue";
import { TeamRadioType } from "../../models/radio.model";

const audioRef = ref<VNodeRef | null>(null);
const intervalRef = ref<NodeJS.Timer | null>(null);
const playing = ref<boolean>(false);
const duration = ref<number>(0);
const progress = ref<number>(0);

defineProps({
  teamRadios: Array as PropType<TeamRadioType[]>,
});

const loadMeta = () => {
  if (!audioRef.value) return;
  duration.value = audioRef.value.duration;
};

const onEnded = () => {
  playing.value = false;
  progress.value = 0;

  if (intervalRef.value) {
    clearInterval(Number(intervalRef.value));
  }
};

const updateProgress = () => {
  if (!audioRef.value) return;
  progress.value = audioRef.value.currentTime;
};

const togglePlayback = () => {
  const updatePlaying = () => {
    if (!audioRef.value) return playing.value;

    console.log(playing.value);
    if (!playing.value) {
      console.log("play", audioRef.value);
      audioRef.value.play();

      intervalRef.value = setInterval(updateProgress, 10);
    } else {
      audioRef.value.pause();

      if (intervalRef.value) {
        clearInterval(Number(intervalRef.value));
      }
    }

    return !playing.value;
  };

  playing.value = updatePlaying();
};
</script>

<template>
  <ul class="flex flex-col gap-2">
    <li class="flex flex-col gap-1" v-for="teamRadio in teamRadios">
      <time
        class="text-sm font-medium leading-none text-gray-500"
        :datetime="moment.utc(teamRadio.utc).local().format('HH:mm:ss')"
      >
        {{ moment.utc(teamRadio.utc).local().format("HH:mm:ss") }}
      </time>

      <div
        className="grid place-items-center items-center gap-4"
        :style="{
          gridTemplateColumns: '2rem 20rem',
        }"
      >
        <div className="w-10 place-self-start">
          <DriverTag
            :team-color="teamRadio.teamColor"
            :short="teamRadio.short"
          />
        </div>

        <div className="flex items-center gap-4">
          <AudioControls :playing="playing" :on-click="togglePlayback" />
          <AudioProgress :duration="duration" :progress="progress" />

          <audio
            :src="teamRadio.audioUrl"
            ref="audioRef"
            :on-ended="onEnded"
            :on-loadedmetadata="loadMeta"
          />
        </div>
      </div>
    </li>
  </ul>
</template>
