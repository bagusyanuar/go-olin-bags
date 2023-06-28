package agent

import (
	"fmt"
	"time"

	request "github.com/bagusyanuar/go-olin-bags/app/http/request/agent"
	repository "github.com/bagusyanuar/go-olin-bags/app/repositories/agent"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/bagusyanuar/go-olin-bags/model"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Purchasing interface {
	Checkout(accountID string, request request.PurchasingRequest) (*model.Purchasing, error)
	List(accountID string, limit, offset int) ([]model.Purchasing, error)
}

type PurchasingService struct {
	PurchasingRepository     repository.Purchasing
	PurchasingItemRepository repository.PurchaseItem
}

// Checkout implements Purchasing.
func (svc *PurchasingService) Checkout(accountID string, request request.PurchasingRequest) (*model.Purchasing, error) {
	carts, err := svc.PurchasingItemRepository.List(accountID)
	if err != nil {
		return nil, err
	}
	var subTotal uint64
	var total uint64
	var discount uint64
	for _, cart := range carts {
		subTotal += cart.Total
	}

	total = subTotal + request.ShippingCost - discount
	productionHouseID, err := uuid.Parse(request.ProductionHouseID)
	if err != nil {
		return nil, err
	}

	purchaserID, err := uuid.Parse(accountID)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	formattedTime := now.Format(common.CodeDateFormat)
	purchaseNumber := fmt.Sprintf("PO-%s", formattedTime)
	e := model.Purchasing{
		ProductionHouseID:   productionHouseID,
		PurchaserID:         purchaserID,
		PurchaseNumber:      purchaseNumber,
		Date:                datatypes.Date(now),
		ShippingDestination: request.ShippingDestination,
		SubTotal:            subTotal,
		ShippingCost:        request.ShippingCost,
		Discount:            discount,
		Total:               total,
		Status:              common.PurchasingOnWaiting,
		PurchaseItems:       carts,
	}

	return svc.PurchasingRepository.Checkout(accountID, e, carts)
}

// List implements Purchasing.
func (svc *PurchasingService) List(accountID string, limit int, offset int) ([]model.Purchasing, error) {
	if limit == 0 {
		limit = common.DefaultLimit
	}
	return svc.PurchasingRepository.List(accountID, limit, offset)
}

func NewPurchasingService(purchasingRepository repository.Purchasing, purchasingItemRepository repository.PurchaseItem) Purchasing {
	return &PurchasingService{
		PurchasingRepository:     purchasingRepository,
		PurchasingItemRepository: purchasingItemRepository,
	}
}
