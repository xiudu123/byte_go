#!/bin/bash

# 并行启动所有服务
/app/frontend &
/app/user &
/app/auth &
/app/product &
/app/cart &
/app/order &
/app/checkout &
/app/payment &

# 保持容器运行
wait