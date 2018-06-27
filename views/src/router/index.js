import Vue from 'vue'
import Router from 'vue-router'
import Index from '@/components/Index'
import Edit from '@/components/Edit'
import Show from '@/components/Show'

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
      component: Edit,
      props: true
    },
    {
      path: '/:id',
      name: 'show',
      component: Show,
      props: true
    }
  ]
})
