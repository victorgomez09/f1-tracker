<script setup lang="ts">
import { defineProps, onMounted, ref } from "vue";

const props = defineProps({
  circuit: Number,
});

const space = 1000;

const rad = (deg) => deg * (Math.PI / 180);
const deg = (rad) => rad / (Math.PI / 180);

const rotate = (x, y, a, px, py) => {
  const c = Math.cos(rad(a));
  const s = Math.sin(rad(a));

  x -= px;
  y -= py;

  const newX = x * c - y * s;
  const newY = y * c + x * s;

  return [newX + px, (newY + py) * -1];
};

const getTrackStatusColour = (status) => {
  switch (status) {
    case "2":
    case "4":
    case "6":
    case "7":
      return "yellow";
    case "5":
      return "red";
    default:
      return "var(--colour-fg)";
  }
};

const sortDriverPosition = (Lines) => (a, b) => {
  const [racingNumberA] = a;
  const [racingNumberB] = b;

  const driverA = Lines[racingNumberA];
  const driverB = Lines[racingNumberB];

  return Number(driverB?.Position) - Number(driverA?.Position);
};

const bearingToCardinal = (bearing) => {
  const cardinalDirections = ["N", "NE", "E", "SE", "S", "SW", "W", "NW"];
  return cardinalDirections[Math.floor(bearing / 45) % 8];
};

const expanded = ref(false);
const data = ref({});
const minX = ref();
const minY = ref();
const widthX = ref();
const widthY = ref();
const stroke = ref(0);

onMounted(async () => {
  try {
    const apiResponse = await fetch(
      `https://api.multiviewer.app/api/v1/circuits/${
        props.circuit
      }/${new Date().getFullYear()}`,
      {
        headers: {
          "User-Agent": "tdjsnelling/monaco",
        },
      }
    );
    if (apiResponse.status !== 200) console.log("status code");

    const data = await apiResponse.json();
    console.log("data from map", data);
  } catch (e) {
    console.log("error", e);
  }
});
</script>

<template></template>
