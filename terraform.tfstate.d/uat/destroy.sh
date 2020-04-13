source variables.sh

read -p 'Username: ' username
read -sp 'Password: ' password
echo
read -p 'MongoDB Username: ' mongodb_username
read -sp 'MongoDB Password: ' mongodb_password
echo

terraform destroy \
 -var "employerdbname=UbidyServicesEmployersDatabase" \
 -var "agencydbname=UbidyServicesAgenciesDatabase" \
 -var "auctiondbname=UbidyServicesAuctionsDatabase" \
 -var "engagementdbname=UbidyServicesEngagementDatabase" \
 -var "username=$username" \
 -var "password=$password" \
 -var "host=210.4.126.35" \
 -var "port=55107" \
 -var "config_map_name=employernotificationapi-config" \
 -var "rolename=employernotificationapirole" \
 -var "policyname=employernotificationapipolicy" \
 -var "namespace=jx-uat" \
 -var "serviceaccount=employernotificationapi-vault" \
 -var "clusterrolebindingname=role-employernotificationapi-binding" \
 -var "sa_jwt_token=$TF_VAR_sa_jwt_token" \
 -var "sa_ca_crt=$TF_VAR_sa_ca_crt" \
 -var "k8s_host=$TF_VAR_k8s_host" \
 -var "config_context_auth_info=clusterUser_Ubidy.IT.Kubernetes.AustraliaEast.Production_ubidy-kube-prod" \
 -var "config_context_cluster=ubidy-kube-prod" \
 -var "vault_mount_path=mssqluat" \
 -var "vault_kv_mount_path=kvuat" \

