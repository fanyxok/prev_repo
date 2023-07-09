MPCFGO_ROOT := $(shell pwd)
# set root path of the project
# this macro is a compromise for the need of project root path
SETROOT := export MPCFGO_ROOT=$(MPCFGO_ROOT) &&  


all: config/config.json 
	@echo $(MPCFGO_ROOT)
	
setup:
	./setup.sh 


ote: config/config.json 
	-make -C cmd/ote clean 
	$(SETROOT)make -C cmd/ote build
	$(SETROOT)make -C cmd/ote run

# generate frontend for code optimization
gen_front: config/config.json 
# make sure we have progress after fail
	-make -C cmd/gtransl clean
	$(SETROOT)make -C cmd/gtransl build
	$(SETROOT)make -C cmd/gtransl run

TEST_SRC := test/yShare_test.go
TEST_REQUIRED := test/general.go
TEST_TARGET := test/yShare.test
TEST_FLAG := 

test: config/config.json $(TEST_TARGET)
	$(SETROOT) $(TEST_TARGET)

$(TEST_TARGET): $(TEST_REQUIRED) $(TEST_SRC)
	$(SETROOT)go test $(TESTFLAG) -c $^ -o $@

test_all: config/config.json 
	$(SETROOT)go test test/

config/config.json: config/config.json.1
	@echo "Generating config/config.json"
# substitute placeholder with project root path
	sed -e 's@PLACEHOLDER_ROOT@$(MPCFGO_ROOT)@g' config/config.json.1  > config/config.json 

clean:
	-make -C cmd/gtransl clean 
	rm config/config.json 

xiaban: 
	./xiaban.sh 


.PHONY: test clean xiaban gen_front ote setup