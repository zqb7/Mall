<template>
	<view class="container">
	    <scroll-view class="topic-list" scroll-y="true" :scroll-top="scrollTop">
	        <navigator class="item" v-for="(item, key) in topicList" :key="key" :url="'../topicDetail/topicDetail?id=' + item.id">
	            <image class="img" :src="item.scene_pic_url"></image>
	            <view class="info">
	                <text class="title">{{item.title}}</text>
	                <text class="desc">{{item.subtitle}}</text>
	                <text class="price">{{item.price_info}}元起</text>
	            </view>
	        </navigator>
	        <view class="page" v-if="showPage">
	            <view class="prev" :class="{'disabled': page <= 1}" @click="prevPage">上一页</view>
	            <view class="next" :class="{'disabled': !hasNextPage}" @click="nextPage">下一页</view>
	        </view>
	    </scroll-view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				showPage:false,
				tscrollTop: 0,
				page: 1,
				size: 10,
				count: 0,
				scrollTop: 0,
				topicList: [],
				hasNextPage:true
			}
		},
		onLoad() {
			this.loadData(this.page,this.size)
			if (this.hasNextPage) {
				this.showPage = true
			}
		},
		methods: {
			async loadData(page,size) {
				uni.request({
					url:this.api +`/topic?page=${page}&size=${size}`,
					success: (res) => {
						let d = res["data"]				
						this.count = d["count"]
						let topicList = d["topic_list"]
						if (topicList.length > 0 ){
							this.topicList = topicList
						}
						if (d["has_next_page"]){
							this.hasNextPage = true
						} else {
							this.hasNextPage = false
						}
						this.page = d["current_page"]
					}
				})
			},
			prevPage(){
				this.loadData(this.page-1,this.size)
			},
			nextPage(){
				this.loadData(this.page+1,this.size)
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
	    // border-right: 0.5px solid #D9D9D9;
	}
	
	.page .disabled{
	    color: #ccc;
		pointer-events: none;
	}
</style>
