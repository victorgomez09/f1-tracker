import { reactive, ref } from "vue";
import { F1 } from "../models/f1.model";
import { F1Driver } from "@/models/driver.model";

export const dashboardData = ref<F1>({} as F1);

export const driversStore = reactive({
  data: [] as F1Driver[],
  increment(item: F1Driver) {
    const i = this.data.findIndex((e) => e.Number === item.Number);
    if (i > -1) this.data[i] = item; // (2)
    else this.data.push(item);
  },
});
