package ctrl

import (
	"net/http"

	"github.com/elvizlai/swagg-example/gateway/model"
	"github.com/elvizlai/swagg-example/pb"
)

// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @Accept json
// @Produce json
// @Param id path int true "Account ID"
// @Success 200 {object} pb.BuildingInfo
// @Failure 404 {object} pb.BuildingInfo
// @Router /accounts/{id} [get]
func ShowAccount(rw http.ResponseWriter, r *http.Request) {
	req := &pb.BuildingInfo{}
	_ = req
	_ = &pb.Coordinate{}
}

// ShowAccount2 godoc
// @Summary Show a account
// @Description get string by ID
// @Accept json
// @Produce json
// @Param id path int true "Account ID"
// @Success 200 {object} model.User
// @Router /accounts2/{id} [get]
func ShowAccount2(rw http.ResponseWriter, r *http.Request) {
	req := model.User{}
	_ = req
}
