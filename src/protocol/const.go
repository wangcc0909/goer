package protocol

type HuPaiType int

//OpType
const (
	OptypeIllegal = 0
	OptypeChu     = 1
	OptypePeng    = 2
	OptypeGang    = 3
	OptypeHu      = 4
	OptypePass    = 5

	OptyMoPai = 500
	//以下三种杠的分类主要用以解决上面的 OptypeGang分类不细致,导致抢杠等操作处理麻烦的问题
	//在判定时必须满足两条件 x % 10 == 4 && x >1000
	OptypeAnGang   = 1004
	OptypeMingGang = 1014
	OptypeBaGang   = 1024
)

const (
	HuTypeDianPao HuPaiType = iota
	HuTypeZiMo
	HuTypePei
)

const (
	ExitTypeExitDeskUI           = -1
	ExitTypeDissolve             = 6
	ExitTypeSelfRequest          = 0
	ExitTypeClassicCoinNotEnough = 1
	ExitTypeDailyMatchEnd        = 2
	ExitTypeNotReadyForStart     = 3
	ExitTypeChangeDesk           = 4
	ExitTypeRepeatLogin          = 5
)
