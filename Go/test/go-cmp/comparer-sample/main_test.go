package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGetUsers(t *testing.T) {
	tests := []struct {
		name string
		want []User
	}{
		{
			name: "3usersを取得すること",
			want: []User{
				{
					ID:   1,
					Name: "hoge",
					Age:  20,
				},
				{
					ID:   1,
					Name: "foo",
					Age:  20,
				},
				{
					ID:   1,
					Name: "bar",
					Age:  20,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := GetUsers(); !reflect.DeepEqual(got, tt.want) {
			//	 t.Errorf("GetUsers() = %v, want %v", got, tt.want)
			// }

			// 上記のような場合、以下のような感じでエラーになる
			// --- FAIL: TestGetUsers (0.00s)
			// --- FAIL: TestGetUsers/3usersを取得すること (0.00s)
			// main_test.go:37: GetUsers() = [{1 hoge 20 2019-08-28 16:40:19.772311 +0900 JST m=+0.000646976} {1 foo 20 2019-08-28 16:40:19.772312 +0900 JST m=+0.000647073} {1 bar 20 2019-08-28 16:40:19.772312 +0900 JST m=+0.000647166}], want [{1 hoge 20 0001-01-01 00:00:00 +0000 UTC} {1 foo 20 0001-01-01 00:00:00 +0000 UTC} {1 bar 20 0001-01-01 00:00:00 +0000 UTC}]


			// 以下のようにしてもテストはちゃんと動く
			//got := GetUsers()
			//opt := cmpopts.IgnoreFields(User{}, "CreatedAt")
			//for i, v := range got {
			//	if diff := cmp.Diff(v, tt.want[i], opt); diff != "" {
			//		t.Errorf("diff (-got +expected)\n%s", diff)
			//	}
			//}


			got := GetUsers()
			opt := cmp.Comparer(func(x, y User) bool {
				return x.ID == y.ID && x.Name == y.Name &&x.Age == y.Age
			})

			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("diff (-got +expected)\n%s", diff)
			}


		})
	}
}
