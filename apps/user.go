package apps

import (
	myjwt "../jwt"
	"../models"
	"bytes"
	"github.com/dchest/captcha"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

//1.获取验证码
func CaptchaId (c *gin.Context) {
	captchaId := captcha.NewLen(4)
	c.SetCookie("captchaId", captchaId, 3600000, "/", "localhost", false, false)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "SUCCESS",
		"CaptchaId": captchaId,
	})
}

//2.获取验证码图片
func CaptchaImage(c *gin.Context) {
	ServeHTTP(c.Writer, c.Request)
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dir, file := path.Split(r.URL.Path)
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	if ext == "" || id == "" {
		http.NotFound(w, r)
		return
	}
	if r.FormValue("reload") != "" {
		captcha.Reload(id)
	}
	lang := strings.ToLower(r.FormValue("lang"))
	download := path.Base(dir) == "download"
	if Serve(w, r, id, ext, lang, download, captcha.StdWidth, captcha.StdHeight) == captcha.ErrNotFound {
		http.NotFound(w, r)
	}
}

func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		width = 120
		height = 40
		captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}
	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}

//静态首页的token
func LoginToken(c *gin.Context)  {
	j := &myjwt.JWT{
		[]byte("newtrekWang"),
	}
	var requestID string
	requestID = uuid.New()
	claims := myjwt.LoginClaims{
		requestID,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),
			ExpiresAt: int64(time.Now().Unix() + 3600),
			Issuer: "newtrekWang",
		},
	}
	token, err := j.CreateLoginToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}
	c.SetCookie("loginToken", token, 3600000, "/", "localhost", false, false)
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登录成功！",
		"data":   token,
	})
	return
}

// 登录
func Login(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	text := c.Request.FormValue("text")
	captchaId := c.Request.FormValue("captchaId")
	if !captcha.VerifyString(captchaId, text) {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "验证码错误",
		})
	}
	login_token := c.Request.Header.Get("AUTHORIZATION")
	cookie_token, err := c.Cookie("loginToken")
	if login_token == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"message": "请求未携带token，无权限访问",
		})
		c.Abort()
		return
	}
	if cookie_token != login_token {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"message": "token被篡改",
		})
		c.Abort()
		return
	}
	j := myjwt.NewJWT()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(login_token)
	if err != nil {
		if err == myjwt.TokenExpired {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"message": "授权已过期",
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"message": err.Error(),
		})
		c.Abort()
		return
	}
	// 继续交由下一个路由处理,并将解析出的信息传递下去
	if claims == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "token已失效",
		})
	}
	mem, err := models.Login(username, password)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
	} else {
		var userID string
		userID = uuid.New()
		session := sessions.Default(c)
		session.Set(userID, username)
		session.Save()
		c.SetCookie("userID", userID, 3600000, "/", "localhost", false, false)
		generateToken(c, mem)
	}
}

// LoginResult 登录结果结构
type LoginResult struct {
	Token string `json:"token"`
	models.User
}

// 生成令牌
func generateToken(c *gin.Context, user models.User) {
	j := &myjwt.JWT{
		[]byte("newtrekWang"),
	}
	claims := myjwt.CustomClaims{
		string(user.Id),
		user.UserName,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer: "newtrekWang",                   //签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"message":    err.Error(),
		})
		return
	}
	log.Println(token)
	data := LoginResult{
		User:  user,
		Token: token,
	}
	c.SetCookie("token", token, 3600000, "/", "localhost", false, false)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message":    "登录成功！",
		"data":   data,
	})
	return
}

// 检查用户
func CheckUser(c *gin.Context) {
	userID := c.Request.FormValue("userID")
	session := sessions.Default(c)
	user := session.Get(userID)
	if user == nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": "用户未登录",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "检验成功",
			"data": user,
		})
	}
}

// 所有管理员展示
func UserList(c *gin.Context) {
	filters := make([]interface{}, 0)
	filters = append(filters, "id", "<>", "0")
	page, _ := strconv.Atoi(c.Request.FormValue("page"))
	pageSize, _ := strconv.Atoi(c.Request.FormValue("limit"))
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	list, n, err := models.UserList(page, pageSize, filters...)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusExpectationFailed,
			"message": err.Error(),
			"data":    "123",
		})
		log.Fatal(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":    http.StatusOK,
			"message":   "SUCCESS",
			"results":      list,
			"count":     n,
		})
	}
}

// 添加管理员
func AddUser(c *gin.Context)  {
	u := new(models.User)
	u.UserName = c.Request.FormValue("username")
	u.Password = c.Request.FormValue("password")
	if id, err := u.AddUser(); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
	} else {
		u.Id = int(id)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data": u,
		})
	}
}

// 删除管理员
func DeleteUser(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Request.FormValue("id"))
	if n, err := models.DeleteUser(uid); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		log.Fatal(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data":    n,
		})
	}
}

// 更改管理员
func UpdateUser(c *gin.Context)  {
	uid, _ := strconv.Atoi(c.Request.FormValue("id"))
	u := new(models.User)
	u.Id = uid
	u.UserName = c.Request.FormValue("username")
	u.Password = c.Request.FormValue("password")
	if n, err := u.UpdateUser(uid); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data":    n,
		})
	}
}

