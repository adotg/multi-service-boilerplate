.PHONY: package run-local

export ENV1=env1_value

THIS_DIR := $(dir $(abspath $(firstword $(MAKEFILE_LIST))))

run-local:

	@echo "ENV1=${ENV1}"
	@echo "SECRET1=${SECRET1}"
	@echo "THIS_DIR=${THIS_DIR}"

	echo "REACT_APP_SECRET1=${SECRET1}" > env.development.tmp
	docker build -t ui .
	docker run --rm \
		-v ${THIS_DIR}env.development.tmp:/source/.env.development \
		-e REACT_APP_ENV1="${ENV1}" \
		-p 3000:3000 \
		ui
	rm env.development.tmp
	

package:
	rm source.tar
	git archive -o source.tar HEAD
	aws s3 cp source.tar s3://constellation-prior/tmp/source.tar

