package v1

import (
	"assignment_7/internal/entity"
	"assignment_7/internal/usecase"
	"assignment_7/utils"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	u usecase.UserInterface
}

func New(r *gin.Engine, u usecase.UserInterface) {
	h := &Handler{u}

	api := r.Group("/users")
	{
		api.POST("/", h.Register)
		api.POST("/login", h.Login)

		protected := api.Group("/")
		protected.Use(utils.JWTAuthMiddleware())
		protected.Use(utils.RateLimiter(5))
		{
			protected.GET("/me", h.GetMe)
			protected.PATCH("/promote/:id",
				utils.RoleMiddleware("admin"),
				h.Promote,
			)
		}
	}
}

func (h *Handler) Register(c *gin.Context) {
	var dto entity.CreateUserDTO
	c.BindJSON(&dto)

	hash, _ := utils.HashPassword(dto.Password)

	user := entity.User{
		Username: dto.Username,
		Email:    dto.Email,
		Password: hash,
		Role:     "user",
	}

	h.u.Register(&user)
	c.JSON(200, user)
}

func (h *Handler) Login(c *gin.Context) {
	var dto entity.LoginUserDTO
	c.BindJSON(&dto)

	token, _ := h.u.Login(&dto)
	c.JSON(200, gin.H{"token": token})
}

func (h *Handler) GetMe(c *gin.Context) {
	id := c.GetString("userID")
	user, _ := h.u.GetMe(id)
	c.JSON(200, user)
}

func (h *Handler) Promote(c *gin.Context) {
	id := c.Param("id")
	h.u.Promote(id)
	c.JSON(200, gin.H{"message": "promoted"})
}
