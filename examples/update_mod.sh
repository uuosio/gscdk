for DIR in callcpp counter database1 database3 hello notification token verifysignature callc codedebugging customtarget database2 finalize inlineaction tablestruct transaction
do
    pushd $DIR
    go get github.com/uuosio/chain
    popd
done
