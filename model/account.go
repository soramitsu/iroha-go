package model

type Account struct {
	Username string
	Signatories []Signature
	Quorum byte
	TargetDomain string
}
