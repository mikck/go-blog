package controller

import (
	"blog/app/middleware"
	"blog/app/model"
	"blog/mysql"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var noUser = "Email or password is error"

type UserRes struct {
	User  model.User
	Token string
}

func Login(c *gin.Context) {
	db := mysql.DbObj()
	defer db.Close()
	email := c.PostForm("email")
	password := c.PostForm("password")
	user := model.User{}
	db.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		c.JSON(300, gin.H{
			"error":  1,
			"message": noUser,
			"data":    nil,
		})
		return
	}
	h := md5.New()
	h.Write([]byte(password + email))
	md5Password := hex.EncodeToString(h.Sum(nil))
	if md5Password != user.Password {
		c.JSON(300, gin.H{
			"error":  1,
			"message": noUser,
			"data":    nil,
		})
		return
	}

	j := &middleware.JWT{
		SigningKey: []byte("ccan.blog"),
	}
	claims := middleware.CustomClaims{
		ID:    int(user.ID),
		Name:  user.UserName,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), //time.Now().Add(24 * time.Hour).Unix()
			Issuer:    "user",
		},
	}
	fmt.Println(claims)
	token, err := j.CreateToken(claims)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"error":  0,
		"message": nil,
		"data":    &UserRes{User: user, Token: token},
	})
}

func Register(c *gin.Context) {
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	username := c.PostForm("username")
	password := c.PostForm("password")
	ConfirmPassword := c.PostForm("confirmPassword")
	if password != ConfirmPassword {
		c.JSON(300, gin.H{
			"error":  1,
			"message": "Password is not equal to ConfirmPassword",
			"data":    nil,
		})
		return
	}
	db := mysql.DbObj()
	defer db.Close()
	user := model.User{}
	db.Where("email = ?", email).First(&user)
	if user.ID != 0 {
		c.JSON(300, gin.H{
			"error":  1,
			"message": "User already exists",
			"data":    nil,
		})
		return
	}
	h := md5.New()
	h.Write([]byte(password + email))
	md5Password := hex.EncodeToString(h.Sum(nil))
	userInfo := model.User{Email: email, UserName: username, Password: md5Password, Phone: phone, LoginTime: time.Now().Local()}
	db.Create(&userInfo)
	c.JSON(http.StatusOK, gin.H{
		"error":  0,
		"message": nil,
		"data":    nil,
	})
}
func RefreshToken(c *gin.Context) {
	token := c.DefaultQuery("token", "")
	if token == "" {
		token = c.Request.Header.Get("Authorization")
		if s := strings.Split(token, " "); len(s) == 2 {
			token = s[1]
		}
	}
	j := &middleware.JWT{
		SigningKey: []byte("ccan.blog"),
	}
	//c.String(http.StatusOK, token+"---------------<br>")
	res, err := j.ParseToken(token)
	fmt.Println(token)
	if err != nil {
		if err == middleware.TokenExpired {
			newToken, err := j.RefreshToken(token)
			if err != nil {
				c.String(http.StatusOK, err.Error())
			} else {
				user, _ := j.ParseToken(newToken)
				c.JSON(http.StatusOK, gin.H{
					"error":  0,
					"message": nil,
					"data":    user,
					"token":   newToken,
				})
			}
		} else {
			c.String(http.StatusOK, err.Error())
		}
	} else {
		c.JSON(http.StatusOK, res)
	}
}
