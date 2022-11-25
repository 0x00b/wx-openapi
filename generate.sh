#!/bin/sh
 
sh ./generate_.sh api api/client go

#
# //Body 请求包体，手动替换的，yaml暂不支持生成这种字段
# Data  --> Body []byte `in:"body" json:"body"`
# sh ./generate_.sh api/callback api/server

 