package persist

import (
	"My_Go_Study/crawler/engine"
	"gopkg.in/olivere/elastic.v5"
	"My_Go_Study/crawler/persist"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {

	err := persist.Save(s.Client, s.Index, item)
	log.Printf("save profile %s", item.Url)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("item:%s save error: %s", item, err)
	}

	return err
}