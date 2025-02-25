import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    allowedHosts: [
      "5173-victorgomez09-f1tracker-a7zkp3fi4k8.ws-eu118.gitpod.io",
    ],
  },
});
