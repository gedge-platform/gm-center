package model


type CredentialCount struct {
	CredentialNames []CredentialNameCount `json:"credential"`
}

type CredentialNameCount struct {
	CredentialName       string             `json:"CredentialName" `
}

type VMCount struct {
	VMCount []VMList `json:"vm"`
}

type VMList struct {
	VMList       string             `json:"IId" `
}

type VMStatusCount struct {
	Vmstatus []VMStatus `json:"vmstatus"`
}

type VMStatus struct {
	VMStatus       string             `json:"Vmstatus" `
}