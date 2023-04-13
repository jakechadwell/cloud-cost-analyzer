import Vue from 'vue'
import App from './App.vue'
import router from './routes.js';
import BootstrapVue, { BootstrapVueIcons } from 'bootstrap-vue';
import 'bootstrap-icons/font/bootstrap-icons.css';


Vue.config.productionTip = false

Vue.use(BootstrapVue, BootstrapVueIcons)
new Vue({
  router,
  render: h => h(App),
}).$mount('#app')