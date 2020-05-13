<template>
	<view class="content">
		<view class="cate-nav" :style="{position:headerPosition,top:headerTop}">
			<scroll-view scroll-x class="cate-nav-body" style="width: 750upx;" :scroll-left="scrollLeft">
				<view   v-for="item in navList"  :class="{active: id ===item.id}" class="item"  @click="switchCate(item)">
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
					<navigator v-for="(item,index) in goodsList" :class="{'item-b': (index+1) % 2 === 0 }" class="item" @click="navToList(item.id)">
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
				currentId: 0,
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
        "id": 1147048,
        "name": "简约知性系列居家地毯 蓝灰格",
        "list_pic_url": "http://yanxuan.nosdn.127.net/fd7920a2eadd10fa10c0c03959a2abe0.png",
        "retail_price": 559
    },
    {
        "id": 1147047,
        "name": "简约知性系列居家地毯 蓝粉拼接",
        "list_pic_url": "http://yanxuan.nosdn.127.net/bda805b0a2464b6ec33c18981565e50e.png",
        "retail_price": 559
    },
    {
        "id": 1147046,
        "name": "清新趣粉系列居家地毯 条纹间粉",
        "list_pic_url": "http://yanxuan.nosdn.127.net/655d718df8107f8e7fd1dc6140e29039.png",
        "retail_price": 599
    },
    {
        "id": 1147045,
        "name": "清新趣粉系列居家地毯 灰黄条纹",
        "list_pic_url": "http://yanxuan.nosdn.127.net/5cda4a0c4c4ff9728d03186bd053c9ca.png",
        "retail_price": 599
    },
    {
        "id": 1138001,
        "name": "北欧简约山形纹绣花抱枕",
        "list_pic_url": "http://yanxuan.nosdn.127.net/dbc5b25b824c3b3d7ff43b56ca35eee9.png",
        "retail_price": 79
    },
    {
        "id": 1138000,
        "name": "日式蓬软太鼓抱枕",
        "list_pic_url": "http://yanxuan.nosdn.127.net/ad953e16ad8c33b714e7af941ce8cd56.png",
        "retail_price": 29
    },
    {
        "id": 1135058,
        "name": "日式多功能手卷午睡枕坐垫",
        "list_pic_url": "http://yanxuan.nosdn.127.net/37bc0fa3524a904ac740340fa92bd515.png",
        "retail_price": 79
    },
    {
        "id": 1135057,
        "name": "梦幻系简约轻透莹白纱帘",
        "list_pic_url": "http://yanxuan.nosdn.127.net/98c5e80b8e328687ce9c949314ebc41d.png",
        "retail_price": 199
    },
    {
        "id": 1135056,
        "name": "糖果色凹凸条纹儿童房窗帘",
        "list_pic_url": "http://yanxuan.nosdn.127.net/536246ca4adb77274a94b18bb2f91f47.png",
        "retail_price": 259
    },
    {
        "id": 1135055,
        "name": "北欧印象几何条纹混色窗帘",
        "list_pic_url": "http://yanxuan.nosdn.127.net/87b6a608b99279ebf1764682e9e5fcec.png",
        "retail_price": 399
    },
    {
        "id": 1135054,
        "name": "美式田园风蜻蜓提花窗帘",
        "list_pic_url": "http://yanxuan.nosdn.127.net/30d7daa0824fbb61b6c36175c8203349.png",
        "retail_price": 559
    },
    {
        "id": 1135053,
        "name": "法式复古山形纹提花窗帘",
        "list_pic_url": "http://yanxuan.nosdn.127.net/1f9e34b1aadd14ea0c9c299c530d86ba.png",
        "retail_price": 429
    },
    {
        "id": 1135052,
        "name": "日式简约素色窗帘",
        "list_pic_url": "http://yanxuan.nosdn.127.net/63f5da1f5363af43aa91864bf66af48e.png",
        "retail_price": 259
    },
    {
        "id": 1135051,
        "name": "日式素雅纯色流星纹窗帘",
        "list_pic_url": "http://yanxuan.nosdn.127.net/9126151f028a8804026d530836b481cb.png",
        "retail_price": 299
    },
    {
        "id": 1134032,
        "name": "趣味粉彩系列记忆棉坐垫",
        "list_pic_url": "http://yanxuan.nosdn.127.net/8b30eeb17c831eba08b97bdcb4c46a8e.png",
        "retail_price": 49
    },
    {
        "id": 1134030,
        "name": "简约知性记忆棉坐垫",
        "list_pic_url": "http://yanxuan.nosdn.127.net/aa49dfe878becf768eddc4c1636643a6.png",
        "retail_price": 46
    },
    {
        "id": 1134022,
        "name": "清新趣粉防滑浴垫",
        "list_pic_url": "http://yanxuan.nosdn.127.net/a2b7489b4a2b1c09b66464cede4dabd7.png",
        "retail_price": 79
    },
    {
        "id": 1130042,
        "name": "皮毛一体多用单张长毛皮垫",
        "list_pic_url": "http://yanxuan.nosdn.127.net/dc9d09334eb201fe9408ed604e549941.png",
        "retail_price": 239
    },
    {
        "id": 1130041,
        "name": "皮毛一体多用长毛坐垫",
        "list_pic_url": "http://yanxuan.nosdn.127.net/442b9d99c0e7f39efd7967e0e5987374.png",
        "retail_price": 109
    },
    {
        "id": 1130037,
        "name": "帆布丝羽绒多用坐垫",
        "list_pic_url": "http://yanxuan.nosdn.127.net/19ecd7c6f6f31219cf75117238d95139.png",
        "retail_price": 39
    },
    {
        "id": 1128002,
        "name": "清新趣粉系列居家地毯 青粉拼接",
        "list_pic_url": "http://yanxuan.nosdn.127.net/a1094a808ffb3a52a6cb13565a283d98.png",
        "retail_price": 599
    },
    {
        "id": 1115052,
        "name": "日式和风蔺草蒲团坐垫",
        "list_pic_url": "http://yanxuan.nosdn.127.net/39dea35a3ea2361e4b054ee2f421af53.png",
        "retail_price": 86
    },
    {
        "id": 1081002,
        "name": "北欧风珊瑚绒多功能抱枕",
        "list_pic_url": "http://yanxuan.nosdn.127.net/380cfcd5d8bc22360de089f0b4eb11da.png",
        "retail_price": 89
    },
    {
        "id": 1081000,
        "name": "北欧风珊瑚绒多功能暖手枕",
        "list_pic_url": "http://yanxuan.nosdn.127.net/cc45baafad00405699552c187c64c512.png",
        "retail_price": 49
    },
    {
        "id": 1072001,
        "name": "色织水洗棉绣花抱枕套",
        "list_pic_url": "http://yanxuan.nosdn.127.net/0e9d5954d7dc2477d9c46b730e05ab42.png",
        "retail_price": 49
    },
    {
        "id": 1072000,
        "name": "手工针织绞花抱枕套",
        "list_pic_url": "http://yanxuan.nosdn.127.net/87cf3a17ad40bfdcdc3314ea4591a5e8.png",
        "retail_price": 89
    },
    {
        "id": 1057036,
        "name": "日式纯色水洗亚麻抱枕",
        "list_pic_url": "http://yanxuan.nosdn.127.net/8a9ee5ba08929cc9e40b973607d2f633.png",
        "retail_price": 79
    },
    {
        "id": 1055016,
        "name": "日式纯棉针织条纹抱枕",
        "list_pic_url": "http://yanxuan.nosdn.127.net/23e0203f1512f33e605f61c28fa03d2d.png",
        "retail_price": 59
    },
    {
        "id": 1055012,
        "name": "300根全棉羽丝绒抱枕芯",
        "list_pic_url": "http://yanxuan.nosdn.127.net/3d437c8d68e2ec3f3dd61001bf98f16e.png",
        "retail_price": 39
    },
    {
        "id": 1048005,
        "name": "日式色织水洗条纹抱枕",
        "list_pic_url": "http://yanxuan.nosdn.127.net/ce980c16810a471dffff6aa8d7bac754.png",
        "retail_price": 59
    },
    {
        "id": 1043005,
        "name": "日式记忆绵坐垫",
        "list_pic_url": "http://yanxuan.nosdn.127.net/2a95b16f5b147cab4845641bee738a2e.png",
        "retail_price": 59
    },
    {
        "id": 1039051,
        "name": "多功能午睡枕",
        "list_pic_url": "http://yanxuan.nosdn.127.net/c8ca0600fa7ba11ca8be6a3173dd38c9.png",
        "retail_price": 79
    },
    {
        "id": 1035006,
        "name": "全棉单面割绒浴室地垫",
        "list_pic_url": "http://yanxuan.nosdn.127.net/ee92704f3b8323905b51fc647823e6e5.png",
        "retail_price": 56
    },
    {
        "id": 1030003,
        "name": "160*230羊毛手工几何地毯",
        "list_pic_url": "http://yanxuan.nosdn.127.net/1d1ab099dc0e254c15e57302e78e200b.png",
        "retail_price": 1469
    },
    {
        "id": 1030002,
        "name": "160*230羊毛圈绒枪刺地毯",
        "list_pic_url": "http://yanxuan.nosdn.127.net/8b9328496990357033d4259fda250679.png",
        "retail_price": 899
    },
    {
        "id": 1030001,
        "name": "160*230羊毛手工地毯",
        "list_pic_url": "http://yanxuan.nosdn.127.net/88dc5d80c6f84102f003ecd69c86e1cf.png",
        "retail_price": 969
    },
    {
        "id": 1020000,
        "name": "升级款记忆绵护椎腰靠",
        "list_pic_url": "http://yanxuan.nosdn.127.net/819fdf1f635a694166bcfdd426416e8c.png",
        "retail_price": 79
    },
    {
        "id": 1015007,
        "name": "典雅美式全棉刺绣抱枕",
        "list_pic_url": "http://yanxuan.nosdn.127.net/a2045004de8a6225289376ad54317fc8.png",
        "retail_price": 59
    },
    {
        "id": 1009024,
        "name": "日式和风懒人沙发",
        "list_pic_url": "http://yanxuan.nosdn.127.net/149dfa87a7324e184c5526ead81de9ad.png",
        "retail_price": 599
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
			this.headerTop = document.getElementsByTagName('uni-page-head')[0].offsetHeight+'px';
		},
		methods: {
			switchCate(item) {
				this.id = item.id;
			},
			navToList(pid) {
				uni.navigateTo({
					url:`/pages/goods/goods?id=${pid}`
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
