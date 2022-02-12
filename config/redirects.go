package config

var LoginRedirect = RedirectAction {
	Url: "/",
	Args: []string{"login"},
}
