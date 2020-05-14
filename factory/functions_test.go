/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2020/5/14 14:00
 */

package main

import "testing"

func TestAli(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "支付宝测试",
			args: args{str: ""},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ali(tt.args.str); got != tt.want {
				t.Errorf("Ali() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeChat(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "微信测试",
			args: args{str: ""},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WeChat(tt.args.str); got != tt.want {
				t.Errorf("WeChat() = %v, want %v", got, tt.want)
			}
		})
	}
}
