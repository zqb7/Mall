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

const checkHttpCode = (code) => {
	switch (code){
		case 200:
			break;
		case 401:
			uni.navigateTo({
				url:"/pages/public/login",
				success: (res) => {
				}
			})
			break;
	}
}

App.mpType = 'app'
Vue.prototype.$store = store;
Vue.prototype.$msg = msg;
Vue.prototype.$checkHttpCode = checkHttpCode;

const app = new Vue({
    ...App
})
app.$mount()
