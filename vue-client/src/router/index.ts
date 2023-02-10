// Composables
import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    component: () => import('@/layouts/default/Default.vue'),
    redirect: '/home',
    children: [
      {
        path: '/home',
        name: 'Home',
        component: () => import('@/views/Home.vue'),
      },
      {
        path: '/about',
        name: 'About',
        component: () => import('@/views/About.vue'),
      },
      {
        path: '/settings',
        name: 'Settings',
        component: () => import('@/views/Settings.vue'),
      },
      {
        path: '/error',
        name: 'Error',
        component: () => import('@/views/Error.vue'),
      },
      {
        // path: "*",
        path: "/:catchAll(.*)",
        name: "NotFound",
        component: () => import("@/views/NotFound.vue"),
        // meta: {
        //   requiresAuth: false
        // }
      }
    ],
  },
  {
    path: '/n',
    component: () => import('@/layouts/nonadmin/Default.vue'),
    children: [
      {
        path: '/n/setup',
        name: 'Setup',
        component: () => import('@/views/Setup.vue'),
      },
      {
        path: '/n/error',
        name: 'AnonError',
        component: () => import('@/views/Error.vue'),
      },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})

export default router
