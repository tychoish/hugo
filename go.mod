module github.com/gohugoio/hugo

go 1.15

require (
	github.com/BurntSushi/locker v0.0.0-20171006230638-a6e239ea1c69
	github.com/BurntSushi/toml v0.3.1
	github.com/PuerkitoBio/purell v1.1.0
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/alecthomas/assert v0.0.0-20170929043011-405dbfeb8e38
	github.com/alecthomas/chroma v0.7.2-0.20200305040604-4f3623dce67a
	github.com/alecthomas/repr v0.0.0-20181024024818-d37bc2a10ba1 // indirect
	github.com/armon/go-radix v1.0.0
	github.com/aws/aws-sdk-go v1.34.5
	github.com/bep/debounce v1.2.0
	github.com/bep/gitmap v1.0.0
	github.com/bep/go-tocss v0.6.0
	github.com/bep/tmc v0.5.1
	github.com/chaseadamsio/goorgeous v1.1.0
	github.com/cpuguy83/go-md2man v1.0.8 // indirect
	github.com/disintegration/gift v1.2.1
	github.com/disintegration/imaging v1.5.0
	github.com/dustin/go-humanize v1.0.0
	github.com/eknkc/amber v0.0.0-20171010120322-cdade1c07385
	github.com/evanw/esbuild v0.6.24
	github.com/fortytw2/leaktest v1.3.0
	github.com/fsnotify/fsnotify v1.4.7
	github.com/getkin/kin-openapi v0.20.0
	github.com/ghodss/yaml v1.0.0
	github.com/gobwas/glob v0.2.3
	github.com/gorilla/websocket v1.4.0
	github.com/hashicorp/go-immutable-radix v1.0.0
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jdkato/prose v1.1.0
	github.com/kr/pretty v0.1.0 // indirect
	github.com/kyokomi/emoji v1.5.1
	github.com/magefile/mage v1.4.0
	github.com/markbates/inflect v1.0.0
	github.com/mattn/go-isatty v0.0.4
	github.com/mattn/go-runewidth v0.0.3 // indirect
	github.com/miekg/mmark v1.3.6
	github.com/mitchellh/hashstructure v1.0.0
	github.com/mitchellh/mapstructure v1.1.2
	github.com/muesli/smartcrop v0.0.0-20180228075044-f6ebaa786a12
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646 // indirect
	github.com/nicksnyder/go-i18n v1.10.0
	github.com/niklasfasching/go-org v1.3.2
	github.com/olekukonko/tablewriter v0.0.0-20180506121414-d4647c9c7a84
	github.com/pkg/errors v0.9.1
	github.com/rogpeppe/go-internal v1.6.1
	github.com/russross/blackfriday v0.0.0-20180804101149-46c73eb196ba
	github.com/rwcarlsen/goexif v0.0.0-20190401172101-9e8deecbddbd
	github.com/sanity-io/litter v1.1.0
	github.com/shurcooL/sanitized_anchor_name v0.0.0-20170918181015-86672fcb3f95 // indirect
	github.com/spf13/afero v1.2.0
	github.com/spf13/cast v1.3.0
	github.com/spf13/cobra v0.0.3
	github.com/spf13/fsync v0.0.0-20170320142552-12a01e648f05
	github.com/spf13/jwalterweatherman v1.0.1-0.20181028145347-94f6ae3ed3bc
	github.com/spf13/nitro v0.0.0-20131003134307-24d7ef30a12d
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.3.1
	github.com/stretchr/testify v1.5.1
	github.com/tdewolff/minify/v2 v2.3.7
	github.com/tychoish/shimgo v0.0.0-20181231152250-dadfa3df4ee9
	github.com/yosssi/ace v0.0.5
	github.com/yuin/goldmark v1.1.32
	github.com/yuin/goldmark-highlighting v0.0.0-20200307114337-60d527fdb691
	gocloud.dev v0.20.0
	golang.org/x/image v0.0.0-20190802002840-cff245a6509b
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208
	golang.org/x/text v0.3.3
	google.golang.org/api v0.30.0
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v2 v2.3.0
)

exclude github.com/chaseadamsio/goorgeous v2.0.0+incompatible

replace github.com/markbates/inflect => github.com/markbates/inflect v0.0.0-20171215194931-a12c3aec81a6
