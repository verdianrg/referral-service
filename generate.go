package referralservice

//go:generate swagger generate server --exclude-main -A referral-server -t gen -f ./api/swagger.yml --principal models.Principal
//go:generate swagger -q generate client -A referral-service -f api/swagger.yml -c pkg/client -m gen/models --principal models.Principal
