<script setup lang="ts">
import { PropType, computed, defineProps, onMounted, ref } from "vue";
import moment from "moment";

import { F1TrackStatus } from "../../models/track-status.model";
import { DriverPositionBatch } from "../../models/position.model";
import { sortPos } from "../../utils/position.utils";

const props = defineProps({
  circuit: Number,
  trackStatus: Object as PropType<F1TrackStatus>,
  windDirection: Number,
  positionBatches: Array as PropType<DriverPositionBatch[]>,
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
const data = ref({} as any);
const minX = ref();
const minY = ref();
const widthX = ref();
const widthY = ref();
const stroke = ref(0);
const points = ref<null | { x: number; y: number }[]>(null);
const rotation = ref<number>(0);
const ogPoints = ref<null | { x: number; y: number }[]>(null);

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

    if (apiResponse.status === 200) {
      const rawData = await apiResponse.json();

      const px = (Math.max(...rawData.x) - Math.min(...rawData.x)) / 2;
      const py = (Math.max(...rawData.y) - Math.min(...rawData.y)) / 2;

      rawData.transformedPoints = rawData.x.map((x: number, i: number) =>
        rotate(x, rawData.y[i], rawData.rotation, px, py)
      );

      const cMinX =
        Math.min(...rawData.transformedPoints.map(([x]: any) => x)) - space;
      const cMinY =
        Math.min(...rawData.transformedPoints.map(([, y]: any) => y)) - space;
      const cWidthX =
        Math.max(...rawData.transformedPoints.map(([x]: any) => x)) -
        cMinX +
        space * 2;
      const cWidthY =
        Math.max(...rawData.transformedPoints.map(([, y]: any) => y)) -
        cMinY +
        space * 2;

      minX.value = cMinX;
      minY.value = cMinY;
      widthX.value = cWidthX;
      widthY.value = cWidthY;

      const cStroke = (cWidthX + cWidthY) / 225;
      stroke.value = cStroke;

      rawData.corners = rawData.corners.map((corner: any) => {
        const transformedCorner = rotate(
          corner.trackPosition.x,
          corner.trackPosition.y,
          rawData.rotation,
          px,
          py
        );

        const transformedLabel = rotate(
          corner.trackPosition.x + 4 * cStroke * Math.cos(rad(corner.angle)),
          corner.trackPosition.y + 4 * cStroke * Math.sin(rad(corner.angle)),
          rawData.rotation,
          px,
          py
        );

        return { ...corner, transformedCorner, transformedLabel };
      });

      rawData.startAngle = deg(
        Math.atan(
          (rawData.transformedPoints[3][1] - rawData.transformedPoints[0][1]) /
            (rawData.transformedPoints[3][0] - rawData.transformedPoints[0][0])
        )
      );

      data.value = rawData;
    }
  } catch (e) {
    console.log("error", e);
  }
});

const xS = ogPoints.value?.map((item) => item.x);
const yS = ogPoints.value?.map((item) => item.y);

const rotatedPos = (pos: any) => rotate(
  pos.x,
  pos.y,
  rotation,
  (Math.max(...xS!) - Math.min(...xS!)) / 2,
  (Math.max(...yS!) - Math.min(...yS!)) / 2,
);

const out = (pos: any) => pos.status === "OUT" || pos.status === "RETIRED" || pos.status === "STOPPED";

const transform = [`translateX(${rotatedPos.x}px)`, `translateY(${rotatedPos.y}px)`].join(" ");


const positions = computed(() => props.positionBatches ? props.positionBatches.sort((a, b) => moment.utc(b.utc).diff(moment.utc(a.utc)))[0].positions : null);
</script>

<template>
    <svg
			:viewBox="minX minY widthX widthY"
			className="h-full w-full xl:max-h-screen"
			xmlns="http://www.w3.org/2000/svg"
		>
			<path
				className="stroke-slate-700"
				strokeWidth={300}
				strokeLinejoin="round"
				fill="transparent"
				:d="`M${points[0].x},${points[0].y} ${points.map((point) => `L${point.x},${point.y}`).join(" ")}`"
			/>

			<path
				stroke="white"
				strokeWidth={60}
				strokeLinejoin="round"
				fill="transparent"
				:d="`M${points[0].x},${points[0].y} ${points.map((point) => `L${point.x},${point.y}`).join(" ")}`"
			/>
								<g
									v-for="pos in positions!.sort(sortPos).reverse()"
									:class="{ 'opacity-30': out(pos) }"
                  class="fill-zinc-700"
									:style="{
                    transition: 'all 1s linear',
										transform,
										...(pos.teamColor && { fill: `#${pos.teamColor}` })
                  }"
								>
									<circle :id="`map.driver.${pos.driverNr}.circle`" r={120} />
									<text
										:id="`map.driver.${pos.driverNr}.text`"
										fontWeight="bold"
										fontSize={120 * 3}
										style={{
											transform: "translateX(150px) translateY(-120px)",
										}}
									>
										{pos.short}
									</text>
								</g>
		</svg>
</template>
