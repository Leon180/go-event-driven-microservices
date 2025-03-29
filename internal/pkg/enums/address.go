package enums

type City string

const (
	CityInvalid   City = "invalid"
	CityTaipei    City = "taipei"     // 臺北市 Táiběi Shì
	CityNewTaipei City = "new_taipei" // 新北市 Xīnběi Shì
	CityTaoyuan   City = "taoyuan"    // 桃園市 Táoyuán Shì
	CityTaichung  City = "taichung"   // 臺中市 Táizhōng Shì
	CityTainan    City = "tainan"     // 臺南市 Táinán Shì
	CityKaohsiung City = "kaohsiung"  // 高雄市 Gāoxióng Shì
	CityHsinchu   City = "hsinchu"    // 新竹市 Xīnzhú Shì
	CityKeelung   City = "keelung"    // 基隆市 Jīlóng Shì
	CityChiayi    City = "chiayi"     // 嘉義市 Jiāyì Shì
	CityMiaoli    City = "miaoli"     // 苗栗市 Miáolì Shì
	CityChanghua  City = "changhua"   // 彰化市 Zhānghuà Shì
	CityNantou    City = "nantou"     // 南投市 Nántóu Shì
	CityDouliu    City = "douliu"     // 斗六市 Dǒuliù Shì
	CityPingtung  City = "pingtung"   // 屏東市 Píngdōng Shì
	CityYilan     City = "yilan"      // 宜蘭市 Yílán Shì
	CityHualien   City = "hualien"    // 花蓮市 Huālián Shì
	CityTaitung   City = "taitung"    // 臺東市 Táidōng Shì
	CityMagong    City = "magong"     // 馬公市 Mǎgōng Shì
	CityPuzi      City = "puzi"       // 朴子市 Púzǐ Shì
	CityTaibao    City = "taibao"     // 太保市 Tàibǎo Shì
	CityToufen    City = "toufen"     // 頭份市 Tóufèn Shì
	CityYuanlin   City = "yuanlin"    // 員林市 Yuánlín Shì
	CityZhubei    City = "zhubei"     // 竹北市 Zhúběi Shì
)

func (c City) ToString() string {
	return string(c)
}

func (c City) ToCityCode() CityCode {
	code, ok := CityToCityCodeMap[c]
	if !ok {
		return CityCodeInvalid
	}
	return code
}

func (c City) IsValid() bool {
	if c == CityInvalid {
		return false
	}
	_, ok := CityToCityCodeMap[c]
	return ok
}

var CityToCityCodeMap = map[City]CityCode{
	CityTaipei:    CityCodeTaipei,
	CityNewTaipei: CityCodeNewTaipei,
	CityTaoyuan:   CityCodeTaoyuan,
	CityTaichung:  CityCodeTaichung,
	CityTainan:    CityCodeTainan,
	CityKaohsiung: CityCodeKaohsiung,
	CityHsinchu:   CityCodeHsinchu,
	CityKeelung:   CityCodeKeelung,
	CityChiayi:    CityCodeChiayi,
	CityMiaoli:    CityCodeMiaoli,
	CityChanghua:  CityCodeChanghua,
	CityNantou:    CityCodeNantou,
	CityDouliu:    CityCodeDouliu,
	CityPingtung:  CityCodePingtung,
	CityYilan:     CityCodeYilan,
	CityHualien:   CityCodeHualien,
	CityTaitung:   CityCodeTaitung,
	CityMagong:    CityCodeMagong,
	CityPuzi:      CityCodePuzi,
	CityTaibao:    CityCodeTaibao,
	CityToufen:    CityCodeToufen,
	CityYuanlin:   CityCodeYuanlin,
	CityZhubei:    CityCodeZhubei,
}

var Cities = []City{
	CityInvalid,
	CityTaipei,
	CityNewTaipei,
	CityTaoyuan,
	CityTaichung,
	CityTainan,
	CityKaohsiung,
	CityHsinchu,
	CityKeelung,
	CityChiayi,
	CityMiaoli,
	CityChanghua,
	CityNantou,
	CityDouliu,
	CityPingtung,
	CityYilan,
	CityHualien,
	CityTaitung,
	CityMagong,
	CityPuzi,
	CityTaibao,
	CityToufen,
	CityYuanlin,
	CityZhubei,
}

type CityCode int

func (c CityCode) ToCity() City {
	if int(c) >= len(Cities) || int(c) < 0 {
		return CityInvalid
	}
	return Cities[c]
}

const (
	CityCodeInvalid CityCode = iota
	CityCodeTaipei
	CityCodeNewTaipei
	CityCodeTaoyuan
	CityCodeTaichung
	CityCodeTainan
	CityCodeKaohsiung
	CityCodeHsinchu
	CityCodeKeelung
	CityCodeChiayi
	CityCodeMiaoli
	CityCodeChanghua
	CityCodeNantou
	CityCodeDouliu
	CityCodePingtung
	CityCodeYilan
	CityCodeHualien
	CityCodeTaitung
	CityCodeMagong
	CityCodePuzi
	CityCodeTaibao
	CityCodeToufen
	CityCodeYuanlin
	CityCodeZhubei
)

type Country string

const (
	CountryInvalid Country = "invalid"
	CountryTaiwan  Country = "taiwan"
)

func (c Country) ToCountryCode() CountryCode {
	code, ok := CountryToCountryCodeMap[c]
	if !ok {
		return CountryCodeInvalid
	}
	return code
}

func (c Country) IsValid() bool {
	if c == CountryInvalid {
		return false
	}
	_, ok := CountryToCountryCodeMap[c]
	return ok
}

var Countries = []Country{
	CountryInvalid,
	CountryTaiwan,
}

func (c Country) ToString() string {
	return string(c)
}

type CountryCode int

const (
	CountryCodeInvalid CountryCode = iota
	CountryCodeTaiwan
)

var CountryToCountryCodeMap = map[Country]CountryCode{
	CountryTaiwan: CountryCodeTaiwan,
}

func (c CountryCode) ToCountry() Country {
	if int(c) >= len(Countries) || int(c) < 0 {
		return CountryInvalid
	}
	return Countries[c]
}
