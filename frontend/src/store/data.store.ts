import { reactive } from "vue";
import { F1 } from "../models/f1.model";

export const dashboardData = reactive<{ data: F1 }>({ data: {} as F1 });
