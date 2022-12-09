package entity

type Product struct { //NOME DE VARIAVEL TEM Q SER MAISUCULA POR CAUSA DO JSON
	ID        uint32 `json:"id"`
	Name      string `json:"name_prod"`
	Price     string `json:"price_prod"`
	Code      string `json:"code_prod"`
	CreatedAt string `json:"created_at,omitempty"`
}
