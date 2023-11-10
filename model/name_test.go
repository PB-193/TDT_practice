package model

import "testing"

func TestNewNameModel(t *testing.T) {
	// NewNameModelを使って新しいNameModelを作成
	nm := NewNameModel("山田", "太郎")

	// テストの開始をログ出力
	t.Log("NewNameModelのテストを開始します。")

	// FamilyNameが正しく設定されているかを確認
	if nm.FamilyName != "山田" {
		t.Errorf("FamilyNameが正しくありません。期待する値: %s, 取得した値: %s", "山田", nm.FamilyName)
	} else {
		t.Logf("FamilyNameが正しく設定されています。値: %s", nm.FamilyName)
	}

	// LastNameが正しく設定されているかを確認
	if nm.LastName != "太郎" {
		t.Errorf("LastNameが正しくありません。期待する値: %s, 取得した値: %s", "太郎", nm.LastName)
	} else {
		t.Logf("LastNameが正しく設定されています。値: %s", nm.LastName)
	}
}
