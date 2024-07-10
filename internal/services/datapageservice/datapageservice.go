package datapageservice

import (
	"context"
	"fmt"

	"github.com/outagelab/outagelab/internal/models"
)

type datapageService struct {
	accountRepo accountRepo
	apiKeyRepo  apiKeyRepo
}

func New(accountRepo accountRepo, apiKeyRepo apiKeyRepo) *datapageService {
	return &datapageService{
		accountRepo: accountRepo,
		apiKeyRepo:  apiKeyRepo,
	}
}

func (ds *datapageService) GenerateDatapage(ctx context.Context, apiKeyStr string, request *GenerateDatapageRequest) (*GenerateDatapageResponse, error) {
	// get api key
	apiKey, err := ds.apiKeyRepo.Get(ctx, apiKeyStr) // TODO: hash?
	if err != nil {
		return nil, err
	}
	if apiKey == nil {
		return nil, nil
	}

	// get account
	account, err := ds.accountRepo.Get(ctx, apiKey.AccountID)
	if err != nil {
		return nil, err
	}
	if !isActiveKey(apiKey, account) {
		return nil, nil
	}

	rules := []*OutageRule{}
	for _, app := range account.Applications {
		if app.ID != request.Application {
			continue
		}

		for _, env := range app.Environments {
			if env.ID != request.Environment || !env.Enabled {
				continue
			}

			for _, rule := range app.Rules {
				if rule.Enabled {
					switch rule.Type {
					case "send-http":
						rules = append(rules, &OutageRule{
							Type: "http-client-request.v1",
							HttpClientRequestV1: &HttpClientRequestV1{
								Host:     rule.Host,
								Status:   rule.Status,
								Duration: rule.Duration,
							},
						})
					default:
						return nil, fmt.Errorf("failed to map outage rule type: %v", rule.Type)
					}
				}
			}
			break
		}
		break
	}

	response := &GenerateDatapageResponse{
		OutageRules: rules,
	}

	return response, nil
}

func (ds *datapageService) GenerateDatapageDeprecated(ctx context.Context, apiKeyStr string, request *GenerateDatapageRequest) (*models.Account, error) {
	// get api key
	apiKey, err := ds.apiKeyRepo.Get(ctx, apiKeyStr) // TODO: hash?
	if err != nil {
		return nil, err
	}
	if apiKey == nil {
		return nil, nil
	}

	// get account
	account, err := ds.accountRepo.Get(ctx, apiKey.AccountID)
	if err != nil {
		return nil, err
	}
	if !isActiveKey(apiKey, account) {
		return nil, nil
	}

	return account, nil
}

func isActiveKey(apiKey *models.ApiKey, account *models.Account) bool {
	for _, key := range account.ApiKeys {
		if key.Key == apiKey.ApiKey {
			return true
		}
	}

	return false
}
