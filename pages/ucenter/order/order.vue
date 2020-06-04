<template>
	<view class="container">
	    <view class="orders">
	        <navigator class="order" v-for="(item, key) in orderList">
	            <view class="h">
	                <view class="l">订单编号：{{item.order_sn}}</view>
	                <view class="r">{{item.order_status_text}}</view>
	            </view>
	            <view class="goods" v-for="(gitem, _) in item.goodsList">
	                <view class="img">
	                    <image :src="gitem.list_pic_url"></image>
	                </view>
	                <view class="info">
	                    <text class="name">{{gitem.goods_name}}</text>
	                    <text class="number">共{{gitem.number}}件商品</text>
	                </view>
	                <view class="status"></view>
	            </view>
	            <view class="b">
	                <view class="l">实付：￥{{item.actual_price}}</view>
	                <view class="r">
	                    <button class="btn" catchtap="payOrder" v-if="item.handleOption.pay">去付款</button>
	                </view>
	            </view>
	        </navigator>
	    </view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				orderList: [
					{
						"order_sn": "1",
						"order_status_text": "ttt",
						"actual_price": "13.14",
						"handleOption": {
							pay: true,
						},
						goodsList: [
							{
								"list_pic_url": "http://yanxuan.nosdn.127.net/655d718df8107f8e7fd1dc6140e29039.png",
								"goods_name": "test",
								"number": 0
							}
						]
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
	
	.orders{
	    height: auto;
	    width: 100%;
	    overflow: hidden;
	}
	
	.order{
	    margin-top: 20rpx;
	    background: #fff;
	}
	
	.order .h{
	    height: 83.3rpx;
	    line-height: 83.3rpx;
	    margin-left: 31.25rpx;
	    padding-right: 31.25rpx;
	    border-bottom: 1px solid #f4f4f4;
	    font-size: 30rpx;
	    color: #333;
	}
	
	.order .h .l{
	    float: left;
	}
	
	.order .h .r{
	    float: right;
	    color: #b4282d;
	    font-size: 24rpx;
	}
	
	.order .goods{
	    display: flex;
	    align-items: center;
	    height: 199rpx;
	    margin-left: 31.25rpx;
	}
	
	.order .goods .img{
	    height:145.83rpx;
	    width:145.83rpx;
	    background: #f4f4f4;
	}
	
	.order .goods .img image{
	    height:145.83rpx;
	    width:145.83rpx;
	}
	
	.order .goods .info{
	    height: 145.83rpx;
	    flex: 1;
	    padding-left: 20rpx;
	}
	
	.order .goods .name{
	    margin-top: 30rpx;
	    display: block;
	    height: 44rpx;
	    line-height: 44rpx;
	    color: #333;
	    font-size: 30rpx;
	}
	
	.order .goods .number{
	    display: block;
	    height: 37rpx;
	    line-height: 37rpx;
	    color: #666;
	    font-size: 25rpx;
	}
	
	.order .goods .status{
	    width:105rpx;
	    color: #b4282d;
	    font-size: 25rpx;
	}
	
	.order .b{
	    height: 103rpx;
	    line-height: 103rpx;
	    margin-left: 31.25rpx;
	    padding-right: 31.25rpx;
	    border-top: 1px solid #f4f4f4;
	    font-size: 30rpx;
	    color: #333;
	}
	
	.order .b .l{
	    float: left;
	}
	
	.order .b .r{
	    float: right;
	}
	
	.order .b .btn{
	    margin-top: 19rpx;
	    height: 64.5rpx;
	    line-height: 64.5rpx;
	    text-align: center;
	    padding: 0 20rpx;
	    border-radius: 5rpx;
	    font-size: 28rpx;
	    color: #fff;
	    background: #b4282d;
	}
</style>
