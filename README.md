创建可以工kubectl使用的patch

```shell
Usage:
  kubepatch from.yaml to.yaml [flags]

Examples:
kubepatch from.yaml to.yaml

Flags:
  -h, --help            help for kubepatch
  -o, --output string   patch输出类型；可选的参数有[json yaml] (default "yaml")
  -t, --type string     生成的patch类型；可选的参数有[json strategic] (default "json")
  -v, --version         version for kubepatch
```

