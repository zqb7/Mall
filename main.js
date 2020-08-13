import Vue from 'vue'
import App from './App'
import store from './store'

Vue.config.productionTip = false
if (process.env.NODE_ENV === "development") {
  Vue.prototype.api = "http://127.0.0.1:8001"
}else {
  Vue.prototype.api = "http://gomall.sadeye.cn"
}
const msg = (title, duration=1500, mask=false, icon='none')=>{
	//统一提示方便全局修改
	if(Boolean(title) === false){
		return;
	}
	uni.showToast({
		title,
		duration,
		mask,
		icon
	});
}

App.mpType = 'app'
Vue.prototype.$store = store;
Vue.prototype.$msg = msg;

const app = new Vue({
    ...App
})
app.$mount()
