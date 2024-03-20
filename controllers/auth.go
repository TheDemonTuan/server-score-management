package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
	"qldiemsv/common"
	"qldiemsv/models/entity"
	"qldiemsv/models/req"
	"strconv"
	"time"
)

func AuthLogin(c *fiber.Ctx) error {
	bodyData, err := common.Validator[req.AuthLogin](c)

	if err != nil || bodyData == nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var userRecord entity.User

	if err := common.DBConn.First(&userRecord, "user_name = ?", bodyData.UserName).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusBadRequest, "Username không tồn tại")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Có lỗi trong khi truy vấn cơ sở dữ liệu")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userRecord.Password), []byte(bodyData.Password)); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Mật khẩu không đúng")
	}

	if err := createJWT(c, userRecord.ID); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(common.NewResponse(fiber.StatusOK, "Đăng nhập thành công", userRecord))
}

func AuthRegister(c *fiber.Ctx) error {
	bodyData, err := common.Validator[req.AuthRegister](c)

	if err != nil || bodyData == nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var userRecord entity.User

	if err := common.DBConn.First(&userRecord, "user_name = ?", bodyData.UserName).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusInternalServerError, "Có lỗi trong khi truy vấn cơ sở dữ liệu")
		}
	}

	if userRecord != (entity.User{}) {
		return fiber.NewError(fiber.StatusBadRequest, "Username đã tồn tại")
	}

	hashPassword, hashPasswordErr := bcrypt.GenerateFromPassword([]byte(bodyData.Password), 11)

	if hashPasswordErr != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Có lỗi trong khi tạo tài khoản")
	}

	newUser := entity.User{
		FirstName: bodyData.FirstName,
		LastName:  bodyData.LastName,
		UserName:  bodyData.UserName,
		Password:  string(hashPassword),
	}

	if err := common.DBConn.Create(&newUser).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Có lỗi trong khi tạo tài khoản")
	}

	if err := createJWT(c, newUser.ID); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(common.NewResponse(fiber.StatusOK, "Đăng ký tài khoản thành công", newUser))
}

func createJWT(c *fiber.Ctx, userId uint) error {
	// Create the Claims
	claims := jwt.MapClaims{
		"uid": strconv.Itoa(int(userId)),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenSignedString, tokenSignedErr := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if tokenSignedErr != nil {
		return errors.New("Có lỗi trong khi tạo token")
	}
	//Create cookie
	cookie := new(fiber.Cookie)
	cookie.Name = os.Getenv("JWT_NAME")
	cookie.Value = tokenSignedString
	cookie.Secure = os.Getenv("APP_ENV") == "production"
	cookie.HTTPOnly = true
	cookie.SameSite = "strict"
	cookie.Domain = "." + os.Getenv("JWT_DOMAIN")
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.Cookie(cookie)

	return nil
}

func AuthVerify(c *fiber.Ctx) error {
	currUserId, currUserIdIsOk := c.Locals("currUserId").(string)

	if !currUserIdIsOk {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	c.Next()
	userRecord := entity.User{}

	if err := common.DBConn.First(&userRecord, "id = ?", currUserId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(fiber.StatusUnauthorized, "Không tìm thấy user")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Có lỗi trong khi truy vấn cơ sở dữ liệu")
	}

	return c.JSON(common.NewResponse(
		fiber.StatusOK,
		"Success",
		userRecord),
	)
}

func AuthLogout(c *fiber.Ctx) error {
	cookie := new(fiber.Cookie)
	cookie.Name = os.Getenv("JWT_NAME")
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-24 * time.Hour)
	c.Cookie(cookie)
	c.ClearCookie(os.Getenv("JWT_NAME"))
	return c.JSON(common.NewResponse(fiber.StatusOK, "Đăng xuất thành công", nil))
}
