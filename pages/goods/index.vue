<template>
	<view>
		<view class="container">
		  <swiper class="goodsimgs" indicator-dots="true" autoplay="true" interval="3000" duration="1000">
		    <swiper-item v-for="(item,key) in gallery" :key="key">
		      <image :src="item.img_url" background-size="cover"></image>
		    </swiper-item>
		  </swiper>
		  <view class="service-policy">
		    <view class="item">30天无忧退货</view>
		    <view class="item">48小时快速退款</view>
		    <view class="item">满88元免邮费</view>
		  </view>
		  <view class="goods-info">
		    <view class="c">
		      <text class="name">{{goods.name}}</text>
		      <text class="desc">{{goods.goods_brief}}</text>
		      <text class="price">￥{{goods.retail_price}}</text>
		      <view class="brand" v-if="brand.name">
		        <navigator :url="'../brandDetail/brandDetail?id=' + brand.brandId">
		          <text>{{brand.name}}</text>
		        </navigator>
		      </view>
		    </view>
		  </view>
		  <view class="section-nav section-attr" @click="toggleSpec">
		    <view class="t">请选择规格数量</view>
		    <image class="i" src="/static/images/goods/address_right.png" background-size="cover"></image>
		  </view>
		  <!--<view class="section-nav section-act">
		      <view class="t">
		        <view class="label">1个促销:</view>
		        <view class="tag">万圣趴</view>
		        <view class="text">全场满499，额外送糖果</view>
		      </view>
		      <image class="i" src="../../static/images/address_right.png" background-size="cover"></image>
		    </view>-->
		  <view class="comments" v-if="comment.count > 0">
		    <view class="h">
		      <navigator :url="'../comment/comment?valueId=' + goods.id + '&typeId=0'">
		        <text class="t">评价({{comment.count > 999 ? '999+' : comment.count}})</text>
		        <text class="i">查看全部</text>
		      </navigator>
		    </view>
		    <view class="b">
		      <view class="item">
		        <view class="info">
		          <view class="user">
		            <image :src="comment.data.avatar"></image>
		            <text>{{comment.data.nickname}}</text>
		          </view>
		          <view class="time">{{comment.data.add_time}}</view>
		        </view>
		        <view class="content">
		          {{comment.data.content}}
		        </view>
		        <view class="imgs" v-if="comment.data.pic_list.length > 0">
		          <image class="img" v-for="(item, key) in comment.data.pic_list" :key="key"  src="item.pic_url"></image>
		        </view>
		        <!-- <view class="spec">白色 2件</view> -->
		      </view>
		    </view>
		  </view>
		  
		  <view class="goods-attr">
		    <view class="t">商品参数</view>
		    <view class="l">
		      <view class="item" v-for="(item,key) in attribute" :key="key">
		        <text class="left">{{item.key}}</text>
		        <text class="right">{{item.value}}</text>
		      </view>
		    </view>
		  </view>
		
		  <view class="detail">
			  <u-parse :content="goods.goods_desc" @preview="preview" @navigate="navigate" />
		  </view>
		
		
		  <view class="common-problem">
		    <view class="h">
		      <view class="line"></view>
		      <text class="title">常见问题</text>
		    </view>
			
		    <view class="b">
		      <view class="item" v-for="(item,key) in issueList" :key="key">
		        <view class="question-box">
		          <text class="spot"></text>
		          <text class="question">{{item.question}}</text>
		        </view>
		        <view class="answer">
		          {{item.answer}}
		        </view>
		      </view>
		    </view>
		  </view>
		
		  <view class="related-goods" v-if="relatedGoods.length > 0">
		    <view class="h">
		      <view class="line"></view>
		      <text class="title">大家都在看</text>
		    </view>
		    <view class="b">
		      <view class="item" v-for="(item,key) in relatedGoods" :key="key">
		        <navigator :url="'/pages/goods/index?id=' + item.id">
		          <image class="img" :src="item.list_pic_url" background-size="cover"></image>
		          <text class="name">{{item.name}}</text>
		          <text class="price">￥{{item.retail_price}}</text>
		        </navigator>
		      </view>
		    </view>
		  </view>
		</view>
		
		<!-- 规格-模态层弹窗 -->
		<view 
			class="popup spec" 
			:class="specClass"
			@touchmove.stop.prevent="stopPrevent"
			@click="toggleSpec"
		>
			<!-- 遮罩层 -->
			<view class="mask"></view>
			<view class="layer attr-content" @click.stop="stopPrevent">
				<view class="a-t">
					<image :src="goods.list_pic_url"></image>
					<view class="right">
						<text class="price">¥{{selectedProduct.retail_price}}</text>
						<text class="stock">库存：{{selectedProduct.goods_number}}件</text>
						<view class="selected">
							已选：
							<text class="selected-text" v-for="(sItem, sIndex) in specSelected" :key="sIndex">
								{{sItem.value}}
							</text>
						</view>
					</view>
				</view>
				<view v-for="(item,index) in specificationList" :key="index" class="attr-list">
					<text>{{item.name}}</text>
					<view class="item-list">
						<text 
							v-for="(childItem, childIndex) in item.value_list" 
							v-if="childItem.specification_key_id === item.id"
							:key="childIndex" class="tit"
							:class="{selected: childItem.selected}"
							@click="selectSpec(index, childIndex, childItem.specification_key_id)"
						>
							{{childItem.value}}
						</text>
					</view>
				</view>
				<button class="btn" @click="toggleSpec">完成</button>
			</view>
		</view>
		
		<!-- 底部操作菜单 -->
		<view class="page-bottom">
			<navigator url="/pages/index/index" open-type="switchTab" class="p-b-btn">
				<text class="yticon icon-xiatubiao--copy"></text>
				<text>首页</text>
			</navigator>
			<navigator url="/pages/shopping/index" open-type="switchTab" class="p-b-btn">
				<text class="yticon icon-gouwuche"></text>
				<text>购物车</text>
			</navigator>
			<view class="p-b-btn" :class="{active: favorite}" @click="toFavorite">
				<text class="yticon icon-shoucang"></text>
				<text>收藏</text>
			</view>
			
			<view class="action-btn-group">
				<button type="primary" class=" action-btn no-border buy-now-btn" @click="buy">立即购买</button>
				<button type="primary" class=" action-btn no-border add-cart-btn" @click="addToCart">加入购物车</button>
			</view>
		</view>
		<uni-popup ref="popupMessage" type="message" @change="change">
			<uni-popup-message type="succes" :message="message" :duration="2000"></uni-popup-message>
		</uni-popup>
	</view>
</template>

<script>
	import {
		mapState
	} from 'vuex'
	import uParse from '@/components/u-parse/u-parse.vue'
	import uniPopupMessage from '@/components/uni-popup/uni-popup-message.vue'
	export default {
		components: {
		    uParse,
			uniPopupMessage
	    },
		computed:{
			...mapState(['token'])
		},
		data() {
			return {
				id: 1181000,
				goods: {},
				gallery: [],
				attribute: [],
				issueList: [],
				comment: [],
				brand: {},
				specificationList: [],
				productList: [],
				relatedGoods: [],
				cartGoodsCount: 0,
				userHasCollect: 0,
				number: 1,
				checkedSpecText: '请选择规格数量',
				openAttr: false,
				noCollectImage: "/static/images/goods/icon_collect.png",
				hasCollectImage: "/static/images/goods/icon_collect_checked.png",
				collectBackImage: "/static/images/goods/icon_collect.png",
				message: '加入成功',
				specClass: 'none',
				specSelected: [],
				favorite: true,
				selectedProduct: {},
			};
		},
		onLoad(option) {
			if (option.id == undefined) {
				option.id = this.id
			}
			this.id=option.id
			this.loadData()
		},
		methods: {
			loadData(){
				uni.request({
					url:this.api + `/goods/${this.id}`,
					success: (res) => {
						let d = res["data"]
						this.gallery = d["gallery"]
						this.goods = d["goods"]
						this.issueList = d["issue"]
						this.attribute = d["attribute"]
						this.specificationList = d["specification_list"]
						this.productList = d["product_list"]
					}
				})
			},
			addToCart() {
				if (this.specSelected.length === 0) {
					this.toggleSpec()
					return
				}
				if (this.specSelected.length !== this.specificationList.length) {
					this.$msg("请选择规格")
					this.toggleSpec()
					return
				}
				uni.request({
					url:this.api + `/user/cart`,
					method:"POST",
					header: {
						'token': this.token,
					},
					data:JSON.stringify({
						goods_id:this.goods.id,
						number:1,
						"goods_sn": this.selectedProduct.goods_sn,
						"goods_name": this.goods.name,
						"goods_pic": this.selectedProduct.goods_pic,
						"goods_specification_value": this.selectedProduct.goods_specification_value,
						"goods_specification_ids": this.selectedProduct.ids,
					}),
					success: (res) => {
						if (res.statusCode == 200){
							this.message = '加入成功'
							this.$refs.popupMessage.open()
						}else {
							console.log(res.data)
						}
					},
					fail: (res) => {
						this.message = '加入失败，请稍后再试'
						this.$refs.popupMessage.open()
					}
				})
			},
			//规格弹窗开关
			toggleSpec() {
				if(this.specClass === 'show'){
					this.specClass = 'hide';
					setTimeout(() => {
						this.specClass = 'none';
					}, 250);
				}else if(this.specClass === 'none'){
					this.specClass = 'show';
				}
			},
			//选择规格
			selectSpec(index,childIndex, specification_key_id){
				let list = this.specificationList[index].value_list;
				list.forEach(item=>{
					if(item.specification_key_id === specification_key_id){
						this.$set(item, 'selected', false);
					}
				})
			
				this.$set(list[childIndex], 'selected', true);
				//存储已选择
				this.specSelected = []; 
				list = this.specificationList
				let goodsSpecificationValue = ""
				let goodsPic = ""
				list.forEach(item=>{ 
					item.value_list.forEach(item2=>{
						if(item2.selected === true){
							this.specSelected.push(item2); 
							goodsSpecificationValue += item2.value + " "
							if (item2.pic_url.trim() !== ""){
								goodsPic = item2.pic_url
							}
						}
					})
 
				})	
				this.productList.forEach(item=>{
					let ids = ""
					this.specSelected.forEach(item2=>{
						ids += `${item2.id}_`
					})
					ids = ids.substring(0, ids.lastIndexOf("_"))
					if (item.goods_specification_ids === ids || item.goods_specification_ids == ids.split('').reverse().join('')){
						this.selectedProduct = item
						this.$set(this.selectedProduct, 'goods_specification_value',goodsSpecificationValue)
						this.$set(this.selectedProduct, 'ids',ids)
						this.$set(this.selectedProduct, 'goods_pic',goodsPic)
						
					}
				})
			},
			stopPrevent(){
				
			},
			change(){
				
			},
			buy(){
				if(JSON.stringify(this.selectedProduct) == '{}'){
					this.toggleSpec()
					return
				}
				console.log(this.selectedProduct)
				let product = this.selectedProduct
				product.goods_pic = escape(product.goods_pic)
				product.goods_name = this.goods.name
				let data = JSON.stringify({goodsData:product})
				uni.navigateTo({
					url: `/pages/order/createOrder?data=${data}`
				})
			},
		}
	}
</script>

<style lang="scss">
	.container {
	  margin-bottom: 100upx;
	}

	.goodsimgs {
	  width: 750upx;
	  height: 750upx;
	}

	.goodsimgs image {
	  width: 750upx;
	  height: 750upx;
	}

	.service-policy {
	  width: 750upx;
	  height: 73upx;
	  background: #f4f4f4;
	  padding: 0 31.25upx;
	  display: flex;
	  flex-flow: row nowrap;
	  align-items: center;
	  justify-content: space-between;
	}

	.service-policy .item {
	  background: url(http://nos.netease.com/mailpub/hxm/yanxuan-wap/p/20150730/style/img/icon-normal/servicePolicyRed-518d32d74b.png) 0 center no-repeat;
	  background-size: 10upx;
	  padding-left: 15upx;
	  display: flex;
	  align-items: center;
	  font-size: 25upx;
	  color: #666;
	}

	.goods-info {
	  width: 750upx;
	  height: 306upx;
	  overflow: hidden;
	  background: #fff;
	}

	.goods-info .c {
	  display: block;
	  width: 718.75upx;
	  height: 100%;
	  margin-left: 31.25upx;
	  padding: 38upx 31.25upx 38upx 0;
	  border-bottom: 1px solid #f4f4f4;
	}

	.goods-info .c text {
	  display: block;
	  width: 687.5upx;
	  text-align: center;
	}

	.goods-info .name {
	  height: 41upx;
	  margin-bottom: 5.208upx;
	  font-size: 41upx;
	  line-height: 41upx;
	}

	.goods-info .desc {
	  height: 43upx;
	  margin-bottom: 41upx;
	  font-size: 24upx;
	  line-height: 36upx;
	  color: #999;
	}

	.goods-info .price {
	  height: 35upx;
	  font-size: 35upx;
	  line-height: 35upx;
	  color: #b4282d;
	}

	.goods-info .brand {
	  margin-top: 23upx;
	  min-height: 40upx;
	  text-align: center;
	}

	.goods-info .brand text {
	  display: inline-block;
	  width: auto;
	  padding: 2px 30upx 2px 10.5upx;
	  line-height: 35.5upx;
	  border: 1px solid #f48f18;
	  font-size: 25upx;
	  color: #f48f18;
	  border-radius: 4px;
	  background: url(http://nos.netease.com/mailpub/hxm/yanxuan-wap/p/20150730/style/img/icon-normal/detailTagArrow-18bee52dab.png) 95% center no-repeat;
	  background-size: 10.75upx 18.75upx;
	}

	.section-nav {
	  width: 750upx;
	  height: 108upx;
	  background: #fff;
	  margin-bottom: 20upx;
	}

	.section-nav .t {
	  float: left;
	  width: 600upx;
	  height: 108upx;
	  line-height: 108upx;
	  font-size: 29upx;
	  color: #333;
	  margin-left: 31.25upx;
	}

	.section-nav .i {
	  float: right;
	  width: 52upx;
	  height: 52upx;
	  margin-right: 16upx;
	  margin-top: 28upx;
	}

	.section-act .t {
	  float: left;
	  display: flex;
	  align-items: center;
	  width: 600upx;
	  height: 108upx;
	  overflow: hidden;
	  line-height: 108upx;
	  font-size: 29upx;
	  color: #999;
	  margin-left: 31.25upx;
	}

	.section-act .label {
	  color: #999;
	}

	.section-act .tag {
	  display: flex;
	  align-items: center;
	  padding: 0 10upx;
	  border-radius: 3px;
	  height: 37upx;
	  width: auto;
	  color: #f48f18;
	  overflow: hidden;
	  border: 1px solid #f48f18;
	  font-size: 25upx;
	  margin: 0 10upx;
	}

	.section-act .text {
	  display: flex;
	  align-items: center;
	  height: 37upx;
	  width: auto;
	  overflow: hidden;
	  color: #f48f18;
	  font-size: 29upx;
	}

	.comments {
	  width: 100%;
	  height: auto;
	  padding-left: 30upx;
	  background: #fff;
	  margin: 20upx 0;
	}

	.comments .h {
	  height: 102.5upx;
	  line-height: 100.5upx;
	  width: 718.75upx;
	  padding-right: 16upx;
	  border-bottom: 1px solid #d9d9d9;
	}

	.comments .h .t {
	  display: block;
	  float: left;
	  width: 50%;
	  font-size: 38.5upx;
	  color: #333;
	}

	.comments .h .i {
	  display: block;
	  float: right;
	  width: 164upx;
	  height: 100.5upx;
	  line-height: 100.5upx;
	  background: url(http://nos.netease.com/mailpub/hxm/yanxuan-wap/p/20150730/style/img/icon-normal/address-right-990628faa7.png) right center no-repeat;
	  background-size: 52upx;
	}

	.comments .b {
	  height: auto;
	  width: 720upx;
	}

	.comments .item {
	  height: auto;
	  width: 720upx;
	  overflow: hidden;
	}

	.comments .info {
	  height: 127upx;
	  width: 100%;
	  padding: 33upx 0 27upx 0;
	}

	.comments .user {
	  float: left;
	  width: auto;
	  height: 67upx;
	  line-height: 67upx;
	  font-size: 0;
	}

	.comments .user image {
	  float: left;
	  width: 67upx;
	  height: 67upx;
	  margin-right: 17upx;
	  border-radius: 50%;
	}

	.comments .user text {
	  display: inline-block;
	  width: auto;
	  height: 66upx;
	  overflow: hidden;
	  font-size: 29upx;
	  line-height: 66upx;
	}

	.comments .time {
	  display: block;
	  float: right;
	  width: auto;
	  height: 67upx;
	  line-height: 67upx;
	  color: #7f7f7f;
	  font-size: 25upx;
	  margin-right: 30upx;
	}

	.comments .content {
	  width: 720upx;
	  padding-right: 30upx;
	  line-height: 45.8upx;
	  font-size: 29upx;
	  margin-bottom: 24upx;
	}

	.comments .imgs {
	  width: 720upx;
	  height: auto;
	  margin-bottom: 25upx;
	}

	.comments .imgs .img {
	  height: 150upx;
	  width: 150upx;
	  margin-right: 28upx;
	}

	.comments .spec {
	  width: 720upx;
	  padding-right: 30upx;
	  line-height: 30upx;
	  font-size: 24upx;
	  color: #999;
	  margin-bottom: 30upx;
	}

	.goods-attr {
	  width: 750upx;
	  height: auto;
	  overflow: hidden;
	  padding: 0 31.25upx 25upx 31.25upx;
	  background: #fff;
	}

	.goods-attr .t {
	  width: 687.5upx;
	  height: 104upx;
	  line-height: 104upx;
	  font-size: 38.5upx;
	}

	.goods-attr .item {
	  width: 687.5upx;
	  height: 68upx;
	  padding: 11upx 20upx;
	  margin-bottom: 11upx;
	  background: #f7f7f7;
	  font-size: 38.5upx;
	}

	.goods-attr .left {
	  float: left;
	  font-size: 25upx;
	  width: 134upx;
	  height: 45upx;
	  line-height: 45upx;
	  overflow: hidden;
	  color: #999;
	}

	.goods-attr .right {
	  float: left;
	  font-size: 36.5upx;
	  margin-left: 20upx;
	  width: 480upx;
	  height: 45upx;
	  line-height: 45upx;
	  overflow: hidden;
	  color: #333;
	}

	.detail {
	  width: 750upx;
	  height: auto;
	  overflow: hidden;
	}

	.detail image {
	  width: 750upx;
	  display: block;
	}

	.common-problem {
	  width: 750upx;
	  height: auto;
	  overflow: hidden;
	}

	.common-problem .h {
	  position: relative;
	  height: 145.5upx;
	  width: 750upx;
	  padding: 56.25upx 0;
	  background: #fff;
	  text-align: center;
	}

	.common-problem .h .line {
	  display: inline-block;
	  position: absolute;
	  top: 72upx;
	  left: 0;
	  z-index: 2;
	  height: 1px;
	  margin-left: 225upx;
	  width: 300upx;
	  background: #ccc;
	}

	.common-problem .h .title {
	  display: inline-block;
	  position: absolute;
	  top: 56.125upx;
	  left: 0;
	  z-index: 3;
	  height: 33upx;
	  margin-left: 285upx;
	  width: 180upx;
	  background: #fff;
	}

	.common-problem .b {
	  width: 750upx;
	  height: auto;
	  overflow: hidden;
	  padding: 0upx 30upx;
	  background: #fff;
	}

	.common-problem .item {
	  height: auto;
	  overflow: hidden;
	  padding-bottom: 25upx;
	}

	.common-problem .question-box .spot {
	  float: left;
	  display: block;
	  height: 8upx;
	  width: 8upx;
	  background: #b4282d;
	  border-radius: 50%;
	  margin-top: 11upx;
	}

	.common-problem .question-box .question {
	  // float: left;
	  line-height: 30upx;
	  padding-left: 8upx;
	  display: block;
	  font-size: 26upx;
	  padding-bottom: 15upx;
	  color: #303030;
	}

	.common-problem .answer {
	  line-height: 36upx;
	  padding-left: 16upx;
	  font-size: 26upx;
	  color: #787878;
	}

	.related-goods {
	  width: 750upx;
	  height: auto;
	  overflow: hidden;
	}

	.related-goods .h {
	  position: relative;
	  height: 145.5upx;
	  width: 750upx;
	  padding: 56.25upx 0;
	  background: #fff;
	  text-align: center;
	  border-bottom: 1px solid #f4f4f4;
	}

	.related-goods .h .line {
	  display: inline-block;
	  position: absolute;
	  top: 72upx;
	  left: 0;
	  z-index: 2;
	  height: 1px;
	  margin-left: 225upx;
	  width: 300upx;
	  background: #ccc;
	}

	.related-goods .h .title {
	  display: inline-block;
	  position: absolute;
	  top: 56.125upx;
	  left: 0;
	  z-index: 3;
	  height: 33upx;
	  margin-left: 285upx;
	  width: 180upx;
	  background: #fff;
	}

	.related-goods .b {
	  width: 750upx;
	  height: auto;
	  overflow: hidden;
	}

	.related-goods .b .item {
	  float: left;
	  background: #fff;
	  width: 375upx;
	  height: auto;
	  overflow: hidden;
	  text-align: center;
	  padding: 15upx 31.25upx;
	  border-right: 1px solid #f4f4f4;
	  border-bottom: 1px solid #f4f4f4;
	}

	.related-goods .item .img {
	  width: 311.45upx;
	  height: 311.45upx;
	}

	.related-goods .item .name {
	  display: block;
	  width: 311.45upx;
	  height: 35upx;
	  margin: 11.5upx 0 15upx 0;
	  text-align: center;
	  overflow: hidden;
	  font-size: 30upx;
	  color: #333;
	}

	.related-goods .item .price {
	  display: block;
	  width: 311.45upx;
	  height: 30upx;
	  text-align: center;
	  font-size: 30upx;
	  color: #b4282d;
	}

	.bottom-btn {
	  position: fixed;
	  left: 0;
	  bottom: 0;
	  z-index: 10;
	  width: 750upx;
	  height: 100upx;
	  display: flex;
	  background: #fff;
	}

	.bottom-btn .l {
	  float: left;
	  height: 100upx;
	  width: 162upx;
	  border: 1px solid #f4f4f4;
	  display: flex;
	  align-items: center;
	  justify-content: center;
	}

	.bottom-btn .l.l-collect {
	  border-right: none;
	  border-left: none;
	  text-align: center;
	}

	.bottom-btn .l.l-cart .box {
	  position: relative;
	  height: 60upx;
	  width: 60upx;
	}

	.bottom-btn .l.l-cart .cart-count {
	  height: 28upx;
	  width: 28upx;
	  z-index: 10;
	  position: absolute;
	  top: 0;
	  right: 0;
	  background: #b4282d;
	  text-align: center;
	  font-size: 18upx;
	  color: #fff;
	  line-height: 28upx;
	  border-radius: 50%;
	}

	.bottom-btn .l.l-cart .icon {
	  position: absolute;
	  top: 10upx;
	  left: 0;
	}

	.bottom-btn .l .icon {
	  display: block;
	  height: 44upx;
	  width: 44upx;
	}

	.bottom-btn .c {
	  float: left;
	  height: 100upx;
	  line-height: 96upx;
	  flex: 1;
	  text-align: center;
	  color: #333;
	  border-top: 1px solid #f4f4f4;
	  border-bottom: 1px solid #f4f4f4;
	}

	.bottom-btn .r {
	  border: 1px solid #b4282d;
	  background: #b4282d;
	  float: left;
	  height: 100upx;
	  line-height: 96upx;
	  flex: 1;
	  text-align: center;
	  color: #fff;
	}

	.attr-pop-box {
	  width: 100%;
	  height: 100%;
	  position: fixed;
	  background: rgba(0, 0, 0, .5);
	  z-index: 8;
	  bottom: 0;
	  /* display: none; */
	}

	.attr-pop {
	  width: 100%;
	  height: auto;
	  max-height: 780upx;
	  padding: 31.25upx;
	  background: #fff;
	  position: fixed;
	  z-index: 9;
	  bottom: 100upx;
	}

	.attr-pop .close {
	  position: absolute;
	  width: 48upx;
	  height: 48upx;
	  right: 31.25upx;
	  overflow: hidden;
	  top: 31.25upx;
	}

	.attr-pop .close .icon {
	  width: 48upx;
	  height: 48upx;
	}

	.attr-pop .img-info {
	  width: 687.5upx;
	  height: 177upx;
	  overflow: hidden;
	  margin-bottom: 41.5upx;
	}

	.attr-pop .img {
	  float: left;
	  height: 177upx;
	  width: 177upx;
	  background: #f4f4f4;
	  margin-right: 31.25upx;
	}

	.attr-pop .info {
	  float: left;
	  height: 177upx;
	  display: flex;
	  align-items: center;
	}

	.attr-pop .p {
	  font-size: 33upx;
	  color: #333;
	  height: 33upx;
	  line-height: 33upx;
	  margin-bottom: 10upx;
	}

	.attr-pop .a {
	  font-size: 29upx;
	  color: #333;
	  height: 40upx;
	  line-height: 40upx;
	}

	.spec-con {
	  width: 100%;
	  height: auto;
	  overflow: hidden;
	}

	.spec-con .name {
	  height: 32upx;
	  margin-bottom: 22upx;
	  font-size: 29upx;
	  color: #333;
	}

	.spec-con .values {
	  height: auto;
	  margin-bottom: 31.25upx;
	  font-size: 0;
	}

	.spec-con .value {
	  display: inline-block;
	  height: 62upx;
	  padding: 0 35upx;
	  line-height: 56upx;
	  text-align: center;
	  margin-right: 25upx;
	  margin-bottom: 16.5upx;
	  border: 1px solid #333;
	  font-size: 25upx;
	  color: #333;
	}

	.spec-con .value.disable {
	  border: 1px solid #ccc;
	  color: #ccc;
	}

	.spec-con .value.selected {
	  border: 1px solid #b4282d;
	  color: #b4282d;
	}

	.number-item .selnum {
	  width: 322upx;
	  height: 71upx;
	  border: 1px solid #ccc;
	  display: flex;
	}

	.number-item .cut {
	  width: 93.75upx;
	  height: 100%;
	  text-align: center;
	  line-height: 65upx;
	}

	.number-item .number {
	  flex: 1;
	  height: 100%;
	  text-align: center;
	  line-height: 68.75upx;
	  border-left: 1px solid #ccc;
	  border-right: 1px solid #ccc;
	  float: left;
	}

	.number-item .add {
	  width: 93.75upx;
	  height: 100%;
	  text-align: center;
	  line-height: 65upx;
	}


	.icon-you{
		font-size: $font-base + 2upx;
		color: #888;
	}
	
	/* 底部操作菜单 */
	.page-bottom{
		position:fixed;
		left: 30upx;
		bottom:30upx;
		z-index: 95;
		display: flex;
		justify-content: center;
		align-items: center;
		width: 690upx;
		height: 100upx;
		background: rgba(255,255,255,.9);
		box-shadow: 0 0 20upx 0 rgba(0,0,0,.5);
		border-radius: 16upx;
		
		.p-b-btn{
			display:flex;
			flex-direction: column;
			align-items: center;
			justify-content: center;
			font-size: $font-sm;
			color: $font-color-base;
			width: 96upx;
			height: 80upx;
			.yticon{
				font-size: 40upx;
				line-height: 48upx;
				color: $font-color-light;
			}
			&.active, &.active .yticon{
				color: $uni-color-primary;
			}
			.icon-fenxiang2{
				font-size: 42upx;
				transform: translateY(-2upx);
			}
			.icon-shoucang{
				font-size: 46upx;
			}
		}
		.action-btn-group{
			display: flex;
			height: 76upx;
			border-radius: 100px;
			overflow: hidden;
			box-shadow: 0 20upx 40upx -16upx #fa436a;
			box-shadow: 1px 2px 5px rgba(219, 63, 96, 0.4);
			background: linear-gradient(to right, #ffac30,#fa436a,#F56C6C);
			margin-left: 20upx;
			position:relative;
			&:after{
				content: '';
				position:absolute;
				top: 50%;
				right: 50%;
				transform: translateY(-50%);
				height: 28upx;
				width: 0;
				border-right: 1px solid rgba(255,255,255,.5);
			}
			.action-btn{
				display:flex;
				align-items: center;
				justify-content: center;
				width: 180upx;
				height: 100%;
				font-size: $font-base ;
				padding: 0;
				border-radius: 0;
				background: transparent;
			}
		}
	}
	
	/* 规格选择弹窗 */
	.attr-content{
		padding: 10upx 30upx;
		.a-t{
			display: flex;
			image{
				width: 170upx;
				height: 170upx;
				flex-shrink: 0;
				margin-top: -40upx;
				border-radius: 8upx;;
			}
			.right{
				display: flex;
				flex-direction: column;
				padding-left: 24upx;
				font-size: $font-sm + 2upx;
				color: $font-color-base;
				line-height: 42upx;
				.price{
					font-size: $font-lg;
					color: $uni-color-primary;
					margin-bottom: 10upx;
				}
				.selected-text{
					margin-right: 10upx;
				}
			}
		}
		.attr-list{
			display: flex;
			flex-direction: column;
			font-size: $font-base + 2upx;
			color: $font-color-base;
			padding-top: 30upx;
			padding-left: 10upx;
		}
		.item-list{
			padding: 20upx 0 0;
			display: flex;
			flex-wrap: wrap;
			text{
				display: flex;
				align-items: center;
				justify-content: center;
				background: #eee;
				margin-right: 20upx;
				margin-bottom: 20upx;
				border-radius: 100upx;
				min-width: 60upx;
				height: 60upx;
				padding: 0 20upx;
				font-size: $font-base;
				color: $font-color-dark;
			}
			.selected{
				background: #fbebee;
				color: $uni-color-primary;
			}
		}
	}
	
	/*  弹出层 */
	.popup {
		position: fixed;
		left: 0;
		top: 0;
		right: 0;
		bottom: 0;
		z-index: 99;
		
		&.show {
			display: block;
			.mask{
				animation: showPopup 0.2s linear both;
			}
			.layer {
				animation: showLayer 0.2s linear both;
			}
		}
		&.hide {
			.mask{
				animation: hidePopup 0.2s linear both;
			}
			.layer {
				animation: hideLayer 0.2s linear both;
			}
		}
		&.none {
			display: none;
		}
		.mask{
			position: fixed;
			top: 0;
			width: 100%;
			height: 100%;
			z-index: 1;
			background-color: rgba(0, 0, 0, 0.4);
		}
		.layer {
			position: fixed;
			z-index: 99;
			bottom: 0;
			width: 100%;
			min-height: 40vh;
			border-radius: 10upx 10upx 0 0;
			background-color: #fff;
			.btn{
				height: 66upx;
				line-height: 66upx;
				border-radius: 100upx;
				background: $uni-color-primary;
				font-size: $font-base + 2upx;
				color: #fff;
				margin: 30upx auto 20upx;
			}
		}
		@keyframes showPopup {
			0% {
				opacity: 0;
			}
			100% {
				opacity: 1;
			}
		}
		@keyframes hidePopup {
			0% {
				opacity: 1;
			}
			100% {
				opacity: 0;
			}
		}
		@keyframes showLayer {
			0% {
				transform: translateY(120%);
			}
			100% {
				transform: translateY(0%);
			}
		}
		@keyframes hideLayer {
			0% {
				transform: translateY(0);
			}
			100% {
				transform: translateY(120%);
			}
		}
	}

</style>
