package mapper

import (
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
)

func MapServiceDbToDomain(sdb *gateway.Service_DB) *models.Service {
	return &models.Service{
		Id:          sdb.Id,
		Name:        sdb.Name,
		TargetUrl:   sdb.TargetUrl,
		Path:        sdb.Path,
		Description: sdb.Description,
		CreatedAt:   sdb.CreatedAt,
		UpdatedAt:   sdb.UpdatedAt,
		Routes:      make([]models.Route, 0),
	}
}
