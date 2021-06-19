package main

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

const (
	fmtJSON = "json"
	fmtStrategic = "strategic"
	fmtYaml = "yaml"
)

var (
	patchType string
	outputType string
)

func main()  {
	var rootCmd = &cobra.Command{
		Use:                        "kubepatch from.yaml to.yaml",
		Short:                      "构造k8s的patch",
		Long:                       "构造可以用于kubectl的patch",
		Example:                    "kubepatch from.yaml to.yaml",
		Version:                    "0.1",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return fmt.Errorf("usage: kubepatch from.yaml to.yaml")
			}
			var patch string
			var err error
			switch patchType {
			case fmtJSON:
				patch, err = GenerateJsonPatchFromFile(args[0], args[1])
			case fmtStrategic:
				patch, err = GenerateStrategicMergePatchFromFile(args[0], args[1])
			default:
				return fmt.Errorf("unknown patch type %s", patchType)
			}
			if err != nil {
				return err
			}

			fmt.Println(patch)
			return nil
		},
	}

	rootCmd.Flags().StringVarP(&patchType,"type", "t", fmtJSON, "生成的patch类型；可选的参数有[json strategic]")
	rootCmd.Flags().StringVarP(&outputType, "output", "o", fmtYaml, "patch输出类型；可选的参数有[json yaml]")
	rootCmd.Flags().AddGoFlagSet(flag.CommandLine)
	utilruntime.Must(flag.CommandLine.Parse([]string{}))

	utilruntime.Must(rootCmd.Execute())
}
