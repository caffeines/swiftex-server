package validators

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/techartificer/swiftex/constants"
	"github.com/techartificer/swiftex/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderCreateReq struct {
	RecipientName         string    `validate:"required" json:"recipientName"`
	RecipientPhone        string    `validate:"required" json:"recipientPhone"`
	RecipientCity         string    `validate:"required" json:"recipientCity"`
	RecipientThana        string    `validate:"required" json:"recipientThana"`
	RecipientArea         string    `validate:"required" json:"recipientArea"`
	RecipientZip          string    `validate:"required" json:"recipientZip"`
	RecipientAddress      string    `validate:"required" json:"recipientAddress"`
	PackageCode           string    `validate:"omitempty" json:"packageCode"`
	PaymentStatus         string    `validate:"required" json:"paymentStatus"`
	Price                 float64   `validate:"required" json:"price"`
	PercelType            string    `validate:"required" json:"percelType"`
	RequestedDeliveryTime time.Time `validate:"omitempty" json:"requestedDeliveryTime"`
	PickAddress           string    `validate:"required" json:"pickAddress"`
	PickHub               string    `validate:"required" json:"pickHub"`
	Comments              string    `validate:"omitempty,max=300" json:"comments"`
	NumberOfItems         int       `validate:"required" json:"numberOfItems"`
	DeliveryType          string    `validate:"required" json:"deliveryType"`
}

func ValidateOrderCreate(ctx echo.Context) (*models.Order, error) {
	body := OrderCreateReq{}
	if err := ctx.Bind(&body); err != nil {
		return nil, err
	}
	if err := GetValidationError(body); err != nil {
		return nil, err
	}
	order := &models.Order{
		ID:                    primitive.NewObjectID(),
		DeliveryBoy:           nil,
		ShopModeratorID:       nil,
		MerchantID:            nil,
		RecipientName:         body.RecipientName,
		RecipientPhone:        body.RecipientPhone,
		RecipientCity:         body.RecipientCity,
		RecipientThana:        body.RecipientThana,
		RecipientZip:          body.RecipientZip,
		RecipientArea:         body.RecipientArea,
		RecipientAddress:      body.RecipientAddress,
		PackageCode:           body.PackageCode,
		PercelType:            body.PercelType,
		RequestedDeliveryTime: body.RequestedDeliveryTime,
		PickAddress:           body.PickAddress,
		PickHub:               body.PickHub,
		Price:                 body.Price,
		NumberOfItems:         body.NumberOfItems,
		Comments:              body.Comments,
		DeliveryType:          body.DeliveryType,
		PaymentStatus:         body.PaymentStatus,
		IsCancelled:           false,
		IsAccepted:            false,
		Status: []models.OrderStatus{
			{
				ID:     primitive.NewObjectID(),
				Text:   constants.Pending,
				Status: constants.Accepted,
				Time:   time.Now().UTC(),
			},
		},
		CreatedAt: time.Now().UTC(),
	}
	return order, nil
}

type OrderStatusUpdateReq struct {
	Text            string             `validate:"required" json:"text"`
	DeleveryBoyID   primitive.ObjectID `validate:"omitempty" json:"deleveryBoy"`
	ShopModeratorID primitive.ObjectID `validate:"omitempty" json:"shopModerator"`
	MerchantID      primitive.ObjectID `validate:"omitempty" json:"merchant"`
	AdminID         primitive.ObjectID `validate:"omitempty" json:"admin"`
	Status          string             `validate:"required" json:"status"`
}

func UpdateOrderStatus(ctx echo.Context) (*models.OrderStatus, error) {
	body := OrderStatusUpdateReq{}
	if err := ctx.Bind(&body); err != nil {
		return nil, err
	}
	if err := GetValidationError(body); err != nil {
		return nil, err
	}

	orderStatus := &models.OrderStatus{
		ID:              primitive.NewObjectID(),
		DeleveryBoyID:   &body.DeleveryBoyID,
		ShopModeratorID: &body.ShopModeratorID,
		MerchantID:      &body.MerchantID,
		AdminID:         &body.AdminID,
		Status:          body.Status,
		Text:            body.Text,
		Time:            time.Now().UTC(),
	}
	return orderStatus, nil
}

type OrderUpdateReq struct {
	DeliveryBoy           primitive.ObjectID `validate:"omitempty" json:"deliveryBoy"`
	RecipientName         string             `validate:"omitempty" json:"recipientName"`
	RecipientPhone        string             `validate:"omitempty" json:"recipientPhone"`
	RecipientCity         string             `validate:"omitempty" json:"recipientCity"`
	RecipientThana        string             `validate:"omitempty" json:"recipientThana"`
	RecipientArea         string             `validate:"omitempty" json:"recipientArea"`
	RecipientZip          string             `validate:"omitempty" json:"recipientZip"`
	RecipientAddress      string             `validate:"omitempty" json:"recipientAddress"`
	PackageCode           string             `validate:"omitempty" json:"packageCode"`
	PaymentStatus         string             `validate:"omitempty" json:"paymentStatus"`
	Price                 float64            `validate:"omitempty" json:"price"`
	PercelType            string             `validate:"omitempty" json:"percelType"`
	RequestedDeliveryTime time.Time          `validate:"omitempty" json:"requestedDeliveryTime"`
	PickAddress           string             `validate:"omitempty" json:"pickAddress"`
	PickHub               string             `validate:"omitempty" json:"pickHub"`
	Comments              string             `validate:"omitempty,max=300" json:"comments"`
	NumberOfItems         int                `validate:"omitempty" json:"numberOfItems"`
	DeliveryType          string             `validate:"omitempty" json:"deliveryType"`
}

func UpdateOrder(ctx echo.Context) (*models.Order, error) {
	body := OrderUpdateReq{}
	if err := ctx.Bind(&body); err != nil {
		return nil, err
	}
	if err := GetValidationError(body); err != nil {
		return nil, err
	}
	UserID := ctx.Get(constants.UserID).(primitive.ObjectID)

	order := &models.Order{
		DeliveryBoy:           &body.DeliveryBoy,
		RecipientName:         body.RecipientName,
		RecipientPhone:        body.RecipientPhone,
		RecipientCity:         body.RecipientCity,
		RecipientThana:        body.RecipientThana,
		RecipientZip:          body.RecipientZip,
		RecipientArea:         body.RecipientArea,
		RecipientAddress:      body.RecipientAddress,
		PackageCode:           body.PackageCode,
		PercelType:            body.PercelType,
		RequestedDeliveryTime: body.RequestedDeliveryTime,
		PickAddress:           body.PickAddress,
		PickHub:               body.PickHub,
		Price:                 body.Price,
		NumberOfItems:         body.NumberOfItems,
		Comments:              body.Comments,
		DeliveryType:          body.DeliveryType,
		PaymentStatus:         body.PaymentStatus,
		DeliveredAt:           nil,
		UpdateBy:              &UserID,
		UpdatedAt:             time.Now().UTC(),
	}
	return order, nil
}
