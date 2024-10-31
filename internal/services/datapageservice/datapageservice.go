package datapageservice

import (
	"context"
	"fmt"

	"github.com/outagelab/outagelab/internal/models"
)

type datapageService struct {
	accountService accountService
	apiKeyRepo     apiKeyRepo
}

func New(accountService accountService, apiKeyRepo apiKeyRepo) *datapageService {
	return &datapageService{
		accountService: accountService,
		apiKeyRepo:     apiKeyRepo,
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
	account, err := ds.accountService.GetAccount(ctx, apiKey.AccountID)
	if err != nil {
		return nil, err
	}
	if !isActiveKey(apiKey, account) {
		return nil, nil
	}

	ds.registerDiscoveredComponents(ctx, account, request)

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
	account, err := ds.accountService.GetAccount(ctx, apiKey.AccountID)
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

func (ds *datapageService) registerDiscoveredComponents(
	ctx context.Context,
	account *models.Account,
	req *GenerateDatapageRequest,
) (*models.Account, error) {
	var needsUpdate bool

	// check application
	var app *models.Application
	for _, a := range account.Applications {
		if a.ID == req.Application {
			app = a
			break
		}
	}

	if app == nil {
		app = &models.Application{
			ID:           req.Application,
			Environments: []*models.Environment{},
			Rules:        []*models.OutageRule{},
		}
		account.Applications = append(account.Applications, app)
		needsUpdate = true
	}

	// check environment
	var env *models.Environment
	for _, e := range app.Environments {
		if e.ID == req.Environment {
			env = e
			break
		}
	}

	if env == nil {
		env = &models.Environment{
			ID:      req.Environment,
			Enabled: false,
		}
		app.Environments = append(app.Environments, env)
		needsUpdate = true
	}

	// update if needed
	if !needsUpdate {
		return account, nil
	}

	return ds.accountService.UpdateAccount(ctx, account.ID, account)
}
