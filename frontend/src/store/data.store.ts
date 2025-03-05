import { Driver } from "@/models/driver.model";
import { Telemetry } from "@/models/telemetry.model";
import { set } from "@vueuse/core";
import { reactive } from "vue";

import { Event } from "@/models/event.model";
import { General } from "@/models/general.model";
import { Information } from "@/models/information.model";
import { Time } from "@/models/time.model";
import { RaceControl } from "@/models/race-control.model";

export const useDriverStore = reactive({
  drivers: [] as Driver[],
  increment(item: Driver) {
    const i = this.drivers.findIndex((e) => e.Number === item.Number);
    set(this.drivers, i, item);
  },
});

export const useTimeStore = reactive({
  time: {} as Time,
  addTime(item: Time) {
    this.time = item;
  },
});

export const useInformationStore = reactive({
  information: {} as Information,
  addInformation(item: string) {
    this.information.CircuitTimezone = item
  },
});

export const useTelemetryStore = reactive({
  telemetry: [] as Telemetry[],
  addTelemetry(item: Telemetry) {
    const i = this.telemetry.findIndex(
      (e) => e.DriverNumber === item.DriverNumber
    );
    if (i === -1) {
      this.telemetry.push(item);
    }
    set(this.telemetry, i, item);
  },
});

export const useEventStore = reactive({
  event: {} as Event,
  addEvent(item: Event) {
    // set(this.event, item);
    this.event = item;
  },
});

export const useSessionStore = reactive({
  session: 0 as number,
  addSession(item: number) {
    this.session = item;
  },
});

export const useGeneralStore = reactive({
  general: {} as General,
  addGeneral(item: General) {
    this.general = item;
  },
});

export const usePausedStore = reactive({
  paused: false,
  setPaused(paused: boolean) {
    this.paused = paused;
  },
});

export const useRaceControlStore = reactive({
  raceControl: [] as RaceControl[],
  addRaceControl(raceControl: RaceControl) {
    this.raceControl.push(raceControl)
  },
});
