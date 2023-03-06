# 当前版本
Version="2.0"

AVGPG=""

# package Name
CLOUDOPS_PACKAGE_NAME="cloudopscmd"
CLOUDTOOL_PACKAGE_NAME="cloudtoolcmd"

# 获取源码最近一次 git commit log，包含 commit sha 值，以及 commit message
GitCommitLog=`git log -1 --pretty=format:%H`
# 检查源码在git commit 基础上，是否有本地修改，且未提交的内容
GitStatus=`git status -s`
# 检索作者
GitAuthor=`git log --pretty="%an"  -1 | base64`
# 获取分支
GitBranch="CloudDevKubernetes.com/dev"
# 获取当前时间
BuildTime=`date +'%Y.%m.%d.%H%M%S'` 
# 获取 Go 的版本
BuildGoVersion=`go version`

# 将以上变量序列化至 LDFlags 变量中
LDFlags=" \
    -X 'CloudDevKubernetes/util.version=${Version}' \
    -X 'CloudDevKubernetes/util.gitCommit=${GitCommitLog}' \
    -X 'CloudDevKubernetes/util.gitAuthor=${GitAuthor}' \
    -X 'CloudDevKubernetes/util.gitBranch=${GitBranch}' \
    -X 'CloudDevKubernetes/util.buildDate=${BuildTime}' \
    -X 'CloudDevKubernetes/util.buildGoVersion=${BuildGoVersion}' \
"


# 程序路径
ROOT_DIR=`pwd`
CLOUDOPS_MAIN_PATH="${ROOT_DIR}/cmds/${CLOUDOPS_PACKAGE_NAME}/${CLOUDOPS_PACKAGE_NAME}.go"
CLOUDTOOL_MAIN_PATH="${ROOT_DIR}/cmds/${CLOUDTOOL_PACKAGE_NAME}/${CLOUDTOOL_PACKAGE_NAME}.go"

CLOUDOPS_READMD_PATH="${ROOT_DIR}/cmds/${CLOUDOPS_PACKAGE_NAME}/README.md"
CLOUDTOOL_README_PATH="${ROOT_DIR}/cmds/${CLOUDTOOL_PACKAGE_NAME}/README.md"

CLOUDOPS_PACKAGE_PATH="${ROOT_DIR}/bin/${CLOUDOPS_PACKAGE_NAME}"
CLOUDTOOL_PACKAGE_PATH="${ROOT_DIR}/bin/${CLOUDTOOL_PACKAGE_NAME}"

MAIN_PATH="${ROOT_DIR}/main.go"


#########################################################################################################################
# 如果可执行程序输出目录不存在，则创建
bindir:
		@if [ ! -d ${ROOT_DIR}/bin ]; then \
			mkdir ${ROOT_DIR}/bin ; \
		fi;
		@if [ ! -d ${CLOUDOPS_PACKAGE_PATH} ]; then \
			mkdir ${CLOUDOPS_PACKAGE_PATH} ; \
		fi;
		@if [ ! -d ${CLOUDTOOL_PACKAGE_PATH} ]; then \
			mkdir ${CLOUDTOOL_PACKAGE_PATH} ; \
		fi;

readme:
		@if [ ! -f ${ROOT_DIR}/README.md ]; then \
			echo >> ${ROOT_DIR}/README.md ; \
		fi;
		
		@if [ ! -f ${CLOUDOPS_READMD_PATH} ]; then \
			touch ${CLOUDOPS_READMD_PATH} ; \
		fi;
		
		@if [ ! -f ${CLOUDTOOL_README_PATH} ]; then \
			touch ${CLOUDTOOL_README_PATH} ; \
		fi;
		

modtidy:
		@go mod tidy

gendoc:
		@go run main.go --doc --docDir="bin/docs" 2>/dev/null 

cleanall:
		@if [ -d ${CLOUDOPS_PACKAGE_PATH} ]; then \
			rm -rf ${CLOUDOPS_PACKAGE_PATH};\
		fi;
		@if [ -d ${CLOUDTOOL_PACKAGE_PATH} ]; then \
			rm -rf ${CLOUDTOOL_PACKAGE_PATH};\
		fi;

codecheck: modtidy
		@golangci-lint run $(CLOUDOPS_MAIN_PATH)
		@golangci-lint run $(CLOUDTOOL_MAIN_PATH)
		@gosec ${ROOT_DIR}/*

####################################################################################################################################

## build cloudopscmd
cloudopscmdbuild_arm:  bindir   modtidy
		@rm -rf ${CLOUDOPS_PACKAGE_PATH}/arm && mkdir ${CLOUDOPS_PACKAGE_PATH}/arm
		@GOOS="linux" GOARCH="arm64" go build -ldflags ${LDFlags}  -o ${CLOUDOPS_PACKAGE_PATH}/arm ${CLOUDOPS_MAIN_PATH}
		@echo "BUILD DONE: ${CLOUDOPS_PACKAGE_PATH}/arm/${CLOUDOPS_PACKAGE_NAME}"

cloudopscmdbuild_amd:  bindir  modtidy
		@rm -rf ${CLOUDOPS_PACKAGE_PATH}/amd && mkdir ${CLOUDOPS_PACKAGE_PATH}/amd
		@GOOS="linux" GOARCH="amd64" go build -ldflags ${LDFlags} -o ${CLOUDOPS_PACKAGE_PATH}/amd ${CLOUDOPS_MAIN_PATH}
		@echo "BUILD DONE: ${CLOUDOPS_PACKAGE_PATH}/amd/${CLOUDOPS_PACKAGE_NAME}"
		
cloudopspackage:   cloudopscmdbuild_arm cloudopscmdbuild_amd  
		@if [ -d bin ]; then \
    	# tar -czf cloudopscmd.tar.gz  bin 2>/dev/null ; \
			# mv cloudopscmd.tar.gz bin; \
			echo "BUILD Over ..." ;\
		else \
		  echo "当前目录不正确，请切换到main同级目录" ;\
		fi;
