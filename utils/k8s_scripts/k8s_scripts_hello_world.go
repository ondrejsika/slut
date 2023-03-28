package k8s_scripts

import "strconv"

func InstallHelloWorld(host string, replicas int, namespace string, dry bool) {
	sh(`helm upgrade --install \
		hello-world hello-world \
	--repo https://helm.sikalabs.io \
	--create-namespace \
	--namespace `+namespace+` \
	--set host=`+host+` \
	--set replicas=`+strconv.Itoa(replicas)+` \
	--wait`, dry)
}
