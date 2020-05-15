<template>
	<view class="container">
	    <scroll-view class="topic-list" scroll-y="true" :scroll-top="scrollTop">
	        <navigator class="item" v-for="(item, key) in topicList" :url="'../topicDetail/topicDetail?id=' + item.id">
	            <image class="img" :src="item.scene_pic_url"></image>
	            <view class="info">
	                <text class="title">{{item.title}}</text>
	                <text class="desc">{{item.subtitle}}</text>
	                <text class="price">{{item.price_info}}元起</text>
	            </view>
	        </navigator>
	        <view class="page" v-if="showPage">
	            <view :class="'prev' + {disabled: page <= 1}" @click="prevPage">上一页</view>
	            <view :class="'next' + {disabled: (count / size) < page +1}" @click="nextPage">下一页</view>
	        </view>
	    </scroll-view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				tscrollTop: 0,
				showPage: true,
				page: 1,
				topicList: [
						{
							"id": 314,
							"title": "关爱他成长的每一个足迹",
							"price_info": 0,
							"scene_pic_url": "https://yanxuan.nosdn.127.net/14943267735961674.jpg",
							"subtitle": "专业运动品牌同厂，毛毛虫鞋买二送一"
						},
						{
							"id": 313,
							"title": "一次解决5个节日送礼难题",
							"price_info": 59.9,
							"scene_pic_url": "https://yanxuan.nosdn.127.net/14942996754171334.jpg",
							"subtitle": "这些就是他们想要的礼物清单"
						},
						{
							"id": 300,
							"title": "秒杀化学洗涤剂的纯天然皂",
							"price_info": 0,
							"scene_pic_url": "https://yanxuan.nosdn.127.net/14939843143621089.jpg",
							"subtitle": "前段时间有朋友跟我抱怨，和婆婆住到一起才发现生活理念有太多不和。别的不提，光是洗..."
						},
						{
							"id": 299,
							"title": "买过的人都说它是差旅神器",
							"price_info": 0,
							"scene_pic_url": "https://yanxuan.nosdn.127.net/14938873919030679.jpg",
							"subtitle": "许多人经历过旅途中内裤洗晾不便的烦恼，尤其与旅伴同居一室时，晾在卫生间里的内裤更..."
						},
						{
							"id": 295,
							"title": "他们在严选遇见的新生活",
							"price_info": 35.8,
							"scene_pic_url": "https://yanxuan.nosdn.127.net/14938092956370380.jpg",
							"subtitle": "多款商品直减中，最高直减400元"
						},
						{
							"id": 294,
							"title": "这只锅，可以从祖母用到孙辈",
							"price_info": 149,
							"scene_pic_url": "https://yanxuan.nosdn.127.net/14937214454750141.jpg",
							"subtitle": "买100年传世珐琅锅送迷你马卡龙色小锅"
						},
						{
							"id": 291,
							"title": "舒适新主张",
							"price_info": 29,
							"scene_pic_url": "https://yanxuan.nosdn.127.net/14939496197300723.jpg",
							"subtitle": "如何挑选适合自己的好物？"
						},
						{
							"id": 289,
							"title": "专业运动袜也可以高性价比",
							"price_info": 0,
							"scene_pic_url": "https://yanxuan.nosdn.127.net/14932840600970609.jpg",
							"subtitle": "越来越多运动人士意识到，运动鞋要购置好的，鞋里的运动袜也不可忽视。专业运动袜帮助..."
						},
						{
							"id": 287,
							"title": "严选新式样板间",
							"price_info": 29.9,
							"scene_pic_url": "https://yanxuan.nosdn.127.net/14931970965550315.jpg",
							"subtitle": "一种软装一个家"
						},
						{
							"id": 286,
							"title": "无“油”无虑的甜蜜酥脆",
							"price_info": 0,
							"scene_pic_url": "https://yanxuan.nosdn.127.net/14931121822100127.jpg",
							"subtitle": "大家都知道，饮食组是严选体重最重的一组，基本上每个新人都能在一个月之内迅速长胖。..."
						}
					],
			}
		},
		onLoad() {

		},
		methods: {
			prevPage(){
				
			},
			nextPage(){
				
			}
		}
	}
</script>

<style lang="scss">
	page ,.container{
	   width: 750rpx;
	    height: 100%;
	    overflow: hidden;
	    background: #f4f4f4;
	}
	.topic-list{
	    width: 750rpx;
	    height: 100%;
	    overflow: hidden;
	    background: #f4f4f4;
	}
	
	.topic-list .item{
	    width: 100%;
	    height: 625rpx;
	    overflow: hidden;
	    background: #fff;
	    margin-bottom: 20rpx;
	}
	
	.topic-list .img{
	    width: 100%;
	    height: 415rpx;
	}
	
	.topic-list .info{
	    width: 100%;
	    height: 210rpx;
	    overflow: hidden;
	}
	
	.topic-list .title{
	    display: block;
	    text-align: center;
	    width: 100%;
	    height: 33rpx;
	    line-height: 35rpx;
	    color: #333;
	    overflow: hidden;
	    font-size: 35rpx;
	    margin-top: 30rpx;
	}
	
	.topic-list .desc{
	    display: block;
	    text-align: center;
	    position: relative;
	    width: auto;
	    height: 24rpx;
	    line-height: 24rpx;
	    overflow: hidden;
	    color: #999;
	    font-size: 24rpx;
	    margin-top: 16rpx;
	    margin-bottom: 30rpx;
	}
	
	.topic-list .price{
	    display: block;
	    text-align: center;
	    width: 100%;
	    height: 27rpx;
	    line-height: 27rpx;
	    overflow: hidden;
	    color: #b4282d;
	    font-size: 27rpx;
	}
	
	
	.page{
	    width: 750rpx;
	    height: 108rpx;
	    background: #fff;
	    margin-bottom: 20rpx;
	}
	
	.page view{
	    height: 108rpx;
	    width: 50%;
	    float: left;
	    font-size: 29rpx;
	    color: #333;
	    text-align: center;
	    line-height: 108rpx;
	}
	
	.page .prev{
	    border-right: 1px solid #D9D9D9;
	}
	
	.page .disabled{
	    color: #ccc;
	}
</style>
