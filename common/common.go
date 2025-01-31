package common

// PayClient 支付客户端接口
type PayClient interface {
	// 用户下单付款
	Pay(charge *Charge) (map[string]string, error)
	// 付款给用户
	PayToClient(charge *Charge) (map[string]string, error)
}

// Charge 支付参数
type Charge struct {
	APPID       string  `json:"-"`
	TradeNum    string  `json:"tradeNum,omitempty"`
	Origin      string  `json:"origin,omitempty"`
	UserID      string  `json:"userId,omitempty"`
	PayMethod   int64   `json:"payMethod,omitempty"`
	MoneyFee    float64 `json:"MoneyFee,omitempty"`
	CallbackURL string  `json:"callbackURL,omitempty"`
	ReturnURL   string  `json:"returnURL,omitempty"`
	ShowURL     string  `json:"showURL,omitempty"`
	Describe    string  `json:"describe,omitempty"`
	OpenID      string  `json:"openid,omitempty"`
	CheckName   bool    `json:"check_name,omitempty"`
	ReUserName  string  `json:"re_user_name,omitempty"`
	// 阿里提现
	AliAccount     string `json:"ali_account"`
	AliAccountType string `json:"ali_account_type"`
}

//PayCallback 支付返回
type PayCallback struct {
	Origin      string `json:"origin"`
	TradeNum    string `json:"trade_num"`
	OrderNum    string `json:"order_num"`
	CallBackURL string `json:"callback_url"`
	Status      int64  `json:"status"`
}

// CallbackReturn 回调业务代码时的参数
type CallbackReturn struct {
	IsSucceed     bool   `json:"isSucceed"`
	OrderNum      string `json:"orderNum"`
	TradeNum      string `json:"tradeNum"`
	UserID        string `json:"userID"`
	MoneyFee      int64  `json:"moneyFee"`
	Sign          string `json:"sign"`
	ThirdDiscount int64  `json:"thirdDiscount"`
}

// BaseResult 支付结果
type BaseResult struct {
	IsSucceed     bool   // 是否交易成功
	TradeNum      string // 交易流水号
	MoneyFee      int64  // 支付金额
	TradeTime     string // 交易时间
	ContractNum   string // 交易单号
	UserInfo      string // 支付账号信息(有可能有，有可能没有)
	ThirdDiscount int64  // 第三方优惠
}

type RefundCharge struct {
	TransactionId string `json:"transactionId"`

	// 相对于微信而言的外部交易号，即正向传递过去的
	OutTradeNo string `json:"outTradeNo"`
	// 内部退款单号，一般来说全额退款可以直接使用ordersn
	RefundSn   string `json:"refundSn"`
	TotalFee   int64  `json:"totalFee"`
	RefundFee  int64  `json:"refundFee"`
	RefundDesc string `json:"refundDesc"` // 退款描述
	NotifyUrl  string `json:"notifyUrl"`  // 退款回调URL

	PayMethod   int64   `json:"payMethod"` // 支付方式
}

