import { createRouter, createWebHistory } from 'vue-router'
import Content from '../views/Content.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/content:indexes',
      name: 'content',
      component: Content
    }
  ]
})

export default router
