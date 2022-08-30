package cmd

import (
	"github.com/spf13/cobra"
)

var fileType string
var fileNum int

func init() {
	generateCmd.Flags().StringVarP(&fileType, "filetype", "f", "txt", "生成文件类型 xml 或者　txt")
	generateCmd.Flags().IntVarP(&fileNum, "fileNum", "n", 50000, "每个文件的url数量,xml默认30000,txt默认50000")
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "生成sitemap",
	Long:  `生成sitemap,指定-f 为xml或者txt可生成对应的sitemap文件`,
	Run: func(cmd *cobra.Command, args []string) {
		generate.GenerateWpSiteMap(fileType, fileNum)
	},
}
