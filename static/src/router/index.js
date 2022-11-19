import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import NotFound from "../views/NotFound.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/browse/:category',
      name: 'category',
      props: true,

      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/Category.vue')
    },
    {
      path: '/book/:isbn',
      name: 'book',
      props: true,

      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/Book.vue')
    },
    {
      path: '/receipt/:id',
      name: 'receipt',
      props: true,

      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/Receipt.vue')
    },
    {
      path: '/cart',
      name: 'cart',
      props: true,

      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/Cart.vue')
    },
    {
      path: '/search',
      name: 'search',
      props: true,

      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/Search.vue')
    },
    {path:'/privacy', component: () => import('../views/Privacy.vue')},
    {path:'/contact', component: () => import('../views/Contact.vue')},
    {path:'/about', component: () => import('../views/About.vue')},
    { path: '/:pathMatch(.*)*', component: NotFound },

  ]
})

export default router
