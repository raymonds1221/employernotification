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
 -var "config_map_name=employernotificationliveapi-config" \
 -var "rolename=employernotificationliveapirole" \
 -var "policyname=employernotificationliveapipolicy" \
 -var "namespace=jx-live" \
 -var "serviceaccount=employernotificationliveapi-vault" \
 -var "clusterrolebindingname=role-employernotificationliveapi-binding" \
 -var "sa_jwt_token=$TF_VAR_sa_jwt_token" \
 -var "sa_ca_crt=$TF_VAR_sa_ca_crt" \
 -var "k8s_host=$TF_VAR_k8s_host" \
 -var "config_context_auth_info=clusterUser_Ubidy.IT.Kubernetes.AustraliaEast.Production_ubidy-kube-prod" \
 -var "config_context_cluster=ubidy-kube-prod" \
 -var "vault_mount_path=mssqllive" \
 -var "vault_kv_mount_path=kvlive"
