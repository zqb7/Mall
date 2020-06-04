<template>
	<view class="container">
	    <view class="address-list" v-if="addressList.length > 0 ">
	        <view class="item" v-for="item in addressList">
	            <view class="l">
	                <view class="name">{{item.name}}</view>
	                <view class="default" v-if="item.is_default">默认</view>
	            </view>
	            <view class="c">
	                <view class="mobile">{{item.mobile}}</view>
	                <view class="address">{{item.full_region + item.address}}</view>
	            </view>
	            <view class="r">
	                <image catchtap="deleteAddress"  class="del" src="/static/images/ucenter/del-address.png"></image>
	            </view>
	        </view>
	    </view>
	    <view class="empty-view" v-if="addressList.length <= 0">
	      <image class="icon" src="http://yanxuan.nosdn.127.net/hxm/yanxuan-wap/p/20161201/style/img/icon-normal/noAddress-26d570cefa.png"></image>
	      <text class="text">收货地址在哪里</text>
	    </view>
	    <view class="add-address">新建</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				addressList: [
					{
						id:1,
						name: "test",
						is_default: true,
						mobile: 1561111111,
						full_region: 1,
						adress: "address"
					}
				]
			}
		},
		methods: {
			
		}
	}
</script>

<style lang="scss">
	page{
	    height: 100%;
	    width: 100%;
	    background: #f4f4f4;
	}
	
	.container{
	  height: 100%;
	  width: 100%;
	}
	
	.address-list{
	    padding-left: 31.25upx;
	    background: #fff url(http://yanxuan.nosdn.127.net/hxm/yanxuan-wap/p/20161201/style/img/icon-normal/address-bg-bd30f2bfeb.png) 0 0 repeat-x;
	    background-size: auto 10.5upx;
	    margin-bottom: 90upx;
	}
	
	.address-list .item{
	    height: 156.55upx;
	    align-items: center;
	    display: flex;
	    border-bottom: 1upx solid #DCD9D9;
	}
	
	.address-list .l{
	    width: 125upx;
	    height: 80upx;
	    overflow: hidden;
	}
	
	.address-list .name{
	    width: 125upx;
	    height: 43upx;
	    font-size: 29upx;
	    color: #333;
	    margin-bottom: 5.2upx;
	        text-overflow: ellipsis;
	    white-space: nowrap;
	    overflow: hidden;
	}
	
	.address-list .default{
	    width: 62.5upx;
	    height: 33upx;
	    line-height: 28upx;
	    text-align: center;
	    font-size: 20upx;
	    color: #b4282d;
	    border: 1upx solid #b4282d;
	    visibility: visible;
	}
	
	
	.address-list .c{
	    flex: 1;
	    height: auto;
	    overflow: hidden;
	}
	
	.address-list .mobile{
	
	    height: 29upx;
	    font-size: 29upx;
	    line-height: 29upx;
	    overflow: hidden;
	    color: #333;
	    margin-bottom: 6.25upx;
	}
	
	.address-list .address{
	    height: 37upx;
	    font-size: 25upx;
	    line-height: 37upx;
	    overflow: hidden;
	    color: #666;
	}
	
	.address-list .r{
	    width: 52upx;
	    height: auto;
	    overflow: hidden;
	    margin-right: 16.5upx;
	}
	
	.address-list .del{
	    display: block;
	    width: 52upx;
	    height: 52upx;
	}
	
	.add-address{
	    background: #b4282d;
	    text-align: center;
	    width: 100%;
	    height: 99upx;
	    line-height: 99upx;
	    position: fixed;
	    border-radius: 0;
	    border: none;
	    color: #fff;
	    font-size: 29upx;
	    bottom: 0;
	    left:0;
	}
	
	.empty-view{
	  height: 100%;
	  width: 100%;
	  display: flex;
	  flex-direction: column;
	  align-items: center;
	  justify-content: center;
	}
	
	.empty-view .icon{
	  height: 248upx;
	  width: 258upx;
	  margin-bottom: 10upx;
	}
	
	.empty-view .text{
	  width: auto;
	  font-size: 28upx;
	  line-height: 35upx;
	  color: #999;
	}
</style>
