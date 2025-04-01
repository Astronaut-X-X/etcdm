import { createMemoryHistory, createRouter } from 'vue-router';
import type { RouteRecordRaw } from 'vue-router';

// components
import Connection from './components/connection/connection.vue';
import Data from './components/data/data.vue';
import Setting from './components/setting/setting.vue';

const routes = [
  { path: '/', redirect: '/data' },
  { path: '/connection', component: Connection },
  { path: '/data', component: Data },
  { path: '/setting', component: Setting },
] as RouteRecordRaw[];

const router = createRouter({
  history: createMemoryHistory(),
  routes,
})

export default router;