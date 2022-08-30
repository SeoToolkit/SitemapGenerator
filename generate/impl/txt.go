package impl

type TxtSitemapBuilder struct {
}

// GetDefaultLimit　default 5w
func (s *TxtSitemapBuilder) GetDefaultLimit() int {
	return 50000
}
func (s *TxtSitemapBuilder) GetName() string {
	return "txt-sitemap"
}

func (s *TxtSitemapBuilder) GetFileType() SiteMapType {
	return Txt
}

func (s *TxtSitemapBuilder) BuildFooter(writer *bufio.Writer) {
	return
}
func (s *TxtSitemapBuilder) BuildHead(writer *bufio.Writer) {
	return
}
func (s *TxtSitemapBuilder) BuildBody(v dto.SiteUrl, writer *bufio.Writer) {
	writer.WriteString(relpaceHost(v.String()) + "\r\n")
}

func (s *TxtSitemapBuilder) GenerateRootIndex(fpath string) {
	return
}
