// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import editor from 'vue2-medium-editor'

Vue.config.productionTip = false

/* eslint-disable no-new */

Vue.component('medium-editor', editor)
new Vue({
  el: '#app',
  router,
  components: {
    App
  },
  template: '<App/>'
})
