package basic

import "core-banking-server/internal/database"

type registerBlockchainService struct{}

func GetNewRegisterBlockchainServiceController() *registerBlockchainService {
	return &registerBlockchainService{}
}

func (c *registerBlockchainService) SaveCustomerPublicKey(customerID, customerPublicKey string) (err error) {

	return c.saveCustomerPublicKey(customerID, customerPublicKey)
}

func (c *registerBlockchainService) saveCustomerPublicKey(customerID, customerPublicKey string) (err error) {
	db, err := database.GetDBConnection()
	if err != nil {
		return
	}
	_, err = db.GetcustomerByID(customerID)
	if err != nil {
		return
	}

	return db.AddCustomerPublicEncryptionKey(customerID, customerPublicKey)
}
