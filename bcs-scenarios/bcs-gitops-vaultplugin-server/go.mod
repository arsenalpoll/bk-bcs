module github.com/Tencent/bk-bcs/bcs-scenarios/bcs-gitops-vaultplugin-server

go 1.19

replace (
        github.com/argoproj/gitops-engine => github.com/argoproj/gitops-engine v0.7.1-0.20230906152414-b0fffe419a0f
        github.com/chai2010/gettext-go => github.com/chai2010/gettext-go v0.1.0

        k8s.io/api => k8s.io/api v0.24.2
        k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.24.2
        k8s.io/apimachinery => k8s.io/apimachinery v0.24.2
        k8s.io/apiserver => k8s.io/apiserver v0.24.2
        k8s.io/cli-runtime => k8s.io/cli-runtime v0.24.2
        k8s.io/client-go => k8s.io/client-go v0.24.2
        k8s.io/cloud-provider => k8s.io/cloud-provider v0.24.2
        k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.24.2
        k8s.io/code-generator => k8s.io/code-generator v0.24.2
        k8s.io/component-base => k8s.io/component-base v0.24.2
        k8s.io/component-helpers => k8s.io/component-helpers v0.24.2
        k8s.io/controller-manager => k8s.io/controller-manager v0.24.2
        k8s.io/cri-api => k8s.io/cri-api v0.24.2
        k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.24.2
        k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.24.2
        k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.24.2
        k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20220627174259-011e075b9cb8
        k8s.io/kube-proxy => k8s.io/kube-proxy v0.24.2
        k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.24.2
        k8s.io/kubectl => k8s.io/kubectl v0.24.2
        k8s.io/kubelet => k8s.io/kubelet v0.24.2
        k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.24.2
        k8s.io/metrics => k8s.io/metrics v0.24.2
        k8s.io/mount-utils => k8s.io/mount-utils v0.24.2
        k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.24.2
        k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.24.2
        sigs.k8s.io/kustomize/api => sigs.k8s.io/kustomize/api v0.11.4
        sigs.k8s.io/kustomize/kyaml => sigs.k8s.io/kustomize/kyaml v0.13.6
)

require (
        github.com/Tencent/bk-bcs/bcs-common v0.0.0-20231124101643-0d25188774c2
        github.com/Tencent/bk-bcs/bcs-common/pkg/otel v0.0.0-20231124101643-0d25188774c2
        github.com/Tencent/bk-bcs/bcs-scenarios/bcs-gitops-manager v0.0.0-20231124101643-0d25188774c2
        github.com/agiledragon/gomonkey/v2 v2.11.0
        github.com/argoproj-labs/argocd-vault-plugin v1.17.0
        github.com/argoproj/argo-cd/v2 v2.9.2
        github.com/avast/retry-go v3.0.0+incompatible
        github.com/google/uuid v1.4.0
        github.com/gorilla/mux v1.8.1
        github.com/hashicorp/vault/api v1.10.0
        github.com/pkg/errors v0.9.1
        github.com/prometheus/client_golang v1.17.0
        github.com/spf13/viper v1.17.0
        github.com/stretchr/testify v1.8.4
        k8s.io/api v0.28.3
        k8s.io/apimachinery v0.28.3
        k8s.io/client-go v0.28.3
)

require (
        cloud.google.com/go/compute v1.23.2 // indirect
        cloud.google.com/go/compute/metadata v0.2.3 // indirect
        cloud.google.com/go/iam v1.1.4 // indirect
        cloud.google.com/go/secretmanager v1.11.3 // indirect
        dario.cat/mergo v1.0.0 // indirect
        filippo.io/age v1.0.0 // indirect
        github.com/1Password/connect-sdk-go v1.5.3 // indirect
        github.com/Azure/azure-sdk-for-go v68.0.0+incompatible // indirect
        github.com/Azure/go-ansiterm v0.0.0-20230124172434-306776ec8161 // indirect
        github.com/Azure/go-autorest v14.2.0+incompatible // indirect
        github.com/Azure/go-autorest/autorest v0.11.29 // indirect
        github.com/Azure/go-autorest/autorest/adal v0.9.23 // indirect
        github.com/Azure/go-autorest/autorest/azure/auth v0.5.12 // indirect
        github.com/Azure/go-autorest/autorest/azure/cli v0.4.6 // indirect
        github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
        github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
        github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
        github.com/Azure/go-autorest/logger v0.2.1 // indirect
        github.com/Azure/go-autorest/tracing v0.6.0 // indirect
        github.com/DelineaXPM/tss-sdk-go/v2 v2.0.0 // indirect
        github.com/IBM/go-sdk-core/v5 v5.14.1 // indirect
        github.com/IBM/secrets-manager-go-sdk v1.2.0 // indirect
        github.com/MakeNowJust/heredoc v1.0.0 // indirect
        github.com/Masterminds/semver/v3 v3.2.1 // indirect
        github.com/Microsoft/go-winio v0.6.1 // indirect
        github.com/ProtonMail/go-crypto v0.0.0-20230717121422-5aa5874ade95 // indirect
        github.com/acomagu/bufpipe v1.0.4 // indirect
        github.com/argoproj/gitops-engine v0.7.3 // indirect
        github.com/argoproj/pkg v0.13.7-0.20230626144333-d56162821bd1 // indirect
        github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
        github.com/aws/aws-sdk-go v1.44.331 // indirect
        github.com/aws/aws-sdk-go-v2 v1.21.2 // indirect
        github.com/aws/aws-sdk-go-v2/config v1.18.45 // indirect
        github.com/aws/aws-sdk-go-v2/credentials v1.13.43 // indirect
        github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.13.13 // indirect
        github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.43 // indirect
        github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.37 // indirect
        github.com/aws/aws-sdk-go-v2/internal/ini v1.3.45 // indirect
        github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.37 // indirect
        github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.21.5 // indirect
        github.com/aws/aws-sdk-go-v2/service/sso v1.15.2 // indirect
        github.com/aws/aws-sdk-go-v2/service/ssooidc v1.17.3 // indirect
        github.com/aws/aws-sdk-go-v2/service/sts v1.23.2 // indirect
        github.com/aws/smithy-go v1.15.0 // indirect
        github.com/beorn7/perks v1.0.1 // indirect
        github.com/bitly/go-simplejson v0.5.0 // indirect
        github.com/blang/semver v3.5.1+incompatible // indirect
        github.com/blang/semver/v4 v4.0.0 // indirect
        github.com/bmatcuk/doublestar/v4 v4.6.0 // indirect
        github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
        github.com/bombsimon/logrusr/v2 v2.0.1 // indirect
        github.com/bradleyfalzon/ghinstallation/v2 v2.6.0 // indirect
        github.com/cenkalti/backoff/v3 v3.2.2 // indirect
        github.com/cespare/xxhash/v2 v2.2.0 // indirect
        github.com/chai2010/gettext-go v1.0.2 // indirect
        github.com/cloudflare/circl v1.3.3 // indirect
        github.com/coreos/go-oidc/v3 v3.6.0 // indirect
        github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
        github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
        github.com/dimchansky/utfbom v1.1.1 // indirect
        github.com/docker/distribution v2.8.2+incompatible // indirect
        github.com/emicklei/go-restful/v3 v3.11.0 // indirect
        github.com/emirpasic/gods v1.18.1 // indirect
        github.com/evanphx/json-patch v5.6.0+incompatible // indirect
        github.com/exponent-io/jsonpath v0.0.0-20151013193312-d6023ce2651d // indirect
        github.com/fatih/camelcase v1.0.0 // indirect
        github.com/fatih/color v1.15.0 // indirect
        github.com/felixge/httpsnoop v1.0.4 // indirect
        github.com/fsnotify/fsnotify v1.6.0 // indirect
        github.com/fvbommel/sortorder v1.1.0 // indirect
        github.com/gabriel-vasile/mimetype v1.4.2 // indirect
        github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32 // indirect
        github.com/go-errors/errors v1.4.2 // indirect
        github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376 // indirect
        github.com/go-git/go-billy/v5 v5.4.1 // indirect
        github.com/go-git/go-git/v5 v5.8.1 // indirect
        github.com/go-jose/go-jose/v3 v3.0.0 // indirect
        github.com/go-logr/logr v1.3.0 // indirect
        github.com/go-logr/stdr v1.2.2 // indirect
        github.com/go-openapi/errors v0.20.3 // indirect
        github.com/go-openapi/jsonpointer v0.20.0 // indirect
        github.com/go-openapi/jsonreference v0.20.2 // indirect
        github.com/go-openapi/strfmt v0.21.7 // indirect
        github.com/go-openapi/swag v0.22.4 // indirect
        github.com/go-playground/locales v0.14.1 // indirect
        github.com/go-playground/universal-translator v0.18.1 // indirect
        github.com/go-playground/validator/v10 v10.14.0 // indirect
        github.com/go-redis/cache/v9 v9.0.0 // indirect
        github.com/gobwas/glob v0.2.3 // indirect
        github.com/gogo/protobuf v1.3.2 // indirect
        github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
        github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
        github.com/golang/protobuf v1.5.3 // indirect
        github.com/google/btree v1.1.2 // indirect
        github.com/google/gnostic v0.6.9 // indirect
        github.com/google/go-cmp v0.6.0 // indirect
        github.com/google/go-github/v53 v53.2.0 // indirect
        github.com/google/go-querystring v1.1.0 // indirect
        github.com/google/gofuzz v1.2.0 // indirect
        github.com/google/s2a-go v0.1.7 // indirect
        github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
        github.com/googleapis/enterprise-certificate-proxy v0.3.1 // indirect
        github.com/googleapis/gax-go/v2 v2.12.0 // indirect
        github.com/goware/prefixer v0.0.0-20160118172347-395022866408 // indirect
        github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
        github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
        github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
        github.com/hashicorp/errwrap v1.1.0 // indirect
        github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
        github.com/hashicorp/go-multierror v1.1.1 // indirect
        github.com/hashicorp/go-retryablehttp v0.7.4 // indirect
        github.com/hashicorp/go-rootcerts v1.0.2 // indirect
        github.com/hashicorp/go-secure-stdlib/parseutil v0.1.7 // indirect
        github.com/hashicorp/go-secure-stdlib/strutil v0.1.2 // indirect
        github.com/hashicorp/go-sockaddr v1.0.4 // indirect
        github.com/hashicorp/hcl v1.0.1-vault-5 // indirect
        github.com/howeyc/gopass v0.0.0-20210920133722-c8aef6fb66ef // indirect
        github.com/imdario/mergo v0.3.16 // indirect
        github.com/inconshreveable/mousetrap v1.1.0 // indirect
        github.com/itchyny/gojq v0.12.13 // indirect
        github.com/itchyny/timefmt-go v0.1.5 // indirect
        github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
        github.com/jmespath/go-jmespath v0.4.0 // indirect
        github.com/jonboulle/clockwork v0.2.2 // indirect
        github.com/josharian/intern v1.0.0 // indirect
        github.com/json-iterator/go v1.1.12 // indirect
        github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
        github.com/keeper-security/secrets-manager-go/core v1.6.2 // indirect
        github.com/kevinburke/ssh_config v1.2.0 // indirect
        github.com/klauspost/compress v1.17.0 // indirect
        github.com/leodido/go-urn v1.2.4 // indirect
        github.com/lib/pq v1.10.9 // indirect
        github.com/liggitt/tabwriter v0.0.0-20181228230101-89fcab3d43de // indirect
        github.com/magiconair/properties v1.8.7 // indirect
        github.com/mailru/easyjson v0.7.7 // indirect
        github.com/mattn/go-colorable v0.1.13 // indirect
        github.com/mattn/go-isatty v0.0.19 // indirect
        github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
        github.com/mitchellh/go-homedir v1.1.0 // indirect
        github.com/mitchellh/go-wordwrap v1.0.1 // indirect
        github.com/mitchellh/mapstructure v1.5.0 // indirect
        github.com/moby/spdystream v0.2.0 // indirect
        github.com/moby/term v0.5.0 // indirect
        github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
        github.com/modern-go/reflect2 v1.0.2 // indirect
        github.com/monochromegane/go-gitignore v0.0.0-20200626010858-205db1a8cc00 // indirect
        github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
        github.com/oklog/ulid v1.3.1 // indirect
        github.com/opencontainers/go-digest v1.0.0 // indirect
        github.com/opencontainers/image-spec v1.1.0-rc4 // indirect
        github.com/opencontainers/selinux v1.10.0 // indirect
        github.com/opentracing/opentracing-go v1.2.1-0.20220228012449-10b1cf09e00b // indirect
        github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
        github.com/pelletier/go-toml/v2 v2.1.0 // indirect
        github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
        github.com/pjbgf/sha1cd v0.3.0 // indirect
        github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
        github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring v0.69.1 // indirect
        github.com/prometheus-operator/prometheus-operator/pkg/client v0.69.1 // indirect
        github.com/prometheus/client_model v0.4.1-0.20230718164431-9a2bf3000d16 // indirect
        github.com/prometheus/common v0.44.0 // indirect
        github.com/prometheus/procfs v0.11.1 // indirect
        github.com/r3labs/diff v1.1.0 // indirect
        github.com/redis/go-redis/v9 v9.0.5 // indirect
        github.com/robfig/cron/v3 v3.0.1 // indirect
        github.com/russross/blackfriday v1.6.0 // indirect
        github.com/ryanuber/go-glob v1.0.0 // indirect
        github.com/sagikazarmark/locafero v0.3.0 // indirect
        github.com/sagikazarmark/slog-shim v0.1.0 // indirect
        github.com/sergi/go-diff v1.2.0 // indirect
        github.com/sirupsen/logrus v1.9.3 // indirect
        github.com/skeema/knownhosts v1.2.0 // indirect
        github.com/sourcegraph/conc v0.3.0 // indirect
        github.com/spf13/afero v1.10.0 // indirect
        github.com/spf13/cast v1.5.1 // indirect
        github.com/spf13/cobra v1.7.0 // indirect
        github.com/spf13/pflag v1.0.5 // indirect
        github.com/subosito/gotenv v1.6.0 // indirect
        github.com/uber/jaeger-client-go v2.30.0+incompatible // indirect
        github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
        github.com/ugorji/go/codec v1.2.11 // indirect
        github.com/vmihailenco/go-tinylfu v0.2.2 // indirect
        github.com/vmihailenco/msgpack/v5 v5.3.4 // indirect
        github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
        github.com/xanzy/ssh-agent v0.3.3 // indirect
        github.com/xlab/treeprint v1.2.0 // indirect
        github.com/yandex-cloud/go-genproto v0.0.0-20231009081144-b948e2f03d1e // indirect
        github.com/yandex-cloud/go-sdk v0.0.0-20231009081448-02cddfe74c51 // indirect
        go.mongodb.org/mongo-driver v1.12.1 // indirect
        go.mozilla.org/gopgagent v0.0.0-20170926210634-4d7ea76ff71a // indirect
        go.mozilla.org/sops/v3 v3.7.3 // indirect
        go.opencensus.io v0.24.0 // indirect
        go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.42.0 // indirect
        go.opentelemetry.io/otel v1.20.0 // indirect
        go.opentelemetry.io/otel/metric v1.20.0 // indirect
        go.opentelemetry.io/otel/trace v1.20.0 // indirect
        go.starlark.net v0.0.0-20230525235612-a134d8f9ddca // indirect
        go.uber.org/atomic v1.11.0 // indirect
        go.uber.org/multierr v1.11.0 // indirect
        golang.org/x/crypto v0.14.0 // indirect
        golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
        golang.org/x/mod v0.12.0 // indirect
        golang.org/x/net v0.17.0 // indirect
        golang.org/x/oauth2 v0.13.0 // indirect
        golang.org/x/sync v0.3.0 // indirect
        golang.org/x/sys v0.13.0 // indirect
        golang.org/x/term v0.13.0 // indirect
        golang.org/x/text v0.14.0 // indirect
        golang.org/x/time v0.3.0 // indirect
        golang.org/x/tools v0.13.0 // indirect
        google.golang.org/api v0.143.0 // indirect
        google.golang.org/appengine v1.6.8 // indirect
        google.golang.org/genproto v0.0.0-20231030173426-d783a09b4405 // indirect
        google.golang.org/genproto/googleapis/api v0.0.0-20231106174013-bbf56f31fb17 // indirect
        google.golang.org/genproto/googleapis/rpc v0.0.0-20231030173426-d783a09b4405 // indirect
        google.golang.org/grpc v1.59.0 // indirect
        google.golang.org/protobuf v1.31.0 // indirect
        gopkg.in/inf.v0 v0.9.1 // indirect
        gopkg.in/ini.v1 v1.67.0 // indirect
        gopkg.in/urfave/cli.v1 v1.20.0 // indirect
        gopkg.in/warnings.v0 v0.1.2 // indirect
        gopkg.in/yaml.v2 v2.4.0 // indirect
        gopkg.in/yaml.v3 v3.0.1 // indirect
        k8s.io/apiextensions-apiserver v0.28.3 // indirect
        k8s.io/apiserver v0.28.3 // indirect
        k8s.io/cli-runtime v0.28.3 // indirect
        k8s.io/cloud-provider v0.24.2 // indirect
        k8s.io/component-base v0.28.3 // indirect
        k8s.io/component-helpers v0.28.0 // indirect
        k8s.io/cri-api v0.0.0 // indirect
        k8s.io/klog/v2 v2.100.1 // indirect
        k8s.io/kube-aggregator v0.24.2 // indirect
        k8s.io/kube-openapi v0.0.0-20230905202853-d090da108d2f // indirect
        k8s.io/kubectl v0.28.0 // indirect
        k8s.io/kubernetes v1.24.2 // indirect
        k8s.io/mount-utils v0.24.2 // indirect
        k8s.io/utils v0.0.0-20230726121419-3b25d923346b // indirect
        oras.land/oras-go/v2 v2.3.0 // indirect
        sigs.k8s.io/controller-runtime v0.16.3 // indirect
        sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
        sigs.k8s.io/kustomize/api v0.13.5-0.20230601165947-6ce0bf390ce3 // indirect
        sigs.k8s.io/kustomize/kyaml v0.14.2 // indirect
        sigs.k8s.io/structured-merge-diff/v4 v4.3.0 // indirect
        sigs.k8s.io/yaml v1.3.0 // indirect
)