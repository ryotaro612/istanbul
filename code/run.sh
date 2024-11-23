mkdir rootfs
# busyboxの/以下をrootfsにコピー
docker export $(docker create busybox) | tar -C rootfs -xvf -
# コンテナを起動するパラメタの雛形を./config.jsonに出力
runc spec
# コンテナIDをcidとしてコンテナを起動
runc --root /tmp/runc run cid
# /bin/shを開始し、/が出力される

