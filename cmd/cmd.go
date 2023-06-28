package cmd

import (
	_ "github.com/sikalabs/slu/cmd/argocd"
	_ "github.com/sikalabs/slu/cmd/argocd/domain"
	_ "github.com/sikalabs/slu/cmd/argocd/initial_password"
	_ "github.com/sikalabs/slu/cmd/argocd/open"
	_ "github.com/sikalabs/slu/cmd/argocd/password_reset"
	_ "github.com/sikalabs/slu/cmd/argocd/port_forward"
	_ "github.com/sikalabs/slu/cmd/argocd/refresh"
	_ "github.com/sikalabs/slu/cmd/argocd/set_image"
	_ "github.com/sikalabs/slu/cmd/argocd/url"
	_ "github.com/sikalabs/slu/cmd/aws"
	_ "github.com/sikalabs/slu/cmd/aws/who_am_i"
	_ "github.com/sikalabs/slu/cmd/awx"
	_ "github.com/sikalabs/slu/cmd/awx/password"
	_ "github.com/sikalabs/slu/cmd/base64"
	_ "github.com/sikalabs/slu/cmd/base64/interactive_decode_clipboard"
	_ "github.com/sikalabs/slu/cmd/chaos_monkey"
	_ "github.com/sikalabs/slu/cmd/chaos_monkey/random_status_code"
	_ "github.com/sikalabs/slu/cmd/check"
	_ "github.com/sikalabs/slu/cmd/check/kubernetes_context"
	_ "github.com/sikalabs/slu/cmd/check/slu_version"
	_ "github.com/sikalabs/slu/cmd/check/version"
	_ "github.com/sikalabs/slu/cmd/cloudflare"
	_ "github.com/sikalabs/slu/cmd/cloudflare/access"
	_ "github.com/sikalabs/slu/cmd/cloudflare/access/service_token/get"
	_ "github.com/sikalabs/slu/cmd/cloudflare/access/service_token/set"
	_ "github.com/sikalabs/slu/cmd/cloudflare/host"
	_ "github.com/sikalabs/slu/cmd/copy_from_cloud"
	_ "github.com/sikalabs/slu/cmd/copy_to_cloud"
	_ "github.com/sikalabs/slu/cmd/ddev"
	_ "github.com/sikalabs/slu/cmd/debug_server"
	_ "github.com/sikalabs/slu/cmd/debug_server/long_response"
	_ "github.com/sikalabs/slu/cmd/debug_server/request_debug"
	_ "github.com/sikalabs/slu/cmd/df"
	_ "github.com/sikalabs/slu/cmd/digitalocean"
	_ "github.com/sikalabs/slu/cmd/digitalocean/all"
	_ "github.com/sikalabs/slu/cmd/digitalocean/all/cleanup"
	_ "github.com/sikalabs/slu/cmd/digitalocean/all/delete"
	_ "github.com/sikalabs/slu/cmd/digitalocean/all/list"
	_ "github.com/sikalabs/slu/cmd/digitalocean/auth"
	_ "github.com/sikalabs/slu/cmd/digitalocean/auth/add"
	_ "github.com/sikalabs/slu/cmd/digitalocean/auth/list"
	_ "github.com/sikalabs/slu/cmd/digitalocean/auth/rm"
	_ "github.com/sikalabs/slu/cmd/digitalocean/auth/use_context"
	_ "github.com/sikalabs/slu/cmd/docker"
	_ "github.com/sikalabs/slu/cmd/docker/create_config"
	_ "github.com/sikalabs/slu/cmd/docker/ping"
	_ "github.com/sikalabs/slu/cmd/docker/registry"
	_ "github.com/sikalabs/slu/cmd/docker/registry/list"
	_ "github.com/sikalabs/slu/cmd/docker/registry/tags"
	_ "github.com/sikalabs/slu/cmd/du"
	_ "github.com/sikalabs/slu/cmd/eck"
	_ "github.com/sikalabs/slu/cmd/eck/elastic_password"
	_ "github.com/sikalabs/slu/cmd/elasticsearch"
	_ "github.com/sikalabs/slu/cmd/elasticsearch/get"
	_ "github.com/sikalabs/slu/cmd/example_server"
	_ "github.com/sikalabs/slu/cmd/expand"
	_ "github.com/sikalabs/slu/cmd/expand/file"
	_ "github.com/sikalabs/slu/cmd/expand/string"
	_ "github.com/sikalabs/slu/cmd/file_templates"
	_ "github.com/sikalabs/slu/cmd/file_templates/editorconfig"
	_ "github.com/sikalabs/slu/cmd/file_templates/gitignore"
	_ "github.com/sikalabs/slu/cmd/file_templates/go_cli_project"
	_ "github.com/sikalabs/slu/cmd/file_templates/incident_response"
	_ "github.com/sikalabs/slu/cmd/file_templates/terraform_project"
	_ "github.com/sikalabs/slu/cmd/file_utils"
	_ "github.com/sikalabs/slu/cmd/file_utils/replace_string"
	_ "github.com/sikalabs/slu/cmd/generate_docs"
	_ "github.com/sikalabs/slu/cmd/generate_files"
	_ "github.com/sikalabs/slu/cmd/generate_files/tree"
	_ "github.com/sikalabs/slu/cmd/git"
	_ "github.com/sikalabs/slu/cmd/git/cleanup"
	_ "github.com/sikalabs/slu/cmd/git/commit"
	_ "github.com/sikalabs/slu/cmd/git/commit/add_charts"
	_ "github.com/sikalabs/slu/cmd/git/commit/editorconfig_and_gitignore"
	_ "github.com/sikalabs/slu/cmd/git/commit/terraform_lock"
	_ "github.com/sikalabs/slu/cmd/git/delete_all_local_branches"
	_ "github.com/sikalabs/slu/cmd/git/get_last_calver"
	_ "github.com/sikalabs/slu/cmd/git/if"
	_ "github.com/sikalabs/slu/cmd/git/if/staged"
	_ "github.com/sikalabs/slu/cmd/git/tag_next_calver"
	_ "github.com/sikalabs/slu/cmd/git/url"
	_ "github.com/sikalabs/slu/cmd/git/url/get"
	_ "github.com/sikalabs/slu/cmd/git/url/open"
	_ "github.com/sikalabs/slu/cmd/git/use_ssh"
	_ "github.com/sikalabs/slu/cmd/github"
	_ "github.com/sikalabs/slu/cmd/github/latest_release"
	_ "github.com/sikalabs/slu/cmd/gitlab"
	_ "github.com/sikalabs/slu/cmd/gitlab/prune_environments"
	_ "github.com/sikalabs/slu/cmd/gitlab_ci"
	_ "github.com/sikalabs/slu/cmd/gitlab_ci/setup_runner"
	_ "github.com/sikalabs/slu/cmd/gitlab_ci/skip_stage"
	_ "github.com/sikalabs/slu/cmd/gitlab_ci/skip_stage/remove_skip"
	_ "github.com/sikalabs/slu/cmd/gitlab_ci/skip_stage/show"
	_ "github.com/sikalabs/slu/cmd/gitlab_ci/skip_stage/skip"
	_ "github.com/sikalabs/slu/cmd/gitlab_ci/terraform_convert_report"
	_ "github.com/sikalabs/slu/cmd/go_code"
	_ "github.com/sikalabs/slu/cmd/go_code/version_bump"
	_ "github.com/sikalabs/slu/cmd/golang"
	_ "github.com/sikalabs/slu/cmd/golang/build_all_platforms"
	_ "github.com/sikalabs/slu/cmd/helm"
	_ "github.com/sikalabs/slu/cmd/helm/version_bump"
	_ "github.com/sikalabs/slu/cmd/host"
	_ "github.com/sikalabs/slu/cmd/install_any_bin"
	_ "github.com/sikalabs/slu/cmd/install_bin"
	_ "github.com/sikalabs/slu/cmd/ip"
	_ "github.com/sikalabs/slu/cmd/ip_local"
	_ "github.com/sikalabs/slu/cmd/k8s"
	_ "github.com/sikalabs/slu/cmd/k8s/create_oidc_user"
	_ "github.com/sikalabs/slu/cmd/k8s/delete_ns"
	_ "github.com/sikalabs/slu/cmd/k8s/delete_pods"
	_ "github.com/sikalabs/slu/cmd/k8s/get"
	_ "github.com/sikalabs/slu/cmd/k8s/get/bad_pods"
	_ "github.com/sikalabs/slu/cmd/k8s/get/configmap"
	_ "github.com/sikalabs/slu/cmd/k8s/get/secret"
	_ "github.com/sikalabs/slu/cmd/k8s/kubeconfig"
	_ "github.com/sikalabs/slu/cmd/k8s/kubeconfig/add"
	_ "github.com/sikalabs/slu/cmd/k8s/lock"
	_ "github.com/sikalabs/slu/cmd/k8s/lock/lock"
	_ "github.com/sikalabs/slu/cmd/k8s/lock/status"
	_ "github.com/sikalabs/slu/cmd/k8s/lock/unlock"
	_ "github.com/sikalabs/slu/cmd/k8s/token"
	_ "github.com/sikalabs/slu/cmd/kdev"
	_ "github.com/sikalabs/slu/cmd/kn"
	_ "github.com/sikalabs/slu/cmd/kv"
	_ "github.com/sikalabs/slu/cmd/kv/get"
	_ "github.com/sikalabs/slu/cmd/kv/set"
	_ "github.com/sikalabs/slu/cmd/kx"
	_ "github.com/sikalabs/slu/cmd/length"
	_ "github.com/sikalabs/slu/cmd/loggen"
	_ "github.com/sikalabs/slu/cmd/login"
	_ "github.com/sikalabs/slu/cmd/mail"
	_ "github.com/sikalabs/slu/cmd/mail/send"
	_ "github.com/sikalabs/slu/cmd/mail/send_vault"
	_ "github.com/sikalabs/slu/cmd/metrics_generator"
	_ "github.com/sikalabs/slu/cmd/metrics_generator/server"
	_ "github.com/sikalabs/slu/cmd/mssql"
	_ "github.com/sikalabs/slu/cmd/mssql/ping"
	_ "github.com/sikalabs/slu/cmd/mysql"
	_ "github.com/sikalabs/slu/cmd/mysql/create"
	_ "github.com/sikalabs/slu/cmd/mysql/drop"
	_ "github.com/sikalabs/slu/cmd/mysql/generate"
	_ "github.com/sikalabs/slu/cmd/mysql/generate/random_data"
	_ "github.com/sikalabs/slu/cmd/mysql/list"
	_ "github.com/sikalabs/slu/cmd/mysql/ping"
	_ "github.com/sikalabs/slu/cmd/ondrejsika"
	_ "github.com/sikalabs/slu/cmd/ondrejsika/clear_dns_cache_mac"
	_ "github.com/sikalabs/slu/cmd/ondrejsika/desktop_cleanup"
	_ "github.com/sikalabs/slu/cmd/ondrejsika/large_desktop_files"
	_ "github.com/sikalabs/slu/cmd/postgres"
	_ "github.com/sikalabs/slu/cmd/postgres/create"
	_ "github.com/sikalabs/slu/cmd/postgres/drop"
	_ "github.com/sikalabs/slu/cmd/postgres/list"
	_ "github.com/sikalabs/slu/cmd/postgres/ping"
	_ "github.com/sikalabs/slu/cmd/proxy"
	_ "github.com/sikalabs/slu/cmd/proxy/smtp"
	_ "github.com/sikalabs/slu/cmd/proxy/tcp"
	_ "github.com/sikalabs/slu/cmd/random"
	_ "github.com/sikalabs/slu/cmd/random/int"
	_ "github.com/sikalabs/slu/cmd/random/password"
	_ "github.com/sikalabs/slu/cmd/random/string"
	_ "github.com/sikalabs/slu/cmd/rke2"
	_ "github.com/sikalabs/slu/cmd/rmline"
	"github.com/sikalabs/slu/cmd/root"
	_ "github.com/sikalabs/slu/cmd/s3"
	_ "github.com/sikalabs/slu/cmd/s3/copy_from_s3"
	_ "github.com/sikalabs/slu/cmd/s3/copy_to_s3"
	_ "github.com/sikalabs/slu/cmd/s3/remove_bucket"
	_ "github.com/sikalabs/slu/cmd/s3/remove_objects_by_age"
	_ "github.com/sikalabs/slu/cmd/scripts"
	_ "github.com/sikalabs/slu/cmd/scripts/atol"
	_ "github.com/sikalabs/slu/cmd/scripts/atol/get_images_from_env_for_values_yaml"
	_ "github.com/sikalabs/slu/cmd/scripts/counter"
	_ "github.com/sikalabs/slu/cmd/scripts/docker_remove_all"
	_ "github.com/sikalabs/slu/cmd/scripts/download"
	_ "github.com/sikalabs/slu/cmd/scripts/gitlab_ci"
	_ "github.com/sikalabs/slu/cmd/scripts/gitlab_ci/update_docker_images"
	_ "github.com/sikalabs/slu/cmd/scripts/gitstats_docker"
	_ "github.com/sikalabs/slu/cmd/scripts/infracost_here"
	_ "github.com/sikalabs/slu/cmd/scripts/kubernetes"
	_ "github.com/sikalabs/slu/cmd/scripts/kubernetes/install_all"
	_ "github.com/sikalabs/slu/cmd/scripts/kubernetes/install_argocd"
	_ "github.com/sikalabs/slu/cmd/scripts/kubernetes/install_cert_manager"
	_ "github.com/sikalabs/slu/cmd/scripts/kubernetes/install_cluster_issuer"
	_ "github.com/sikalabs/slu/cmd/scripts/kubernetes/install_cluster_issuer_cloudflare"
	_ "github.com/sikalabs/slu/cmd/scripts/kubernetes/install_hello_world"
	_ "github.com/sikalabs/slu/cmd/scripts/kubernetes/install_ingress"
	_ "github.com/sikalabs/slu/cmd/scripts/kubernetes/install_maildev"
	_ "github.com/sikalabs/slu/cmd/scripts/kubernetes/install_metrics_server"
	_ "github.com/sikalabs/slu/cmd/scripts/kubernetes/install_prometheus_operator_crd"
	_ "github.com/sikalabs/slu/cmd/scripts/rename"
	_ "github.com/sikalabs/slu/cmd/scripts/rename/spaces_and_commas"
	_ "github.com/sikalabs/slu/cmd/scripts/run_tcp_proxy_in_docker"
	_ "github.com/sikalabs/slu/cmd/serve_files"
	_ "github.com/sikalabs/slu/cmd/shell_scripts"
	_ "github.com/sikalabs/slu/cmd/shell_scripts/required_environment_variables"
	_ "github.com/sikalabs/slu/cmd/sleep"
	_ "github.com/sikalabs/slu/cmd/sleep/forever"
	_ "github.com/sikalabs/slu/cmd/sleep/random"
	_ "github.com/sikalabs/slu/cmd/sqlite"
	_ "github.com/sikalabs/slu/cmd/sqlite/read"
	_ "github.com/sikalabs/slu/cmd/ssh"
	_ "github.com/sikalabs/slu/cmd/ssh/keygen"
	_ "github.com/sikalabs/slu/cmd/static_api"
	_ "github.com/sikalabs/slu/cmd/static_api/version"
	_ "github.com/sikalabs/slu/cmd/systemd"
	_ "github.com/sikalabs/slu/cmd/systemd/create_service"
	_ "github.com/sikalabs/slu/cmd/time"
	_ "github.com/sikalabs/slu/cmd/time/prefix"
	_ "github.com/sikalabs/slu/cmd/time/unix"
	_ "github.com/sikalabs/slu/cmd/time/yyyy_mm_dd_hh_mm_ss"
	_ "github.com/sikalabs/slu/cmd/tls"
	_ "github.com/sikalabs/slu/cmd/tls/parse"
	_ "github.com/sikalabs/slu/cmd/tls/parse_file"
	_ "github.com/sikalabs/slu/cmd/tls/parse_k8s_secret"
	_ "github.com/sikalabs/slu/cmd/tls/pem_to_pfx"
	_ "github.com/sikalabs/slu/cmd/up"
	_ "github.com/sikalabs/slu/cmd/upload"
	_ "github.com/sikalabs/slu/cmd/version"
	_ "github.com/sikalabs/slu/cmd/wait_for"
	_ "github.com/sikalabs/slu/cmd/wait_for/docker"
	_ "github.com/sikalabs/slu/cmd/wait_for/k8s"
	_ "github.com/sikalabs/slu/cmd/wait_for/k8s/job"
	_ "github.com/sikalabs/slu/cmd/wait_for/k8s/pod"
	_ "github.com/sikalabs/slu/cmd/wait_for/tcp"
	_ "github.com/sikalabs/slu/cmd/wait_for/tls"
	_ "github.com/sikalabs/slu/cmd/wait_for_it"
	_ "github.com/sikalabs/slu/cmd/wait_for_tls"
	_ "github.com/sikalabs/slu/cmd/watch"
	_ "github.com/sikalabs/slu/cmd/windows"
	_ "github.com/sikalabs/slu/cmd/windows/scoop_install"
	_ "github.com/sikalabs/slu/cmd/wireguard"
	_ "github.com/sikalabs/slu/cmd/wireguard/genkey"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.RootCmd.Execute())
}
