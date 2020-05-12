<template>
	<view class="content">
		<view class="cate-nav">
			<scroll-view scroll-y class="cate-nav-body" style="width: 750upx;" :scroll-left="scrollLeft">
				<view   v-for="item in navList"  :class="{active: id ===item.id}" class="item">
					<view class="name">{{item.name}}</view>
				</view>
			</scroll-view>
		</view>
		<scroll-view scroll-y="true" :scroll-top="scrollTop" style="height:0;">
	
			<view class="cate-item">
				<view class="h">
					<text class="name">{{currentCategory.name}}</text>
					<text class="desc">{{currentCategory.front_name}}</text>
				</view>
				<view class="b">
					<navigator v-for="(item,index) in goodsList" :class="{'item-b': (index+1) % 2 === 0 }" class="item"   :url="'/pages/goods/goods?id=' + item.id">
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
				id: 0,
				currentId: 1008002,
				scrollLeft: 0,
				scrollTop: 0,
				scrollHeight: 0,
				categoryList: [
					{
						"id": 1008002,
						"name": "居家"
					},{
						"id": 1008003,
						"name": "餐厨"
					},{
						"id": 1008004,
						"name": "饮食"
					},{
						"id": 1008005,
						"name": "配件"
					},{
						"id": 1008006,
						"name": "服装"
					},{
						"id": 1008007,
						"name": "婴童"
					},{
						"id": 1008008,
						"name": "杂货"
					},{
						"id": 1008009,
						"name": "洗护"
					},{
						"id": 1008010,
						"name": "志趣"
					}
				],
				currentCategory: {
					"name": "居家",
					"front_name": "志趣"
				},
				goodsList: [
					{
					"id": 1008002,
					"name": "布艺软装",
					"keywords": "",
					"front_desc": "换个软装，换个家",
					"parent_id": 1005000,
					"sort_order": 6,
					"show_index": 4,
					"is_show": 1,
					"banner_url": "http://yanxuan.nosdn.127.net/8bbcd7de60a678846664af998f57e71c.png",
					"icon_url": "",
					"img_url": "",
					"wap_banner_url": "http://yanxuan.nosdn.127.net/2e2fb4f2856a021bbcd1b4c8400f2b06.png",
					"level": "L2",
					"type": 0,
					"front_name": "各种风格软装装点你的家"
				  },
				  {
					"id": 1008008,
					"name": "被枕",
					"keywords": "",
					"front_desc": "选用优质材料，确保好芯",
					"parent_id": 1005000,
					"sort_order": 2,
					"show_index": 2,
					"is_show": 1,
					"banner_url": "http://yanxuan.nosdn.127.net/927bc33f7ae2895dd6c11cf91f5e3228.png",
					"icon_url": "",
					"img_url": "",
					"wap_banner_url": "http://yanxuan.nosdn.127.net/b43ef7cececebe6292d2f7f590522e05.png",
					"level": "L2",
					"type": 0,
					"front_name": "守护你的睡眠时光"
				  },
				  {
					"id": 1008009,
					"name": "床品件套",
					"keywords": "",
					"front_desc": "严格用料，亲肤舒适",
					"parent_id": 1005000,
					"sort_order": 4,
					"show_index": 3,
					"is_show": 1,
					"banner_url": "http://yanxuan.nosdn.127.net/243e5bf327a87217ad1f54592f0176ec.png",
					"icon_url": "",
					"img_url": "",
					"wap_banner_url": "http://yanxuan.nosdn.127.net/81f671bd36bce05d5f57827e5c88dd1b.png",
					"level": "L2",
					"type": 0,
					"front_name": "MUJI等品牌制造商出品"
				  },
				  {
					"id": 1008016,
					"name": "灯具",
					"keywords": "",
					"front_desc": "极简主义，贴近生活的设计",
					"parent_id": 1005000,
					"sort_order": 8,
					"show_index": 6,
					"is_show": 1,
					"banner_url": "http://yanxuan.nosdn.127.net/c48e0d9dcfac01499a437774a915842b.png",
					"icon_url": "",
					"img_url": "",
					"wap_banner_url": "http://yanxuan.nosdn.127.net/f702dc399d14d4e1509d5ed6e57acd19.png",
					"level": "L2",
					"type": 0,
					"front_name": "一盏灯，温暖一个家"
				  }
				],
				navList: [
    {
        "id": 1008002,
        "name": "布艺软装",
        "keywords": "",
        "front_desc": "换个软装，换个家",
        "parent_id": 1005000,
        "sort_order": 6,
        "show_index": 4,
        "is_show": 1,
        "banner_url": "http://yanxuan.nosdn.127.net/8bbcd7de60a678846664af998f57e71c.png",
        "icon_url": "",
        "img_url": "",
        "wap_banner_url": "http://yanxuan.nosdn.127.net/2e2fb4f2856a021bbcd1b4c8400f2b06.png",
        "level": "L2",
        "type": 0,
        "front_name": "各种风格软装装点你的家"
    },
    {
        "id": 1008008,
        "name": "被枕",
        "keywords": "",
        "front_desc": "选用优质材料，确保好芯",
        "parent_id": 1005000,
        "sort_order": 2,
        "show_index": 2,
        "is_show": 1,
        "banner_url": "http://yanxuan.nosdn.127.net/927bc33f7ae2895dd6c11cf91f5e3228.png",
        "icon_url": "",
        "img_url": "",
        "wap_banner_url": "http://yanxuan.nosdn.127.net/b43ef7cececebe6292d2f7f590522e05.png",
        "level": "L2",
        "type": 0,
        "front_name": "守护你的睡眠时光"
    },
    {
        "id": 1008009,
        "name": "床品件套",
        "keywords": "",
        "front_desc": "严格用料，亲肤舒适",
        "parent_id": 1005000,
        "sort_order": 4,
        "show_index": 3,
        "is_show": 1,
        "banner_url": "http://yanxuan.nosdn.127.net/243e5bf327a87217ad1f54592f0176ec.png",
        "icon_url": "",
        "img_url": "",
        "wap_banner_url": "http://yanxuan.nosdn.127.net/81f671bd36bce05d5f57827e5c88dd1b.png",
        "level": "L2",
        "type": 0,
        "front_name": "MUJI等品牌制造商出品"
    },
    {
        "id": 1008016,
        "name": "灯具",
        "keywords": "",
        "front_desc": "极简主义，贴近生活的设计",
        "parent_id": 1005000,
        "sort_order": 8,
        "show_index": 6,
        "is_show": 1,
        "banner_url": "http://yanxuan.nosdn.127.net/c48e0d9dcfac01499a437774a915842b.png",
        "icon_url": "",
        "img_url": "",
        "wap_banner_url": "http://yanxuan.nosdn.127.net/f702dc399d14d4e1509d5ed6e57acd19.png",
        "level": "L2",
        "type": 0,
        "front_name": "一盏灯，温暖一个家"
    },
    {
        "id": 1010003,
        "name": "地垫",
        "keywords": "",
        "front_desc": "手工编织，时尚环保",
        "parent_id": 1005000,
        "sort_order": 5,
        "show_index": 4,
        "is_show": 1,
        "banner_url": "http://yanxuan.nosdn.127.net/83d4c87f28c993af1aa8d3e4d30a2fa2.png",
        "icon_url": "",
        "img_url": "",
        "wap_banner_url": "http://yanxuan.nosdn.127.net/1611ef6458e244d1909218becfe87c4d.png",
        "level": "L2",
        "type": 0,
        "front_name": "家里的第“五”面墙"
    },
    {
        "id": 1011003,
        "name": "床垫",
        "keywords": "",
        "front_desc": "助你拥有舒眠好梦",
        "parent_id": 1005000,
        "sort_order": 3,
        "show_index": 3,
        "is_show": 1,
        "banner_url": "http://yanxuan.nosdn.127.net/316afeb3948b295dfe073e4c51f77a42.png",
        "icon_url": "",
        "img_url": "",
        "wap_banner_url": "http://yanxuan.nosdn.127.net/d6e0e84961032fc70fd52a8d4d0fb514.png",
        "level": "L2",
        "type": 0,
        "front_name": "承托你的好时光"
    },
    {
        "id": 1011004,
        "name": "家饰",
        "keywords": "",
        "front_desc": "点缀美好生活，品质在于细节",
        "parent_id": 1005000,
        "sort_order": 9,
        "show_index": 7,
        "is_show": 1,
        "banner_url": "http://yanxuan.nosdn.127.net/ab0df9445d985bf6719ac415313a8e88.png",
        "icon_url": "",
        "img_url": "",
        "wap_banner_url": "http://yanxuan.nosdn.127.net/79275db76b5865e6167b0fbd141f2d7e.png",
        "level": "L2",
        "type": 0,
        "front_name": "装饰你的家"
    },
    {
        "id": 1015000,
        "name": "家具",
        "keywords": "",
        "front_desc": "一级原木，严苛工艺",
        "parent_id": 1005000,
        "sort_order": 7,
        "show_index": 5,
        "is_show": 1,
        "banner_url": "http://yanxuan.nosdn.127.net/4f00675caefd0d4177892ad18bfc2df6.png",
        "icon_url": "",
        "img_url": "",
        "wap_banner_url": "http://yanxuan.nosdn.127.net/d5d41841136182bf49c1f99f5c452dd6.png",
        "level": "L2",
        "type": 0,
        "front_name": "大师级工艺"
    },
    {
        "id": 1017000,
        "name": "宠物",
        "keywords": "",
        "front_desc": "出口品质，严选贴合萌宠生活习惯用品。",
        "parent_id": 1005000,
        "sort_order": 10,
        "show_index": 8,
        "is_show": 1,
        "banner_url": "http://yanxuan.nosdn.127.net/a0352c57c60ce4f68370ecdab6a30857.png",
        "icon_url": "",
        "img_url": "",
        "wap_banner_url": "http://yanxuan.nosdn.127.net/dae4d6e89ab8a0cd3e8da026e4660137.png",
        "level": "L2",
        "type": 0,
        "front_name": "抑菌除味，打造宠物舒适空间"
    },
    {
        "id": 1036000,
        "name": "夏凉",
        "keywords": "",
        "front_desc": "夏凉床品，舒适一夏",
        "parent_id": 1005000,
        "sort_order": 1,
        "show_index": 1,
        "is_show": 1,
        "banner_url": "http://yanxuan.nosdn.127.net/13ff4decdf38fe1a5bde34f0e0cc635a.png",
        "icon_url": "",
        "img_url": "",
        "wap_banner_url": "http://yanxuan.nosdn.127.net/bd17c985bacb9b9ab1ab6e9d66ee343c.png",
        "level": "L2",
        "type": 0,
        "front_name": "夏凉床品，舒适一夏"
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
