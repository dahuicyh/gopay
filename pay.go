package gopay

import (
	"errors"
	"github.com/gotomicro/gopay/client"
	"github.com/gotomicro/gopay/common"
	"github.com/gotomicro/gopay/constant"
)

// 用户下单支付接口
func Pay(charge *common.Charge) (map[string]string, error) {
	err := checkCharge(charge)
	if err != nil {
		return map[string]string{}, err
	}

	ct := getPayType(charge.PayMethod)
	re, err := ct.Pay(charge)
	return re, err
}

// 付款给用户接口
func PayToClient(charge *common.Charge) (map[string]string, error) {
	err := checkCharge(charge)
	if err != nil {
		return nil, err
	}
	ct := getPayType(charge.PayMethod)
	re, err := ct.PayToClient(charge)
	return re, err
}

// 验证支付内容
func checkCharge(charge *common.Charge) error {
	if charge.PayMethod <= 0 {
		return errors.New("PayMethod不能少于等于0")
	}
	if charge.MoneyFee <= 0 {
		return errors.New("MoneyFee不能少于等于0")
	}
	return nil
}

func Refund(charge *common.RefundCharge) (map[string]string, error) {
	if charge.RefundFee <= 0 || charge.TotalFee <= 0 {
		return nil, errors.New("refund fee and total fee must > 0")
	}

	if charge.PayMethod != constant.WECHAT_MINI_PROGRAM {
		return nil, errors.New("only support wechat mini program now")
	}

	//only support wechat mini program now
	return client.DefaultWechatMiniProgramClient().Refund(charge)
}

// getPayType 得到需要支付的类型
func getPayType(payMethod int64) common.PayClient {
	//如果使用余额支付
	switch payMethod {
	case constant.ALI_WEB:
		return client.DefaultAliWebClient()
	case constant.ALI_APP:
		return client.DefaultAliAppClient()
	case constant.WECHAT_WEB:
		return client.DefaultWechatWebClient()
	case constant.WECHAT_APP:
		return client.DefaultWechatAppClient()
	case constant.WECHAT_MINI_PROGRAM:
		return client.DefaultWechatMiniProgramClient()
	}
	return nil
}
