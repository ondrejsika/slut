package install_bin_tool

var Tools = []Tool{
	{
		Name: "kubectl",
		// Get Version: `curl -L -s https://dl.k8s.io/release/stable.txt`
		Version:     "v1.22.1",
		UrlTemplate: "https://dl.k8s.io/release/{{.Version}}/bin/{{.Os}}/{{.Arch}}/kubectl",
	},
	{
		Name:        "terraform",
		Version:     "1.0.5",
		UrlTemplate: hashicorpUrlTemplate("terraform"),
	},
	{
		Name:        "vault",
		Version:     "1.8.2",
		UrlTemplate: hashicorpUrlTemplate("vault"),
	},
}

func hashicorpUrlTemplate(name string) string {
	return "https://releases.hashicorp.com/" + name + "/{{.Version}}/" + name +
		"_{{.Version}}_{{.Os}}_{{.Arch}}.zip"
}
