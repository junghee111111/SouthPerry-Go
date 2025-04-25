/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package enum

type AccountRespCode int

type checkAccountResp struct {
	Success            AccountRespCode
	BanAccount         AccountRespCode
	WrongPassword      AccountRespCode
	WrongID            AccountRespCode
	SystemError        AccountRespCode
	AlreadyLoggedIn    AccountRespCode
	ServiceUnavailable AccountRespCode
	Older20            AccountRespCode
}

var CheckAccountResp = checkAccountResp{
	Success:            1,
	BanAccount:         3,
	WrongPassword:      4,
	WrongID:            5,
	SystemError:        6,
	AlreadyLoggedIn:    7,
	ServiceUnavailable: 10,
	Older20:            11,
}
