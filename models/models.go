package models

type Network struct {
	Id                           int `json:"id"`
	Hostname, IPAddress, Gateway string
}

type Storage struct {
	Id         int `json:"id"`
	Name, Type string
	Size       float64
}

type VM struct {
	Id        int    `json:"id"`
	Name      string `json:"nome"`
	Vcpu      int
	Mem       float64
	StorageId uint    // Chave estrangeira que se relaciona com a tabela User
	Storage   Storage `gorm:"foreignKey:StorageId"` // Relacionamento com a entidade Storage
	NetworkId uint
	Network   Network `gorm:"foreignKey:NetworkId"` // Relacionamento com a entidade Storage
}

var Vms []VM
