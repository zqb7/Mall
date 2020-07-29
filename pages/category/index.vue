<template>
	<view class="content">
		<view class="cate-nav" :style="{position:headerPosition,top:headerTop}">
			<scroll-view scroll-x class="cate-nav-body" style="width: 750upx;" :scroll-left="scrollLeft">
				<view   v-for="(item,key) in navList"  :class="{active: id ===item.id}" class="item"  @click="switchCate(item)">
					<view class="name">{{item.name}}</view>
				</view>
			</scroll-view>
		</view>
		<scroll-view scroll-y="true" :scroll-top="scrollTop">
	
			<view class="cate-item">
				<view class="h">
					<text class="name">{{currentCategory.name}}</text>
					<text class="desc">{{currentCategory.front_name}}</text>
				</view>
				<view class="b">
					<navigator v-for="(item,index) in goodsList" :class="{'item-b': (index+1) % 2 === 0 }" class="item" :url="'/pages/goods/index?id=' + item.id">
						<image class="img" :src="item.list_pic_url" background-size="cover"></image>
						<text class="name">{{item.name}}</text>
						<text class="price">￥{{item.retail_price}}</text>
					</navigator>
				</view>
			</view>
		</scroll-view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				headerPosition:"fixed",
				headerTop:"0px",
				id: 1008002,
		
				scrollLeft: 0,
				scrollTop: 0,
				scrollHeight: 0,
				currentCategory: {},
				goodsList: [],
				navList: []
			}
		},
		onLoad(option) {
			console.log(option.id)
			if (option.id == undefined){
				option.id = this.id
			}
			// this.headerTop = document.getElementsByTagName('uni-page-head')[0].offsetHeight+'px';
			this.loadData(option.id)
			this.getGoods(option.id)
		},
		methods: {
			switchCate(item) {
				this.loadData(item.id)
				this.id = item.id;
				this.getGoods(item.id)
			},
			navToList(pid) {
				uni.navigateTo({
					url:`/pages/goods/goods?id=${pid}`
				})
			},
			loadData(categoryID){
				uni.request({
					url:this.api +`/category/${categoryID}`,
					success: (res) => {
						let d = res["data"]
						this.currentCategory = d["current_category"]
						this.navList = d["brother_category_list"]
						this.id = this.currentCategory.id
					}
				})
			},
			getGoods(categoryID){
				uni.request({
					url:this.api +`/category/${categoryID}/goods`,
					success: (res) => {
						let d = res["data"]
						this.goodsList = d["goods_list"]
					}
				})
			}
		}
	}
</script>

<style lang="scss">
	.content{
		background: #f9f9f9;
	}
	.cate-nav{
		position: fixed;
		left:0;
		top:0;
		z-index: 1000;
	}

	.cate-nav-body{
		height: 84upx;
		white-space: nowrap;   
		background: #fff;
		border-top: 1px solid rgba(0,0,0,.15);
		overflow: hidden;
	}

	.cate-nav .item{
		display: inline-block;
		height: 84upx;
		min-width: 130upx;
		padding: 0 15upx;
	}

	.cate-nav .item .name{
		display: block;
		height: 84upx;
		padding: 0 20upx;
		line-height: 84upx;
		color: #333;
		font-size: 30upx;
		width: auto;
	}

	.cate-nav .item.active .name{
		color: #ab2b2b;
		border-bottom: 2px solid #ab2b2b;
	}

	.cate-item{
		margin-top: 94upx;
		height: auto;
		overflow: hidden;
	}

	.cate-item .h{
		height: 145upx;
		width: 750upx;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}

	.cate-item .h .name{
		display: block;
		height: 35upx;
		margin-bottom: 18upx;
		font-size: 30upx;
		color: #333;
	}

	.cate-item .h .desc{
		display: block;
		height: 24upx;
		font-size: 24upx;
		color: #999;
	}

	.cate-item .b{
	  width: 750upx;
	  padding: 0 6.25upx;
	  height: auto;
	  overflow: hidden;
	}

	.cate-item .b .item{
	  float: left;
	  background: #fff;
	  width: 365upx;
	  margin-bottom: 6.25upx;
	  padding-bottom: 33.333upx;
	  height: auto;
	  overflow: hidden;
	  text-align: center;
	}

	.cate-item .b .item-b{
	  margin-left: 6.25upx;
	}

	.cate-item .item .img{
	  width: 302upx;
	  height: 302upx;
	}

	.cate-item .item .name{
	  display: block;
	  width: 365.625upx;
	  height: 35upx;
	  margin: 11.5upx 0 22upx 0;
	  text-align: center;
	  overflow: hidden;
	  padding: 0 20upx;
	  font-size: 30upx;
	  color: #333;
	}

	.cate-item .item .price{
	  display: block;
	  width: 365.625upx;
	  height: 30upx;
	  text-align: center;
	  font-size: 30upx;
	  color: #b4282d;
	}
</style>
