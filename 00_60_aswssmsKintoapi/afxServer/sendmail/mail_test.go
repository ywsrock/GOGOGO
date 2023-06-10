package sendmail

import "testing"

func TestMail_Send(t *testing.T) {
	mail := &Mail{
		UserName:   "",
		Password:   "",
		Subject:    "メールテスト",
		Body:       "テスト　テスト",
		CC:         false,
		BCC:        false,
		From:       "",
		To:         "",
		HostName:   "",
		Port:       587,
		Recipients: []string{""},
	}

	tests := []struct {
		name    string
		m       *Mail
		wantErr bool
	}{
		{name: "Test01", m: mail, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.Send(); (err != nil) != tt.wantErr {
				t.Errorf("Mail.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
