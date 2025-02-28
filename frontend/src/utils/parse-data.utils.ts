import { ServerResponse } from "@/models/server.model";
import {
  useDriverStore,
  useEventStore,
  useGeneralStore,
  useSessionStore,
  useTelemetryStore,
} from "@/store/data.store";

export const parseData = (data: ServerResponse) => {
  switch (data.dataType) {
    case "TIMING":
      useDriverStore.increment(data.data);

      break;

    case "DRIVERS":
      useDriverStore.drivers = data.data.Drivers;

      break;

    case "EVENT":
      useEventStore.addEvent(data.data);

      break;

    case "TELEMETRY":
      useTelemetryStore.addTelemetry(data.data);

      break;

    case "SESSION":
      useSessionStore.addSession(data.data);

      break;

    case "GENERAL":
      useGeneralStore.addGeneral(data.data);

      break;

    default:
      break;
  }
};
