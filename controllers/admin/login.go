package admin

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"net/http"
	"shopping_mall/models"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})

}
func (con LoginController) DoLogin(c *gin.Context) {

	//获取表单传过来的数据
	captchaId := c.PostForm("captchaId")
	username := c.PostForm("username")
	password := c.PostForm("password")
	verifyValue := c.PostForm("verifyValue")
	fmt.Println(username, password)
	//1、验证验证码是否正确
	if flag := models.VerifyCaptcha(captchaId, verifyValue); flag {
		//2、查询数据库 判断用户以及密码是否存在
		userinfoList := []models.Manager{}
		//var userinfoList []models.Manager
		password = models.Md5(password)

		models.DB.Where("username=? AND password=?", username, password).Find(&userinfoList)

		if len(userinfoList) > 0 {
			//3、执行登录 保存用户信息 执行跳转
			session := sessions.Default(c)
			//注意：session.Set没法直接保存结构体对应的切片 把结构体转换成json字符串
			userinfoSlice, _ := json.Marshal(userinfoList)
			session.Set("userinfo", string(userinfoSlice))
			session.Save()
			con.success(c, "登录成功", "/admin")

		} else {
			con.error(c, "用户名或者密码错误", "/admin/login")
		}

	} else {
		con.error(c, "验证码验证失败", "/admin/login")
	}

}
func (con LoginController) Captcha(c *gin.Context) {
	id, b64s, err := models.MakeCaptcha()

	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId":    id,
		"captchaImage": b64s,
	})
}
