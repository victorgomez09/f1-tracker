<script setup lang="ts">
import moment from "moment";
import { computed } from "vue";

import { Driver, DriverLocation, PitStop } from "@/models/driver.model";
import { HistoricalSession, Session } from "@/models/session.model";
import {
  useDriverStore,
  useEventStore,
  useGeneralStore,
  useSessionStore,
  useTelemetryStore,
} from "@/store/data.store";

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
  <div class="overflow-x-auto">
    <table class="table table-zebra table-xs gap-2">
      <thead>
        <tr>
          <th>Position</th>
          <th>Micro-sectors</th>
          <th>Fastest</th>
          <th>Gap</th>
          <th>S1</th>
          <th>S2</th>
          <th>S3</th>
          <th>Last lap</th>
          <th>DRS</th>
          <th>Tire</th>
          <th>Laps</th>
          <th>Pits</th>
          <th
            v-if="useSessionStore.session === Session.Race || useSessionStore.session === HistoricalSession.RaceSession">
            Pit time</th>
          <th
            v-if="useSessionStore.session === Session.Race || useSessionStore.session === HistoricalSession.RaceSession">
            Pit pos</th>
          <th>Spd trp</th>
          <th>Location</th>
        </tr>
      </thead>

      <tbody>
        <tr v-for="driver in sorted">
          <!-- Position -->
          <td :style="{ color: driver.HexColor }"> {{ driver.Position }} {{ driver.ShortName }}</td>
          <!-- Micro sectors -->
          <td class="flex items-center gap-0.5">
            <div v-for="(status, index) in driver.Segment" class="flex items-center">
              <div class="size-2" :class="getMinisectorColor(status)" :style="{ borderRadius: '.15em' }"></div>
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
          </td>
          <!-- fastest lap -->
          <td>{{ parseDuration(driver.FastestLap) }}</td>
          <!-- Gap -->
          <td>
            <span class="flex items-center justify-between gap-0.5 font-semibold">
              {{ `${parseDuration(driver.TimeDiffToPositionAhead) !== '' ? '+' +
                parseDuration(driver.TimeDiffToPositionAhead) : parseDuration(driver.TimeDiffToPositionAhead)}` }}
            </span>

            <!-- <span class="flex items-center justify-between gap-0.5 font-thin">
              {{ `${parseDuration(driver.TimeDiffToFastest) !== '' ? '+' +
                parseDuration(driver.TimeDiffToFastest) : parseDuration(driver.TimeDiffToFastest)}` }}
            </span> -->
          </td>
          <!-- s1 -->
          <td :class="parseTimeColor(
            driver.Sector1PersonalFastest,
            driver.Sector1OverallFastest
          )">{{ parseDuration(driver.Sector1) }}</td>
          <!-- s2 -->
          <td :class="parseTimeColor(
            driver.Sector2PersonalFastest,
            driver.Sector2OverallFastest
          )
            ">{{ parseDuration(driver.Sector2) }}</td>
          <!-- s3 -->
          <td :class="parseTimeColor(
            driver.Sector3PersonalFastest,
            driver.Sector3OverallFastest
          )
            ">{{ parseDuration(driver.Sector3) }}</td>
          <!-- personal best -->
          <td :class="parseTimeColor(
            driver.LastLapPersonalFastest,
            driver.LastLapOverallFastest
          )
            ">{{ parseDuration(driver.LastLap) }}</td>
          <!-- drs -->
          <td :class="{
            'text-green-500': getDriverTelemetry(driver)?.DRS,
            'text-red-500': !getDriverTelemetry(driver)?.DRS,
          }">DRS</td>
          <!-- tire -->
          <td :class="parseTireColor(driver.Tire)">{{ parseTireType(driver.Tire) }}</td>
          <!-- tire lap -->
          <td>{{ driver.LapsOnTire }}</td>
          <!-- pits -->
          <td>{{ driver.Pitstops }}</td>
          <!-- last pit time -->
          <td
            v-if="useSessionStore.session === Session.Race || useSessionStore.session === HistoricalSession.RaceSession">
            {{ parseLastPitTime(driver.PitStopTimes) }}</td>
          <!-- pit position -->
          <td
            v-if="useSessionStore.session === Session.Race || useSessionStore.session === HistoricalSession.RaceSession"
            :class="getPilanePosition(driver).positionColor">

            {{ getPilanePosition(driver).potentialPositionChange }}
          </td>
          <!-- speed trap -->
          <td>{{ driver.SpeedTrap }} KM/h</td>
          <!-- Location -->
          <td>{{ getCarPosition(driver.Location) }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
