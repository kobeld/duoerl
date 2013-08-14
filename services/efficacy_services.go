package services

import (
	"github.com/kobeld/duoerl/models/efficacies"
	"github.com/kobeld/duoerlapi"
)

// -------

func toApiEfficacy(efficacy *efficacies.Efficacy) *duoerlapi.Efficacy {
	apiEfficacy := new(duoerlapi.Efficacy)
	if efficacy != nil {
		apiEfficacy = &duoerlapi.Efficacy{
			Id:       efficacy.Id.Hex(),
			Name:     efficacy.Name,
			ParentId: efficacy.ParentId.Hex(),
		}
	}

	return apiEfficacy
}
