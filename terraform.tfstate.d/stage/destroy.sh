source variables.sh

read -p 'Username: ' username
read -sp 'Password: ' password
echo

terraform destroy \
 -var "employerdbname=UbidyServicesEmployersDatabase" \
 -var "agencydbname=UbidyServicesAgenciesDatabase" \
 -var "auctiondbname=UbidyServicesAuctionsDatabase" \
 -var "engagementdbname=UbidyServicesEngagementDatabase" \
 -var "username=$username" \
 -var "password=$password" \
 -var "host=ubidyaustraliaeastprod.database.windows.net" \
 -var "port=1433" \
 -var "config_map_name=employernotificationstageapi-config" \
 -var "rolename=employernotificationstageapirole" \
 -var "policyname=employernotificationstageapipolicy" \
 -var "namespace=jx-stage" \
 -var "serviceaccount=employernotificationstageapi-vault" \
 -var "clusterrolebindingname=role-employernotificationstageapi-binding" \
 -var "sa_jwt_token=$TF_VAR_sa_jwt_token" \
 -var "sa_ca_crt=$TF_VAR_sa_ca_crt" \
 -var "k8s_host=$TF_VAR_k8s_host" \
 -var "config_context_auth_info=clusterUser_Ubidy.IT.Kubernetes.AustraliaEast.Production_ubidy-kube-prod" \
 -var "config_context_cluster=ubidy-kube-prod" \
 -var "vault_mount_path=mssqlstage" \
 -var "vault_kv_mount_path=kvstage"
