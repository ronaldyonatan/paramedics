package smtp

import "testing"

func TestSendMail(t *testing.T) {
	type args struct {
		req Mail
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Send Email",
			args: args{
				req: Mail{
					From:    "Unit Testing",
					To:      []string{"fernando.riyo@jec.co.id"},
					Subject: "Unit Test Praktek Email",
					Body:    "<h1>Testing email</h1>",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendMail(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("SendMail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
