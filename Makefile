# vdc-session Makefile
#
# DEBUG_MODE = Debug | Release | MinSizeRel | RelWithDebInfo
# 

CURRENT_DIR=$(shell pwd)
BUILD_MODE=Debug
BUILD_PATH=$(CURRENT_DIR)/build
CMAKE=$(shell which cmake)

all:
	@rm -rf $(BUILD_PATH)
	@mkdir -p $(BUILD_PATH)
	@pushd $(BUILD_PATH) && $(CMAKE) .. && popd
	@$(CMAKE) --build $(BUILD_PATH) --config $(BUILD_MODE) --target all -- -j 10

sessiond:
	cd $(BUILD_PATH) && make sessiond -j 10

run_redis_master:
	docker run \
	-v $(CURRENT_DIR)/etc/redis:/etc/redis \
	-p 16379:16379 \
	--name redis_master \
	-d \
	redis:latest \
	redis-server /etc/redis/redis_master.conf 

run_redis_slave:
	docker run \
	-v $(CURRENT_DIR)/etc/redis:/etc/redis \
	-p 26379:26379 \
	--name redis_slave \
	-d \
	redis:latest \
	redis-server /etc/redis/redis_slave.conf

clean:
	@rm -rf build
