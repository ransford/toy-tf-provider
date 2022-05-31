package provider

type FooClient struct {
	hostport  string
	accessKey string

	beep string
}

func NewClient(hostport, accessKey string) *FooClient {
	return &FooClient{
		hostport:  hostport,
		accessKey: accessKey,
	}
}

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
