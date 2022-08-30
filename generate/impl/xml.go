package impl

// <?xml version="1.0" encoding="UTF-8"?>
// <urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
//    <url>
//       <loc>http://www.example.com/</loc>
//       <lastmod>2005-01-01</lastmod>
//       <changefreq>monthly</changefreq>
//       <priority>0.8</priority>
//    </url>
// </urlset>
type XmlSitemapBuilder struct {
}

// GetDefaultLimit 默认为3w
func (s *XmlSitemapBuilder) GetDefaultLimit() int {
	return 30000
}
func (s *XmlSitemapBuilder) GetName() string {
	return "xml-sitemap"
}

func (s *XmlSitemapBuilder) GetFileType() SiteMapType {
	return Xml
}

func (s *XmlSitemapBuilder) BuildHead(writer *bufio.Writer) {
	writer.WriteString(strings.TrimSpace(`
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">	
	`) + "\r\n")
}

func (s *XmlSitemapBuilder) BuildFooter(writer *bufio.Writer) {
	writer.WriteString("</urlset>" + "\r\n")
}

func (s *XmlSitemapBuilder) BuildBody(pageUrl dto.SiteUrl, writer *bufio.Writer) {
	var priority float32 = 1.0
	if strings.HasSuffix(pageUrl.String(), ".shtml") || strings.HasSuffix(pageUrl.String(), ".html") {
		priority = 0.8
	}

	writer.WriteString(strings.TrimSpace(fmt.Sprintf(`
	<url>
		<loc>%s</loc>
		<lastmod>%s</lastmod>
		<priority>%.2f</priority>
   	</url>
	`, relpaceHost(pageUrl.String()), time.Now().Format("2006-01-02"), priority)) + "\r\n")
}

// <?xml version="1.0" encoding="UTF-8"?>
// <sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
//   <sitemap>
//     <loc>http://www.example.com/sitemap1.xml.gz</loc>
//   </sitemap>
//   <sitemap>
//     <loc>http://www.example.com/sitemap2.xml.gz</loc>
//   </sitemap>
// </sitemapindex>
func (s *XmlSitemapBuilder) GenerateRootIndex(fpath string) {
	files, err := ioutil.ReadDir(fpath)
	if err != nil {
		panic(fmt.Sprintf("open path %s occure err:%v", fpath, err))
	}
	filelist := make([]string, 10)
	for _, f := range files {
		log.Printf(f.Name())
		filelist = append(filelist, f.Name())

	}
	dstFileName := fpath + "/sitemap.xml"
	// get xml file list
	dstFd, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer dstFd.Close()
	if err != nil {
		panic(fmt.Sprintf("write xml root index [%s] occure err:%v", dstFileName, err))
	}
	writer := bufio.NewWriter(dstFd)
	writer.WriteString(strings.TrimSpace(`
<?xml version="1.0" encoding="UTF-8"?>
<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
	`) + "\r\n")

	sort.Slice(filelist, func(i, j int) bool {
		return filelist[i] < filelist[j]
	})
	now := time.Now().Format("2006-01-02")
	for _, v := range filelist {
		if len(v) == 0 {
			continue
		}
		v = UrlPrefix + "/sitemap/xml/" + v
		writer.WriteString(strings.TrimSpace(fmt.Sprintf(`
<sitemap>
    <loc>%s</loc>
	<lastmod>%s</lastmod>
</sitemap>
		`, v, now)) + "\r\n")
	}

	writer.WriteString("</sitemapindex>")
	writer.Flush()
	return
}
