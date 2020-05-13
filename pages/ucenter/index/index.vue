<template>
	<view>
		<view class="container">
		  <view class="profile-info">
			<image @click="onUserInfoClick" class="avatar" :src="userInfo.avatar"></image>
			<view class="info">
			  <text class='name' @click='onUserInfoClick'>{{ userInfo.nickname || '点击登录' }}</text>
			  <!-- <text class='level' @click='onUserInfoClick'></text> -->
			</view>
			<image @click="onUserInfoClick" class='btn' src='/static/images/address_right.png'></image>
		  </view>
		
		  <view class="user-menu">
			<view class="item">
			  <navigator url="/pages/ucenter/order/order" class="a">
				<text class="icon order"></text>
				<text class="txt">我的订单</text>
			  </navigator>
			</view>
			<view class="item">
			  <navigator url="/pages/ucenter/coupon/coupon" class="a">
				<text class="icon coupon"></text>
				<text class="txt">优惠券</text>
			  </navigator>
			</view>
			<view class="item no-border">
			  <navigator url="url" class="a">
				<text class="icon gift"></text>
				<text class="txt">礼品卡</text>
			  </navigator>
			</view>
			<view class="item">
			  <navigator url="/pages/ucenter/collect/collect" class="a">
				<text class="icon address"></text>
				<text class="txt">我的收藏</text>
			  </navigator>
			</view>
			<view class="item">
			  <navigator url="/pages/ucenter/footprint/footprint" class="a">
				<text class="icon security"></text>
				<text class="txt">我的足迹</text>
			  </navigator>
			</view>
			<view class="item no-border">
			  <navigator url="url" class="a">
				<text class="icon kefu"></text>
				<text class="txt">会员福利</text>
			  </navigator>
			</view>
			<view class="item">
			  <navigator url="../address/address" class="a">
				<text class="icon address"></text>
				<text class="txt">地址管理</text>
			  </navigator>
			</view>
			<view class="item">
			  <navigator url="url" class="a">
				<text class="icon security"></text>
				<text class="txt">账号安全</text>
			  </navigator>
			</view>
			<view class="item no-border">
			  <navigator url="url" class="a">
				<text class="icon kefu"></text>
				<text class="txt">联系客服</text>
			  </navigator>
			</view>
			<view class="item item-bottom">
			  <navigator url="url" class="a">
				<text class="icon help"></text>
				<text class="txt">帮助中心</text>
			  </navigator>
			</view>
			<view class="item item-bottom">
			  <navigator url="/pages/ucenter/feedback/feedback" class="a">
				<text class="icon feedback"></text>
				<text class="txt">意见反馈</text>
			  </navigator>
			</view>
		  </view>
		
		  <!-- <view class="logout" @click="exitLogin">退出登录</view> -->
		</view>
		
		<view class='dialog-login' v-if="showLoginDialog" @click='onCloseLoginDialog'>
		  <view class='dialog-body' catchtap='onDialogBody'>
			<view class='title'>请选择登录方式</view>
			<view class='content'>
			  <!-- <button type="primary" open-type="getUserInfo" @getuserinfo="onWechatLogin">微信登录</button> -->
			  <!-- <button open-type="getUserInfo" lang="zh_CN" @getuserinfoo="onWechatLogin">手机号登录</button> -->
			</view>
		  </view>
		</view>
	</view>

</template>

<script>
	export default {
		data() {
			return {
				userInfo: {},
				showLoginDialog: false
			}
		},
		onLoad() {

		},
		methods: {
			onUserInfoClick(){
				
			},
			onCloseLoginDialog(){
				
			},
			onWechatLogin(){
				
			},
			exitLogin(){
				
			}
		}
	}
</script>

<style lang="scss">
	.container{
	    background: #f4f4f4;
	    height: auto;
	    overflow: hidden;
	    width: 100%;
	}
	.profile-info {
	  width: 100%;
	  height: 280upx;
	  display: flex;
	  flex-wrap: wrap;
	  align-items: center;
	  justify-content: flex-start;
	  padding: 0 30.25upx;
	  background: #333;
	}
	
	.profile-info .avatar {
	  height: 148upx;
	  width: 148upx;
	  border-radius: 50%;
	  border: 4upx solid #fff;
	}
	
	.profile-info .info {
	  flex: 1;
	  height: 85upx;
	  margin-left: 30upx;
	  display: flex;
	  flex-direction: column;
	  justify-content: center;
	}
	
	.profile-info .name {
	  height: 45upx;
	  line-height: 45upx;
	  color: #fff;
	  font-size: 37.5upx;
	}
	
	.profile-info .level {
	  height: 30upx;
	  line-height: 30upx;
	  margin-top: 10upx;
	  color: #7f7f7f;
	  font-size: 30upx;
	}
	
	.profile-info .btn {
	  width: 50upx;
	  height: 50upx;
	  margin-left: 10upx;
	  border-radius: 50%;
	}
	
	.user-menu{
	    width: 100%;
	    height: auto;
	    overflow: hidden;
	    background: #fff;
	}
	
	.user-menu .item{
	    float: left;
	    width: 33%;
	    height: 187.5upx;
	    border-right: 1px solid rgba(0,0,0,.15);
	    border-bottom: 1px solid rgba(0,0,0,.15);
	    text-align: center;
	}
	
	.user-menu .item .a{
	  display: flex;
	  width: 100%;
	  height: 100%;
	  flex-direction: column;
	  align-items: center;
	  justify-content: center;
	}
	
	.user-menu .item.no-border {
	    border-right: 0;
	}
	
	.user-menu .item.item-bottom {
	    border-bottom: none;
	}
	
	.user-menu .icon{
	    margin: 0 auto;
	    display: block;
	    height: 52.803upx;
	    width: 52.803upx;
	    margin-bottom: 16upx;
	}
	
	.user-menu .icon.order{
	    background: url(http://yanxuan.nosdn.127.net/hxm/yanxuan-wap/p/20161201/style/img/sprites/ucenter-sdf6a55ee56-f2c2b9c2f0.png) 0 -437.5upx no-repeat;
	    background-size: 52.803upx;
	}
	
	.user-menu .icon.coupon{
	    background: url(http://yanxuan.nosdn.127.net/hxm/yanxuan-wap/p/20161201/style/img/sprites/ucenter-sdf6a55ee56-f2c2b9c2f0.png) 0 -62.4997upx no-repeat;
	    background-size: 52.803upx;
	}
	
	.user-menu .icon.gift{
	    background: url(http://yanxuan.nosdn.127.net/hxm/yanxuan-wap/p/20161201/style/img/sprites/ucenter-sdf6a55ee56-f2c2b9c2f0.png) 0 -187.5upx no-repeat;
	    background-size: 52.803upx;
	}
	
	.user-menu .icon.address{
	    background: url(http://yanxuan.nosdn.127.net/hxm/yanxuan-wap/p/20161201/style/img/sprites/ucenter-sdf6a55ee56-f2c2b9c2f0.png) 0 0 no-repeat;
	    background-size: 52.803upx;
	}
	
	.user-menu .icon.security{
	    background: url(http://yanxuan.nosdn.127.net/hxm/yanxuan-wap/p/20161201/style/img/sprites/ucenter-sdf6a55ee56-f2c2b9c2f0.png) 0 -500upx no-repeat;
	    background-size: 52.803upx;
	}
	
	.user-menu .icon.kefu{
	    background: url(http://yanxuan.nosdn.127.net/hxm/yanxuan-wap/p/20161201/style/img/sprites/ucenter-sdf6a55ee56-f2c2b9c2f0.png) 0 -312.5upx no-repeat;
	    background-size: 52.803upx;
	}
	
	.user-menu .icon.help{
	    background: url(http://yanxuan.nosdn.127.net/hxm/yanxuan-wap/p/20161201/style/img/sprites/ucenter-sdf6a55ee56-f2c2b9c2f0.png) 0 -250upx no-repeat;
	    background-size: 52.803upx;
	}
	
	.user-menu .icon.feedback{
	    background: url(http://yanxuan.nosdn.127.net/hxm/yanxuan-wap/p/20161201/style/img/sprites/ucenter-sdf6a55ee56-f2c2b9c2f0.png) 0 -125upx no-repeat;
	    background-size: 52.803upx;
	}
	
	
	.user-menu .txt{
	    display: block;
	    height: 24upx;
	    width: 100%;
	    font-size: 24upx;
	    color:#333;
	}
	
	.logout{
	    margin-top: 50upx;
	    height: 101upx;
	    width: 100%;
	    line-height: 101upx;
	    text-align: center;
	    background: #fff;
	    color: #333;
	    font-size: 30upx;
	}
	
	
	.dialog-login{
	  width: 100%;
	  height: 100%;
	  position: fixed;
	  top: 0;
	  left: 0;
	  background: rgba(0, 0, 0, 0.5); 
	  display: flex;
	  justify-content: center;
	  align-items: center;
	}
	
	
	.dialog-login .dialog-body{
	  width: 540upx;
	  height: auto;
	  overflow: hidden;
	  background: #fff;
	  border-radius: 5px;
	  padding: 40upx;
	}
	
	.dialog-login .dialog-body .title{
	  width: 100%;
	  text-align: center;
	  line-height: 60upx;
	}
	
	.dialog-login .dialog-body .content{
	  width: 100%;
	}
	
	.dialog-login .dialog-body button{
	  margin-top: 30upx;
	}
</style>
