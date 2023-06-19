package resources

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/sunil494/cq-source-test/internal/xkcd"
)

type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func Comics() *schema.Table {
	return &schema.Table{
		Name:      "xkcd_comics",
		Resolver:  fetchComics,
		Transform: transformers.TransformWithStruct(&xkcd.Comic{}),
	}
}

func fetchComics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	client := meta.(*Client)
	latest, err := client.XKCD.GetLatestComic(ctx)
	if err != nil {
		return err
	}
	res <- latest

	for i := 1; i < latest.Num; i++ {
		comic, err := client.XKCD.GetComic(ctx, i)
		if err != nil {
			return err
		}
		res <- comic
	}
	return nil
}
