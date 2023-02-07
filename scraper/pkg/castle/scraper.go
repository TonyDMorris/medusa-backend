package castle

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

const scrapedDir = "./scraped/"
const Domain = "https://www.castlefineart.com"

func Scrape() {

	c := colly.NewCollector()
	onlyOnce := make(map[string]bool)
	mu := sync.Mutex{}
	// Find and visit all links
	c.OnHTML("a", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		if strings.HasPrefix(href, "/artists") {

			split := strings.Split(href, "/")
			if len(split) > 2 {
				artist := split[2]
				mu.Lock()
				if _, ok := onlyOnce[artist]; !ok {
					onlyOnce[artist] = true
					go func() {
						err := ScrapeArtistProfile(Domain, artist)
						if err != nil {
							fmt.Print(err)
						}
					}()

				}
				mu.Unlock()

			}

		}
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Print(string(r.Body))
		panic(err)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(Domain + "/artists")
}

func ScrapeArtistProfile(domain, artist string) error {
	fmt.Printf("scraping artist %s \n", artist)
	err := PrepareDirs(artist)
	if err != nil {
		return err
	}
	c := colly.NewCollector()
	profileDir := scrapedDir + "artists/" + artist + "/profile/"
	imageNo := 2
	c.OnHTML("img", func(e *colly.HTMLElement) {
		src := e.Attr("src")
		if strings.Contains(src, "/resized/fullscreen/") {
			err := SaveImage(profileDir, src, "fullscreen")
			if err != nil {
				fmt.Print(err)
			}
			return
		}
		if strings.Contains(src, "/thumbnail-1x1-w100/") {
			err := SaveImage(profileDir, src, "1_of_3")
			if err != nil {
				fmt.Print(err)
			}
			return
		}
		if strings.Contains(src, "/thumbnail-4x3-w50/") {
			err := SaveImage(profileDir, src, fmt.Sprintf("%d_of_3", imageNo))
			if err != nil {
				fmt.Print(err)
			}
			imageNo++
			return
		}
		if strings.Contains(src, "/artist_hero/") {
			err := SaveImage(profileDir, src, "hero")
			if err != nil {
				fmt.Print(err)
			}
			return
		}
		if strings.Contains(src, "/icon-128/") {
			err := SaveImage(profileDir, src, "icon")
			if err != nil {
				fmt.Print(err)
			}
			return
		}

	})
	bio := make(map[string]string)
	bioMu := sync.Mutex{}
	c.OnHTML("span[class=text-black]", func(e *colly.HTMLElement) {
		bioMu.Lock()
		split := strings.Split(e.DOM.Parent().Text(), ":")
		key := strings.TrimSpace(split[0])
		value := strings.TrimSpace(split[1])
		bio[key] = value
		bioMu.Unlock()
	})
	inspiration := []string{}

	c.OnHTML("h2", func(e *colly.HTMLElement) {
		if strings.TrimSpace(e.Text) == "Inspiration" {
			e.DOM.Siblings().Children().Each(func(i int, g *goquery.Selection) {
				inspiration = append(inspiration, g.Text())
			})
			fmt.Print(inspiration)
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Print(string(r.Body))
		fmt.Print(err)

	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err = c.Visit(domain + "/artists/" + artist)
	if err != nil {
		return err
	}
	bioBytes, err := json.Marshal(bio)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(profileDir+"bio.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	f.Close()
	_, err = f.Write(bioBytes)
	if err != nil {
		return err
	}
	inspirationBytes, err := json.Marshal(inspiration)
	if err != nil {
		return err
	}
	f, err = os.OpenFile(profileDir+"inspiration.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(inspirationBytes)
	if err != nil {
		return err
	}
	return nil

}

func SaveImage(dir, url string, filename string) error {
	ext1 := filepath.Ext(url)
	_, err := os.Stat(dir + filename + ext1)
	if os.IsNotExist(err) {
		res, err := http.Get(url)

		if err != nil {
			return err
		}
		if res.StatusCode != http.StatusOK {
			fmt.Printf("unable to download image %s", url)
			return nil
		}
		bytes, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return os.WriteFile(dir+filename+ext1, bytes, 0644)
	}
	return nil
}

func PrepareDirs(artist string) error {
	if err := os.Mkdir(scrapedDir+"artists/"+artist, os.ModePerm); err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}
	if err := os.Mkdir(scrapedDir+"artists/"+artist+"/profile", os.ModePerm); err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}
	if err := os.Mkdir(scrapedDir+"artists/"+artist+"/pieces", os.ModePerm); err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}
	return nil

}
