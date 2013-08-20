package global

import (
	"github.com/kobeld/duoerlapi"
)

// Product Got From
const (
	FROM_STROE  string = "from_01"
	FROM_ONLINE        = "from_02"
)

var GotFromOptions = map[string]string{
	FROM_STROE:  "实体店铺",
	FROM_ONLINE: "网上购买",
}

// User Skin Texture
const (
	SKIN_01 string = "skin_01"
	SKIN_02        = "skin_02"
	SKIN_03        = "skin_03"
	SKIN_04        = "skin_04"
	SKIN_05        = "skin_05"
)

var SkinTextureOptions = map[string]string{
	SKIN_01: "中性肤质",
	SKIN_02: "干性肤质",
	SKIN_03: "油性肤质",
	SKIN_04: "混合肤质",
	SKIN_05: "敏感肤质",
}

// User Hair Texture
const (
	HAIR_01 string = "hair_01"
	HAIR_02        = "hair_02"
	HAIR_03        = "hair_03"
	HAIR_04        = "hair_04"
	HAIR_05        = "hair_05"
	HAIR_06        = "hair_06"
	HAIR_07        = "hair_07"
	HAIR_08        = "hair_08"
)

var HairTextureOptions = map[string]string{
	HAIR_01: "中性发质",
	HAIR_02: "干性发质",
	HAIR_03: "油性发质",
	HAIR_04: "混合发质",
	HAIR_05: "受损发质",
	HAIR_06: "头屑发质",
	HAIR_07: "暗哑发质",
	HAIR_08: "脱落发质",
}

// Review Rating
const (
	RATING_01 string = "1"
	RATING_02        = "2"
	RATING_03        = "3"
	RATING_04        = "4"
	RATING_05        = "5"
)

var RatingOptions = map[string]string{
	RATING_01: "很差",
	RATING_02: "较差",
	RATING_03: "还行",
	RATING_04: "推荐",
	RATING_05: "力荐",
}

// Cached data
var (
	Categories     []*duoerlapi.Category
	CategoryMap    map[string]*duoerlapi.Category
	SubCategoryMap map[string]*duoerlapi.SubCategory
	EfficacyMap    map[string]*duoerlapi.Efficacy
)
