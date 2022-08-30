package builder

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"time"
	"wp/sitemap/dto"
)

type SiteMapType string

func (s SiteMapType) Str() string {
	return string(s)
}

const (
	Xml       SiteMapType = "xml"
	Txt       SiteMapType = "txt"
	UrlPrefix             = "https://bizhi.cheetahfun.com"
)

type SiteMapBuilder interface {
	GetDefaultLimit() int
	GetName() string
	GetFileType() SiteMapType
	BuildHead(writer *bufio.Writer)
	BuildBody(seoUrlsInfo dto.SiteUrl, writer *bufio.Writer)
	BuildFooter(writer *bufio.Writer)
	GenerateRootIndex(fpath string)
}

// host替换
func relpaceHost(url string) string {
	return strings.ReplaceAll(url, "ijinshan", "cheetahfun")
}
