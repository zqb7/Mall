import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
	state: {
		hasLogin: false,
		userInfo: {},
		token:"",
	},
	mutations: {
		login(state, provider) {

			state.hasLogin = true;
			if (provider.hasOwnProperty("user")){
				state.userInfo = provider["user"];
				uni.setStorage({//缓存用户登陆状态
				    key: 'userInfo',  
				    data: provider["user"]  
				})
			}
			state.token = provider["token"]
			uni.setStorage({
			    key: 'token',  
			    data: provider["token"]  
			})
			uni.setStorage({
				key: "hasLogin",
				data:true
			})
			 
			console.log(state.token,state.hasLogin);
		},
		logout(state) {
			state.hasLogin = false;
			state.userInfo = {};
			uni.removeStorage({  
                key: 'userInfo'  
            })
			uni.removeStorage({
				key:"token"
			})
		}
	},
	actions: {
	
	}
})

export default store
