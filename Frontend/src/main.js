import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './assets/css/style.less'
import './assets/font/iconfont.css'
require('./Mock')
import {parseTime} from './utils'
import VueCoreVideoPlayer from 'vue-core-video-player'
import HLSCore from '@core-player/playcore-hls'
import VuePlyr from 'vue-plyr'
import 'vue-plyr/dist/vue-plyr.css'

Vue.use(VueCoreVideoPlayer)
Vue.use(VuePlyr, {
  plyr: {}
})

Vue.config.productionTip = false
Vue.filter('parseTime', (v) => parseTime(v,'{y}-{m}-{d}'))
new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
