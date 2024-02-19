package logic

import (
	"jcz-backend/internal/engine"
	"jcz-backend/model"
	"jcz-backend/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func getUserByPhone(Phone int64) uint {
	var users model.User
	engine.GetSqlCli().Model(model.User{}).
		Select("id").
		Where("phone = ?", Phone).First(&users)
	return users.ID
}

func UserRegister(ctx *gin.Context) {
	type RegisterParam struct {
		Phone    string `json:"phone" binding:"required"`
		Password string `json:"password"`
	}
	param := RegisterParam{}
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
	}
	// 验证手机号合法性
	if len(param.Phone) != 11 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "手机号不合法",
		})
		return
	}
	NumPhone, err := strconv.Atoi(param.Phone)
	if err != nil {
		logrus.Errorf("[logic] Create Bind Json %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "手机号不合法",
		})
		return
	}
	if getUserByPhone(int64(NumPhone)) != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "手机号已存在!",
		})
		return
	}

	u := model.User{
		Phone:    int64(NumPhone),
		Password: utils.Hash256(param.Password),
	}
	if err = engine.GetSqlCli().Model(model.User{}).Create(&u).Error; err != nil {
		logrus.Errorf("[logic] Create %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "用户创建失败",
		})
		return
	}

	token, err := utils.GenerateToken(u.Phone, u.Nickname, u.ID)
	if err != nil {
		logrus.Errorf("[logic] Create %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
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

func UserLogin(ctx *gin.Context) {
	type LoginParam struct {
		Account  string `json:"account,required" binding:"required"`
		Password string `json:"password,required" binding:"required"`
	}
	param := LoginParam{}
	if err := ctx.ShouldBindJSON(&param); err != nil || param.Account == "" || param.Password == "" {
		logrus.Errorf("[logic.UserLogin] %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "登录失败, 参数错误",
		})
		return
	}

	// 判断是手机号还是昵称登录
	var user model.User

	if param.Account[0] >= '0' && param.Account[0] <= '9' && len(param.Account) == 11 {
		//手机号登录
		phone, err := strconv.Atoi(param.Account)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "手机号不合法",
			})
			return
		}

		if err = engine.GetSqlCli().Model(model.User{}).Where("phone = ?", phone).First(&user).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "用户不存在或密码错误",
			})
			return
		}
	} else {
		//昵称登录
		if err := engine.GetSqlCli().Model(model.User{}).Where("nickname = ?", param.Account).First(&user).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "用户不存在或密码错误",
			})
			return
		}
	}

	// check password
	if utils.Hash256(param.Password) != user.Password {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "密码错误",
		})
		return
	}

	token, err := utils.GenerateToken(user.Phone, user.Nickname, user.ID)
	if err != nil {
		logrus.Errorf("[logic.UserLogin] %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "登录成功",
		"data":  user,
		"token": token,
	})
}

func UserInfoUpdate(ctx *gin.Context) {
	type UpdateInfoParam struct {
		Nickname string `json:"nickname"`
		//AvatarID      string `json:"avatar_id"`
		Name          string `json:"name"`
		Email         string `json:"email"`
		Birthday      string `json:"birthday"`
		Gender        string `json:"gender"`
		HearingLevel  int    `json:"hearing_level"`
		HearingDevice bool   `json:"hearing_device"`
		Phone         string `json:"phone"`
	}
	info := &UpdateInfoParam{}

	if err := ctx.ShouldBindJSON(info); err != nil {
		logrus.Errorf("[api.UpdateInfo] %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}

	// check birthday
	_, err := time.Parse("2006-01-02", info.Birthday)
	if err != nil {
		logrus.Errorf("[api.UpdateInfo] %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "生日格式错误",
		})
		return
	}

	// check phone
	if len(info.Phone) != 11 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "手机号不合法",
		})
		return
	}

	NumPhone, err := strconv.ParseInt(info.Phone, 10, 64)
	if err != nil {
		logrus.Errorf("[api.UpdateInfo] %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "手机号不合法",
		})
		return
	}

	userID := ctx.GetUint("user_id")
	if err = engine.GetSqlCli().Model(model.User{}).Where("id = ?", userID).
		Updates(model.User{
			Email:         info.Email,
			Nickname:      info.Nickname,
			Name:          info.Name,
			HearingLevel:  info.HearingLevel,
			HearingDevice: info.HearingDevice,
			Gender:        info.Gender,
			BirthDay:      info.Birthday,
			Phone:         NumPhone,
		}).Error; err != nil {
		logrus.Errorf("[api.UpdateInfo] %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "更新失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}

// UserVerify 用户验证
// 通过手机号, 姓名, 邮箱验证用户得到token
// 规避了用户忘记密码的情况
func UserVerify(ctx *gin.Context) {
	type UserVerifyParam struct {
		Phone string `json:"phone" form:"phone" binding:"required"`
		Name  string `json:"name" form:"name"`
		Email string `json:"email" form:"email"`
	}
	param := UserVerifyParam{}
	if ctx.ShouldBind(&param) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}

	if err := engine.GetSqlCli().Model(model.User{}).Where("phone = ? AND name = ? AND email = ?",
		param.Phone, param.Name, param.Email).First(&model.User{}).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "验证失败, 信息错误, 请检查",
		})
		return
	}

	phoneUint, err := strconv.ParseUint(param.Phone, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "手机号不合法",
		})
		return
	}

	user := model.User{}
	if err := engine.GetSqlCli().Model(model.User{}).Where("phone = ?", phoneUint).First(&user).Error; err != nil {
		logrus.Errorf("[logic.UserVerify] %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据库查询错误",
		})
		return
	}

	token, err := utils.GenerateToken(user.Phone, user.Nickname, user.ID)

	if err != nil {
		logrus.Errorf("[logic.UserVerify] %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "验证成功",
		"data":  user,
		"token": token,
	})
}

func UserUpdatePhone(ctx *gin.Context) {
	type UpdatePhoneParam struct {
		Phone string `json:"phone,omitempty"`
	}
	param := UpdatePhoneParam{}
	if err := ctx.ShouldBind(&param); err != nil {
		logrus.Errorf("[api.UpdatePhone] %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}

	NumPhone, err := strconv.Atoi(param.Phone)
	if err != nil {
		logrus.Errorf("[api.UserVerify] %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "手机号不合法",
		})
		return
	}
	userID := ctx.GetUint("user_id")
	if engine.GetSqlCli().Model(model.User{}).Where("id = ?", userID).
		Update("phone", int64(NumPhone)).Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "修改失败",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "修改成功",
			"data": gin.H{
				"new_phone": param.Phone,
			},
		})
	}
}

func UserUpdatePasswordByToken(ctx *gin.Context) {
	type UpdatePasswordParam struct {
		NewPassword string `json:"new_password" form:"new_password" binding:"required"`
	}
	param := UpdatePasswordParam{}

	if err := ctx.ShouldBind(&param); err != nil {
		logrus.Errorf("[logic.UpdatePassword] %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}

	userID := ctx.GetUint("user_id")
	user := model.User{}
	if err := engine.GetSqlCli().Model(model.User{}).Where("id = ?", userID).First(&user).Error; err != nil {
		logrus.Errorf("[logic.UpdatePassword] %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "内部错误",
		})
		return
	}

	if err := engine.GetSqlCli().Model(model.User{}).Where("id = ?", userID).
		Update("password", utils.Hash256(param.NewPassword)).Error; err != nil {
		logrus.Errorf("[logic.UpdatePassword] %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据库更新错误",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "修改成功",
	})
}

func UserProfile(ctx *gin.Context) {
	userID := ctx.GetUint("user_id")

	user := model.User{}
	if err := engine.GetSqlCli().Model(model.User{}).Where("id = ?", userID).First(&user).Error; err != nil {
		logrus.Errorf("[logic.UserProfile] %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "数据库查询错误",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": user,
	})
}
