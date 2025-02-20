package v1

import (
	"gin/pkg/util"
	"gin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAddress(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createAddressService := service.AddressService{}
	ctx := c.Request.Context()
	if err := c.ShouldBind(&createAddressService); err == nil {
		res := createAddressService.Create(ctx, claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func ListAddress(c *gin.Context) {
	listAddressService := service.AddressService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listAddressService); err == nil {
		res := listAddressService.List(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func DeleteAddress(c *gin.Context) {
	deleteAddressService := service.AddressService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteAddressService); err == nil {
		res := deleteAddressService.Delete(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func ShowAddress(c *gin.Context) {
	showAddressService := service.AddressService{}
	if err := c.ShouldBind(&showAddressService); err == nil {
		res := showAddressService.Show(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func UpdateAddress(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	updateAddressService := service.AddressService{}
	ctx := c.Request.Context()
	if err := c.ShouldBind(&updateAddressService); err == nil {
		res := updateAddressService.Update(ctx, claim.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
