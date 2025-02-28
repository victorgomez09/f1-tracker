import { MotionPlugin } from "@vueuse/motion";
import { createApp } from "vue";
import { createPinia } from "pinia";

import App from "./App.vue";
import routes from "./routes";
import "./style.css";

const pinia = createPinia();
const app = createApp(App);
app.use(pinia);
app.use(MotionPlugin);
app.use(routes);

app.mount("#app");
