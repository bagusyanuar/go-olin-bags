package admin

import (
	request "github.com/bagusyanuar/go-olin-bags/app/http/request/admin"
	repository "github.com/bagusyanuar/go-olin-bags/app/repositories/admin"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/bagusyanuar/go-olin-bags/model"
	"github.com/google/uuid"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Item interface {
	FindAll(q string, limit, offset int) ([]model.Item, error)
	FindByID(id string) (*model.Item, error)
	Create(request request.ItemRequest) (*model.Item, error)
}

type ItemService struct {
	ItemRepository repository.Item
}

// Create implements Item.
func (svc *ItemService) Create(request request.ItemRequest) (*model.Item, error) {
	materialID, err := uuid.Parse(request.MaterialID)
	if err != nil {
		return nil, err
	}
	e := model.Item{
		MaterialID:  materialID,
		Name:        cases.Title(language.Indonesian, cases.Compact).String(request.Name),
		Description: request.Description,
		Price:       request.Price,
	}
	return svc.ItemRepository.Create(e)
}

// FindAll implements Item.
func (svc *ItemService) FindAll(q string, limit int, offset int) ([]model.Item, error) {
	if limit == 0 {
		limit = common.DefaultLimit
	}
	return svc.ItemRepository.FindAll(q, limit, offset)
}

// FindByID implements Item.
func (svc *ItemService) FindByID(id string) (*model.Item, error) {
	return svc.ItemRepository.FindByID(id)
}

func NewItemService(itemRepository repository.Item) Item {
	return &ItemService{
		ItemRepository: itemRepository,
	}
}
