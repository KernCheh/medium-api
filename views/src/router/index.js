import Vue from 'vue'
import Router from 'vue-router'
import Index from '@/components/Index'
import Edit from '@/components/Edit'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'index',
      component: Index
    },
    {
      path: '/new',
      name: 'new',
      component: Edit
    },
    {
      path: '/:id/edit',
      name: 'edit',
      component: Edit
    }
  ]
})
