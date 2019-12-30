// Copyright 2019 Enseada authors
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package auth

type User struct {
	ID             string `json:"_id,omitempty"`
	Rev            string `json:"_rev,omitempty"`
	Username       string `json:"username"`
	Password       string `json:"-"`
	HashedPassword []byte `json:"hashed_password"`
}

func Root(pwd string) *User {
	return &User{
		Username: "root",
		Password: pwd,
	}
}
