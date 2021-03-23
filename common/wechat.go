package common

// WeChatResult 微信支付返回
type WeChatReResult struct {
	PrepayID string `xml:"prepay_id" json:"prepay_id,omitempty"`
	CodeURL  string `xml:"code_url" json:"code_url,omitempty"`
}

// WechatBaseResult 基本信息
type WechatBaseResult struct {
	ReturnCode string `xml:"return_code" json:"return_code,omitempty"`
	ReturnMsg  string `xml:"return_msg" json:"return_msg,omitempty"`
}

// WechatReturnData 返回通用数据
type WechatReturnData struct {
	AppID      string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchID      string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	MchAppid   string `xml:"mch_appid,omitempty" json:"mch_appid,omitempty"`
	DeviceInfo string `xml:"device_info,omitempty" json:"device_info,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty"`
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
}

// WechatResultData 结果通用数据
type WechatResultData struct {
	OpenID         string `xml:"openid,omitempty" json:"openid,omitempty"`
	IsSubscribe    string `xml:"is_subscribe,omitempty" json:"is_subscribe,omitempty"`
	TradeType      string `xml:"trade_type,omitempty" json:"trade_type,omitempty"`
	BankType       string `xml:"bank_type,omitempty" json:"bank_type,omitempty"`
	FeeType        string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`
	TotalFee       string `xml:"total_fee,omitempty" json:"total_fee,omitempty"`
	CashFeeType    string `xml:"cash_fee_type,omitempty" json:"cash_fee_type,omitempty"`
	CashFee        string `xml:"cash_fee,omitempty" json:"cash_fee,omitempty"`
	TransactionID  string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`
	OutTradeNO     string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
	Attach         string `xml:"attach,omitempty" json:"attach,omitempty"`
	TimeEnd        string `xml:"time_end,omitempty" json:"time_end,omitempty"`
	PartnerTradeNo string `xml:"partner_trade_no,omitempty" json:"partner_trade_no,omitempty"`
	PaymentNo      string `xml:"payment_no,omitempty" json:"payment_no,omitempty"`
	PaymentTime    string `xml:"payment_time,omitempty" json:"payment_time,omitempty"`
	DetailId       string `xml:"detail_id,omitempty" json:"detail_id,omitempty"`
}

type WeChatPayResult struct {
	WechatBaseResult
	WechatReturnData
	WechatResultData
	WechatRefundResult
	OutTradeNO     string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
}

type WechatRefundResult struct {
	RefundStatus string `xml:"refund_status" json:"refund_status"`
	SuccessTime  string `xml:"success_time" json:"success_time"`
	RefundId     string `xml:"refund_id" json:"refund_id"` // wechat refund id
	OutRefundNo  string `xml:"out_refund_no"json:"out_refund_no"`
	OutTradeNO     string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`
}

type WeChatQueryResult struct {
	WechatBaseResult
	WeChatReResult
	WechatReturnData
	WechatResultData
	TradeState     string `xml:"trade_state" json:"trade_state,omitempty"`
	TradeStateDesc string `xml:"trade_state_desc" json:"trade_state_desc,omitempty"`

	RefundId string `json:"refund_id"` // wechat refund id
}

type WechatRefundCallbackResp struct {
	WechatBaseResult
	WechatRefundResult
}

type WechatRefundResultOriginalResp struct {
	WechatBaseResult
	AppID    string `xml:"appid,omitempty" json:"appid,omitempty"`
	MchID    string `xml:"mch_id,omitempty" json:"mch_id,omitempty"`
	NonceStr string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty"`
	ReqInfo  string `xml:"req_info" json:"req_info"`
}
