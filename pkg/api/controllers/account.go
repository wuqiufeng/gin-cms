package controllers

import (
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	BaseController
}

func (a AccountController) Info(c *gin.Context) {
	//userId := int(c.Value("userId").(float64))
	//data := userService.InfoOfId(dto.GeneralGetDto{Id: userId})
	//resp(c, map[string]interface{}{
	//	"result": data,
	//})
}

func (a AccountController) GetPermissionsWithMenu(c *gin.Context) {
	//userId := strconv.Itoa(int(c.Value("userId").(float64)))
	//resp(c, map[string]interface{}{
	//	"result": userService.GetAllPermissions(userId),
	//	"info": map[string]interface{}{
	//		"id":       c.Value("userId"),
	//		"username": c.GetString("userName"),
	//	},
	//})
}

func (a *AccountController) EditPassword(c *gin.Context) {
	//// simulate value in query
	//c.Params = []gin.Param{
	//	{
	//		Key:   "id",
	//		Value: fmt.Sprintf("%d", int(c.Value("userId").(float64))),
	//	},
	//}
	//var accountDto dto.AccountEditPasswordDto
	//var userDto dto.UserEditPasswordDto
	//accountDto.Id = int(c.Value("userId").(float64))
	//
	//if a.BindAndValidate(c, &accountDto) {
	//	if accountDto.NewPassword != accountDto.RePassword {
	//		fail(c, ErrDifferentPasswords)
	//		return
	//	}
	//	//check if equal to old password
	//	userModel := userService.InfoOfId(dto.GeneralGetDto{Id: accountDto.Id})
	//	if login.VerifyPassword(accountDto.NewPassword, userModel) {
	//		fail(c, ErrSamePasswords)
	//		return
	//	}
	//	userDto.Id = accountDto.Id
	//	userDto.Password = accountDto.RePassword
	//	affected := userService.UpdatePassword(userDto)
	//	if affected > 0 {
	//		// 修改密码只有成功时才进入日志
	//		// 因为需要这条日志判断用户上次更新密码的时间
	//		// 故特殊处理
	//		b, _ := json.Marshal(accountDto)
	//		orLogDto := dto.OperationLogDto{
	//			UserId:           userDto.Id,
	//			RequestUrl:       c.Request.URL.Path,
	//			OperationMethod:  c.Request.Method,
	//			Params:           string(b),
	//			Ip:               c.ClientIP(),
	//			IpLocation:       "", //TODO...待接入获取ip位置服务
	//			OperationResult:  "success",
	//			OperationSuccess: 1,
	//			OperationContent: "Bind third account",
	//		}
	//		_ = logService.InsertOperationLog(orLogDto)
	//	}
	//	ok(c, "ok.UpdateDone")
	//}
}

func (AccountController) GetDomains(c *gin.Context) {
	//userId := int(c.Value("userId").(float64))
	//domains := userService.GetRelatedDomains(strconv.Itoa(userId))
	//resp(c, map[string]interface{}{
	//	"result": domains,
	//})
}

func (a *AccountController) AccountInfo(c *gin.Context) {
	//userId := int(c.Value("userId").(float64))
	//myAccountService := service.MyAccountService{}
	//userSecretQuery, err := myAccountService.GetSecret(userId)
	//if err != nil {
	//	fail(c, ErrInvalidUser)
	//	return
	//}
	//account := userSecretQuery.Account_name
	//issuer := "宙斯"
	//URL, err := url.Parse("otpauth://totp")
	//if err != nil {
	//	fail(c, ErrInvalidParams)
	//	return
	//}
	//
	//URL.Path += "/" + url.PathEscape(issuer) + ":" + url.PathEscape(account)
	//params := url.Values{}
	//params.Add("secret", userSecretQuery.Secret)
	//params.Add("issuer", issuer)
	//URL.RawQuery = params.Encode()
	//p, errs := qrcode.New(URL.String(), qrcode.Medium)
	//img := p.Image(256)
	//if errs != nil {
	//	fail(c, ErrInvalidParams)
	//	return
	//}
	//out := new(bytes.Buffer)
	//errx := png.Encode(out, img)
	//if errx != nil {
	//	fail(c, ErrInvalidParams)
	//}
	//resp(c, map[string]interface{}{
	//	"result": dto.UserSecretRetDto{
	//		Is_open: userSecretQuery.Is_open,
	//		Account: account,
	//		Code:    "data:image/png;base64," + base64.StdEncoding.EncodeToString(out.Bytes()),
	//		Secret:  userSecretQuery.Secret,
	//	},
	//})
}

func (a *AccountController) BindGoogle2faCode(c *gin.Context) {
	//userId := int(c.Value("userId").(float64))
	//myAccountService := service.MyAccountService{}
	//userSecretQuery, err := myAccountService.GetSecret(userId)
	//if err != nil {
	//	fail(c, ErrInvalidUser)
	//	return
	//}
	//secretBase32 := userSecretQuery.Secret
	//bindCodeDto := &dto.BindCodeDto{}
	//if !a.BindAndValidate(c, &bindCodeDto) {
	//	fail(c, ErrInvalidParams)
	//	return
	//}
	//otpc := &dgoogauth.OTPConfig{
	//	Secret:      secretBase32,
	//	WindowSize:  3,
	//	HotpCounter: 0,
	//	// UTC:         true,
	//}
	//val, err := otpc.Authenticate(bindCodeDto.Google2faToken)
	//if err != nil {
	//	fail(c, ErrGoogleBindCode)
	//	return
	//}
	//if !val {
	//	fail(c, ErrGoogleBindCode)
	//	return
	//}
	//
	//myAccountService.Update2FaStatus(userId, 1) //更新状态
	//resp(c, map[string]interface{}{
	//	"result": "bind success!",
	//})
}

func (a *AccountController) Close2fa(c *gin.Context) {
	//userId := int(c.Value("userId").(float64))
	//myAccountService := service.MyAccountService{}
	//userSecretQuery, err := myAccountService.GetSecret(userId)
	//if err != nil {
	//	fail(c, ErrInvalidUser)
	//	return
	//}
	//secretBase32 := userSecretQuery.Secret
	//bindCodeDto := &dto.BindCodeDto{}
	//if !a.BindAndValidate(c, &bindCodeDto) {
	//	fail(c, ErrInvalidParams)
	//	return
	//}
	//otpc := &dgoogauth.OTPConfig{
	//	Secret:      secretBase32,
	//	WindowSize:  3,
	//	HotpCounter: 0,
	//	// UTC:         true,
	//}
	//val, err := otpc.Authenticate(bindCodeDto.Google2faToken)
	//if err != nil {
	//	fail(c, ErrGoogleBindCode)
	//	return
	//}
	//if !val {
	//	fail(c, ErrGoogleBindCode)
	//	return
	//}
	//
	//myAccountService.Update2FaStatus(userId, 0) //更新状态
	//
	//resp(c, map[string]interface{}{
	//	"result": "update success!",
	//})
}

func (a *AccountController) CheckGoogle2faCode(c *gin.Context) {
	//userId := int(c.Value("userId").(float64))
	//myAccountService := service.MyAccountService{}
	//userSecretQuery, err := myAccountService.GetSecret(userId)
	//if err != nil {
	//	fail(c, ErrInvalidUser)
	//	return
	//}
	//secretBase32 := userSecretQuery.Secret
	//bindCodeDto := &dto.BindCodeDto{}
	//if !a.BindAndValidate(c, &bindCodeDto) {
	//	fail(c, ErrInvalidParams)
	//	return
	//}
	//otpc := &dgoogauth.OTPConfig{
	//	Secret:      secretBase32,
	//	WindowSize:  3,
	//	HotpCounter: 0,
	//}
	//val, err := otpc.Authenticate(bindCodeDto.Google2faToken)
	//if err != nil {
	//	fail(c, ErrGoogleBindCode)
	//	return
	//}
	//if !val {
	//	fail(c, ErrGoogleBindCode)
	//	return
	//}
	//resp(c, map[string]interface{}{
	//	"result": "Authenticated!",
	//})
}

func (a *AccountController) FindCodeOpen(c *gin.Context) {
	//userId := int(c.Value("userId").(float64))
	//myAccountService := service.MyAccountService{}
	//userSecretQuery, err := myAccountService.GetSecret(userId)
	//if err != nil {
	//	fail(c, ErrInvalidUser)
	//	return
	//}
	//resp(c, map[string]interface{}{
	//	"is_open": userSecretQuery.Is_open,
	//})
}

func (a *AccountController) ThirdList(c *gin.Context) {
	//var listDto dto.GeneralListDto
	////userId := int(c.Value("userId").(float64))
	//if a.BindAndValidate(c, &listDto) {
	//	myAccountService := service.MyAccountService{}
	//	data, total := myAccountService.GetThirdList(listDto)
	//	resp(c, map[string]interface{}{
	//		"result": data,
	//		"total":  total,
	//	})
	//}
}

func (a *AccountController) SendVerifymail(c *gin.Context) {
	//var verifyDto dto.VerifyEmailDto
	//if a.BindAndValidate(c, &verifyDto) {
	//	username := viper.GetString("email.smtp.user")
	//	password := viper.GetString("email.smtp.password")
	//	host := viper.GetString("email.smtp.server")
	//	port := viper.GetInt("email.smtp.port")
	//	from := viper.GetString("email.smtp.address")
	//	if port == 0 {
	//		port = 25
	//	}
	//	config := fmt.Sprintf(`{"username":"%s","password":"%s","host":"%s","port":%d,"from":"%s"}`, username, password, host, port, from)
	//	temail := utils.NewEMail(config)
	//	temail.To = []string{verifyDto.Email} //指定收件人邮箱地址
	//	temail.From = from                    //指定发件人的邮箱地址
	//	temail.Subject = "验证账号邮件"             //指定邮件的标题
	//	temail.HTML = mailTemplate.MailBody()
	//	err := temail.Send()
	//	if err != nil {
	//		fail(c, ErrSendMail)
	//		return
	//	}
	//	resp(c, map[string]interface{}{
	//		"result": "email send success！",
	//	})
	//}

}

func (a *AccountController) EmailVerify(c *gin.Context) {
	//emailVerificationDto := &dto.EmailVerificationDto{}
	//if a.BindAndValidate(c, &emailVerificationDto) {
	//	resp(c, map[string]interface{}{
	//		"result": "email verify success！",
	//	})
	//}
}

func (a *AccountController) ThirdBind(c *gin.Context) {
	//bindThirdDto := &dto.BindThirdDto{}
	//if a.BindAndValidate(c, &bindThirdDto) {
	//	from := bindThirdDto.From
	//	if from == 0 {
	//		from = 1
	//	}
	//	userId := int(c.Value("userId").(float64))
	//	myAccountService := service.MyAccountService{} //switch case from  1 钉钉 2 微信 TODO
	//	_, err := myAccountService.BindDingtalk(bindThirdDto.Code, userId, from)
	//	if err != nil {
	//		log.Error(err.Error())
	//		fail(c, ErrBindDingtalk)
	//		return
	//	}
	//	//data := map[string]string{
	//	//	"openid": openid,
	//	//}
	//	//resp(c, map[string]interface{}{
	//	//	"result": data,
	//	//})
	//	c.Redirect(301, "/#/my/third")
	//}
}

func (a *AccountController) ThirdUnbind(c *gin.Context) {
	//UnBindDingtalkDto := &dto.UnBindThirdDto{}
	//if a.BindAndValidate(c, &UnBindDingtalkDto) {
	//	userId := int(c.Value("userId").(float64))
	//	oauthType := UnBindDingtalkDto.OAuthType
	//	if oauthType == 0 {
	//		oauthType = 1
	//	}
	//	userService := service.UserService{} //switch case from  1 钉钉 2 微信 TODO
	//	errs := userService.UnBindUserDingtalk(oauthType, userId)
	//	if errs != nil {
	//		fail(c, ErrUnBindDingtalk)
	//		return
	//	}
	//	data := map[string]bool{
	//		"state": true,
	//	}
	//	resp(c, map[string]interface{}{
	//		"result": data,
	//	})
	//}

}

// check if need to send sms code
func (a *AccountController) SmsSendCheck(c *gin.Context) {
	//twoFaDto := dto.TwoFaDto{}
	//if a.BindAndValidate(c, &twoFaDto) {
	//	resp(c, map[string]interface{}{
	//		"show": userService.VerifySmsCodeIfNeedToShow(twoFaDto),
	//	})
	//}
}

// SmsSendCheck send sms code
func (a *AccountController) SmsSendCode(c *gin.Context) {
	//twoFaDto := dto.TwoFaDto{}
	//if a.BindAndValidate(c, &twoFaDto) {
	//	if mobile, err := userService.Verify2FaHandler(twoFaDto); err != nil {
	//		ErrSmsSendCode.Moreinfo = err.Error()
	//	} else {
	//		secureMobileShow := mobile[:3] + "****" + mobile[8:]
	//		rawOk(c, "已向"+secureMobileShow+"下发短信验证码，请注意查收，短信十分钟内有效！")
	//		return
	//	}
	//}
	//fail(c, ErrSmsSendCode)
	//ErrSmsSendCode.Moreinfo = ""
}

func (a *AccountController) LdapAddUser(c *gin.Context) {
	//ldapConn := ldap.GetLdap()
	//rr := ldapConn.Add("zeus", "zeus@bullteam.cn", "10111", "10111", "123456")
	//resp(c, map[string]interface{}{
	//	"result": rr,
	//})
}

//upload avatar
func (a *AccountController) UploadAvatar(c *gin.Context) {
	//file, header, err := c.Request.FormFile("file")
	//if err != nil {
	//	//c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
	//	fail(c, ErrUploadAvatar)
	//	return
	//}
	//filename := header.Filename
	//Account := service.MyAccountService{}
	//result, err := Account.UploadAvatar(file, filename, int(c.Value("userId").(float64)))
	//if err != nil {
	//	fail(c, ErrUploadAvatar)
	//	return
	//}
	//resp(c, map[string]interface{}{
	//	"result": result,
	//})
}

// CheckIdle check if user is idle
func (a *AccountController) CheckIdle(c *gin.Context) {
	//userId := int(c.Value("userId").(float64))
	//resp(c, map[string]interface{}{
	//	"idle": logService.CheckAccountIdleTooLong(dto.GeneralGetDto{Id: userId}),
	//})
}

// CheckIfNeedToChangePwd check if user need to change password
func (a *AccountController) CheckIfNeedToChangePwd(c *gin.Context) {
	//userId := int(c.Value("userId").(float64))
	//needed, days := userService.GetLastPwdChangeDaySinceNow(dto.GeneralGetDto{Id: userId})
	//resp(c, map[string]interface{}{
	//	"needed": needed,
	//	"days":   days,
	//})
}
