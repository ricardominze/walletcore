package gateway

import "github.com/ricardominze/ms-wallet/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(Id string) (*entity.Account, error)
}
