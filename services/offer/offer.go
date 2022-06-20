package offer

import (
	"errors"
	"strconv"
	"trading/services/db"
)

func ProcessOrder(id string, quantity float64, buyerId string) (err error) {
	offer := db.OfferById(id)
	if offer.Completed {
		return errors.New("offer already completed")
	}
	if offer.Quantity < quantity {
		return errors.New("offer is not of requested quantity")
	}
	total := quantity * offer.Value
	updateUsersAccounts(buyerId, offer, total, quantity)

	offer.Quantity = offer.Quantity - quantity
	if offer.Quantity <= 0 {
		offer.Completed = true
	}
	db.UpdateOffer(offer)
	return nil
}

func updateUsersAccounts(buyerId string, offer db.Offer, total float64, quantity float64) {
	updateBuyer(buyerId, total, offer.GoodID, quantity)
	updateSeller(offer, total, offer.GoodID, quantity)
}

func updateSeller(offer db.Offer, total float64, goodId uint, quantity float64) {
	userSeller := db.GetUserById(strconv.Itoa(int(offer.UserID)))
	newCredit := *userSeller.Credit + total
	userSeller.Credit = &newCredit
	db.UpdateUser(userSeller)
	db.AddToAccount(-quantity, goodId, offer.UserID)
}

func updateBuyer(buyerId string, total float64, goodId uint, quantity float64) {
	userBuyer := db.GetUserById(buyerId)
	newCredit := *userBuyer.Credit - total
	userBuyer.Credit = &newCredit
	db.UpdateUser(userBuyer)
	db.AddToAccount(quantity, goodId, userBuyer.ID)
}
