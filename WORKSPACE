workspace(name = "react_ts_bazel")

# ================================================================
# Repositories
# ================================================================
http_archive(
    name = "build_bazel_rules_typescript",
    url = "https://github.com/bazelbuild/rules_typescript/archive/0.20.2.zip",
    strip_prefix = "rules_typescript-0.20.2",
)

http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.16.1/rules_go-0.16.1.tar.gz",
    sha256 = "f5127a8f911468cd0b2d7a141f17253db81177523e4429796e14d429f5444f5f",
)

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.15.0/bazel-gazelle-0.15.0.tar.gz"],
    sha256 = "6e875ab4b6bf64a38c352887760f21203ab054676d9c1b274963907e0768740d",
)

http_archive(
    name = "build_stack_rules_proto",
    urls = ["https://github.com/stackb/rules_proto/archive/844077a71597f91c41b02d4509c5f79d51588552.tar.gz"],
    sha256 = "867d09bf45515cb3ddeb06f7bdd2182eecf171ae3cd6b716b3b9d2fce50f292f",
    strip_prefix = "rules_proto-844077a71597f91c41b02d4509c5f79d51588552",
)

# ================================================================
# Typescript extensions
# ================================================================
# https://github.com/bazelbuild/rules_typescript/blob/master/README.md

# Fetch our Bazel dependencies that aren't distributed on npm
load("@build_bazel_rules_typescript//:package.bzl", "rules_typescript_dependencies")
rules_typescript_dependencies()

load("@build_bazel_rules_typescript//:defs.bzl", "ts_setup_workspace")
ts_setup_workspace()

load("@build_bazel_rules_nodejs//:defs.bzl", "node_repositories", "yarn_install")
node_repositories()

# Setup Bazel managed npm dependencies with the `yarn_install` rule.
# The name of this rule should be set to `npm` so that `ts_library` and `ts_web_test_suite`
# can find your npm dependencies by default in the `@npm` workspace. You may
# also use the `npm_install` rule with a `package-lock.json` file if you prefer.
# See https://github.com/bazelbuild/rules_nodejs#dependencies for more info.
yarn_install(
  name = "npm",
  package_json = "//:package.json",
  yarn_lock = "//:yarn.lock",
)

# Setup web testing, choose browsers we can test on
load(
	"@io_bazel_rules_webtesting//web:repositories.bzl",
	"browser_repositories",
	"web_test_repositories",
)
web_test_repositories()
browser_repositories(
    chromium = True,
)

# ================================================================
# Go extensions
# ================================================================
load(
	"@io_bazel_rules_go//go:def.bzl",
	"go_rules_dependencies",
	"go_register_toolchains",
)
go_rules_dependencies()
go_register_toolchains()

load(
	"@bazel_gazelle//:deps.bzl",
	"gazelle_dependencies",
	"go_repository"
)
gazelle_dependencies()

# ================================================================
# Protobuf extensions
# ================================================================
load("@build_stack_rules_proto//go:deps.bzl", "go_proto_library")
go_proto_library()
