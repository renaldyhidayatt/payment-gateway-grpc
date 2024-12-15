package api

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/emptypb"
)

type transferHandleApi struct {
	client pb.TransferServiceClient
}

func NewHandlerTransfer(client pb.TransferServiceClient, router *echo.Echo) *transferHandleApi {
	transferHandler := &transferHandleApi{
		client: client,
	}
	routerTransfer := router.Group("/api/transfer")

	routerTransfer.GET("", transferHandler.FindAll)
	routerTransfer.GET("/:id", transferHandler.FindById)
	routerTransfer.GET("/transfer_from/:transfer_from", transferHandler.FindByTransferByTransferFrom)
	routerTransfer.GET("/transfer_to/:transfer_to", transferHandler.FindByTransferByTransferTo)

	routerTransfer.GET("/active", transferHandler.FindByActiveTransfer)
	routerTransfer.GET("/trashed", transferHandler.FindByTrashedTransfer)

	routerTransfer.POST("/create", transferHandler.CreateTransfer)
	routerTransfer.POST("/update/:id", transferHandler.UpdateTransfer)
	routerTransfer.POST("/trashed/:id", transferHandler.TrashTransfer)
	routerTransfer.POST("/restore/:id", transferHandler.RestoreTransfer)
	routerTransfer.DELETE("/:id", transferHandler.DeleteTransferPermanent)

	return transferHandler

}

func (h *transferHandleApi) FindAll(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	ctx := c.Request().Context()

	req := &pb.FindAllTransferRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindAllTransfer(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transfer data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *transferHandleApi) FindById(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})

	}

	ctx := c.Request().Context()

	res, err := h.client.FindByIdTransfer(ctx, &pb.FindByIdTransferRequest{
		TransferId: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transfer data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *transferHandleApi) FindByTransferByTransferFrom(c echo.Context) error {
	transfer_from := c.Param("transfer_from")

	ctx := c.Request().Context()

	res, err := h.client.FindTransferByTransferFrom(ctx, &pb.FindTransferByTransferFromRequest{
		TransferFrom: transfer_from,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transfer data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *transferHandleApi) FindByTransferByTransferTo(c echo.Context) error {
	transfer_to := c.Param("transfer_to")

	ctx := c.Request().Context()

	res, err := h.client.FindTransferByTransferTo(ctx, &pb.FindTransferByTransferToRequest{
		TransferTo: transfer_to,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transfer data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *transferHandleApi) FindByActiveTransfer(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.FindByActiveTransfer(ctx, &emptypb.Empty{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transfer data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *transferHandleApi) FindByTrashedTransfer(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.FindByTrashedTransfer(ctx, &emptypb.Empty{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve transfer data: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *transferHandleApi) CreateTransfer(c echo.Context) error {
	var body requests.CreateTransferRequest

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Validation Error: " + err.Error(),
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.CreateTransfer(ctx, &pb.CreateTransferRequest{
		TransferFrom:   body.TransferFrom,
		TransferTo:     body.TransferTo,
		TransferAmount: int32(body.TransferAmount),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create transfer: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *transferHandleApi) UpdateTransfer(c echo.Context) error {
	var body requests.UpdateTransferRequest

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Validation Error: " + err.Error(),
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.UpdateTransfer(ctx, &pb.UpdateTransferRequest{
		TransferId:     int32(body.TransferID),
		TransferFrom:   body.TransferFrom,
		TransferTo:     body.TransferTo,
		TransferAmount: int32(body.TransferAmount),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update transfer: ",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *transferHandleApi) TrashTransfer(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})

	}

	ctx := c.Request().Context()

	res, err := h.client.TrashedTransfer(ctx, &pb.FindByIdTransferRequest{
		TransferId: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trash transfer:",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *transferHandleApi) RestoreTransfer(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.RestoreTransfer(ctx, &pb.FindByIdTransferRequest{
		TransferId: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore transfer:",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *transferHandleApi) DeleteTransferPermanent(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	ctx := c.Request().Context()

	res, err := h.client.DeleteTransferPermanent(ctx, &pb.FindByIdTransferRequest{
		TransferId: int32(idInt),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete transfer:",
		})
	}

	return c.JSON(http.StatusOK, res)
}
