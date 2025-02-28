package enums

type BanksBranch string

const (
	BanksBranchInvalid         BanksBranch = "invalid"
	BanksBranchTaipeiSongshan  BanksBranch = "台北市松山區"
	BanksBranchTaipeiZhongshan BanksBranch = "台北市中山區"
	BanksBranchTaipeiXinyi     BanksBranch = "台北市信義區"
	BanksBranchTaipeiWenshan   BanksBranch = "台北市文山區"
	BanksBranchTaipeiNangang   BanksBranch = "台北市南港區"
	BanksBranchTaipeiBeitou    BanksBranch = "台北市北投區"
	BanksBranchTaipeiWanhua    BanksBranch = "台北市萬華區"
)

func (b BanksBranch) ToString() string {
	return string(b)
}

func (b BanksBranch) IsValid() bool {
	active, ok := BanksBranchActiveMap[b]
	return ok && active
}

var BanksBranchActiveMap = map[BanksBranch]bool{
	BanksBranchTaipeiSongshan:  true,
	BanksBranchTaipeiZhongshan: true,
	BanksBranchTaipeiXinyi:     true,
	BanksBranchTaipeiWenshan:   true,
	BanksBranchTaipeiNangang:   true,
	BanksBranchTaipeiBeitou:    true,
	BanksBranchTaipeiWanhua:    true,
}

func (b BanksBranch) ToBanksBranchCode() BanksBranchCode {
	code, ok := BanksBranchToCodeMap[b]
	if !ok {
		return BanksBranchCodeInvalid
	}
	return code
}

var BanksBranchToCodeMap = map[BanksBranch]BanksBranchCode{
	BanksBranchTaipeiSongshan:  BanksBranchCodeTaipeiSongshan,
	BanksBranchTaipeiZhongshan: BanksBranchCodeTaipeiZhongshan,
	BanksBranchTaipeiXinyi:     BanksBranchCodeTaipeiXinyi,
	BanksBranchTaipeiWenshan:   BanksBranchCodeTaipeiWenshan,
	BanksBranchTaipeiNangang:   BanksBranchCodeTaipeiNangang,
	BanksBranchTaipeiBeitou:    BanksBranchCodeTaipeiBeitou,
	BanksBranchTaipeiWanhua:    BanksBranchCodeTaipeiWanhua,
}

type BanksBranchCode int

const (
	BanksBranchCodeInvalid BanksBranchCode = iota - 1
	BanksBranchCodeTaipeiSongshan
	BanksBranchCodeTaipeiZhongshan
	BanksBranchCodeTaipeiXinyi
	BanksBranchCodeTaipeiWenshan
	BanksBranchCodeTaipeiNangang
	BanksBranchCodeTaipeiBeitou
	BanksBranchCodeTaipeiWanhua
)

func (b BanksBranchCode) ToBanksBranch() BanksBranch {
	banksBranch, ok := BanksBranchCodeToBanksBranchMap[b]
	if !ok {
		return BanksBranchInvalid
	}
	return banksBranch
}

var BanksBranchCodeToBanksBranchMap = map[BanksBranchCode]BanksBranch{
	BanksBranchCodeTaipeiSongshan:  BanksBranchTaipeiSongshan,
	BanksBranchCodeTaipeiZhongshan: BanksBranchTaipeiZhongshan,
	BanksBranchCodeTaipeiXinyi:     BanksBranchTaipeiXinyi,
	BanksBranchCodeTaipeiWenshan:   BanksBranchTaipeiWenshan,
	BanksBranchCodeTaipeiNangang:   BanksBranchTaipeiNangang,
	BanksBranchCodeTaipeiBeitou:    BanksBranchTaipeiBeitou,
	BanksBranchCodeTaipeiWanhua:    BanksBranchTaipeiWanhua,
}
