HELM_URL=https://storage.googleapis.com/kubernetes-helm
HELM_TARBALL=helm-v2.4.1-linux-amd64.tar.gz
curl -O $HELM_URL/$HELM_TARBALL
tar xzfv $HELM_TARBALL -C /home/jenkins && rm $HELM_TARBALL
export PATH=/home/jenkins/linux-amd64/:$PATH
helm init --client-only
