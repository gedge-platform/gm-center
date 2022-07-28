package model


type CredentialCount struct {
	CredentialNames []CredentialNameCount `json:"credential"`
}

type CredentialNameCount struct {
	CredentialName       string             `json:"CredentialName" `
}