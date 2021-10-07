VERSION=v0.1.3
# git push origin :refs/tags/$VERSION
git tag -d $VERSION
git tag $VERSION -F release.txt
git push origin $VERSION

