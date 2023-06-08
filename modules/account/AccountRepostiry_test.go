package account

import (
	// "crm-service/mocks"
	// "crm-service/mocks"

	// mocks "crm-service/modules"
	// "crm-service/modules/account/mocks"

	// "crm-service/modules/mocks"

	"reflect"
	"testing"
)

func Test_accountRepository_FindByUsername(t *testing.T) {
	type args struct {
		username string
	}

	// MockFindByUsername := mocks.NewAccountRepositoryInterface(t)
	// mockRepo := mocks.NewMockAccountRepositoryInterface()

	tests := []struct {
		name    string
		c       accountRepository
		args    args
		want    Actor
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TestByusername",
			c:    accountRepository{},
			args: args{
				username: "roby",
			},
			want: Actor{
				Username: "roby",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.FindByUsername(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("accountRepository.FindByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("accountRepository.FindByUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}
