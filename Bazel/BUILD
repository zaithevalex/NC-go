load("@io_bazel_rules_go//go:def.bzl", "go_binary")

#filegroup(
#    name = "templates",
#    srcs = glob(["templates/*"]),
#)

go_binary(
    name = "main",
    srcs = ["src/bazelrun.go"],
    data = [":templates"],
)