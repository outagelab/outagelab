import { createRouter, createWebHistory } from 'vue-router'
import ApplicationListView from '../views/ApplicationListView.vue'
import ApplicationDetailsView from '../views/ApplicationDetailsView.vue'
import ApiKeyView from '@/views/ApiKeyView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: { path: 'applications' }
    },
    {
      path: '/applications',
      name: 'application-list',
      component: ApplicationListView
    },
    {
      path: '/applications/details/:id',
      name: 'application-details',
      component: ApplicationDetailsView
    },
    {
      path: '/api-keys',
      name: 'api-keys-list',
      component: ApiKeyView
    }
  ]
})

export default router
