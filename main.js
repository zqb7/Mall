import Vue from 'vue'
import App from './App'

Vue.config.productionTip = false
Vue.prototype.api = "http://127.0.0.1:8001"

App.mpType = 'app'

const app = new Vue({
    ...App
})
app.$mount()
