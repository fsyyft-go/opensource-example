# 编码输出的应用程序名称。
APP_OUTPUT=./bin
APP_LOGS=./logs

# 设置命令路径。
CMD_CAT=cat
CMD_CD=cd
CMD_CP=cp
CMD_CLEAR=clear
CMD_DATE=date
CMD_DIRNAME=dirname
CMD_GREP=grep
CMD_ECHO=echo
CMD_FILD=find
CMD_GO=go
CMD_GOLDS=golds
# 安装方式：https://golangci-lint.run/usage/install/#local-installation 。
CMD_GOLANGCI-LINT=golangci-lint
CMD_GOVULNCHECK=govulncheck
CMD_GO_ASSETS_BUILDER=go-assets-builder
CMD_GO_BINDATA=go-bindata
CMD_GO_CALLVIS=go-callvis
CMD_GIT=git
CMD_GMCHART=gmchart
CMD_LS=ls
CMD_MKDIR=mkdir
CMD_MV=mv
CMD_PROTOC=protoc
CMD_PWD=pwd
CMD_RM=rm
CMD_SED=sed
CMD_SWAG=swag
CMD_TINYGO=tinygo
CMD_UPX=upx
CMD_UNZIP=unzip
CMD_ZIP=zip

.PHONY: test
test: test-cmd/example
	# 准备单元测试后置工作。
	-@$(CMD_RM) -rf *.test
	@$(CMD_GO) mod tidy
	# 准备单元测试后置工作。

.PHONY: test-pre
test-pre: clean
	# 准备单元测试前置工作。
	# 准备对依赖包进行梳理。
	@$(CMD_GO) mod tidy
	# 完成对依赖包进行梳理。
	# 准备静态代码检查。
	-@$(CMD_GOLANGCI-LINT) run --verbose ./...
	# 结束静态代码检查。
	# 准备代码安全检查。
	-@${CMD_GOVULNCHECK} ./...
	# 完成代码安全检查。
	# 结束单元测试前置工作。

.PHONY: test-cmd/example
test-cmd/example:
	# 准备测试 cmd/example。
	@$(CMD_MKDIR) -p $(APP_OUTPUT)/out/$(subst test-,,$@)

	@$(CMD_GO) test -v -bench=. ./$(subst test-,,$@)                        \
		 	-coverprofile=$(APP_OUTPUT)/out/$(subst test-,,$@)/conver.out   \
			-cpuprofile=$(APP_OUTPUT)/out/$(subst test-,,$@)/cpu.out        \
			-memprofile=$(APP_OUTPUT)/out/$(subst test-,,$@)/mem.out        \
			-blockprofile=$(APP_OUTPUT)/out/$(subst test-,,$@)/block.out

	@$(CMD_GO) tool pprof -pdf $(APP_OUTPUT)/out/$(subst test-,,$@)/cpu.out     > $(APP_OUTPUT)/out/$(subst test-,,$@)/cpu.pdf
	@$(CMD_GO) tool pprof -pdf $(APP_OUTPUT)/out/$(subst test-,,$@)/mem.out     > $(APP_OUTPUT)/out/$(subst test-,,$@)/mem.pdf
	@$(CMD_GO) tool pprof -pdf $(APP_OUTPUT)/out/$(subst test-,,$@)/block.out   > $(APP_OUTPUT)/out/$(subst test-,,$@)/block.pdf
	# 完成测试 cmd/example。

.PHONY: clean
clean:
	@$(CMD_CLEAR)
	# 准备清理已编译文件。
	-@$(CMD_RM) -rf $(APP_OUTPUT)
	-@$(CMD_RM) -rf $(APP_LOGS)
	@$(CMD_MKDIR) $(APP_OUTPUT)
	@$(CMD_MKDIR) $(APP_OUTPUT)/out
	@$(CMD_MKDIR) $(APP_OUTPUT)/test
	@$(CMD_MKDIR) $(APP_OUTPUT)/darwin_amd64
	@$(CMD_MKDIR) $(APP_OUTPUT)/linux_amd64
	@$(CMD_MKDIR) $(APP_OUTPUT)/linux_arm
	@$(CMD_MKDIR) $(APP_OUTPUT)/windows_amd64
	@$(CMD_MKDIR) $(APP_LOGS)
	@$(CMD_GO) mod tidy
	# 完成清理已编译文件。

.PHONY: tidy
tidy:
	# 准备清理无效引用。
	@$(CMD_CAT) go.mod | $(CMD_GREP) -v "// indirect" > go.mod.tmp
	@$(CMD_RM) -rf go.mod
	@$(CMD_MV) go.mod.tmp go.mod
	@$(CMD_GO) mod tidy
	# 完成清理无效引用。