import { createRouter, createWebHashHistory } from "vue-router";

const router = createRouter({
  // 4. Provide the history implementation to use. We
  // are using the hash history for simplicity here.
  history: createWebHashHistory(),
  routes: [
    {
      path: "/",
      name: "F1 Live Timming",
      component: () => import("../views/Dashboard.vue"),
    },
  ], // short for `routes: routes`
});

export default router;
