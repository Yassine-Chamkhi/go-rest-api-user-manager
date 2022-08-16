package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"target/onboarding-assignment/models"
	"target/onboarding-assignment/services"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

type UserHandlerInterface interface {
	Greet() func(c *gin.Context)
	GetUserFromPath() func(c *gin.Context)
	GetAllUsers() func(c *gin.Context)
	PostUser() func(c *gin.Context)
	DeleteUser() func(c *gin.Context)
}

type UserHandler struct {
	UserSvc     services.UserServiceInterface
	ReqsCounter prometheus.Counter
}

func (h *UserHandler) Greet() func(c *gin.Context) {
	return func(c *gin.Context) {
		h.ReqsCounter.Inc()
		c.JSON(http.StatusOK, "Welcome, app has started")
	}
}

func (h *UserHandler) GetUserFromPath() func(c *gin.Context) {
	return func(c *gin.Context) {
		h.ReqsCounter.Inc()
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil || id <= 0 {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Id must be an integer value superior to 0"})
			return
		}

		user, err := h.UserSvc.GetUserById(id)

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}

}

func (h *UserHandler) GetAllUsers() func(c *gin.Context) {
	return func(c *gin.Context) {
		h.ReqsCounter.Inc()
		users, err := h.UserSvc.GetAllUsers()

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

func (h *UserHandler) PostUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		h.ReqsCounter.Inc()
		var user models.User

		err := c.ShouldBindJSON(&user)

		if err != nil {
			err = errors.New("could not bind user data from request body")
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		user, err = h.UserSvc.AddUser(user)

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func (h *UserHandler) DeleteUser() func(c *gin.Context) {
	type idInt struct {
		Id int `json:"id"`
	}
	return func(c *gin.Context) {
		h.ReqsCounter.Inc()
		var idObj idInt

		if err := c.ShouldBindJSON(&idObj); err != nil {
			err = errors.New("could not bind id from request body")
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		err := h.UserSvc.DeleteUserById(idObj.Id)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("successfully deleted user id:%d", idObj.Id)})

	}
}
