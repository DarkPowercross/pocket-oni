package view

type View struct {
	Header  HeaderView
	Content Content
}

type LayoutProperties struct {
	Title   string
	Border  bool
	Install func(p LayoutProperties)
}
