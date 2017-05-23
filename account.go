package iroha

import (
	"github.com/soramitsu/iroha-go/model"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func (c *Client) AddAccount(ctx context.Context, account model.Account, opts ...grpc.CallOption) error {
	client, err := c.getClient()
	if err != nil {
		return err
	}
	// TODO: Account →　FBS
	res, err := client.Torii(ctx, nil, opts...)
	c.Logger.Printf("[INFO]\tAddAccount\tMessage\t%s", res.Message())

	return err
}

func (c *Client) GetAccounts(ctx context.Context, query model.AccountsQuery) ([]model.Account, error) {
	return []model.Account{}, nil
}

func (c *Client) GetAccount(ctx context.Context, query model.AccountQuery) (model.Account, error) {
	var err error
	a := model.Account{}
	if query.UUID != "" {
		a, err = c.getAccountByUUID(ctx, query.UUID)
	} else if query.UserName != "" && query.Domain != "" {
		a, err = c.getAccountByUserName(ctx, query.UserName, query.Domain)
	} else {
		a, err = c.getAccountByUserNameFromDefaultDomain(ctx, query.UserName)
	}

	return a, err
}

func (c *Client) getAccountByUUID(ctx context.Context, uuid string) (model.Account, error) {
	return model.Account{}, nil
}
func (c *Client) getAccountByUserName(ctx context.Context, userName string, domain string) (model.Account, error) {
	return model.Account{}, nil
}
func (c *Client) getAccountByUserNameFromDefaultDomain(ctx context.Context, userName string) (model.Account, error) {
	return model.Account{}, nil
}

func (c *Client) UpdateAccountQuorum(ctx context.Context, account model.Account) error {
	return nil
}
