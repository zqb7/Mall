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
</style>
