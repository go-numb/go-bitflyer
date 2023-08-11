package bitflyer

import (
	"github.com/rluisr/go-bitflyer/private"
)

func (p *Client) Permissions(req *private.Permissions) ([]string, *APILimit, error) {
	res := new(private.ResponseForPermissions)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}

	return []string(*res), apiLimit, nil
}

func (p *Client) Balance(req *private.Balance) ([]private.Currency, *APILimit, error) {
	res := new(private.ResponseForBalance)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}

	return []private.Currency(*res), apiLimit, nil
}

func (p *Client) Collateral(req *private.Collateral) (*private.ResponseForCollateral, *APILimit, error) {
	res := new(private.ResponseForCollateral)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}

	return res, apiLimit, nil
}

func (p *Client) CollateralAccounts(req *private.CollateralAccounts) ([]private.Currency, *APILimit, error) {
	res := new(private.ResponseForCollateralAccounts)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}

	return []private.Currency(*res), apiLimit, nil
}

func (p *Client) Addresses(req *private.Addresses) ([]private.Address, *APILimit, error) {
	res := new(private.ResponseForAddresses)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}

	return []private.Address(*res), apiLimit, nil
}

func (p *Client) Coinins(req *private.Coinins) ([]private.Coinin, *APILimit, error) {
	res := new(private.ResponseForCoinins)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}

	return []private.Coinin(*res), apiLimit, nil
}

func (p *Client) Coinouts(req *private.Coinouts) ([]private.Coinout, *APILimit, error) {
	res := new(private.ResponseForCoinouts)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}

	return []private.Coinout(*res), apiLimit, nil
}

func (p *Client) Banks(req *private.Banks) ([]private.Bank, *APILimit, error) {
	res := new(private.ResponseForBanks)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}

	return []private.Bank(*res), apiLimit, nil
}

func (p *Client) Deposits(req *private.Deposits) ([]private.Deposit, *APILimit, error) {
	res := new(private.ResponseForDeposits)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}

	return []private.Deposit(*res), apiLimit, nil
}

func (p *Client) Withdrawals(req *private.Withdrawals) ([]private.Withdrawal, *APILimit, error) {
	res := new(private.ResponseForWithdrawals)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}

	return []private.Withdrawal(*res), apiLimit, nil
}

func (p *Client) ChildOrder(req *private.ChildOrder) (*private.ResponseForChildOrder, *APILimit, error) {
	res := new(private.ResponseForChildOrder)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}

	return res, apiLimit, nil
}

func (p *Client) CancelChildOrder(req *private.CancelChildOrder) (*private.ResponseForCancelChildOrder, *APILimit, error) {
	res := new(private.ResponseForCancelChildOrder)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}

	return res, apiLimit, nil
}

func (p *Client) ParentOrder(req *private.ParentOrder) (*private.ResponseForParentOrder, *APILimit, error) {
	res := new(private.ResponseForParentOrder)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}

	return res, apiLimit, nil
}

func (p *Client) CancelParentOrder(req *private.CancelParentOrder) (*private.ResponseForCancelParentOrder, *APILimit, error) {
	res := new(private.ResponseForCancelParentOrder)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}

	return res, apiLimit, nil
}

func (p *Client) Cancel(req *private.Cancel) (*private.ResponseForCancel, *APILimit, error) {
	res := new(private.ResponseForCancel)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}

	return res, apiLimit, nil
}

func (p *Client) ChildOrders(req *private.ChildOrders) ([]private.COrder, *APILimit, error) {
	res := new(private.ResponseForChildOrders)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}

	return []private.COrder(*res), apiLimit, nil
}

func (p *Client) DetailParentOrder(req *private.DetailParentOrder) (*private.ResponseForDetailParentOrder, *APILimit, error) {
	res := new(private.ResponseForDetailParentOrder)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}

	return res, apiLimit, nil
}

func (p *Client) MyExecutions(req *private.Executions) ([]private.Execution, *APILimit, error) {
	res := new(private.ResponseForExecutions)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}

	return []private.Execution(*res), apiLimit, nil
}

func (p *Client) BalanceHistory(req *private.BalanceHistory) ([]private.BalanceHis, *APILimit, error) {
	res := new(private.ResponseForBalanceHistory)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}

	return []private.BalanceHis(*res), apiLimit, nil
}

func (p *Client) Positions(req *private.Positions) ([]private.Position, *APILimit, error) {
	res := new(private.ResponseForPositions)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}

	return []private.Position(*res), apiLimit, nil
}

func (p *Client) CollateralHistory(req *private.CollateralHistory) ([]private.CollateralHis, *APILimit, error) {
	res := new(private.ResponseForCollateralHistory)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}

	return []private.CollateralHis(*res), apiLimit, nil
}

func (p *Client) Commission(req *private.Commission) (*private.ResponseForCommission, *APILimit, error) {
	res := new(private.ResponseForCommission)
	apiLimit, err := p.request(req, res)
	if err != nil {
		return nil, apiLimit, err
	}

	return res, apiLimit, nil
}
