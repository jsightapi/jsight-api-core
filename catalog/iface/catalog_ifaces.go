package iface

type Catalog interface {
	GetInfo() Info
}

type Info interface {
	GetTitle() string
}
