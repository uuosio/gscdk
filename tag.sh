VERSION=v0.2.0
REMOTE=eosio
# git push $REMOTE :refs/tags/$VERSION
git tag -d $VERSION
git tag $VERSION -F release.txt
git push -f $REMOTE $VERSION

