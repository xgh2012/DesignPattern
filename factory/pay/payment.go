/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2020/5/14 11:42
 */

package pay

import (
	"github.com/xgh2012/DesignPattern/factory/pay/alipay"
	"github.com/xgh2012/DesignPattern/factory/pay/wechatpay"
)

type Methods interface {
	H5()          //h5支付
	QrCodeFront() //正扫
	QrCodePass()  //被扫
	App()         //APP支付
	JsPay()       //公众号支付
}

type ShapeFactory struct {
	Method string
}

func (this ShapeFactory) GetPayWayFactory() Methods {
	switch this.Method {
	case "WeChat":
		return wechatpay.Pay{}
	case "Ali":
		return alipay.Pay{}
	}
	return nil
}
