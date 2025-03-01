<script setup lang="ts">
import { computed } from "vue";

import {
  useDriverStore,
  useEventStore,
  useGeneralStore,
  useSessionStore,
} from "@/store/data.store";
import { useTelemetryStore } from "@/store/data.store";
import { Driver, DriverLocation, PitStop } from "@/models/driver.model";
import { Session } from "@/models/session.model";
import moment from "moment";

const sorted = computed(() =>
  useDriverStore.drivers.sort((a, b) => a.Position - b.Position)
);

const getMinisectorColor = (segment: number) => {
  // None SegmentType = iota
  // YellowSegment
  // GreenSegment
  // InvalidSegment // Doesn't get displayed, cut corner/boundaries or invalid segment time?
  // PurpleSegment
  // RedSegment     // After chequered flag/stopped on track
  // PitlaneSegment // In pitlane
  // Mystery
  // Mystery2 // ??? 2021 - Turkey Practice_2
  // Mystery3 // ??? 2020 - Italy Race
  switch (segment) {
    case 1:
      return "bg-yellow-500";
    case 2:
      return "bg-green-500";
    case 3:
      return "bg-black-500";
    case 4:
      return "bg-violet-500";
    case 5:
      return "bg-red-500";
    case 6:
      return "bg-white";
    case 7:
    case 8:
    case 9:
      return "bg-blue-500";
    default:
      return "bg-white";
  }
};

const parseDuration = (duration: number = 0): string => {
  if (duration == 0) {
    return "";
  }

  const parsed = moment(duration)
  if (parsed.minutes() <= 0) {
    return moment(duration).format('ss.SSS');
  }

  return moment(duration).format('mm:ss.SSS');
};

const getCarPosition = (carPosition: number) => {
  // NoLocation CarLocation = iota
  // Pitlane
  // PitOut
  // OutLap
  // OnTrack
  // OutOfRace
  // Stopped
  switch (carPosition) {
    case 1:
      return "Pit lane";
    case 2:
      return "Pit Exit";
    case 3:
      return "Out Lap";
    case 4:
      return "On Track";
    case 5:
      return "Out";
    case 6:
      return "Stopped";
    default:
      break;
  }
};

const getDriverTelemetry = (driver: Driver) => {
  return useTelemetryStore.telemetry.find(
    (t) => t.DriverNumber === driver.Number
  );
};

const parseTimeColor = (personalFastest: boolean, overallFastest: boolean) => {
  if (overallFastest) {
    return "text-purple-500";
  } else if (personalFastest) {
    return "text-green-500";
  } else {
    return "text-yellow-500";
  }
};

const parseTireColor = (tire: number) => {
  switch (tire) {
    case 1:
      return "text-red-500";
    case 2:
      return "text-yellow-500";
    case 3:
      return "text-white";
    case 4:
      return "text-green-500";
    case 5:
      return "text-blue-500";
    default:
      return "text-orange-500";
  }
};

const parseTireType = (tire: number) => {
  // Unknown TireType = iota
  // Soft
  // Medium
  // Hard
  // Intermediate
  // Wet
  // Test
  // HYPERSOFT
  // ULTRASOFT
  // SUPERSOFT
  switch (tire) {
    case 0:
      return "Unknown"
    case 1:
      return "Soft"
    case 2:
      return "Medium"
    case 3:
      return "Hard"
    case 4:
      return "Intermediate"
    case 5:
      return "Wet"
    case 6:
      return "Test"
    case 7:
      return "HYPERSOFT"
    case 8:
      return "ULTRASOFT"
    case 9:
      return "SUPERSOFT"
    default:
      break;
  }
}

const parseLastPitTime = (pitStopTimes: PitStop[]) => {
  if (pitStopTimes && pitStopTimes.length > 0) {
    let lastPitlane = pitStopTimes[pitStopTimes.length - 1];

    if (lastPitlane.PitlaneTime != 0) {
      return parseDuration(lastPitlane.PitlaneTime);
    }
  }
};

const getPilanePosition = (driver: Driver) => {
  let potentialPositionChange: number = 0;
  let positionColor: string = "text-green-500";

  if (
    driver.Location != DriverLocation.Pitlane &&
    driver.Location != DriverLocation.PitOut
  ) {
    let newPosition = driver.Position;
    let pitTimeLost =
      useGeneralStore.general.TimeLostInPitlane +
      useGeneralStore.general.TimeLostInPitlane;
    // Default value for the last car because we won't enter the loop
    let timeToCarAhead = pitTimeLost;
    let timeToCarBehind = 0;

    let timeToClearPitlane = 1000 * 1000 * 10;
    let gapToCar = 1000 * 1000 * 0;
    let minGap = pitTimeLost - timeToClearPitlane;

    for (
      let driverBehind = 1;
      driverBehind < useDriverStore.drivers.length;
      driverBehind++
    ) {
      //timeToCarAhead = pitTimeLost

      // Can't drop below stopped cars
      if (
        useDriverStore.drivers[driverBehind].Location ==
        DriverLocation.Stopped ||
        useDriverStore.drivers[driverBehind].Location ==
        DriverLocation.OutOfRace
      ) {
        break;
      }

      newPosition++;

      gapToCar += useDriverStore.drivers[driverBehind].TimeDiffToPositionAhead;

      // If the gap to the prediction car is less than the time to drive past the pitlane then keep looking
      if (gapToCar < minGap) {
        continue;
      }

      timeToCarBehind = gapToCar - minGap;
      timeToCarAhead =
        useDriverStore.drivers[driverBehind].TimeDiffToPositionAhead -
        timeToCarBehind;
      break;
    }

    if (newPosition == driver.Position) {
      timeToCarAhead += driver.TimeDiffToPositionAhead;
    }

    if (
      driver.Location != DriverLocation.Stopped &&
      driver.Location != DriverLocation.OutOfRace
    ) {
      // potentialPositionChange = fmt.Sprintf("%02d", newPosition)
      potentialPositionChange = newPosition;
    }

    if (newPosition != driver.Position) {
      // positionColor = colornames.Red
      positionColor = "text-red-500";
    }
  } else {
    positionColor = "text-white";
    potentialPositionChange = 0;
  }

  return { potentialPositionChange, positionColor };
};
</script>

<template>
  <div class="flex flex-col flex-1 w-full">
    <div class="flex w-full">
      <div className="grid items-center gap-2 p-1 px-2 text-sm font-medium text-zinc-500 w-full" :style="{
        gridTemplateColumns:
          useSessionStore.session === Session.RaceSession
            ? '5.5rem 100rem 5.5rem 4rem 5rem 5.5rem 5rem 5rem 5rem 5rem 5rem 5rem 5rem 5rem 5rem 5rem'
            : '5rem 42rem 5.5rem 4rem 5rem 5.5rem 5rem 5rem 5rem 5rem 5rem 5rem 5rem 5rem 5rem',
      }">
        <p>Position</p>
        <p>Micro-sectors</p>
        <p>Fastest</p>
        <p>Gap</p>
        <p>S1</p>
        <p>S2</p>
        <p>S3</p>
        <p>Last lap</p>
        <p>DRS</p>
        <p>Tire / Laps</p>
        <p>Pits</p>
        <p v-if="useSessionStore.session === Session.RaceSession">
          Last pit stop
        </p>
        <p v-if="useSessionStore.session === Session.RaceSession">
          Pit position
        </p>
        <p>Telemetry</p>
        <p>Speed trap</p>
        <p>Location</p>
      </div>
    </div>

    <div class="flex flex-col">
      <div v-for="driver in sorted" class="flex select-none flex-col gap-1 p-1.5 border-b">
        <div class="grid items-center gap-2 w-full" :style="{
          gridTemplateColumns:
            useSessionStore.session === Session.RaceSession
              ? '5.5rem 100rem 5.5rem 4rem 5rem 5.5rem 5rem 5rem 5rem 5rem 5rem 5rem 5rem 5rem 5rem 5rem'
              : '5rem 42rem 5.5rem 4rem 5rem 5.5rem 5rem 5rem 5rem 5rem 5rem 5rem 5rem 5rem 5rem 5rem',
        }">
          <!-- POSITION -->
          <span class="flex w-fit items-center justify-between gap-0.5 px-1 py-1 font-black"
            :style="{ color: driver.HexColor }">
            {{ driver.Position }} {{ driver.ShortName }}
          </span>

          <!-- MICRO-SECTORS -->
          <div class="flex gap-2">
            <!-- {{ driver.Segment }} -->
            <div v-for="(status, index) in driver.Segment" class="flex items-center gap-2">
              <div class="border" :class="getMinisectorColor(status)"
                :style="{ height: '10px', width: '8px', borderRadius: '3.2px' }"></div>
              <div v-if="
                index == useEventStore.event.Sector1Segments - 1 ||
                index ==
                useEventStore.event.Sector1Segments +
                useEventStore.event.Sector2Segments -
                1
              " class="text-gray-600">
                |
              </div>
            </div>
          </div>

          <!-- FASTEST LAP -->
          <span class="flex w-fit items-center justify-between gap-0.5 px-1 py-1 font-black">
            {{ parseDuration(driver.FastestLap) }}
          </span>

          <!-- GAP -->
          <div class="flex flex-col justify-between gap-0.5 px-1 py-1">
            <span class="flex items-center justify-between gap-0.5 px-1 py-1 font-semibold">
              {{ `${parseDuration(driver.TimeDiffToPositionAhead) !== '' ? '+' +
                parseDuration(driver.TimeDiffToPositionAhead) : parseDuration(driver.TimeDiffToPositionAhead)}` }}
            </span>

            <span class="flex items-center justify-between gap-0.5 px-1 py-1 font-thin">
              {{ `${parseDuration(driver.TimeDiffToFastest) !== '' ? '+' +
                parseDuration(driver.TimeDiffToFastest) : parseDuration(driver.TimeDiffToFastest)}` }}
            </span>
          </div>

          <!-- S1 -->
          <span class="flex w-fit items-center justify-between gap-0.5 px-1 py-1 font-black" :class="parseTimeColor(
            driver.Sector1PersonalFastest,
            driver.Sector1OverallFastest
          )
            ">
            {{ parseDuration(driver.Sector1) }}
          </span>

          <!-- S2 -->
          <span class="flex w-fit items-center justify-between gap-0.5 px-1 py-1 font-black" :class="parseTimeColor(
            driver.Sector2PersonalFastest,
            driver.Sector2OverallFastest
          )
            ">
            {{ parseDuration(driver.Sector2) }}
          </span>

          <!-- S3 -->
          <span class="flex w-fit items-center justify-between gap-0.5 px-1 py-1 font-black" :class="parseTimeColor(
            driver.Sector3PersonalFastest,
            driver.Sector3OverallFastest
          )
            ">
            {{ parseDuration(driver.Sector3) }}
          </span>

          <!-- Last lap -->
          <span class="flex w-fit items-center justify-between gap-0.5 px-1 py-1 font-black" :class="parseTimeColor(
            driver.LastLapPersonalFastest,
            driver.LastLapOverallFastest
          )
            ">
            {{ parseDuration(driver.LastLap) }}
          </span>

          <!-- DRS -->
          <span class="flex w-fit items-center justify-between gap-0.5 rounded-md px-1 py-1 font-black" :class="{
            'text-green-500': getDriverTelemetry(driver)?.DRS,
            'text-red-500': !getDriverTelemetry(driver)?.DRS,
          }">
            DRS
          </span>

          <!-- Tire / Laps -->
          <span class="flex w-fit items-center justify-between gap-0.5 rounded-md px-1 py-1 font-black"
            :class="parseTireColor(driver.Tire)">
            {{ parseTireType(driver.Tire) }} / {{ driver.LapsOnTire }}
          </span>

          <!-- Pits -->
          <span class="flex w-fit items-center justify-between gap-0.5 rounded-md px-1 py-1 font-black">
            {{ driver.Pitstops }}
          </span>

          <!-- Last pit stop time -->
          <span class="flex w-fit items-center justify-between gap-0.5 rounded-md px-1 py-1 font-black"
            v-if="useSessionStore.session === Session.RaceSession">
            {{ parseLastPitTime(driver.PitStopTimes) }}
          </span>

          <!-- Pit position -->
          <span class="flex w-fit items-center justify-between gap-0.5 rounded-md px-1 py-1 font-black"
            :class="getPilanePosition(driver).positionColor" v-if="useSessionStore.session === Session.RaceSession">
            {{ getPilanePosition(driver).potentialPositionChange }}
          </span>

          <!-- Telemetry -->
          <div className="flex flex-col items-center gap-2 place-self-start">
            <p className="flex h-8 w-8 items-center justify-center font-mono text-lg">
              {{ getDriverTelemetry(driver)?.Gear }}
            </p>

            <div>
              <p className="text-right font-mono font-medium leading-none">
                {{ getDriverTelemetry(driver)?.Speed }}
              </p>
              <p className="text-sm leading-none text-zinc-600">KM/h</p>
            </div>

            <div className="flex flex-col">
              <div className="flex flex-col gap-1">
                <!-- <DriverPedals
                  className="bg-red-500"
                  value="{carData[5]}"
                  maxValue="{1}"
                />
                <DriverPedals
                  className="bg-emerald-500"
                  value="{carData[4]}"
                  maxValue="{100}"
                />
                <DriverPedals
                  className="bg-blue-500"
                  value="{carData[0]}"
                  maxValue="{15000}"
                  /> -->
                <!-- Throttle -->
                <div className="h-1.5 w-20 overflow-hidden rounded-xl bg-zinc-800">
                  <div class="h-1.5 bg-green-500" :style="{
                    width: `${getDriverTelemetry(driver)?.Throttle || 0 / 100
                      }%`,
                  }" :animate="{ transitionDuration: '0.1s' }" layout></div>
                </div>
                <!-- Brake -->
                <div className="h-1.5 w-20 overflow-hidden rounded-xl bg-zinc-800">
                  <div class="h-1.5 bg-red-500" :style="{
                    width: `${getDriverTelemetry(driver)?.Brake || 0 / 100}%`,
                  }" :animate="{ transitionDuration: '0.1s' }" layout></div>
                </div>
                <!-- RPM -->
                {{ getDriverTelemetry(driver)?.RPM || 0 / 15000 }}
                {{ (getDriverTelemetry(driver)?.RPM || 0 / 15000) * 100 }}
                <div className="h-1.5 w-20 overflow-hidden rounded-xl bg-zinc-800">
                  <div class="h-1.5 bg-blue-500" :style="{
                    width: `${getDriverTelemetry(driver)?.RPM || 0 / 15000}%`,
                  }" :animate="{ transitionDuration: '0.1s' }" layout></div>
                </div>
              </div>
            </div>
          </div>
          <!-- {{ getDriverTelemetry(driver)?.Throttle }} -->

          <!-- Speed trap -->
          <span class="flex w-fit items-center justify-between gap-0.5 rounded-md px-1 py-1 font-black">
            {{ driver.SpeedTrap }} KM/h
          </span>

          <!-- Location -->
          <span class="flex w-fit items-center justify-between gap-0.5 rounded-md px-1 py-1 font-black">
            {{ getCarPosition(driver.Location) }}
          </span>
        </div>
      </div>
    </div>
  </div>
</template>
