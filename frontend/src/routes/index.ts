import { createRouter, createWebHistory } from "vue-router";

const router = createRouter({
  // 4. Provide the history implementation to use. We
  // are using the hash history for simplicity here.
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      component: () => import("../views/Landing.vue"),
    },
    {
      path: "/dashboard",
      component: () => import("../views/Dashboard.vue"),
    },
    {
      path: "/telemetry",
      component: () => import("../views/Telemetry.vue"),
    },
  ], // short for `routes: routes`
});

export default router;
