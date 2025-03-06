package enums

type Category string

const (
	CategoryInvalid    Category = "invalid"
	CategoryBreakfast  Category = "breakfast"
	CategoryLunch      Category = "lunch"
	CategoryDinner     Category = "dinner"
	CategorySnack      Category = "snack"
	CategoryDessert    Category = "dessert"
	CategoryDrink      Category = "drink"
	CategoryOther      Category = "other"
	CategoryItalian    Category = "italian"
	CategoryJapanese   Category = "japanese"
	CategoryKorean     Category = "korean"
	CategoryChinese    Category = "chinese"
	CategoryAmerican   Category = "american"
	CategoryMexican    Category = "mexican"
	CategoryThai       Category = "thai"
	CategoryVietnamese Category = "vietnamese"
	CategoryIndian     Category = "indian"
	CategoryFrench     Category = "french"
	CategoryGerman     Category = "german"
	CategorySpanish    Category = "spanish"
	CategoryBrazilian  Category = "brazilian"
	CategoryTaiwanese  Category = "taiwanese"
)

func (c Category) String() string {
	return string(c)
}

func (c Category) ToCategoryCode() CategoryCode {
	categoryCode, ok := CategoryToCategoryCodeMap[c]
	if !ok {
		return CategoryCodeInvalid
	}
	return categoryCode
}

func (c Category) IsValid() bool {
	if c == CategoryInvalid {
		return false
	}
	_, ok := CategoryToCategoryCodeMap[c]
	return ok
}

var Categories = []Category{
	CategoryInvalid,
	CategoryBreakfast,
	CategoryLunch,
	CategoryDinner,
	CategorySnack,
	CategoryDessert,
	CategoryDrink,
	CategoryOther,
	CategoryItalian,
	CategoryJapanese,
	CategoryKorean,
	CategoryChinese,
	CategoryAmerican,
	CategoryMexican,
	CategoryThai,
	CategoryVietnamese,
	CategoryIndian,
	CategoryFrench,
	CategoryGerman,
	CategorySpanish,
	CategoryBrazilian,
	CategoryTaiwanese,
}

type CategoryCode int

const (
	CategoryCodeInvalid CategoryCode = iota
	CategoryCodeBreakfast
	CategoryCodeLunch
	CategoryCodeDinner
	CategoryCodeSnack
	CategoryCodeDessert
	CategoryCodeDrink
	CategoryCodeOther
	CategoryCodeItalian
	CategoryCodeJapanese
	CategoryCodeKorean
	CategoryCodeChinese
	CategoryCodeAmerican
	CategoryCodeMexican
	CategoryCodeThai
	CategoryCodeVietnamese
	CategoryCodeIndian
	CategoryCodeFrench
	CategoryCodeGerman
	CategoryCodeSpanish
	CategoryCodeBrazilian
	CategoryCodeTaiwanese
)

func (c CategoryCode) ToCategory() Category {
	if int(c) >= len(Categories) || int(c) < 0 {
		return CategoryInvalid
	}
	return Categories[c]
}

var CategoryToCategoryCodeMap = map[Category]CategoryCode{
	CategoryInvalid:    CategoryCodeInvalid,
	CategoryBreakfast:  CategoryCodeBreakfast,
	CategoryLunch:      CategoryCodeLunch,
	CategoryDinner:     CategoryCodeDinner,
	CategorySnack:      CategoryCodeSnack,
	CategoryDessert:    CategoryCodeDessert,
	CategoryDrink:      CategoryCodeDrink,
	CategoryOther:      CategoryCodeOther,
	CategoryItalian:    CategoryCodeItalian,
	CategoryJapanese:   CategoryCodeJapanese,
	CategoryKorean:     CategoryCodeKorean,
	CategoryChinese:    CategoryCodeChinese,
	CategoryAmerican:   CategoryCodeAmerican,
	CategoryMexican:    CategoryCodeMexican,
	CategoryThai:       CategoryCodeThai,
	CategoryVietnamese: CategoryCodeVietnamese,
	CategoryIndian:     CategoryCodeIndian,
	CategoryFrench:     CategoryCodeFrench,
	CategoryGerman:     CategoryCodeGerman,
	CategorySpanish:    CategoryCodeSpanish,
	CategoryBrazilian:  CategoryCodeBrazilian,
	CategoryTaiwanese:  CategoryCodeTaiwanese,
}

var CategoryCodes = []CategoryCode{
	CategoryCodeInvalid,
	CategoryCodeBreakfast,
	CategoryCodeLunch,
	CategoryCodeDinner,
	CategoryCodeSnack,
	CategoryCodeDessert,
	CategoryCodeDrink,
	CategoryCodeOther,
	CategoryCodeItalian,
	CategoryCodeJapanese,
	CategoryCodeKorean,
	CategoryCodeChinese,
	CategoryCodeAmerican,
	CategoryCodeMexican,
	CategoryCodeThai,
	CategoryCodeVietnamese,
	CategoryCodeIndian,
	CategoryCodeFrench,
	CategoryCodeGerman,
	CategoryCodeSpanish,
	CategoryCodeBrazilian,
	CategoryCodeTaiwanese,
}
