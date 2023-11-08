package model

import (
	"reflect"
	"testing"
)

func TestNewHumanModel(t *testing.T) {
	type args struct {
		name string
		age  int
	}
	tests := []struct {
		name    string
		args    args        // NewHumanModel関数に渡す引数(今回はnameとage)
		want    *humanModel // 期待する結果
		wantErr bool        // エラーが発生するかどうか true: エラーが発生する, false: エラーが発生しない
		errMsg  string      // wantErr:trueの場合、期待されるエラーメッセージ
	}{
		{
			name: "正常系",
			args: args{
				name: "テストさん",
				age:  10,
			},
			want: &humanModel{
				Name: "テストさん",
				Age:  10,
			},
			wantErr: false,
		},
		{
			name: "異常系: nameが空文字",
			args: args{
				name: "",
				age:  20,
			},
			want:    nil,
			wantErr: true,
			errMsg:  "nameは空文字にできません。", // 期待するエラーメッセージを追加
		},
		{
			name: "異常系: ageが0未満",
			args: args{
				name: "テストさん",
				age:  -1,
			},
			want:    nil,
			wantErr: true,
			errMsg:  "ageは0以上の値でなければなりません。", // 期待するエラーメッセージを追加
		},
	}
	for _, tt := range tests { // テストケースを繰り返し実行
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewHumanModel(tt.args.name, tt.args.age)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHumanModel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// エラーメッセージを検証するためのコードを追加
			if tt.wantErr && err.Error() != tt.errMsg {
				t.Errorf("NewHumanModel() error message は %v, want %v", err.Error(), tt.errMsg)
			}
			if !reflect.DeepEqual(got, tt.want) { //
				t.Errorf("NewHumanModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
