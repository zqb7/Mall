<template>
	<view class="content">
		<view class="row b-b">
			<text class="tit">联系人</text>
			<input class="input" type="text" v-model="addressData.consignee" placeholder="收货人姓名" placeholder-class="placeholder" />
		</view>
		<view class="row b-b">
			<text class="tit">手机号</text>
			<input class="input" type="number" v-model="addressData.mobile" placeholder="收货人手机号码" placeholder-class="placeholder" />
		</view>
		<view class="row b-b"  @click="handleTap('picker')">
			<text class="tit">收货地址</text>
			<input class="input" type="text" v-model="address"  placeholder="收请选择收货地址" placeholder-class="placeholder" />
			<view>
			  <view placeholder="请选择地址"></view>
			  <lb-picker ref="picker"
			    v-model="value"
			    mode="multiSelector"
			    :list="list"
			    :level="3"
				:props="myProps"
			    @change="handleChange"
			    @confirm="handleConfirm"
			    @cancel="handleCancel">
			  </lb-picker>
			</view>
		</view>
		<view class="row b-b"> 
			<text class="tit">详细地址</text>
			<input class="input" type="text" v-model="addressData.address" placeholder="如小区、街道" placeholder-class="placeholder" />
		</view>
		
		<view class="row default-row">
			<text class="tit">设为默认</text>
			<switch :checked="addressData.is_default" color="#fa436a" @change="switchChange" />
		</view>
		<button class="add-btn" @click="confirm">提交</button>
	</view>
</template>

<script>
	import LbPicker from '@/components/lb-picker'
	import {
		mapState
	} from 'vuex'
	export default {
		components: {
		      LbPicker
		},
		computed:{
			...mapState(['token'])
		},
		data() {
			return {
				addressData: {
					consignee: '',
					country: "中国",
					province:"",
					city:'',
					district:  '',
					mobile: '',
					is_default: false,
				},
				myProps:{
					label: 'name',
				    value: 'id',
				    children: 'child'
				},
				address: "",
				value: [],
				list: [],
				manageType: "add"
			}
		},
		onReady () {
		},
		onLoad(option){
			let title = '新增收货地址';
			if(option.type==='edit'){
				title = '编辑收货地址'
				this.addressData = JSON.parse(option.data)
				console.log(this.addressData)
				let item = this.addressData
				this.address = `${item.country} ${item.province} ${item.city} ${item.district} ${item.address}`
			}
			this.manageType = option.type;
			uni.setNavigationBarTitle({
				title
			})
		},
		methods: {
			switchChange(e){
				this.addressData.is_default = e.detail.value;
			},
			
			//提交
			confirm(){
				let data = this.addressData;
				if(!data.consignee){
					this.$msg('请填写收货人姓名');
					return;
				}
				if(!/(^1[3|4|5|7|8][0-9]{9}$)/.test(data.mobile)){
					this.$msg('请输入正确的手机号码');
					return;
				}
				if(!data.address){
					this.$msg('请填写详细地址');
					return;
				}
				this.addressData.country = "中国"
				let api = ""
				let method = "POST"
				switch (this.manageType){
					case "edit":
						api = this.api + `/user/address/${this.addressData.id}`
						method = "PUT"
						break
					case "add":
						api = this.api + "/user/address"
						break
					default:
						return
				}
				uni.request({
					url:api,
					method:method,
					data:JSON.stringify(this.addressData),
					header:{
						'token': this.token,
					},
					success: (res) => {
						if (res.statusCode == 200 ){
							uni.navigateBack()
						}else {
							this.$msg("添加失败")
						}
					}
				})
			},
			handleTap (picker) {
				uni.request({
					url:this.api + "/region" + `?parent_id=1`,
					success: (res) => {
						this.list = this.list.concat(res.data)
					}
				})
				this.$refs[picker].show()
			},
			handleChange (item) {
				let itemLength = item.index.length
				let valueLength = item.value.length
				let valueLast = 0
				for(let i = 0,len=valueLength; i < len; i++) {
					if (item.value[i] != 0){
						valueLast = item.value[i]
					}else{
						break
					}
				}
				switch (itemLength){
					case 1:
						if (this.list[item.index[0]].hasOwnProperty('child')){
							return
						}
						break
					case 2:
						if (this.list[item.index[0]].hasOwnProperty('child')){
							if (this.list[item.index[0]].child[item.index[1]].hasOwnProperty('child')){
								return
							}
						}
						break
					case 3:
						// 第三级 滚动不请求数据，直接返回
						if (item.index[2] != 0){
							return
						}
						break
				}
				uni.request({
					url:this.api + "/region" + `?parent_id=${valueLast}`,
					success: (res) => {
						if (res.statusCode == 200 && res.data.length != 0 ){
							switch (itemLength){
								case 0:
									break
								case 1:
									// this.list[item.index[0]].child = res.data
									this.$set(this.list[item.index[0]],'child',res.data)
									break
								case 2:
									// this.list[item.index[0]].child[item.index[1]].child = res.data
									this.$set(this.list[item.index[0]].child[item.index[1]],'child',res.data)
									break
							}
						}
					}
				})
			},
			handleConfirm (item) {
				let valueLengthIndex = item.value.length -1
				switch (valueLengthIndex) {
					case 0:
						this.addressData['province'] = this.list[item.index[0]].name
						break
					case 1:
						this.addressData['province'] = this.list[item.index[0]].name
						this.addressData['city'] = this.list[item.index[0]].child[item.index[1]].name
						break
					case 2:
						this.addressData['province'] = this.list[item.index[0]].name
						this.addressData['city'] = this.list[item.index[0]].child[item.index[1]].name
						this.addressData['district'] = this.list[item.index[0]].child[item.index[1]].child[item.index[2]].name
						break
				}
				this.address  = this.addressData.province + " " + this.addressData.city + " "+ this.addressData.district
			},
			handleCancel (item) {
				console.log('cancel::', item)
			}
		}
	}
</script>

<style lang="scss">
	page{
		background: $page-color-base;
		padding-top: 16upx;
	}

	.row{
		display: flex;
		align-items: center;
		position: relative;
		padding:0 30upx;
		height: 110upx;
		background: #fff;
		
		.tit{
			flex-shrink: 0;
			width: 120upx;
			font-size: 30upx;
			color: $font-color-dark;
		}
		.input{
			flex: 1;
			font-size: 30upx;
			color: $font-color-dark;
		}
		.icon-shouhuodizhi{
			font-size: 36upx;
			color: $font-color-light;
		}
	}
	.default-row{
		margin-top: 16upx;
		.tit{
			flex: 1;
		}
		switch{
			transform: translateX(16upx) scale(.9);
		}
	}
	.add-btn{
		display: flex;
		align-items: center;
		justify-content: center;
		width: 690upx;
		height: 80upx;
		margin: 60upx auto;
		font-size: $font-lg;
		color: #fff;
		background-color: $base-color;
		border-radius: 10upx;
		box-shadow: 1px 2px 5px rgba(219, 63, 96, 0.4);
	}
</style>
