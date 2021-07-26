package client

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/gotomicro/gopay/common"
	"github.com/gotomicro/gopay/util"
	"io"
	"net/http"
	"time"
)

var defaultWechatMiniProgramClient *WechatMiniProgramClient

const contentTypeJson = "application/json"

func InitWxMiniProgramClient(c *WechatMiniProgramClient) {
	if len(c.PrivateKey) != 0 && len(c.PublicKey) != 0 {
		c.httpsClient = NewHTTPSClient(c.PublicKey, c.PrivateKey)
	}

	defaultWechatMiniProgramClient = c
}

func DefaultWechatMiniProgramClient() *WechatMiniProgramClient {
	return defaultWechatMiniProgramClient
}

// WechatMiniProgramClient 微信小程序
type WechatMiniProgramClient struct {
	AppID       string       // 公众账号ID
	MchID       string       // 商户号ID
	Key         string       // 密钥

	Secret      string       // APP Secret 用于获得 token
	Spappid string // 第三方开票平台 ID
	Phone string // 商家联系方式

	PrivateKey  []byte       // 私钥文件内容
	PublicKey   []byte       // 公钥文件内容
	httpsClient *HTTPSClient // 双向证书链接
}

func GetInvoiceAuthUrl(orderId string,
	amount int64, redirectUrl string) (*common.WechatMiniProgramGetInvoiceAuthUrlResp, error) {
	client := DefaultWechatMiniProgramClient()
	token, err := client.GetToken()
	if err != nil {
		return nil, fmt.Errorf("GetInvoiceAuthUrl get token got error %w", err)
	}
	if !token.Ok() {
		return nil, errors.New(fmt.Sprintf("could not get token: %d, %s", token.ErrCode, token.ErrMsg))
	}

	ticket, err := client.GetTicket(token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("GetInvoiceAuthUrl get ticket got error %w", err)
	}
	return client.GetInvoiceAuthUrl(orderId, amount, redirectUrl, ticket.Ticket)
}

func (this *WechatMiniProgramClient) GetInvoiceAuthUrl(orderId string,
	amount int64, redirectUrl string, ticket string) (*common.WechatMiniProgramGetInvoiceAuthUrlResp, error){
	req := &common.WechatMiniProgramGetInvoiceAuthUrlReq{
		Spappid: this.Spappid,
		OrderId: orderId,
		Money: amount,
		Timestamp: time.Now().Unix(),
		Source: "wxa",
		RedirectUrl: redirectUrl,
		Ticket: ticket,
		Type: 1,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post("", contentTypeJson, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	authUrlResp := &common.WechatMiniProgramGetInvoiceAuthUrlResp{}
	err = json.Unmarshal(respBody, authUrlResp)
	return authUrlResp, err
}

func (this *WechatMiniProgramClient) GetToken() (*common.WechatMiniProgramToken, error) {
	const tokenUrlPattern = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"

	url := fmt.Sprintf(tokenUrlPattern, this.AppID, this.Secret)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	token := &common.WechatMiniProgramToken{}
	err = json.Unmarshal(data, token)
	return token, err
}

func (this *WechatMiniProgramClient) GetTicket(accessToken string) (*common.WechatMiniProgramTicket, error) {
	const urlPattern = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token==%s&type=wx_card"
	resp, err := http.Get(fmt.Sprintf(urlPattern, accessToken))
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ticket := &common.WechatMiniProgramTicket{}
	err = json.Unmarshal(body, ticket)
	return ticket, err
}

// SetContact 设置商家联系方式
func (this *WechatMiniProgramClient) SetContact(accessToken string) (*common.BaseResp, error) {
	const urlPattern = "https://api.weixin.qq.com/card/invoice/setbizattr?action=set_contact&access_token=%s"
	contact := &common.SetContactReq{
		Contact: common.Contact{
			Phone: this.Phone,
			// 十分钟超时
			TimeOut: 600,
		},
	}

	body, err := json.Marshal(contact)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(fmt.Sprintf(urlPattern, accessToken),
		contentTypeJson,
		bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	baseResp := &common.BaseResp{}
	respData, err  := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(respData, baseResp)
	return baseResp, err
}

// Pay 支付
func (this *WechatMiniProgramClient) Pay(charge *common.Charge) (map[string]string, error) {
	var m = make(map[string]string)
	appId := this.AppID
	if charge.APPID != "" {
		appId = charge.APPID
	}
	m["appid"] = appId
	m["mch_id"] = this.MchID
	m["nonce_str"] = util.RandomStr()
	m["body"] = TruncatedText(charge.Describe, 32)
	m["out_trade_no"] = charge.TradeNum
	m["total_fee"] = WechatMoneyFeeToString(charge.MoneyFee)
	m["spbill_create_ip"] = util.LocalIP()
	m["notify_url"] = charge.CallbackURL
	m["trade_type"] = "JSAPI"
	m["openid"] = charge.OpenID
	m["sign_type"] = "MD5"

	sign, err := WechatGenSign(this.Key, m)
	if err != nil {
		return map[string]string{}, err
	}
	m["sign"] = sign

	// 转出xml结构
	xmlRe, err := PostWechat("https://api.mch.weixin.qq.com/pay/unifiedorder", m, nil)
	if err != nil {
		return map[string]string{}, err
	}

	var c = make(map[string]string)
	c["appId"] = appId
	c["timeStamp"] = fmt.Sprintf("%d", time.Now().Unix())
	c["nonceStr"] = util.RandomStr()
	c["package"] = fmt.Sprintf("prepay_id=%s", xmlRe.PrepayID)
	c["signType"] = "MD5"
	sign2, err := WechatGenSign(this.Key, c)
	if err != nil {
		return map[string]string{}, errors.New("WechatWeb: " + err.Error())
	}
	c["paySign"] = sign2
	delete(c, "appId")
	return c, nil
}

// 关闭订单
func (this *WechatMiniProgramClient) CloseOrder(outTradeNo string) (common.WeChatQueryResult, error) {
	return WachatCloseOrder(this.AppID, this.MchID, this.Key, outTradeNo)
}

// 支付到用户的微信账号
func (this *WechatMiniProgramClient) PayToClient(charge *common.Charge) (map[string]string, error) {
	return WachatCompanyChange(this.AppID, this.MchID, this.Key, this.httpsClient, charge)
}

// QueryOrder 查询订单
func (this *WechatMiniProgramClient) QueryOrder(tradeNum string) (common.WeChatQueryResult, error) {
	return WachatQueryOrder(this.AppID, this.MchID, this.Key, tradeNum)
}

func (this *WechatMiniProgramClient) Refund(charge *common.RefundCharge) (map[string]string, error) {
	var m = make(map[string]string)
	m["appid"] = this.AppID
	m["mch_id"] = this.MchID
	m["nonce_str"] = util.RandomStr()
	//m["body"] = TruncatedText(charge.Describe, 32)
	m["out_trade_no"] = charge.OutTradeNo
	m["transaction_id"] = charge.TransactionId
	m["out_refund_no"] = charge.RefundSn
	m["total_fee"] = fmt.Sprintf("%d", charge.TotalFee)
	m["refund_fee"] = fmt.Sprintf("%d", charge.RefundFee)
	m["notify_url"] = charge.NotifyUrl
	m["sign_type"] = "MD5"

	sign, err := WechatGenSign(this.Key, m)
	if err != nil {
		return nil, err
	}
	m["sign"] = sign

	// 转出xml结构
	xmlRe, err := PostWechat("https://api.mch.weixin.qq.com/secapi/pay/refund", m, this.httpsClient)
	if err != nil {
		return map[string]string{}, err
	}

	var c = make(map[string]string)
	c["transaction_id"] = xmlRe.TransactionID
	c["return_code"] = xmlRe.ReturnCode
	c["return_msg"] = xmlRe.ReturnMsg
	c["refund_id"] = xmlRe.RefundId

	return c, nil
}

func (this *WechatMiniProgramClient) AuthInvoiceCallbackHandle(request *http.Request) (*common.WechatAuthInvoiceResult, error){
	respBody, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}
	result := &common.WechatAuthInvoiceResult{}
	err = xml.Unmarshal(respBody, result)
	return result, err
}