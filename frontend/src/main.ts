import { createApp } from "vue";

import "./style.css";
import routes from "./routes";
import App from "./App.vue";

const app = createApp(App);
app.use(routes);

app.mount("#app");
