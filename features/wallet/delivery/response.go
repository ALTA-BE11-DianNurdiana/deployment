package delivery

import "be11/apiclean/features/wallet"

type WalletResponse struct {
	Jenis  string       `json:"jenis" form:"jenis"`
	Nomor  string       `json:"nomor" form:"nomor"`
	UserID uint         `json:"user_id" form:"user_id"`
	User   UserResponse `json:"user"`
}

type UserResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email,omitempty"`
	Address string `json:"address,omitempty"`
}

func fromCore(data wallet.Core) WalletResponse {
	return WalletResponse{
		Jenis:  data.Jenis,
		Nomor:  data.Nomor,
		UserID: data.UserID,
		User: UserResponse{
			Name: data.User.Name,
		},
	}
}

func fromCoreList(data []wallet.Core) []WalletResponse {
	var dataResponse []WalletResponse
	for _, v := range data {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
