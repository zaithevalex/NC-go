load("@io_bazel_rules_go//go:def.bzl", "go_binary")
load("//:functions.bzl", "test_rule")

build_status = "running"

#filegroup(
#    name = "templates",
#    srcs = glob(["templates/*"]),
#)

go_binary(
    name = "main",
    srcs = ["src/bazelrun.go"],
    data = [":templates"],
)
build_status = "success"
load("//:functions.bzl", "get_status")
print(get_status(build_status))

# CONSTS:
CONST_1 = "const"

load("//:functions.bzl", "sum")
sum = sum(1, 2)
print("sum is", sum)

print("go is %s" % "cool lang")

#load("@rules_jvm_external//:defs.bzl", "maven_install")
#maven_install(
# ...
#)

test_rule(
    name = "test_rule",
    version = "123123v",
)