package transaction

import "wardfunding/user"

type GetCampaignTransactionsInput struct {
	ID 		int `uri:"id" binding:"required"`
	User	user.User
}