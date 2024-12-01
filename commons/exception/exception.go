package exception

import "errors"

var (
	ErrInternalServerError   = errors.New("error - 500 Internal Server Error - please try again later")
	ErrBadRequest            = errors.New("error - 400 bad Request - please try again later")
	ErrUnauthorized          = errors.New("error - 401 ErrUnauthorized")
	ErrAccountNotFound       = errors.New("Akun dengan email dan password yang Anda input tidak ditemukan")
	ErrExpToken              = errors.New("Token tidak ditemukan atau sudah kadaluarsa")
	ErrTokenRequired         = errors.New("The token is required in the request header and cannot be empty.")
	ErrDataNotFound          = errors.New("Data tidak ditemukan.")
	ErrAccountNonActived     = errors.New("Akun Anda saat ini sedang nonaktif.")
	ErrEmailAlreadyExists    = errors.New("email sudah digunakan.")
	ErrNitsnAlreadyExists    = errors.New("NTISN sudah digunakan.")
	ErrNikAlreadyExists      = errors.New("NIK sudah digunakan.")
	ErrOrderRefAlreadyExists = errors.New("Order reference sudah digunakan.")
	ErrHkTrx                 = errors.New("jumlah transaksi tidak sesuai dengan total bayar")
	ErrStockNotAvailable     = errors.New("jumlah stock yang diminta tidak mencukupi")
	ErrMinQTY                = errors.New("quantitiy tidak boleh 0 (nol)")
	Success                  = ("success")
)
