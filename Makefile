BINARY_PATH=/var/apps/qrapp
BINARY_NAME=survey
SERVICE_CONFIG_FILE=/lib/systemd/system/${BINARY_NAME}.service

# build : will build the app for linux
build:
	@echo "building linux binary started ..."
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}
	@echo "build done."

# run : builds and run the service
run: build
	@echo executing the serviec binary
	./${BINARY_NAME}

# service : creates configuration file and start the app as a service in linux OS
service:
	@echo "creating service config file"
	> ${SERVICE_CONFIG_FILE}
	echo "[Unit]" >> ${SERVICE_CONFIG_FILE}
	echo "Description=qrcode" >> ${SERVICE_CONFIG_FILE}
	echo "[Service]" >> ${SERVICE_CONFIG_FILE}
	echo "Type=simple" >> ${SERVICE_CONFIG_FILE}
	echo "Restart=always" >> ${SERVICE_CONFIG_FILE}
	echo "RestartSec=5s" >> ${SERVICE_CONFIG_FILE}
	echo "ExecStart=${BINARY_PATH}/${BINARY_NAME}" >> ${SERVICE_CONFIG_FILE}
	echo "[Install]" >> ${SERVICE_CONFIG_FILE}
	echo "WantedBy=multi-user.target" >> ${SERVICE_CONFIG_FILE}
	@echo "service configuration done"
	@echo "************************* starting service .."
	service ${BINARY_NAME} start
	@echo "************************* service status .."
	service ${BINARY_NAME} status

clean:
	@echo clean started ...
	go clean
	rm ${BINARY_NAME}
	@echo cleaning done.


