package agent

import (
	request "github.com/bagusyanuar/go-olin-bags/app/http/request/agent"
	repository "github.com/bagusyanuar/go-olin-bags/app/repositories/agent"
	"github.com/bagusyanuar/go-olin-bags/model"
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
	for _, cart := range carts {
		subTotal += cart.Total
	}

	// e := model.Purchasing{

	// }
	return nil, err
}

// List implements Purchasing.
func (svc *PurchasingService) List(accountID string, limit int, offset int) ([]model.Purchasing, error) {
	panic("unimplemented")
}

func NewPurchasingService(purchasingRepository repository.Purchasing, purchasingItemRepository repository.PurchaseItem) Purchasing {
	return &PurchasingService{
		PurchasingRepository:     purchasingRepository,
		PurchasingItemRepository: purchasingItemRepository,
	}
}
