package agent

import (
	request "github.com/bagusyanuar/go-olin-bags/app/http/request/agent"
	repository "github.com/bagusyanuar/go-olin-bags/app/repositories/agent"
	"github.com/bagusyanuar/go-olin-bags/model"
	"github.com/google/uuid"
)

type PurchaseItem interface {
	List(accountID string) ([]model.PurchaseItem, error)
	AddItem(accountID string, request request.PurchaseItemRequest) (*model.PurchaseItem, error)
}

type PurchaseItemService struct {
	PurchaseItemRepository repository.PurchaseItem
	ItemRepository         repository.Item
}

// AddItem implements PurchaseItem.
func (svc *PurchaseItemService) AddItem(accountID string, request request.PurchaseItemRequest) (*model.PurchaseItem, error) {

	purchaserID, err := uuid.Parse(accountID)
	if err != nil {
		return nil, err
	}

	item, err := svc.ItemRepository.GetItemByID(request.ItemID)
	if err != nil {
		return nil, err
	}

	price := item.Price
	total := (price * int64(request.Qty))
	e := model.PurchaseItem{
		PurchaserID: purchaserID,
		ItemID:      &item.ID,
		Price:       uint64(price),
		Qty:         request.Qty,
		Total:       uint64(total),
	}
	return svc.PurchaseItemRepository.AddItem(e)
}

// List implements PurchaseItem.
func (svc *PurchaseItemService) List(accountID string) ([]model.PurchaseItem, error) {
	return svc.PurchaseItemRepository.List(accountID)
}

func NewPurchaseItemService(purchaseItemRepository repository.PurchaseItem, itemRepository repository.Item) PurchaseItem {
	return &PurchaseItemService{
		PurchaseItemRepository: purchaseItemRepository,
		ItemRepository:         itemRepository,
	}
}
