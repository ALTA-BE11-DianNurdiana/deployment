package delivery

import "be11/apiclean/features/wallet"

type WalletRequest struct {
	Jenis string `json:"jenis" form:"jenis"`
	Nomor string `json:"nomor" form:"nomor"`
}

func toCore(data WalletRequest) wallet.Core {
	return wallet.Core{
		Jenis: data.Jenis,
		Nomor: data.Nomor,
	}
}
