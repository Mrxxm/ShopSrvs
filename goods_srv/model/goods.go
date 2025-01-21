package model

type Category struct {
	Name        string `gorm:"column:name;not null;default:'';type:varchar(20) comment '分类名称'"`
	Level       int32  `gorm:"column:level;not null;default:1;type:int comment '级别'"`
	IsTab       int8   `gorm:"column:is_tab;not null;default:0;type:tinyint comment '是否展示在tab栏 0否 1是'"`
	PCategoryID int32  `gorm:"column:pid;type:int comment '父级id'"`
	PCategory   *Category
	BaseModel
}

type Brands struct {
	Name string `gorm:"column:name;not null;type:varchar(30) comment '品牌名称'"`
	Logo string `gorm:"column:logo;not null;default:'';type:varchar(255) comment 'logo链接'"`

	BaseModel
}

// 品牌和分类是多对多关系
type GoodsCategoryBrand struct {
	CategoryID int32 `gorm:"type:int;index:index_category_brand,unique"`
	Category   Category
	BrandsID   int32 `gorm:"type:int;index:index_category_brand,unique"`
	Brands     Brands

	BaseModel
}

func (GoodsCategoryBrand) TableName() string {
	return "goods_category_brand"
}

type Banner struct {
	Image string `gorm:"column:image;not null;type:varchar(200) comment '轮播图链接'"`
	Url   string `gorm:"column:url;not null;type:varchar(200) comment '轮播图调整至商品链接'"`
	Index int32  `gorm:"column:index;not null;default:1;type:int comment '排序'"`

	BaseModel
}

type Goods struct {
	CategoryID int32 `gorm:"type:int comment '分类id'"`
	Category   Category
	BrandsID   int32 `gorm:"type:int comment '品牌id'"`
	Brands     Brands

	OnSale int32 `gorm:"column:on_sale;default:0;not null;type:tinyint(1) comment '是否在售 0 否 1 是'"`
	IsFree int32 `gorm:"column:is_free;default:0;not null;type:tinyint(1) comment '是否免运费 0 否 1 是'"`
	IsNew  int32 `gorm:"column:is_new;default:0;not null;type:tinyint(1) comment '是否新品 0 否 1 是'"`
	IsHot  int32 `gorm:"column:is_hot;default:0;not null;type:tinyint(1) comment '是否热卖商品 0 否 1 是'"`

	Name             string   `gorm:"not null;type:varchar(100) comment '商品名称'"`
	GoodsSn          string   `gorm:"not null;default:'';type:varchar(50) comment '商品编码'"`
	ClickNum         int32    `gorm:"not null;default:0;type:int comment '点击数量'"`
	SoldNum          int32    `gorm:"not null;default:0;type:int comment '销售数量'"`
	FavNum           int32    `gorm:"not null;default:0;type:int comment '收藏数量'"`
	MarketPrice      float32  `gorm:"not null;type:float comment '市场价'"`
	ShopPrice        float32  `gorm:"not null;type:float comment '销售价'"`
	GoodsBrief       string   `gorm:"not null;default:'';type:varchar(100) comment '商品简介'"`
	Images           GormList `gorm:"not null;type:json comment '商品详情页轮播图'"`
	DescImages       GormList `gorm:"not null;type:json comment '商品详情页下拉图'"`
	GoodsFrontImages GormList `gorm:"not null;type:varchar(1000) comment '商品首图'"`

	BaseModel
}
