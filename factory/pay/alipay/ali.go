/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2020/5/14 11:49
 */

package alipay

import "fmt"

type Pay struct {
}

func (this Pay) H5() {
	fmt.Println("alipay H5 ")
}

func (this Pay) QrCodeFront() {
	fmt.Println("alipay QrCodeFront ")
}

func (this Pay) QrCodePass() {
	fmt.Println("alipay QrCodePass ")
}
func (this Pay) App() {
	fmt.Println("alipay App ")
}
func (this Pay) JsPay() {
	fmt.Println("alipay JsPay")
}
