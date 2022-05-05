package algorithm

import (
	"fmt"
	"reflect"
	"sort"

	log "github.com/sirupsen/logrus"
	"github.com/z9fr/greensforum-backend/internal/database"
	"github.com/z9fr/greensforum-backend/internal/question"
	"github.com/z9fr/greensforum-backend/internal/user"
	"gorm.io/gorm"
)

type App struct {
	Name    string
	Version string
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

func (app *App) Run() error {
	//log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    app.Name,
			"AppVersion": app.Version,
		}).Info("Setting up Application")

	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	if err := database.MigrateDB(db); err != nil {
		return err
	}

	teststuff(db)

	return nil
}

func Start() {
	app := App{
		Name:    "api-greenforum-staging.dasith.works",
		Version: "1.0.0",
	}

	if err := app.Run(); err != nil {
		log.Error(err)
	}
}

func teststuff(db *gorm.DB) {
	commonwords := [100]string{"the", "of", "and", "a", "to", "in", "is", "you", "that", "it", "he", "was", "for", "on", "are", "as", "with", "his", "they", "I", "at", "be", "this", "have", "from", "or", "one", "had", "by", "word", "but", "not", "what", "all", "were", "we", "when", "your", "can", "said", "there", "use", "an", "each", "which", "she", "do", "how", "their", "if", "will", "up", "other", "about", "out", "many", "then", "them", "these", "so", "some", "her", "would", "make", "like", "him", "into", "time", "has", "look", "two", "more", "write", "go", "see", "number", "no", "way", "could", "people", "my", "than", "first", "water", "been", "call", "who", "oil", "its", "now", "find", "long", "down", "day", "did", "get", "come", "made", "may", "part"}

	vals := []string{}

	var u user.User
	db.Preload("Interests").Where("id =?", 2).First(&u)
	in := []string{}

	for _, interest := range u.Interests {
		in = append(in, interest.Word)
	}

	dup := dup_count(in)

	p := make(PairList, len(dup))

	i := 0
	for k, v := range dup {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p) // p is sorted

	fmt.Println(in)
	fmt.Println(p)

	for _, value := range p {
		if !itemExists(commonwords, value.Key) {

			vals = append(vals, value.Key)
		}
	}

	fmt.Println("=============== vals finally ==============")
	fmt.Println(vals, u.ID)

	GetRelatedQuestions(db, vals)

}

func GetRelatedQuestions(db *gorm.DB, vals []string) {
	var questions []question.Question
	db.Debug().Preload("Related", "Word in ? ", vals).Find(&questions)
	fmt.Println(questions)
}

func dup_count(list []string) map[string]int {
	duplicate_frequency := make(map[string]int)
	for _, item := range list {
		_, exist := duplicate_frequency[item]

		if exist {
			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicate_frequency[item] = 1 // else start counting from 1
		}
	}
	return duplicate_frequency
}

func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func RemoveIndex(s []Pair, index int) []Pair {
	return append(s[:index], s[index+1:]...)
}
