import { createApp } from "vue";
import { MotionPlugin } from "@vueuse/motion";

import "./style.css";
import routes from "./routes";
import App from "./App.vue";

const app = createApp(App);
app.use(MotionPlugin);
app.use(routes);

app.mount("#app");
