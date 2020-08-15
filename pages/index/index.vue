<template>
	<view class="container">
		<view class="search">
		    <navigator url="/pages/search/search" class="input">
		      <image class="icon"></image>
		      <text class="txt">商品搜索, 共{{goodsCount}}款好物</text>
		    </navigator>
	    </view>
		
		<!-- 轮播图 -->
		<view>
			<swiper  class="banner"  :indicator-dots="indicatorDots" :autoplay="autoplay" :interval="interval" :duration="duration">
				<swiper-item v-for="item in adList" :key="item.id">
					<image :src="item.image_url" />
				</swiper-item>
			</swiper>
		</view>
		
		<!-- 分类 -->
		<view class="m-menu">
			<navigator class="item" v-for="item in channelList" :key="item.id" :url="'/pages/category/index'">
				<image :src="item.icon_url" />
				<text>{{item.name}}</text>
			</navigator>
		</view>
		
		<!-- 品牌 -->
		<view class="a-section a-brand">
		    <view class="h">
		      <navigator url="../brand/brand">
		        <text class="txt">品牌制造商直供</text>
		      </navigator>
		    </view>
		    <view class="b">
		      <view  class="item item-1" v-for="item in brandList" :key="item.id">
		        <navigator url="/pages/brandDetail/brandDetail?id=item.id">
		          <view class="wrap">
		            <image class="img" :src="item.pic_url" mode="aspectFill"></image>
		            <view class="mt">
		              <text class="brand">{{item.name}}</text>
		              <text class="price">{{item.price}}</text>
		              <text class="unit">元起</text>
		            </view>
		          </view>
		        </navigator>
		      </view>
		    </view>
		  </view>
		
		<!-- 新品首发 -->
	    <view class="a-section a-new" v-if="newGoods.length > 0">
		    <view class="h">
		      <view>
		        <navigator url="../newGoods/newGoods">
		          <text class="txt">周一周四 · 新品首发</text>
		        </navigator>
		      </view>
		    </view>
		    <view class="b">
		      <view class="item" v-for="item in newGoods" :key="item.id">
		        <navigator url="../goods/goods?id=item.id">
		          <image class="img" :src="item.list_pic_url" background-size="cover"></image>
		          <text class="name">{{item.name}}</text>
		          <text class="price">￥{{item.retail_price}}</text>
		        </navigator>
		      </view>
		    </view>
	    </view>
		
		<!-- 人气推荐 -->
		<view class="a-section a-popular" v-if="hotGoods.length > 0">
		    <view class="h">
		      <view>
		        <navigator url="../hotGoods/hotGoods">
		          <text class="txt">人气推荐</text>
		        </navigator>
		      </view>
		    </view>
		    <view class="b">
		      <view class="item" v-for="item in hotGoods" :key="item.id">
		        <navigator url="/pages/goods/goods?id=item.id">
		          <image class="img" :src="item.list_pic_url" background-size="cover"></image>
		          <view class="right">
		            <view class="text">
		              <text class="name">{{item.name}}</text>
		              <text class="desc">{{item.goods_brief}}</text>
		              <text class="price">￥{{item.retail_price}}</text>
		            </view>
		          </view>
		        </navigator>
		      </view>
		    </view>
		  </view>
		
		<!-- 专题 -->
		<view class="a-section a-topic" v-if="topics.length > 0">
		    <view class="h">
		      <view>
		        <navigator url="../topic/topic" open-type="switchTab">
		          <text class="txt">专题精选</text>
		        </navigator>
		      </view>
		    </view>
		    <view class="b">
		      <scroll-view scroll-x="true" class="list">
		        <view class="item" v-for="item in topics" :key="item.id">
		          <navigator url="../topicDetail/topicDetail?id=item.id">
		            <image class="img" :src="item.scene_pic_url" background-size="cover"></image>
		            <view class="np">
		              <text class="name">{{item.title}}</text>
		              <text class="price">￥{{item.price_info}}元起</text>
		            </view>
		            <text class="desc">{{item.subtitle}}</text>
		          </navigator>
		        </view>
		      </scroll-view>
		    </view>
		  </view>
		
	    <view class="good-grid" v-for="(item,key) in floorGoodsList" :key="item.id">
		    <view class="h">
		      <view>
		        <text>{{item.name}}</text>
		      </view>
		    </view>
		    <view class="b">
		      <block v-for="iitem in item.goods_list">
		        <navigator url="../goods/goods?id=iitem.id" class="item">
		          <image class="img" :src="iitem.list_pic_url" background-size="cover"></image>
		          <view class="name">{{iitem.name}}</view>
		          <view class="price">￥{{iitem.retail_price}}</view>
		        </navigator>
		      </block>
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
				goodsCount: 0,
				adList:[],
				channelList: [],
				brandList: [],
				categoryList:[],
				floorGoodsList:[],
				topics:[],
				newGoods:[],
				hotGoods:[]
			}
		},
		onLoad() {
			this.loadData()
		},
		onPullDownRefresh:function(){
			this.loadData()
			uni.stopPullDownRefresh()
		},
		methods: {
			async loadData() {
				uni.request({
					url:this.api +"/index",
					success: (res) => {
						let d = res["data"]
						this.adList = d["ads"]
						this.brandList = d["brands"]
						this.channelList= d["channel"]
						this.categoryList = d["category_list"]
						this.newGoods = d["new_goods_list"]
						this.hotGoods = d["hot_goods_list"]
						this.topics = d["topic_list"]
						this.floorGoodsList = d["floor_list"]
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

	.search {
	  height: 88rpx;
	  width: 100%;
	  padding: 0 30rpx;
	  background: #fff;
	  display: flex;
	  align-items: center;
	}

	.search .input {
	  width: 690rpx;
	  height: 56rpx;
	  background: #ededed;
	  border-radius: 8rpx;
	  display: flex;
	  align-items: center;
	  justify-content: center;
	}

	.search .icon {
	  background: url(http://yanxuan.nosdn.127.net/hxm/yanxuan-wap/p/20161201/style/img/icon-normal/search2-2fb94833aa.png) center no-repeat;
	  background-size: 100%;
	  width: 28rpx;
	  height: 28rpx;
	}

	.search .txt {
	  height: 42rpx;
	  line-height: 42rpx;
	  color: #666;
	  padding-left: 10rpx;
	  font-size: 30rpx;
	}

	.banner {
	  width: 750rpx;
	  height: 417rpx;
	}

	.banner image {
	  width: 100%;
	  height: 417rpx;
	}

	.m-menu {
	  display: flex;
	  height: 181rpx;
	  width: 750rpx;
	  flex-flow: row nowrap;
	  align-items: center;
	  justify-content: space-between;
	  background-color: #fff;
	}

	.m-menu .item {
	  flex: 1;
	  display: block;
	  padding: 20rpx 0;
	}

	.m-menu image {
	  display: block;
	  width: 58rpx;
	  height: 58rpx;
	  margin: 0 auto;
	  margin-bottom: 12rpx;
	}

	.m-menu text {
	  display: block;
	  font-size: 24rpx;
	  text-align: center;
	  margin: 0 auto;
	  line-height: 1;
	  color: #333;
	}

	.a-section {
	  width: 750rpx;
	  height: auto;
	  overflow: hidden;
	  background: #fff;
	  color: #333;
	  margin-top: 20rpx;
	}

	.a-section .h {
	  display: flex;
	  flex-flow: row nowrap;
	  align-items: center;
	  justify-content: center;
	  height: 130rpx;
	}

	.a-section .h .txt {
	  padding-right: 30rpx;
	  background:  right 4rpx no-repeat;
	  background-size: 16.656rpx 27rpx;
	  display: inline-block;
	  height: 36rpx;
	  font-size: 33rpx;
	  line-height: 36rpx;
	}

	.a-brand .b {
	  width: 750rpx;
	  height: auto;
	  overflow: hidden;
	  position: relative;
	}

	.a-brand .wrap {
	  position: relative;
	}

	.a-brand .img {
	  position: absolute;
	  left: 0;
	  top: 0;
	}

	.a-brand .mt {
	  position: absolute;
	  z-index: 2;
	  padding: 27rpx 31rpx;
	  left: 0;
	  top: 0;
	}

	.a-brand .mt .brand {
	  display: block;
	  font-size: 33rpx;
	  height: 43rpx;
	  color: #333;
	}

	.a-brand .mt .price, .a-brand .mt .unit {
	  font-size: 25rpx;
	  color: #999;
	}

	.a-brand .item-1 {
	  float: left;
	  width: 375rpx;
	  height: 252rpx;
	  overflow: hidden;
	  border-top: 1rpx solid #fff;
	  margin-left: 1rpx;
	}

	.a-brand .item-1:nth-child(2n+1){
	  margin-left: 0;
	  width: 374rpx;
	}

	.a-brand .item-1 .img {
	  width: 375rpx;
	  height: 253rpx;
	}

	.a-new .b {
	  width: 750rpx;
	  height: auto;
	  overflow: hidden;
	  padding: 0 31rpx 45rpx 31rpx;
	}

	.a-new .b .item {
	  float: left;
	  width: 302rpx;
	  margin-top: 10rpx;
	  margin-left: 21rpx;
	  margin-right: 21rpx;
	}

	.a-new .b .item-b {
	  margin-left: 42rpx;
	}

	.a-new .b .img {
	  width: 302rpx;
	  height: 302rpx;
	}

	.a-new .b .name {
	  text-align: center;
	  display: block;
	  width: 302rpx;
	  height: 35rpx;
	  margin-bottom: 14rpx;
	  overflow: hidden;
	  font-size: 30rpx;
	  color: #333;
	}

	.a-new .b .price {
	  display: block;
	  text-align: center;
	  line-height: 30rpx;
	  font-size: 30rpx;
	  color: #b4282d;
	}
	 
	.a-popular {
	  width: 750rpx;
	  height: auto;
	  overflow: hidden;
	}

	.a-popular .b .item {
	  border-top: 1px solid #d9d9d9;
	  margin: 0 20rpx;
	  height: 264rpx;
	  width: 710rpx;
	}

	.a-popular .b .img {
	  margin-top: 12rpx;
	  margin-right: 12rpx;
	  float: left;
	  width: 240rpx;
	  height: 240rpx;
	}

	.a-popular .b .right {
	  float: left;
	  height: 264rpx;
	  width: 456rpx;
	  display: flex;
	  flex-flow: row nowrap;
	}

	.a-popular .b .text {
	  display: flex;
	  flex-wrap: nowrap;
	  flex-direction: column;
	  justify-content: center;
	  overflow: hidden;
	  height: 264rpx;
	  width: 456rpx;
	}

	.a-popular .b .name {
	  width: 456rpx;
	  display: block;
	  color: #333;
	  line-height: 50rpx;
	  font-size: 30rpx;
	}

	.a-popular .b .desc {
	  width: 456rpx;
	  display: block;
	  color: #999;
	  line-height: 50rpx;
	  font-size: 25rpx;
	}

	.a-popular .b .price {
	  width: 456rpx;
	  display: block;
	  color: #b4282d;
	  line-height: 50rpx;
	  font-size: 33rpx;
	}

	.a-topic .b {
	  height: 533rpx;
	  width: 750rpx;
	  padding: 0 0 48rpx 0;
	}

	.a-topic .b .list {
	  height: 533rpx;
	  width: 750rpx;
	  white-space: nowrap;
	}

	.a-topic .b .item {
	  display: inline-block;
	  height: 533rpx;
	  width: 680.5rpx;
	  margin-left: 30rpx;
	  overflow: hidden;
	}

	.a-topic .b .item:last-child {
	  margin-right: 30rpx;
	}

	.a-topic .b .img {
	  height: 387.5rpx;
	  width: 680.5rpx;
	  margin-bottom: 30rpx;
	}

	.a-topic .b .np {
	  height: 35rpx;
	  margin-bottom: 13.5rpx;
	  color: #333;
	  font-size: 30rpx;
	}

	.a-topic .b .np .price {
	  margin-left: 20.8rpx;
	  color: #b4282d;
	}

	.a-topic .b .desc {
	  display: block;
	  height: 30rpx;
	  color: #999;
	  font-size: 24rpx;
	  white-space: nowrap;
	  overflow: hidden;
	  text-overflow:ellipsis;
	}

	.good-grid {
	  width: 750rpx;
	  height: auto;
	  overflow: hidden;
	  background: #fff;
	  margin-top: 20rpx;
	}

	.good-grid .h {
	  display: flex;
	  flex-flow: row nowrap;
	  align-items: center;
	  justify-content: center;
	  height: 130rpx;
	  font-size: 33rpx;
	  color: #333;
	}

	.good-grid .b {
	  width: 750rpx;
	  height: auto;
	  overflow: hidden;
	  padding: 0 20rpx;
	}

	.good-grid .b .item {
	  display: block;
	  float: left;
	  width: 345rpx;
	  margin-bottom: 30rpx;
	  height: auto;
	  overflow: hidden;
	  text-align: center;
	}

	.good-grid .b .item:nth-child(2n + 1) {
	  margin-right: 20rpx;
	}

	.good-grid .item .img {
	  width: 345rpx;
	  height: 345rpx;
	  background: #f8f8f8;
	  margin-bottom: 10rpx;
	}

	.good-grid .item .name {
	  width: 100%;
	  overflow: hidden;
	  height: 40rpx;
	  line-height: 40rpx;
	  text-align: center;
	  font-size: 30rpx;
	  color: #333;
	}

	.good-grid .item .price {
	  width: 100%;
	  height: 46rpx;
	  line-height: 46rpx;
	  text-align: center;
	  font-size: 30rpx;
	  color: #b4282d;
	}


</style>
