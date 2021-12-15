package Config

type Linx struct {
	AddressLinhao string `sql:"address" json:"address" form:"address"`
	AmountLinhao int64  `sql:"amount" json:"amount" form:"amount"`
	IpLinhao     string `sql:"ip" json:"ip" form:"ip"`

}