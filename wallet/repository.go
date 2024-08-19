package wallet

import "github.com/hamstag/fintech/core/db"

type (
	UserRepo interface {
		FindByID(userID string) (*UserAccount, error)
		FindMerchantAccountByName(name string) (*UserAccount, error)
		Save(account *UserAccount) error
	}

	WalletRepo interface {
		FindByID(walletID string) (*Wallet, error)
		FindByUserID(userID string) (*Wallet, error)
		Save(wallet *Wallet) error
	}

	TransactionRepo interface {
		FindByID(txnID string) (*Transaction, error)
		Save(txn *Transaction) error
	}

	UserRepoImpl struct {
		db *db.Database
	}
	WalletRepoImpl struct {
		db *db.Database
	}
	TransactionRepoImpl struct {
		db *db.Database
	}
)

func (r *UserRepoImpl) FindByID(userID string) (*UserAccount, error) {
	panic("not implemented") // TODO: Implement
}

func (r *UserRepoImpl) FindMerchantAccountByName(name string) (*UserAccount, error) {
	panic("not implemented") // TODO: Implement
}

func (r *UserRepoImpl) Save(account *UserAccount) error {
	panic("not implemented") // TODO: Implement
}

func (r *WalletRepoImpl) FindByID(walletID string) (*Wallet, error) {
	panic("not implemented") // TODO: Implement
}

func (r *WalletRepoImpl) FindByUserID(userID string) (*Wallet, error) {
	panic("not implemented") // TODO: Implement
}

func (r *WalletRepoImpl) Save(wallet *Wallet) error {
	panic("not implemented") // TODO: Implement
}

func (r *TransactionRepoImpl) FindByID(txnID string) (*Transaction, error) {
	panic("not implemented") // TODO: Implement
}

func (r *TransactionRepoImpl) Save(txn *Transaction) error {
	panic("not implemented") // TODO: Implement
}
