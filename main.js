import Vue from 'vue'
import App from './App'
import store from './store'

Vue.config.productionTip = false
if (process.env.NODE_ENV === "development") {
  Vue.prototype.api = "http://127.0.0.1:8001"
}else {
  Vue.prototype.api = "http://gomall.sadeye.cn"
}

App.mpType = 'app'
Vue.prototype.$store = store;

const app = new Vue({
    ...App
})
app.$mount()
