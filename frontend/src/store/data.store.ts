import { Driver } from "@/models/driver.model";
import { Telemetry } from "@/models/telemetry.model";
import { set } from "@vueuse/core";
import { reactive, ref } from "vue";

import { F1 } from "../models/f1.model";
import { Event } from "@/models/event.model";
import { General } from "@/models/general.model";

export const dashboardData = ref<F1>({} as F1);

export const useDriverStore = reactive({
  drivers: [] as Driver[],
  increment(item: Driver) {
    const i = this.drivers.findIndex((e) => e.Number === item.Number);
    set(this.drivers, i, item);
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
