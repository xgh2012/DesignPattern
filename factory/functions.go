package main

import "github.com/xgh2012/DesignPattern/factory/pay"

//微信支付
func WeChat(str string) bool {
	factory := pay.ShapeFactory{
		Method: "WeChat",
	}.GetPayWayFactory()
	factory.H5()
	factory.App()
	return true
}

//支付宝支付
func Ali(str string) bool {
	factory := pay.ShapeFactory{
		Method: "WeChat",
	}.GetPayWayFactory()
	factory.H5()
	factory.App()
	return true
}
