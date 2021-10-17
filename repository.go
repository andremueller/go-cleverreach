package main

type Repository interface {
	Groups() []Group
	Mailings() Mailings
	Mailing(id int) Mailing
	AddMailing(Mailing)
}
