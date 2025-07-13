helm-install:
	helm install mercury ./helm-mercury --namespace mercury

helm-uninstall:
	helm uninstall mercury --namespace mercury

helm-upgrade:
	helm upgrade mercury ./helm-mercury --namespace mercury

helm-rollback:
	helm rollback mercury 1 --namespace mercury