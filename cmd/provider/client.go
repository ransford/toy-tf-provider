package provider

type FooClient struct {
	accessKey string

	beep string
}

// TODO NewClient that takes connection info

func (f *FooClient) CreateFoo() string {
	// Doesn't actually create anything
	f.beep = "twelve"
	return "newfoo"
}

func (f *FooClient) GetFoo(id string) string {
	return f.beep // TODO replace with JSON fetch
}

func (f *FooClient) SetFoo(id string, newval string) error {
	f.beep = newval
	return nil
}
