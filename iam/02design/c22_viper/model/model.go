package model

type Config struct {
	Host      Host      `json:"host"`
	Datastore Datastore `json:"datastore"`
}
type Host struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}
type Metric struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}
type Warehouse struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}
type Datastore struct {
	Metric    Metric    `json:"metric"`
	Warehouse Warehouse `json:"warehouse"`
}
