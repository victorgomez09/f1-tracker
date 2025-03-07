<script setup lang="ts">
import { useDriverStore, useEventStore, useLocationStore, useTimeStore } from "@/store/data.store";
import moment from "moment";
import { computed, onMounted, ref } from "vue";
import { Location } from "@/models/location.model"

const sortPos = (a: Location, b: Location) => {
  const aPos = useDriverStore.drivers.find(driver => driver.Number === a.DriverNumber)
  const bPos = useDriverStore.drivers.find(driver => driver.Number === b.DriverNumber)
  return (aPos?.Position || 0) - (bPos?.Position || 0);
};

const space = 1000;
const rad = (deg: number) => deg * (Math.PI / 180);
const rotate = (x: number, y: number, a: number, px: number, py: number) => {
  const c = Math.cos(rad(a));
  const s = Math.sin(rad(a));

  x -= px;
  y -= py;

  const newX = x * c - y * s;
  const newY = y * c + x * s;

  return { y: newX + px, x: newY + py };
};
const rotationFIX = 90;

// REF
const points = ref<null | { x: number; y: number }[]>(null);
const rotation = ref<number>(0);
const ogPoints = ref<null | { x: number; y: number }[]>(null);
const minX = ref<number | null>(null);
const minY = ref<number | null>(null);
const widthX = ref<number | null>(null);
const widthY = ref<number | null>(null);
const positions = computed(() =>
  useLocationStore.locations
    ? useLocationStore.locations.sort((a, b) =>
      moment.utc(b.Timestamp).diff(moment.utc(a.Timestamp))
    )
    : null
);
const xS = computed(() => ogPoints.value?.map((item) => item.x));
const yS = computed(() => ogPoints.value?.map((item) => item.y));

const rotatedPos = (pos: Location) =>
  rotate(
    pos.X,
    pos.Y,
    rotation.value,
    (Math.max(...xS.value!) - Math.min(...xS.value!)) / 2,
    (Math.max(...yS.value!) - Math.min(...yS.value!)) / 2
  )
// const out = (pos: any) =>
//   () =>
//     pos.status === "OUT" ||
//     pos.status === "RETIRED" ||
//     pos.status === "STOPPED"
const transformTanslate = (pos: any) =>
  [
    `translateX(${rotatedPos(pos).x}px)`,
    `translateY(${rotatedPos(pos).y}px)`,
  ].join(" ");

onMounted(async () => {
  try {
    const apiResponse = await fetch(
      `https://api.multiviewer.app/api/v1/circuits/${useEventStore.event.Meeting.Circuit.Key}/${moment(useTimeStore.time.Timestamp).year()}`,
      {
        headers: {
          "User-Agent": "tdjsnelling/monaco",
        },
      }
    );
    if (apiResponse.status !== 200) console.log("status code");

    if (apiResponse.status === 200) {
      const rawData = await apiResponse.json();

      const centerX = (Math.max(...rawData.x) - Math.min(...rawData.x)) / 2;
      const centerY = (Math.max(...rawData.y) - Math.min(...rawData.y)) / 2;

      const fixedRotation = rawData.rotation + rotationFIX;

      const rotatedPoints = rawData.x.map((x: any, index: number) =>
        rotate(x, rawData.y[index], fixedRotation, centerX, centerY)
      );

      const pointsX = rotatedPoints.map((item: any) => item.x);
      const pointsY = rotatedPoints.map((item: any) => item.y);

      const cMinX = Math.min(...pointsX) - space;
      const cMinY = Math.min(...pointsY) - space;
      const cWidthX = Math.max(...pointsX) - cMinX + space * 2;
      const cWidthY = Math.max(...pointsY) - cMinY + space * 2;

      minX.value = cMinX;
      minY.value = cMinY;
      widthX.value = cWidthX;
      widthY.value = cWidthY;
      points.value = rotatedPoints;
      rotation.value = fixedRotation;
      ogPoints.value = rawData.x.map((xItem: number, index: number) => ({
        x: xItem,
        y: rawData.y[index],
      }));
    }
  } catch (e) {
    console.log("error", e);
  }
});
</script>

<template>
  <div v-if="!points || !minX || !minY || !widthX || !widthY" class="flex h-full w-full items-center justify-center">
    <div class="h-5/6 w-5/6 animate-pulse rounded-lg bg-gray-700" />
  </div>
  <svg v-else :viewBox="minX + ' ' + minY + ' ' + widthX + ' ' + widthY" class="w-full h-full"
    xmlns="http://www.w3.org/2000/svg">
    <path class="bg-secondary" stroke-width="300" stroke-linecap="round" fill="transparent"
      :d="'M' + points![0].x + ',' + points![0].y + ' ' + points?.map((point) => 'L' + point.x + ',' + point.y).join(' ')" />
    <path stroke="white" stroke-width="60" stroke-linecap="round" fill="transparent"
      :d="'M' + points![0].x + ',' + points![0].y + ' ' + points!.map((point) => 'L' + point.x + ',' + point.y).join(' ')" />
    <!-- <g v-for="pos in positions!.sort(sortPos).reverse()" :class="{ 'opacity-30': out(pos) }" class="fill-zinc-700" -->
    <g v-for="pos in positions!.sort(sortPos).reverse()" class="fill-zinc-700" :style="{
      transition: 'all 1s linear',
      transform: transformTanslate(pos),
      // ...(pos.teamColor && { fill: '#' + pos.teamColor }),
    }">
      <circle :id="`map.driver.${pos.DriverNumber}.circle`" r="120" />
      <text :id="`map.driver.${pos.DriverNumber}.text`" font-weight="bold" :font-size="120 * 3" :style="{
        transform: 'translateX(150px) translateY(-120px)',
      }">
        {{ pos.DriverNumber }}
      </text>
    </g>
  </svg>
</template>
