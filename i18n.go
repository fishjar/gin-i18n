package i18n

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"strings"
	"sync"

	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

type message map[string]string

type localize struct {
	defaultLang  language.Tag
	supportLangs []language.Tag
	messages     map[language.Tag]message
}

func (l localize) matchUserLang(s string) language.Tag {
	langs, _, err := language.ParseAcceptLanguage(s)
	if err != nil {
		return l.defaultLang
	}

	for _, lang := range langs {
		for _, support := range l.supportLangs {
			if lang == support {
				return lang
			}
		}
	}

	return l.defaultLang
}

func newLocalize(defaultLang language.Tag, supportLangs []language.Tag, filePath string) *localize {
	if !tagContains(supportLangs, defaultLang) {
		log.Fatal("supportLangs must contains defaultLang")
	}

	messages := make(map[language.Tag]message, len(supportLangs))
	for _, lang := range supportLangs {
		var message message
		f, err := ioutil.ReadFile(path.Join(filePath, fmt.Sprintf("%s.yml", lang)))
		if err != nil {
			log.Fatal("read lang file: %w", err)
		}
		if err := yaml.Unmarshal(f, &message); err != nil {
			log.Fatal("unmarshal lang file: %w", err)
		}
		messages[lang] = message
	}

	return &localize{defaultLang, supportLangs, messages}
}

func tagContains(tags []language.Tag, t language.Tag) bool {
	for _, tag := range tags {
		if tag == t {
			return true
		}
	}
	return false
}

var localizer *localize
var once sync.Once

func LocalizerInit(defaultLang string, supportLang string, filePath string) {
	once.Do(func() {
		var supportLangs []language.Tag
		for _, s := range strings.Split(supportLang, ",") {
			supportLangs = append(supportLangs, language.Make(strings.TrimSpace(s)))
		}
		localizer = newLocalize(language.Make(defaultLang), supportLangs, filePath)
	})
}

type UserLocalize struct {
	localize *localize
	userLang language.Tag
}

func (u *UserLocalize) GetMsg(messageID string, a ...interface{}) string {
	var format string

	if u.localize == nil {
		return format
	}

	if s, ok := u.localize.messages[u.userLang][messageID]; ok {
		format = s
	} else {
		format = u.localize.messages[u.localize.defaultLang][messageID]
	}
	if len(format) > 0 && len(a) > 0 {
		return fmt.Sprintf(format, a...)
	}

	return format
}

func NewUserLocalize(acceptLang string) *UserLocalize {
	userLang := localizer.matchUserLang(acceptLang)
	return &UserLocalize{localizer, userLang}
}
