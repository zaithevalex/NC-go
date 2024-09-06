def get_status(status):
    if status == "success":
        return "program is success"
    else:
        return "program is failed"

def sum(num1, num2):
    return num1 + num2

def _test_rule_impl(ctx):
    version = ctx.attr.version
    print("version: ", version)

test_rule = rule(
    implementation= _test_rule_impl,
    attrs = {
        "version": attr.string(
            doc = "output file name",
            # if we forgot to pass arg version, will use default meaning of this attribute:
            default = "000000v"
        ),
    },
)