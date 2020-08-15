<template>
	<view class="container">
		 <view class="search">
		    <navigator url="/pages/search/search" class="input">
		      <image class="icon"></image>
		      <text class="txt">商品搜索, 共{{goodsCount}}款好物</text>
		    </navigator>
		  </view>
		  <view class="catalog">
			<scroll-view class="nav" scroll-y>
				<view v-for="(item,key) in navList" :key="key" :class="{active: item.id === currentCategory.id}" class="item"  @click="tabtap(item)">{{item.name}}</view>
			</scroll-view>
			
			<scroll-view scroll-y class="cate" style="height: 1024upx;">
				<navigator url="url" class="banner">
					<image class="image" :src="currentCategory.wap_banner_url"></image>
					<view class="txt">{{currentCategory.front_name}}</view>
				</navigator>
				<view class="hd">
					<text class="line"></text>
					<text class="txt">{{currentCategory.name}}分类</text>
					<text class="line"></text>
				</view>
				<view class="bd">
					<navigator v-for="(item, index) in currentCategory.sub_category_list" :key="index" class="item" :class="{last: (index+1) %3 ===0}" :url="'/pages/category/index?id=' + item.id">
						<image class="icon" :src="item.wap_banner_url"></image>
						<text class="txt">{{item.name}}</text>
					</navigator>
				</view>
			</scroll-view>
			
			</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				offsetHeight:0,
				goodsCount: 0,
				navList: [],
				goodsList: [],
				id: 0,
				currentCategory: {},
				scrollLeft: 0,
				scrollTop: 0,
				scrollHeight: 0,
				page: 1,
				size: 10000
			}
		},
		onLoad() {
			this.loadData()
		},
		methods: {
			//一级分类点击
			tabtap(item){
				console.log(item);
				this.currentCategory.id = item.id;
				uni.request({
					url:this.api + "/catalog/" + item.id,
					success: (res) => {
						let d = res["data"]
						this.currentCategory = d["current_category"]
					}
				})
			},
			async loadData(){
				uni.request({
					url:this.api + "/catalog",
					success: (res) => {
						let d = res["data"]
						this.navList = d["category_list"]
						this.currentCategory = d["current_category"]
					}
				})
				uni.request({
					url:this.api+ "/count/goods",
					success: (res) => {
						if (res.data.hasOwnProperty("goods_count")){
							this.goodsCount = res.data.goods_count
						}
					}
				})
			}
		}
	}
</script>

<style lang="scss">
	page {
	  height: 100%;
	}

	.container {
	  background: #f9f9f9;
	  height: 100%;
	  width: 100%;
	  display: flex;
	  flex-direction: column;
	}

	.search {
	  height: 88upx;
	  width: 100%;
	  padding: 0 30upx;
	  background: #fff;
	  display: flex;
	  align-items: center;
	}

	.search .input {
	  width: 690upx;
	  height: 56upx;
	  background: #ededed;
	  border-radius: 8upx;
	  display: flex;
	  align-items: center;
	  justify-content: center;
	}

	.search .icon {
	  background: url(http://yanxuan.nosdn.127.net/hxm/yanxuan-wap/p/20161201/style/img/icon-normal/search2-2fb94833aa.png) center no-repeat;
	  background-size: 100%;
	  width: 28upx;
	  height: 28upx;
	}

	.search .txt {
	  height: 42upx;
	  line-height: 42upx;
	  color: #666;
	  padding-left: 10upx;
	  font-size: 30upx;
	}

	.catalog {
	  flex: 1;
	  width: 100%;
	  background: #fff;
	  display: flex;
	  border-top: 1px solid #fafafa;
	}

	.catalog .nav {
	  width: 162upx;
	  height: 100%;
	}

	.catalog .nav .item {
	  text-align: center;
	  line-height: 90upx;
	  width: 162upx;
	  height: 90upx;
	  color: #333;
	  font-size: 28upx;
	  border-left: 6upx solid #fff;
	}

	.catalog .nav .item.active {
	  color: #ab2b2b;
	  font-size: 36upx;
	  border-left: 6upx solid #ab2b2b;
	}

	.catalog .cate {
	  border-left: 1px solid #fafafa;
	  flex: 1;
	  overflow: hidden;
	  height: 100%;
	  padding: 0 30upx 0 30upx;
	}

	.banner {
	  display: block;
	  height: 222upx;
	  width: 100%;
	  position: relative;
	}

	.banner .image {
	  position: absolute;
	  top: 30upx;
	  left: 0;
	  border-radius: 4upx;
	  height: 192upx;
	  width: 100%;
	}

	.banner .txt {
	  position: absolute;
	  top: 30upx;
	  text-align: center;
	  color: #fff;
	  font-size: 28upx;
	  left: 0;
	  height: 192upx;
	  line-height: 192upx;
	  width: 100%;
	}

	.catalog .hd {
	  height: 108upx;
	  width: 100%;
	  display: flex;
	  justify-content: center;
	  align-items: center;
	}

	.catalog .hd .txt {
	  font-size: 24upx;
	  text-align: center;
	  color: #333;
	  padding: 0 10upx;
	  width: auto;
	}

	.catalog .hd .line {
	  width: 40upx;
	  height: 1px;
	  background: #d9d9d9;
	}

	.catalog .bd {
	  height: auto;
	  width: 100%;
	  overflow: hidden;
	}

	.catalog .bd .item {
	  display: block;
	  float: left;
	  height: 216upx;
	  width: 144upx;
	  margin-right: 34upx;
	}

	.catalog .bd .item.last {
	  margin-right: 0;
	}

	.catalog .bd .item .icon {
	  height: 144upx;
	  width: 144upx;
	}

	.catalog .bd .item .txt {
	  display: block;
	  text-align: center;
	  font-size: 24upx;
	  color: #333;
	  height: 72upx;
	  width: 144upx;
	}
</style>
