<template>
	<view class="container">
	  <view class="coupon-form">
	    <view class="input-box">
	      <input class="coupon-sn" placeholder="请输入优惠码" />
	      <image class="clear-icon" src="http://yanxuan.nosdn.127.net/hxm/yanxuan-wap/p/20161201/style/img/icon-normal/clear-fb-dd9d604f86.png"></image>
	    </view>
	    <button class="add-btn disabled">兑换</button>
	  </view>
	  <view class="help">使用说明</view>
	  <view class="coupon-list">
	    <view class="item">
	      <view class="tag">新人专享</view>
	      <view class="content">
	        <view class="left">
	          <view class="name">限时免单券</view>
	          <view class="time">2017.06.08-2017.06.11</view>
	        </view>
	        <view class="right">
	          <button class="go">去使用</button>
	        </view>
	      </view>
	      <view class="condition">
	        <text class="txt">简约陶瓷马克杯专享；小米用户福利；限时购、三石福利价、礼品卡及详情页标注不可用券特殊商品除外</text>
	        <image src="https://yanxuan.nosdn.127.net/hxm/yanxuan-wap/p/20161201/style/img/icon-normal/couponArrow-71315b4145.png" class="icon"></image>
	      </view>
	    </view>
	
	     <view class="item">
	      <view class="tag">新人专享</view>
	      <view class="content">
	        <view class="left">
	          <view class="name">限时免单券</view>
	          <view class="time">2017.06.08-2017.06.11</view>
	        </view>
	        <view class="right">
	          <button class="go">去使用</button>
	        </view>
	      </view>
	      <view class="condition">
	        <text class="txt">简约陶瓷马克杯专享；小米用户福利；限时购、三石福利价、礼品卡及详情页标注不可用券特殊商品除外</text>
	        <image src="https://yanxuan.nosdn.127.net/hxm/yanxuan-wap/p/20161201/style/img/icon-normal/couponArrow-71315b4145.png" class="icon"></image>
	      </view>
	    </view>
	  </view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				
			}
		},
		methods: {
			
		}
	}
</script>

<style lang="scss">
	page{
	    background: #f4f4f4;
	    min-height: 100%;
	}
	
	.container{
	    background: #f4f4f4;
	    min-height: 100%;
	    padding-top: 30upx;
	}
	
	.coupon-form{
	  height: 110upx;
	  width: 100%;
	  background: #fff;
	  padding-left: 30upx;
	  padding-right: 30upx;
	  padding-top: 20upx;
	  display: flex;
	}
	
	.input-box{
	  flex: 1;
	  height: 70upx;
	  color: #333;
	  font-size: 24upx;
	  background: #fff;
	  position: relative;
	  border: 1px solid rgba(0, 0, 0, 0.15);
	  border-radius: 4upx;
	  margin-right: 30upx;
	}
	
	.input-box .coupon-sn{
	  position: absolute;
	  top: 10upx;
	  left: 30upx;
	  height: 50upx;
	  width: 100%;
	  color: #000;
	  line-height: 50upx;
	  font-size: 24upx;
	}
	
	.clear-icon{
	  position: absolute;
	  top: 21upx;
	  right: 30upx;
	  width: 28upx;
	  height: 28upx;
	}
	
	.add-btn{
	  height: 70upx;
	  border:none;
	  width: 168upx;
	  background: #b4282d;
	  border-radius: 0;
	  line-height: 70upx;
	  color: #fff;
	  font-size: 28upx;
	}
	
	.add-btn.disabled{
	  background: #ccc;
	}
	
	.help{
	  height: 72upx;
	  line-height: 72upx;
	  text-align: right;
	  padding-right: 30upx;
	  background: url(https://yanxuan.nosdn.127.net/hxm/yanxuan-wap/p/20161201/style/img/icon-normal/couponHelp-4768607555.png) 590upx center no-repeat;
	  background-size: 28upx;
	  color: #999;
	  font-size: 24upx;
	}
	
	.coupon-list{
	  width: 90%;
	  height: auto;
	  overflow: hidden;
	  padding-left: 30upx;
	}
	
	.item{
	  position: relative;
	  height: 290upx;
	  width: 100%;
	  background: linear-gradient(to right,#cfa568,#e3bf79);
	  margin-bottom: 30upx;
	  padding-top: 52upx;
	}
	
	.tag{
	  height: 32upx;
	  background: #A48143;
	  padding-left: 16upx;
	  padding-right: 16upx;
	  position: absolute;
	  left: 20upx;
	  color: #fff;
	  top: 20upx;
	  font-size: 20upx;
	  text-align: center;
	  line-height: 32upx;
	}
	
	
	.content{
	  margin-top: 24upx;
	  margin-left: 40upx;
	  display: flex;
	  margin-right: 40upx;
	  flex-direction: row;
	  align-items: center;
	}
	
	.content .left{
	  flex: 1;
	}
	
	.name{
	  font-size: 44upx;
	  color: #fff;
	  margin-bottom: 14upx;
	}
	
	.time{
	  font-size: 24upx;
	  color: rgba(255,255,255, 0.8);
	  line-height: 30upx;
	}
	
	.content .right{
	  width: 162upx;
	}
	
	.go{
	  height: 48upx;
	  border:none;
	  width: 162upx;
	  background: rgba(255,255,255, 0.8);
	  border-radius: 4upx;
	  line-height: 48upx;
	  color: #b69150;
	  font-size: 24upx;
	}
	
	.condition{
	  position: absolute;
	  width: 100%;
	  bottom: 0;
	  left:0;
	  height: 78upx;
	  background: rgba(0,0,0,.08);
	  padding: 24upx 40upx;
	  display: flex;
	  flex-direction: row;
	}
	
	.condition .txt{
	  display: block;
	  height: 30upx;
	  flex: 1;
	  overflow: hidden;
	  font-size: 24upx;
	  line-height: 30upx;
	  color: #fff;
	}
	
	
	.condition .icon{
	  margin-left: 30upx;
	  width: 24upx;
	  height: 24upx;
	}
</style>
