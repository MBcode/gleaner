package objects

import (
	//	config "github.com/earthcubearchitecture-project418/gleaner/internal/config"
	"github.com/spf13/viper"
	"log"
)

type Sources struct {
	Name       string
	Logo       string
	URL        string
	Headless   bool
	PID        string
	ProperName string
	Domain     string
	// SitemapFormat string
	// Active        bool
}

// Return all sources and sitegraph domains
func SourcesAndGraphs(v1 *viper.Viper) []Sources {
	var domains []Sources
	var sm []Sources
	var sg []Sources

	//err := v1.UnmarshalKey("sitemaps", &sm)
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//err = v1.UnmarshalKey("sitegraphs", &sg)
	//if err != nil {
	//	log.Println(err)
	//}
	err := v1.UnmarshalKey("sources", &sm)
	if err != nil {
		log.Println(err)
	}

	domains = append(sg, sm...)

	return domains

}
