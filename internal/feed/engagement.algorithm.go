package feed

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/z9fr/greensforum-backend/internal/question"
	"github.com/z9fr/greensforum-backend/internal/user"
	"github.com/z9fr/greensforum-backend/internal/utils"
)

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

func (s *Service) GetUserInterestedQuestionsEngagement(u user.User) []question.Question {
	commonwords := [100]string{"the", "of", "and", "a", "to", "in", "is", "you", "that", "it", "he", "was", "for", "on", "are", "as", "with", "his", "they", "I", "at", "be", "this", "have", "from", "or", "one", "had", "by", "word", "but", "not", "what", "all", "were", "we", "when", "your", "can", "said", "there", "use", "an", "each", "which", "she", "do", "how", "their", "if", "will", "up", "other", "about", "out", "many", "then", "them", "these", "so", "some", "her", "would", "make", "like", "him", "into", "time", "has", "look", "two", "more", "write", "go", "see", "number", "no", "way", "could", "people", "my", "than", "first", "water", "been", "call", "who", "oil", "its", "now", "find", "long", "down", "day", "did", "get", "come", "made", "may", "part"}
	vals := []string{}
	var questions []question.Question

	var userwithinterests user.User
	s.DB.Debug().Preload("Interests").Where("id = ?", u.ID).Find(&userwithinterests)

	in := []string{}

	for _, interest := range userwithinterests.Interests {
		in = append(in, interest.Word)
	}

	dup := dup_count(in)

	pairlist := make(PairList, len(dup))

	i := 0
	for k, v := range dup {
		pairlist[i] = Pair{k, v}
		i++
	}

	sort.Sort(pairlist) // pair list is sorted

	for _, value := range pairlist {
		if !itemExists(commonwords, value.Key) {
			vals = append(vals, value.Key)
		}
	}

	fmt.Println("=============== vals finally ==============")
	fmt.Println(vals, u.ID)

	s.DB.Debug().Preload("Related", "Word in ? ", vals).Preload("Tags").Order("id DESC").Find(&questions)

	return questions
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
		utils.LogWarn("Invalid data-type")
		return false
		// panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}
