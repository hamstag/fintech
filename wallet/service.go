package wallet

import (
	"fmt"
	"time"

	"github.com/hamstag/fintech/core/db"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"go.uber.org/zap"
)

type (
	WalletService interface {
		RegisterCustomer(cmd *RegisterCustomer) (string, error)
		RegisterMerchant(cmd *RegisterMerchant) (string, error)
		Topup(cmd *TopUp) (string, error)
		Pay(cmd *Payment) (string, error)
	}

	WalletServiceImpl struct {
		userRepo        UserRepo
		walletRepo      WalletRepo
		transactionRepo TransactionRepo

		log *zap.Logger
	}
)

func NewWalletService(db db.Database, log *zap.Logger) *WalletServiceImpl {
	return &WalletServiceImpl{
		userRepo:        &UserRepoImpl{db: db},
		walletRepo:      &WalletRepoImpl{db: db},
		transactionRepo: &TransactionRepoImpl{db: db},

		log: log,
	}
}

func (s *WalletServiceImpl) RegisterCustomer(cmd *RegisterCustomer) (string, error) {
	userID := gonanoid.Must()

	account := UserAccount{
		UserID:      userID,
		Name:        cmd.Name,
		Email:       cmd.Email,
		Phonenumber: cmd.Phonenumber,
		UserType:    CUSTOMER,
	}

	return s.createWallet(&account)
}

func (s *WalletServiceImpl) RegisterMerchant(cmd *RegisterMerchant) (string, error) {
	userID := gonanoid.Must()

	account := UserAccount{
		UserID:   userID,
		Name:     cmd.Name,
		Email:    cmd.Email,
		UserType: MERCHANT,
	}

	return s.createWallet(&account)
}

func (s *WalletServiceImpl) Topup(cmd *TopUp) (string, error) {
	wallet, err := s.walletRepo.FindByID(cmd.WalletID)
	if err != nil {
		s.log.Error(fmt.Sprintf("error topup on find wallet %s \n", err.Error()))
		return "", err
	}

	txnID := gonanoid.Must()
	txn := Transaction{
		TransactionID:   txnID,
		ReferenceID:     cmd.ReferenceID,
		CreditWallet:    cmd.WalletID,
		Description:     cmd.Description,
		Amount:          cmd.Amount,
		TransactionDate: time.Now(),
		TransactionType: TXN_TOPUP,
	}

	if err := s.transactionRepo.Save(&txn); err != nil {
		s.log.Error(fmt.Sprintf("error topup %s \n", err.Error()))
		return "", err
	}

	wallet.creditBalance(cmd.Amount)
	if err := s.walletRepo.Save(wallet); err != nil {
		s.log.Error(fmt.Sprintf("error topup %s \n", err.Error()))
		return "", err
	}

	return txnID, nil
}

func (s *WalletServiceImpl) Pay(cmd *Payment) (string, error) {
	merchantWallet, err := s.findMerchantWallet(cmd.Merchant)
	if err != nil {
		s.log.Error(fmt.Sprintf("error Pay cannot find merchant %s \n", err.Error()))
		return "", err
	}

	customerWallet, err := s.walletRepo.FindByID(cmd.WalletID)
	if err != nil {
		s.log.Error(fmt.Sprintf("error Pay cannot find wallet %s \n", err.Error()))
		return "", err
	}

	txnID := gonanoid.Must()
	txn := Transaction{
		TransactionID:   txnID,
		ReferenceID:     cmd.ReferenceID,
		CreditWallet:    merchantWallet.WalletID,
		DebitedWallet:   cmd.WalletID,
		Description:     cmd.Description,
		Amount:          cmd.Amount,
		TransactionDate: time.Now(),
		TransactionType: TXN_PAYMENT,
	}

	if err := s.transactionRepo.Save(&txn); err != nil {
		s.log.Error(fmt.Sprintf("error Pay %s \n", err.Error()))
		return "", err
	}

	if err := merchantWallet.creditBalance(cmd.Amount); err != nil {
		s.log.Error(fmt.Sprintf("error Pay %s \n", err.Error()))
		return "", err
	}

	if err := s.walletRepo.Save(merchantWallet); err != nil {
		s.log.Error(fmt.Sprintf("error Pay %s \n", err.Error()))
		return "", err
	}

	if err := customerWallet.debitBalance(cmd.Amount); err != nil {
		s.log.Error(fmt.Sprintf("error Pay %s \n", err.Error()))
		return "", err
	}

	if err := s.walletRepo.Save(customerWallet); err != nil {
		s.log.Error(fmt.Sprintf("error Pay %s \n", err.Error()))
		return "", err
	}

	return txnID, nil
}

func (s *WalletServiceImpl) createWallet(account *UserAccount) (string, error) {
	if err := s.userRepo.Save(account); err != nil {
		s.log.Error(fmt.Sprintf("error creating account %s \n", err.Error()))
		return "", err
	}

	walletID := gonanoid.Must()
	wallet := Wallet{
		WalletID: walletID,
		UserID:   account.UserID,
	}

	if err := s.walletRepo.Save(&wallet); err != nil {
		s.log.Error(fmt.Sprintf("error creating wallet %s \n", err.Error()))
		return "", err
	}

	return walletID, nil
}

func (s *WalletServiceImpl) findMerchantWallet(name string) (*Wallet, error) {
	merchant, err := s.userRepo.FindMerchantAccountByName(name)
	if err != nil {
		s.log.Error(fmt.Sprintf("error cannot find merchant %s \n", err.Error()))
		return nil, err
	}

	return s.walletRepo.FindByUserID(merchant.UserID)
}
