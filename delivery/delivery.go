package delivery

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/ilamazuliaf/desaku/models"
	"github.com/ilamazuliaf/desaku/usecase"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type ResponError struct {
	Message string `json:"message"`
}

type newDeliveryHandler struct {
	Ds usecase.Usecase
}

func NewHandler(e *echo.Echo, u usecase.Usecase) {
	handler := &newDeliveryHandler{Ds: u}
	e.GET("/login", handler.Login)

	// r := e.Group("")
	s := e.Group("setting")

	e.GET("/person", handler.GetPerson)
	e.POST("/person", handler.InsertPerson)
	e.GET("/person/:uuid", handler.GetDetailPerson)

	e.GET("/total", handler.TotalPerson)
	e.GET("/persentase", handler.Persentase)
	// Setting
	s.GET("/pendidikan", handler.SettingPendidikan)
	s.POST("/pendidikan", handler.AddPendidikan)
	s.GET("/pekerjaan", handler.SettingPekerjaan)
	s.POST("/pekerjaan", handler.AddPekerjaan)
	s.GET("/penghasilan", handler.SettingPenghasilan)
	s.POST("/penghasilan", handler.AddPenghasilan)
}

func (h *newDeliveryHandler) GetPerson(c echo.Context) error {
	ctx := c.Request().Context()
	l := c.QueryParam("limit")
	p := c.QueryParam("page")
	var maxLimit int = 1000

	if l == "" {
		l = "25"
	}
	if p == "" {
		p = "1"
	}
	limit, err := strconv.Atoi(l)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponError{Message: "Limit harus berupa angka"})
	} else if limit > maxLimit {
		return c.JSON(400, ResponError{Message: "Maksimal limit 1000"})
	}
	page, err := strconv.Atoi(p)
	if err != nil {
		return c.JSON(400, ResponError{Message: "Page harus berupa angka"})
	}

	listPerson, totalData, err := h.Ds.GetPerson(ctx, page, limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponError{Message: err.Error()})
	}
	totalPage := float64(totalData) / float64(limit)
	total := int(totalPage)
	if float64(total) < totalPage {
		total++
	}
	c.Response().Header().Set("x-total-pagination-page", strconv.Itoa(total))
	c.Response().Header().Set("x-pagination-limit-perpage", l)
	c.Response().Header().Set("x-total-data", strconv.Itoa(totalData))
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(200)
	return json.NewEncoder(c.Response()).Encode(listPerson)
	// return c.JSON(200, listPerson)
}

func (h *newDeliveryHandler) GetDetailPerson(c echo.Context) error {
	ctx := c.Request().Context()
	uuid := c.Param("uuid")
	result, err := h.Ds.GetDetailPerson(ctx, uuid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponError{Message: err.Error()})
	}
	return c.JSON(200, result)
}

func (h *newDeliveryHandler) SettingPekerjaan(c echo.Context) error {
	ctx := c.Request().Context()
	listPekerjaan, err := h.Ds.SettingPekerjaan(ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponError{Message: err.Error()})
	}
	return c.JSON(200, listPekerjaan)
}

func (h *newDeliveryHandler) SettingPendidikan(c echo.Context) error {
	ctx := c.Request().Context()
	listPendidikan, err := h.Ds.SettingPendidikan(ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponError{Message: err.Error()})
	}
	return c.JSON(200, listPendidikan)
}

func (h *newDeliveryHandler) SettingPenghasilan(c echo.Context) error {
	ctx := c.Request().Context()
	listPenghasilan, err := h.Ds.SettingPenghasilan(ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponError{Message: err.Error()})
	}
	return c.JSON(200, listPenghasilan)
}

func (h *newDeliveryHandler) InsertPerson(c echo.Context) error {
	var person models.Person
	err := c.Bind(&person)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err = c.Validate(&person); err != nil {
		return c.JSON(http.StatusBadRequest, ResponError{Message: err.Error()})
	}
	ctx := c.Request().Context()
	pembuat := ctx.Value("userInfo").(jwt.MapClaims)
	if ctx == nil {
		ctx = context.Background()
	}

	if err := h.Ds.InsertPerson(ctx, &person, pembuat["uuid"].(string)); err != nil {
		return c.JSON(http.StatusBadRequest, ResponError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, person)
}

func (h *newDeliveryHandler) AddPekerjaan(c echo.Context) error {
	var pekerjaan models.Pekerjaan
	if err := c.Bind(&pekerjaan); err != nil {
		return c.JSON(400, ResponError{Message: err.Error()})
	}
	if err := c.Validate(&pekerjaan); err != nil {
		return c.JSON(400, ResponError{Message: err.Error()})
	}
	ctx := c.Request().Context()
	if err := h.Ds.AddPekerjaan(ctx, &pekerjaan); err != nil {
		return c.JSON(400, ResponError{Message: err.Error()})
	}
	return c.JSON(201, pekerjaan)
}

func (h *newDeliveryHandler) AddPendidikan(c echo.Context) error {
	var pendidikan models.Pendidikan
	if err := c.Bind(&pendidikan); err != nil {
		return c.JSON(400, ResponError{Message: err.Error()})
	}
	if err := c.Validate(&pendidikan); err != nil {
		return c.JSON(400, ResponError{Message: err.Error()})
	}
	ctx := c.Request().Context()
	if err := h.Ds.AddPendidikan(ctx, &pendidikan); err != nil {
		return c.JSON(400, ResponError{Message: err.Error()})
	}
	return c.JSON(201, pendidikan)
}
func (h *newDeliveryHandler) AddPenghasilan(c echo.Context) error {
	var penghasilan models.Penghasilan
	if err := c.Bind(&penghasilan); err != nil {
		return c.JSON(400, ResponError{Message: err.Error()})
	}
	if err := c.Validate(&penghasilan); err != nil {
		return c.JSON(400, ResponError{Message: err.Error()})
	}
	ctx := c.Request().Context()
	if err := h.Ds.AddPenghasilan(ctx, &penghasilan); err != nil {
		return c.JSON(400, ResponError{Message: err.Error()})
	}
	return c.JSON(201, penghasilan)
}

func (l *newDeliveryHandler) Login(c echo.Context) error {
	username, password, ok := c.Request().BasicAuth()
	if !ok {
		return c.JSON(http.StatusBadRequest, ResponError{Message: "Invalid Login"})
	}
	ctx := c.Request().Context()
	userInfo, err := l.Ds.GetUser(ctx, username, password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponError{Message: err.Error()})
	}
	claims := models.MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    models.APP_NAME,
			ExpiresAt: time.Now().Add(models.LOGIN_EXPIRATION_DURATION).Unix(),
		},
		UUID:     userInfo.UUID,
		Username: userInfo.Username,
		Email:    userInfo.Email,
		Group:    userInfo.Group,
	}

	token := jwt.NewWithClaims(
		models.JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(models.JWT_SIGNATURE_KEY)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponError{Message: err.Error()})
	}
	c.Response().Header().Set("x-token", signedToken)
	c.Response().Header().Set("Content-Type", "application/json")
	return c.JSON(http.StatusOK, ResponError{Message: "Sukses Broo"})
}

func (h *newDeliveryHandler) TotalPerson(c echo.Context) error {
	ctx := c.Request().Context()
	total, err := h.Ds.GetTotal(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, total)
}

func (h *newDeliveryHandler) Persentase(c echo.Context) error {
	ctx := c.Request().Context()
	persentase, _ := h.Ds.Persentase(ctx)

	return c.JSON(200, persentase)
}
