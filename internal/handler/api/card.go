package api

import (
	"MamangRust/paymentgatewaygrpc/internal/domain/requests"
	"MamangRust/paymentgatewaygrpc/internal/domain/response"
	"MamangRust/paymentgatewaygrpc/internal/pb"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type cardHandleApi struct {
	card pb.CardServiceClient
}

func NewHandlerCard(card pb.CardServiceClient, router *echo.Echo) *cardHandleApi {
	cardHandler := &cardHandleApi{
		card: card,
	}
	routerCard := router.Group("/api/card")

	routerCard.GET("/hello", cardHandler.handleHello)

	routerCard.GET("/", cardHandler.FindAll)
	routerCard.GET("/:id", cardHandler.FindById)
	routerCard.GET("/user/:id", cardHandler.FindByUserID)
	routerCard.GET("/active/:id", cardHandler.FindByActive)
	routerCard.GET("/trashed", cardHandler.FindByTrashed)
	routerCard.GET("/card_number/:card_number", cardHandler.FindByCardNumber)

	routerCard.POST("/create", cardHandler.CreateCard)
	routerCard.POST("/update/:id", cardHandler.UpdateCard)
	routerCard.POST("/trashed/:id", cardHandler.TrashedCard)
	routerCard.POST("/restore/:id", cardHandler.RestoreCard)

	routerCard.DELETE("/permanent/:id", cardHandler.DeleteCardPermanent)

	return cardHandler
}

// handleHello menangani permintaan GET "/hello" dan mengembalikan pesan "Hello".
// @Summary Mengembalikan pesan "Hello"
// @Tags Auth
// @Description Mengembalikan pesan "Hello"
// @Produce json
// @Success 200 {string} string "Hello"
// @Router /auth/hello [get]
func (h *cardHandleApi) handleHello(c echo.Context) error {
	return c.String(200, "Hello")
}

// @Summary Retrieve all cards
// @Tags Card
// @Description Retrieve all cards with pagination
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Number of data per page"
// @Param search query string false "Search keyword"
// @Success 200 {object} pb.ApiResponsePaginationCard "Card data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve card data"
// @Router /api/card [get]
func (h *cardHandleApi) FindAll(c echo.Context) error {
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

	req := &pb.FindAllCardRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	cards, err := h.card.FindAllCard(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch card records: ",
		})
	}

	return c.JSON(http.StatusOK, cards)
}

// @Summary Retrieve card by ID
// @Tags Card
// @Description Retrieve a card by its ID
// @Accept json
// @Produce json
// @Param id path int true "Card ID"
// @Success 200 {object} pb.ApiResponseCard "Card data"
// @Failure 400 {object} response.ErrorResponse "Invalid card ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve card record"
// @Router /api/card/{id} [get]
func (h *cardHandleApi) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid card ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdCardRequest{
		CardId: int32(id),
	}

	card, err := h.card.FindByIdCard(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch card record: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, card)
}

// @Summary Retrieve cards by user ID
// @Tags Card
// @Description Retrieve a list of cards associated with a user by their ID
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} pb.ApiResponseCards "Card data"
// @Failure 400 {object} response.ErrorResponse "Invalid user ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve card record"
// @Router /api/card/user/{id} [get]
func (h *cardHandleApi) FindByUserID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid user ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByUserIdCardRequest{
		UserId: int32(id),
	}

	card, err := h.card.FindByUserIdCard(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch card record: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, card)
}

// @Summary Retrieve active card by Saldo ID
// @Tags Card
// @Description Retrieve an active card associated with a Saldo ID
// @Accept json
// @Produce json
// @Param id path string true "Saldo ID"
// @Success 200 {object} pb.ApiResponseCard "Card data"
// @Failure 400 {object} response.ErrorResponse "Invalid Saldo ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve card record"
// @Router /api/card/active/{id} [get]
func (h *cardHandleApi) FindByActive(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByActiveCardRequest{
		SaldoId: int32(idInt),
	}

	card, err := h.card.FindByActiveCard(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch card record: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, card)
}

// @Summary Retrieve trashed cards
// @Tags Card
// @Description Retrieve a list of trashed cards
// @Accept json
// @Produce json
// @Success 200 {object} pb.ApiResponseCards "Card data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve card record"
// @Router /api/card/trashed [get]
func (h *cardHandleApi) FindByTrashed(c echo.Context) error {

	res, err := h.card.FindByTrashedCard(c.Request().Context(), &emptypb.Empty{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch card record: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

// @Summary Retrieve card by card number
// @Tags Card
// @Description Retrieve a card by its card number
// @Accept json
// @Produce json
// @Param card_number path string true "Card number"
// @Success 200 {object} pb.ApiResponseCard "Card data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve card record"
// @Router /api/card/{card_number} [get]
func (h *cardHandleApi) FindByCardNumber(c echo.Context) error {
	cardNumber := c.Param("card_number")

	ctx := c.Request().Context()

	req := &pb.FindByCardNumberRequest{
		CardNumber: cardNumber,
	}

	card, err := h.card.FindByCardNumber(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch card record: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, card)
}

// @Summary Create a new card
// @Tags Card
// @Description Create a new card for a user
// @Accept json
// @Produce json
// @Param CreateCardRequest body requests.CreateCardRequest true "Create card request"
// @Success 200 {object} pb.ApiResponseCard "Created card"
// @Failure 400 {object} response.ErrorResponse "Bad request or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to create card"
// @Router /api/card [post]
func (h *cardHandleApi) CreateCard(c echo.Context) error {
	var body requests.CreateCardRequest

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: " + err.Error(),
		})
	}

	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Validation Error: " + err.Error(),
		})
	}

	ctx := c.Request().Context()

	req := &pb.CreateCardRequest{
		UserId:       int32(body.UserID),
		CardType:     body.CardType,
		ExpireDate:   timestamppb.New(body.ExpireDate),
		Cvv:          body.CVV,
		CardProvider: body.CardProvider,
	}

	card, err := h.card.CreateCard(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create card: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, card)
}

// @Summary Update a card
// @Tags Card
// @Description Update a card for a user
// @Accept json
// @Produce json
// @Param id path int true "Card ID"
// @Param UpdateCardRequest body requests.UpdateCardRequest true "Update card request"
// @Success 200 {object} pb.ApiResponseCard "Updated card"
// @Failure 400 {object} response.ErrorResponse "Bad request or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to update card"
// @Router /api/card/{id} [post]
func (h *cardHandleApi) UpdateCard(c echo.Context) error {
	var body requests.UpdateCardRequest

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: " + err.Error(),
		})
	}

	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Validation Error: " + err.Error(),
		})
	}

	ctx := c.Request().Context()

	req := &pb.UpdateCardRequest{
		CardId:       int32(body.CardID),
		UserId:       int32(body.UserID),
		CardType:     body.CardType,
		ExpireDate:   timestamppb.New(body.ExpireDate),
		Cvv:          body.CVV,
		CardProvider: body.CardProvider,
	}

	card, err := h.card.UpdateCard(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update card: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, card)
}

// @Summary Trashed a card
// @Tags Card
// @Description Trashed a card by its ID
// @Accept json
// @Produce json
// @Param id path int true "Card ID"
// @Success 200 {object} pb.ApiResponseCard "Trashed card"
// @Failure 400 {object} response.ErrorResponse "Bad request or invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to trashed card"
// @Router /api/card/trashed/{id} [post]
func (h *cardHandleApi) TrashedCard(c echo.Context) error {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdCardRequest{
		CardId: int32(idInt),
	}

	card, err := h.card.TrashedCard(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trashed card: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, card)
}

// @Summary Restore a card
// @Tags Card
// @Description Restore a card by its ID
// @Accept json
// @Produce json
// @Param id path int true "Card ID"
// @Success 200 {object} pb.ApiResponseCard "Restored card"
// @Failure 400 {object} response.ErrorResponse "Bad request or invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore card"
// @Router /api/card/restore/{id} [post]
func (h *cardHandleApi) RestoreCard(c echo.Context) error {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdCardRequest{
		CardId: int32(idInt),
	}

	card, err := h.card.RestoreCard(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore card: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, card)
}

// @Summary Delete a card permanently
// @Tags Card
// @Description Delete a card by its ID permanently
// @Accept json
// @Produce json
// @Param id path int true "Card ID"
// @Success 200 {object} pb.ApiResponseCard "Deleted card"
// @Failure 400 {object} response.ErrorResponse "Bad request or invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete card"
// @Router /api/card/delete/{id} [delete]
func (h *cardHandleApi) DeleteCardPermanent(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Bad Request: Invalid ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdCardRequest{
		CardId: int32(idInt),
	}

	card, err := h.card.DeleteCardPermanent(ctx, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete card: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, card)
}
