<template>
	<view class="content">
		<view class="search">
			<input class="search-input" type="text" placeholder="输入关键字搜索" />
		</view>
		
		<!-- 轮播图 -->
		<view>
			<swiper  class="banner"  :indicator-dots="indicatorDots" :autoplay="autoplay" :interval="interval" :duration="duration">
				<swiper-item v-for="item in carouseList">
					<image :src="item.src" />
				</swiper-item>
			</swiper>
		</view>
		
		<!-- 分类 -->
		<view class="channel">
			<view class="channel-item" v-for="item in channelList">
				<image :src="item.img_src" />
				<text>{{item.name}}</text>
			</view>
		</view>
		
		<!-- 品牌 -->
		<view class="brand">
			<view  class="h-text">
				<text>品牌制造商直供</text>
			</view>
			<view class="b">
				<view class="brand-item" v-for="item in brandList">
					<view class="wrap">
						<image :src="item.img_src"></image>
						<view class="mt">
							<text class="name">{{item.name}}</text>
							<text class="price">{{item.price}}元起</text>
						</view>
					</view>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				title: 'Hello',
				indicatorDots: true,
				autoplay: true,
				interval: 2000,
				duration: 500,
				carouseList: [
					{
						"src": "http://yanxuan.nosdn.127.net/65091eebc48899298171c2eb6696fe27.jpg"
					},
					{
						"src": "http://yanxuan.nosdn.127.net/bff2e49136fcef1fd829f5036e07f116.jpg"
					},{
						"src": "http://yanxuan.nosdn.127.net/8e50c65fda145e6dd1bf4fb7ee0fcecc.jpg"
					}
				],
				channelList: [
					{
						"name": "居家",
						"img_src": "/static/images/channel/c1.png"
					},{
						"name": "餐厨",
						"img_src": "/static/images/channel/c2.png"
					},{
						"name": "配件",
						"img_src": "/static/images/channel/c3.png"
					}
				],
				brandList: [
					{
						"name": "CK制造商",
						"price": "27",
						"img_src": "http://yanxuan.nosdn.127.net/fb139c9f11a10deb41beeb42aaad488e.png"
					},
					{
						"name": "MUJI制造商",
						"price": "12.9",
						"img_src": "http://yanxuan.nosdn.127.net/1541445967645114dd75f6b0edc4762d.png"
					},
					{
						"name": "WMF制造商",
						"price": "9.9",
						"img_src": "http://yanxuan.nosdn.127.net/2a2ae0d49043e525db8d0d8d84460ac8.png"
					},
					{
						"name": "Coach制造商",
						"price": "49",
						"img_src": "http://yanxuan.nosdn.127.net/c933a662bb79b2a47280363e35ab994b.png"
					}
				]
			}
		},
		onLoad() {

		},
		methods: {

		}
	}
</script>

<style lang="scss">
	.search {
		top:30upx;
		width: 100%;
		margin: 20upx 0;
		.search-input {
			background: rgba(255,255,255,.6);
			text-align: center;
			font-size: 30upx;
			border-radius: 20upx;
		}
	};
	.banner {
		height: 417rpx;
		uni-image {
			width: 100%;
		}
	}
	.channel {
		display: flex;
		align-items: center;
		flex-wrap: wrap;
		justify-content: space-around;
		.channel-item {
			display: flex;
			flex-direction: column;
			align-items: center;
		}
		uni-image {
			width: 58upx;
			height: 58upx;
			border-radius: 50%;
			margin-bottom: 10upx;
		}
	}
	.h-text {
		display: flex;
		align-items: center;
		justify-content: center;
		height: 130rpx;
		uni-text{
			font-size: 36upx;
			line-height: 36upx;
			height: 36upx;
		}
	}
	.brand {
		.b {
			width: 750upx;
			height: auto;
			overflow: hidden;
			position: relative;
		}
		.wrap {
			position: relative;
			uni-image {
				position: absolute;
				left:0;
				top:0px;
			};
			.mt {
				position: absolute;
				left:0;
				top:0;
				z-index: 2;
				padding: 27upx 31upx;
			}
			.price {
			   font-size: 25upx;
			   color: #999;
			}
			.name {
				display: block;
			    font-size: 33rpx;
			    height: 43rpx;
			    color: #333;
			}
		}
		.brand-item {
		  float: left;
		  width: 375upx;
		  height: 252upx;
		  overflow: hidden;
		  border-top: 1upx solid #fff;
		  margin-left: urpx;
		}
	}

</style>
