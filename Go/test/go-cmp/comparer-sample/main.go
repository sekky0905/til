package main

import "time"

type User struct {
	ID uint32
	Name string
	Age uint8
	CreatedAt time.Time
}

func GetUsers()[]User {
	// ここで仮にDBの操作が存在するものとする
	// そして、CreatedAt は DB 側で NOW() で自動で追加されているものとする


	return []User {
		{
			ID:        1,
			Name:      "hoge",
			Age:       20,
			CreatedAt:time.Now(), // 擬似的にここで入力
		},
		{
			ID:        1,
			Name:      "foo",
			Age:       20,
			CreatedAt:time.Now(),
		},
		{
			ID:        1,
			Name:      "bar",
			Age:       20,
			CreatedAt:time.Now(),
		},
	}
}
