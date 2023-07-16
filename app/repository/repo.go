package app

import model "github.com/xvbnm48/go-credit/app/model"

type (
	UserRepo interface {
		FindByID(userID string) (*model.UserAccount, error)
		FindMerchantByID(merchantID string) (*model.UserAccount, error)
		Save(account *model.UserAccount) error
	}

	WalletRepo interface {
		FindByID(walletID string) (*model.Wallet, error)
		FindByUserID(userID string) (*model.Wallet, error)
		Save(wallet *model.Wallet) error
	}
	TransactionRepo interface {
		FindByID(transactionID string) (*model.Transaction, error)
		Save(txn *model.Transaction) error
	}
)
