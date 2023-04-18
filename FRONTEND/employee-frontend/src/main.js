import Vue from 'vue'
import App from './App.vue'
import router from './routes.js';
import BootstrapVue, { BootstrapVueIcons } from 'bootstrap-vue';
/* import the fontawesome core */
import { library } from '@fortawesome/fontawesome-svg-core'

/* import font awesome icon component */
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

/* import specific icons */
import { faRightLong } from '@fortawesome/free-solid-svg-icons'

/* add icons to the library */
library.add(faRightLong)

/* add font awesome icon component */
Vue.component('font-awesome-icon', FontAwesomeIcon)

Vue.config.productionTip = false

Vue.use(BootstrapVue, BootstrapVueIcons)
new Vue({
  router,
  render: h => h(App),
}).$mount('#app')