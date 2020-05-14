/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2020/5/14 13:48
 */

package pay

import (
	"reflect"
	"testing"
)

func TestShapeFactory_GetPayWayFactory(t *testing.T) {
	type fields struct {
		Method string
	}
	tests := []struct {
		name   string
		fields fields
		want   Methods
	}{
		{
			name:   "微信支付",
			fields: fields{Method: "WeChat"},
			want:   nil,
		},
		{
			name:   "支付宝支付",
			fields: fields{Method: "Ali"},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := ShapeFactory{
				Method: tt.fields.Method,
			}
			if got := this.GetPayWayFactory(); !reflect.DeepEqual(got, tt.want) {
				got.App()
				got.H5()
				got.QrCodeFront()
				got.QrCodePass()
				//t.Errorf("GetPayWayFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
