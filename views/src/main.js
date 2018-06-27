// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import editor from 'vue2-medium-editor'
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

Vue.use(BootstrapVue);

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
