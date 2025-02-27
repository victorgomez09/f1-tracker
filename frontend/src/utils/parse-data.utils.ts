import { F1Driver } from "@/models/driver.model";
import { ServerResponse } from "@/models/server.model";
import { driversStore } from "@/store/data.store";

export const parseData = (data: ServerResponse) => {
  console.log("type", data.dataType);
  switch (data.dataType) {
    case "TIMING":
      const driver = data.data as F1Driver;
      driversStore.increment(driver);
      break;

    default:
      break;
  }
};
