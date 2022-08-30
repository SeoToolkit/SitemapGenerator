package cmd

import (
	"wp/sitemap/generate"

	"github.com/spf13/cobra"
)

var host string
var maxDepth int

func init() {
	crawlCmd.Flags().StringVarP(&host, "host", "t", "", "必填，抓取的目标网站的host,如https://www.baidu.com，就填www.baidu.com")
	crawlCmd.Flags().StringVarP(&fileType, "filetype", "f", "txt", "生成文件类型 xml 或者　txt")
	crawlCmd.Flags().IntVarP(&fileNum, "fileNum", "n", 50000, "每个文件的url数量,xml默认30000,txt默认50000")
	crawlCmd.Flags().IntVarP(&maxDepth, "maxDepth", "d", 3, "设置抓取深度，默认3")
	rootCmd.AddCommand(crawlCmd)
}

var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "生成sitemap",
	Long:  `生成sitemap,指定-f 为xml或者txt可生成对应的sitemap文件`,
	Run: func(cmd *cobra.Command, args []string) {
		generate.GetLinks(host, fileType, fileNum, maxDepth)
	},
}
