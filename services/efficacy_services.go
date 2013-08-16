package services

import (
	"github.com/kobeld/duoerl/models/efficacies"
	"github.com/kobeld/duoerlapi"
)

func GetEfficacyById(efficacyId string) *duoerlapi.Efficacy {
	return GetEfficacyMap()[efficacyId]
}

func GetEfficaciesByIds(efficacyIds []string) (apiEfficacies []*duoerlapi.Efficacy) {
	for _, efficacyId := range efficacyIds {
		if efficacyId == "" {
			continue
		}
		apiEfficacies = append(apiEfficacies, GetEfficacyById(efficacyId))
	}
	return
}

// ---- Private ----

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
