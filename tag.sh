VERSION=v0.1.5
REMOTE=eosio
git push $REMOTE :refs/tags/$VERSION
git tag -d $VERSION
git tag $VERSION -F release.txt
git push -f $REMOTE $VERSION

