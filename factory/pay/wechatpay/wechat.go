/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2020/5/14 11:49
 */

package wechatpay

import "fmt"

type Pay struct {
}

func (this Pay) H5() {
	fmt.Println("wechatpay H5Pay.")
}

func (this Pay) QrCodeFront() {
	fmt.Println("wechatpay QrCodeFront")
}

func (this Pay) QrCodePass() {
	fmt.Println("wechatpay QrCodePass")
}
func (this Pay) App() {
	fmt.Println("wechatpay App")
}
func (this Pay) JsPay() {
	fmt.Println("wechatpay JsPay")
}
