<script setup lang="ts">
import { PropType, ref } from 'vue'
import moment from 'moment';

import DriverTag from '../driver/DriverTag.vue';
import AudioControls from './audio/AudioControls.vue';
import AudioProgress from './audio/AudioProgress.vue';
import { TeamRadioType } from "../../models/radio.model";

defineProps({
    radio: Object as PropType<TeamRadioType>
})

const audioRef = ref<HTMLAudioElement | null>(null);
const intervalRef = ref<NodeJS.Timer | null>(null);
const playing = ref<boolean>(false);
const duration = ref<number>(0);
const progress = ref<number>(0);

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

    if (!playing.value) {
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
<li class="flex flex-col gap-1">
      <time
        class="text-sm font-medium leading-none text-gray-500"
        :datetime="moment.utc(radio?.utc).local().format('HH:mm:ss')"
      >
        {{ moment.utc(radio?.utc).local().format("HH:mm:ss") }}
      </time>

      <div
        className="grid items-center gap-1"
        :style="{
          gridTemplateColumns: '3.5rem 20rem',
        }"
      >
        <div className="w-10 place-self-start">
          <DriverTag
            :team-color="radio?.teamColor"
            :short="radio?.short"
          />
        </div>

        <div className="flex items-center gap-4">
          <AudioControls :playing="playing" :on-click="togglePlayback" />
          <AudioProgress :duration="duration" :progress="progress" />

          <audio
            hidden="true"
            :src="radio?.audioUrl"
            ref="audioRef"
            @ended="onEnded"
            @loadedmetadata="loadMeta"
          />
        </div>
      </div>
    </li>
</template>