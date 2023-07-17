import { createWebHistory, createRouter } from "vue-router";
import { RouteRecordRaw } from "vue-router";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    alias: "/events",
    name: "events",
    component: () => import("./components/EventList.vue"),
  },
  {
    path: "/events/:id",
    name: "event-details",
    component: () => import("./components/EventDetails.vue"),
  },
  {
    path: "/add",
    name: "add",
    component: () => import("./components/AddEvent.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
