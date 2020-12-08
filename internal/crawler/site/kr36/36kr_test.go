package kr36

import (
	"context"
	"io/ioutil"
	"testing"
)

func TestNewKr36(t *testing.T) {
	kr36 := NewKr36()
	feed, err := kr36.GenRssFeed(context.Background())
	if err != nil {
		t.Fatalf("GenRssFeed failed %v", err)
	}

	if len(feed.Items) == 0 {
		t.Fatal("can not generate rss feed, please checkout your code")
	}
	rssjson, err := feed.ToJSON()
	if err != nil {
		t.Fatalf("feed to rss failed %v", err)
	}
	rssxml, err := feed.ToRss()
	if err != nil {
		t.Fatalf("feed to rss failed %v", err)
	}
	t.Logf("HupuBXJ total Rss %d", len(feed.Items))
	ioutil.WriteFile("../../../../rss.json", []byte(rssjson), 0755)
	ioutil.WriteFile("../../../../rss.xml", []byte(rssxml), 0755)
}
