# Installing on Kubernetes

Note: if running on GKE, you need to configure permission to create rbac: 
```bash
kubectl create clusterrolebinding --user <gcloud-email> <crb-name> --clusterrole=<any role with RBAC create permission>
```