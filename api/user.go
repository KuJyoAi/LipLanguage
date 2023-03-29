package api

import (
	"LipLanguage/dao"
	"LipLanguage/dao/user"
	"LipLanguage/model"
	"LipLanguage/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

func Register(ctx *gin.Context) {
	param := model.RegisterParam{}
	err := ctx.ShouldBindJSON(&param)
	logrus.Infof("Create %v", param)
	// 验证手机号合法性
	if len(param.Phone) != 11 || param.Phone[0] != '1' {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "手机号不合法",
		})
		return
	}
	NumPhone, err := strconv.Atoi(param.Phone)
	if err != nil {
		logrus.Errorf("[api] Create Bind Json %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}
	if user.Exists(int64(NumPhone)) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "手机号已存在!",
		})
		return
	}
	//todo 生成token
	token, err := service.Register(int64(NumPhone), param.Password)

	if err != nil {
		logrus.Errorf("[api.Create] %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "服务端错误",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
		"data": gin.H{
			"token": token,
		},
	})
}

func Login(ctx *gin.Context) {
	param := model.LoginParam{}
	err := ctx.ShouldBindJSON(&param)
	logrus.Infof("[api.Login]: %v", param)
	if err != nil || param.Account == "" || param.Password == "" {
		logrus.Errorf("[api.Login] %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "登录失败, 参数错误",
		})
		return
	}

	// 判断是手机号还是昵称登录
	if param.Account[0] >= '0' && param.Account[0] <= '9' && len(param.Account) == 11 {
		//手机号登录
		phone, err := strconv.Atoi(param.Account)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "用户不存在或密码错误",
			})
			return
		}
		token, err := service.LoginByPhone(int64(phone), param.Password)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "用户不存在或密码错误",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "登录成功",
			"data": gin.H{
				"token": token,
				"time":  time.Now(),
			},
		})
	} else {
		//昵称登录
		token, err := service.LoginByNickname(param.Account, param.Password)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "用户不存在或密码错误",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "登录成功",
			"data": gin.H{
				"token": token,
				"time":  time.Now(),
			},
		})
	}

}

func ResetPassword(ctx *gin.Context) {
	param := model.ResetPasswordParam{}
	err := ctx.ShouldBindJSON(&param)
	logrus.Infof("[api.ResetPassword] %v", param)
	NumPhone, err := strconv.Atoi(param.Phone)
	if err != nil {
		logrus.Errorf("[api.UserVerify] %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}

	if service.UserResetPassword(int64(NumPhone), param.Password) {
		Response(ctx, http.StatusOK, "重置成功", nil)
		//重置成功, 删除掉redis里的token, 防止重放攻击
		_ = dao.DeleteRedisToken(int64(NumPhone))
		return
	} else {
		Response(ctx, http.StatusBadRequest, "重置失败", nil)
	}
}

func UserInfoUpdate(ctx *gin.Context) {
	info := &model.UpdateInfoParam{}
	err := ctx.ShouldBindJSON(info)
	logrus.Infof("[api.UpdateInfo] %+v", info)
	if err != nil {
		logrus.Errorf("[api.UpdateInfo] %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}

	token, _ := ctx.Cookie("auth")
	err = service.UserInfoUpdate(token, info)

	if err != nil {
		logrus.Errorf("[api.UpdateInfo] %v", err)
		Response(ctx, http.StatusInternalServerError, "服务端错误", nil)
		return
	} else {
		Response(ctx, http.StatusOK, "修改成功", nil)
	}
}

func UserVerify(ctx *gin.Context) {
	param := model.UserVerifyParam{}
	err := ctx.ShouldBindJSON(&param)
	logrus.Infof("[api.UserVerify] %v", param)
	if err != nil {
		logrus.Errorf("[api.UserVerify] %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}
	NumPhone, err := strconv.Atoi(param.Phone)
	token, ok := service.UserVerify(int64(NumPhone), param.Name, param.Email)
	if ok {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "验证成功",
			"data": gin.H{
				"token": token,
			},
		})
		return
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证失败",
		})
		return
	}
}

func UserUpdatePhone(ctx *gin.Context) {
	param := model.UpdatePhoneParam{}
	err := ctx.ShouldBindJSON(&param)
	logrus.Infof("[api.UpdatePhone] %v", param)
	if err != nil {
		logrus.Errorf("[api.UpdatePhone] %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}
	NumPhone, err := strconv.Atoi(param.Phone)
	token, _ := ctx.Cookie("auth")
	if service.UserUpdatePhone(token, int64(NumPhone)) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "改绑成功",
		})
		return
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "改绑失败",
		})
		return
	}
}

func UserUpdatePassword(ctx *gin.Context) {
	param := model.UpdatePasswordParam{}
	err := ctx.ShouldBindJSON(&param)
	logrus.Infof("[api.UpdatePassword] %v", param)
	if err != nil {
		logrus.Errorf("[api.UpdatePassword] %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}

	token, _ := ctx.Cookie("auth")
	if service.UserUpdatePassword(token, param.OldPassword, param.NewPassword) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "修改成功",
		})
		return
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "修改失败",
		})
		return
	}
}

func UserProfile(ctx *gin.Context) {
	token, _ := ctx.Cookie("auth")
	claim, _ := dao.ParseToken(token)

	logrus.Infof("[api.UserProfile] %v", claim)
	User, err := service.UserGetProfile(claim.Phone)
	if err != nil {
		logrus.Errorf("[api.UserProfile] UserGetProfile%v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "获取失败",
		})
		return
	}
	registerTime := User.CreateAt.Format("2006-01-02 15:04:05")
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "获取成功",
		"data": gin.H{
			"nickname":       User.Nickname,
			"name":           User.Name,
			"email":          User.Email,
			"birthday":       User.BirthDay,
			"gender":         User.Gender,
			"phone":          User.Phone,
			"Avatar":         User.AvatarUrl,
			"hearing_device": User.HearingDevice,
			"register_time":  registerTime,
		},
	})
}
