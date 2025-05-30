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

// CheckAccountResp
/*
 * -1,6,8,9 : 시스템 오류로 접속할 수 없습니다.
 * 2,3 : 지워지거나 접속 중지된 아이디 입니다.
 * 5 : 등록되지 않은 아이디 입니다.
 * 4 : 비밀번호가 틀립니다.
 * 7 : 현재 접속중인 아이디 입니다.
 * 10 : 현재 서버에 접속 요청이 많이 처리하지 못했습니다.
 * 11 : 20세 이상만 접속할 수 있습니다.
 * 13 : 현재 IP로는 마스터로그인이 불가능하니 다시 확인해보시기 바랍니다.
 * 14 : 이 계정은 '임시가입'기간이 만료되어 게임이용이 중지되었습니다.
 * 15 : 이 계정의 넥슨 아이디는 존재하지 않습니다.
 * 17 : 웹 인증정보가 일치하지 않습니다.
 */
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
