package cmd


type ProxyIn interface {
	Create(name, dni, card string) map[string]string
	Cancel() map[string]string
}



func NewProxy() ProxyIn {
	return &proxy{}
}