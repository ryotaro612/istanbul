cd rootfs & mkdir put_old
# 新しいマウントポイントの名前空間に移動
unshare -mpfr /bin/sh
mount --bind $(pwd) $(pwd)
# put_oldにもとのルートをおく。後でunmount可
pivot_root $(pwd) $(pwd)/put_old
