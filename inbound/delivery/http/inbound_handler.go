package http

import (
	"context"
	"encoding/json"
	"encoding/xml"
	httpresponse "github.com/gomarkho/demo-callpicker/network/response"
	"net/http"

	"github.com/gomarkho/demo-callpicker/inbound"
	"github.com/gomarkho/demo-callpicker/model"

	"github.com/labstack/echo"
)

type ResponseError struct {
	Message string `json:"message"`
}

type InboundHandler struct {
	InboundUseCase inbound.UseCase
}

func NewInboundHandler(pub *echo.Group, priv *echo.Group, cu inbound.UseCase) *InboundHandler {
	handler := &InboundHandler{cu}

	pub.POST("/dialplan", handler.InsertDialPlan)
	pub.POST("/xml", handler.GetXML)
	pub.DELETE("/dialplan", handler.DeleteDialPlan)
	pub.GET("/dialplan", handler.GetDialPlan)
	//pub.GET("/dialplan", handler.UpdateDialPlan)

	return handler
}

func (i *InboundHandler) InsertDialPlan(c echo.Context) error {
	req := inbound.InsertDialPlan{}
	err := c.Bind(&req)
	if err != nil {
		return httpresponse.CreateBadResponseWithCode(&c, http.StatusBadRequest, model.ErrIncorrectRequest.Error(), err.Error(), model.CodeIncorrectRequest)
	}
	// validate input request body
	if err := c.Validate(req); err != nil {
		return httpresponse.CreateBadResponseWithCode(&c, http.StatusBadRequest, model.ErrIncorrectRequest.Error(), err.Error(), model.CodeIncorrectRequest)
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	result, code, err := i.InboundUseCase.InsertDialPlan(ctx, &inbound.InsertDialPlan{
		Source:      req.Source,
		Destination: req.Destination,
		Say:         req.Say,
		Play:        req.Play,
		Record:      req.Record,
		Dial:        req.Dial,
	})
	if code == 1 {
		return httpresponse.CreateBadResponseWithCode(&c, http.StatusConflict, model.ErrConflict.Error(), err.Error(), model.CodeAlreadyExists)
	}
	if err != nil {
		return httpresponse.CreateBadResponseWithCode(&c, http.StatusBadRequest, model.ErrTryAgain.Error(), err.Error(), model.CodeTryAgain)
	}
	data, err := json.Marshal(result)
	if err != nil {
		return httpresponse.CreateBadResponseWithCode(&c, http.StatusBadRequest, model.ErrTryAgain.Error(), err.Error(), model.CodeTryAgain)
	}
	return httpresponse.CreateSuccessResponse(&c, http.StatusCreated, model.MsgSuccess, model.SubMsgInsert, data)
}

func (i *InboundHandler) GetXML(c echo.Context) error {

	src := c.FormValue("From")
	dest := c.FormValue("To")
	call_sid := c.FormValue("CallSid")
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	result, err := i.InboundUseCase.GetXml(ctx, &inbound.XMLOptions{
		Source:      src,
		Destination: dest,
		CallSid:     call_sid,
	})
	if err != nil {
		return httpresponse.CreateBadResponseWithCode(&c, http.StatusBadRequest, model.ErrTryAgain.Error(), err.Error(), model.CodeTryAgain)
	}
	data, err := xml.Marshal(result)
	if err != nil {
		return httpresponse.CreateBadResponseWithCode(&c, http.StatusBadRequest, model.ErrTryAgain.Error(), err.Error(), model.CodeTryAgain)
	}
	return httpresponse.CreateXmlSuccessResponse(&c, http.StatusOK, model.MsgSuccess, model.SubMsgGetXml, data)
}

func (i *InboundHandler) DeleteDialPlan(c echo.Context) error {

	req := inbound.DeleteDialPlan{}
	err := c.Bind(&req)
	if err != nil {
		return httpresponse.CreateBadResponseWithCode(&c, http.StatusBadRequest, model.ErrIncorrectRequest.Error(), err.Error(), model.CodeIncorrectRequest)
	}
	if err := c.Validate(req); err != nil {
		return httpresponse.CreateBadResponseWithCode(&c, http.StatusBadRequest, model.ErrIncorrectRequest.Error(), err.Error(), model.CodeIncorrectRequest)
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	code, err := i.InboundUseCase.DeleteDialPlan(ctx, inbound.DeleteDialPlan{
		Source:      req.Source,
		Destination: req.Destination,
	})
	if code == 6 {
		return httpresponse.CreateBadResponseWithCode(&c, http.StatusNotFound, model.ErrTryAgain.Error(), err.Error(), model.CodeNotExists)
	}
	if err != nil {
		return httpresponse.CreateBadResponseWithCode(&c, http.StatusBadRequest, model.ErrTryAgain.Error(), err.Error(), model.CodeTryAgain)
	}
	return httpresponse.CreateSuccessResponseWithoutData(&c, http.StatusOK, model.MsgSuccess, model.SubMsgDelete)

}

func (i *InboundHandler) GetDialPlan(c echo.Context) error {

	req := inbound.GetDialPlan{}
	err := c.Bind(&req)
	if err != nil {
		return httpresponse.CreateBadResponseWithCode(&c, http.StatusBadRequest, model.ErrIncorrectRequest.Error(), err.Error(), model.CodeIncorrectRequest)
	}
	if err := c.Validate(req); err != nil {
		return httpresponse.CreateBadResponseWithCode(&c, http.StatusBadRequest, model.ErrIncorrectRequest.Error(), err.Error(), model.CodeIncorrectRequest)
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	result, _, err := i.InboundUseCase.GetDialPlan(ctx, inbound.GetDialPlan{
		Source:      req.Source,
		Destination: req.Destination,
	})
	if err != nil {
		return httpresponse.CreateBadResponseWithCode(&c, http.StatusNotFound, model.ErrTryAgain.Error(), err.Error(), model.CodeTryAgain)
	}
	data, err := json.Marshal(result)
	if err != nil {
		return httpresponse.CreateBadResponseWithCode(&c, http.StatusBadRequest, model.ErrTryAgain.Error(), err.Error(), model.CodeTryAgain)
	}
	return httpresponse.CreateSuccessResponse(&c, http.StatusOK, model.MsgSuccess, model.SubMsgGetDialPlan, data)

}
