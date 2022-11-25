#!/bin/sh

# docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
#     -i https://raw.githubusercontent.com/openapitools/openapi-generator/master/modules/openapi-generator/src/test/resources/3_0/petstore.yaml \
#     -g go \
#     -o /local/out/go

git_user_id=0x00b
git_repo_id=wx-openapi
generator=openapitools/openapi-generator-cli

if [ $# -lt 2 ] ; then
echo "USAGE: $0 relative_in_dir relative_out_dir agoserver/agoclient/generator"
echo " e.g.: $0 . ./tmp" 
echo "       $0 api interfaces OR $0 api interfaces agoserver"
echo "       $0 api api agoclient"
echo "       $0 api api/js javascript"
exit 1;
fi

work_path=$(dirname $(readlink -f $0))


for file in $(ls $work_path/$1); do
  file=`basename $file`
  if [ "${file##*.}"x = "yaml"x ] || [ "${file##*.}"x = "yml"x ] || [ "${file##*.}"x = "json"x ]; then

    fileName=$(echo $file | cut -d . -f1)

    # 清除原来的文件
    rm -rf $work_path/$2/$fileName


    if [ $# == 2 ] || [ $3 == 'agoserver' ] ; then
      docker run -v $work_path:/local $generator \
      java -Xmx1024M -DloggerPath=conf/log4j.properties -DdebugOperations -DdebugModels -jar /openapi-generator-cli.jar \
      generate --package-name $fileName -i /local/$1/$file -o /local/$2/$fileName -t /go-ago-server -g go-ago-server \
      --git-user-id $git_user_id --git-repo-id $git_repo_id
    elif [ $3 == 'agoclient' ]; then
      docker run -v $work_path:/local $generator \
      java -Xmx1024M -DloggerPath=conf/log4j.properties -DdebugOperations -DdebugModels -jar /openapi-generator-cli.jar \
      generate --package-name $fileName -i /local/$1/$file -o /local/$2/$fileName -t /go-ago-client -g go-ago-client \
      --git-user-id $git_user_id --git-repo-id $git_repo_id
    else
      docker run -v $work_path:/local $generator \
      generate --package-name $fileName -i /local/$1/$file -o /local/$2/$fileName  -g $3 \
      --git-user-id $git_user_id --git-repo-id $git_repo_id
    fi

    rm $work_path/$2/$fileName/go.mod
    rm $work_path/$2/$fileName/go.sum
    rm $work_path/$2/$fileName/test -rf
  fi
done

cd $work_path/$2 && goimports -w ./ && go fmt ./...
cd -
