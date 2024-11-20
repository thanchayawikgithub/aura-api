package test

type ServiceTestCase[Args any, Want any] struct {
	Name    string
	Mock    func()
	Args    Args
	Want    Want
	WantErr bool
	Err     error
}
