package i18n

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

type message map[string]string

type localize struct {
	defaultLang  language.Tag
	supportLangs []language.Tag
	messages     map[language.Tag]message
}

func (l *localize) matchUserLang(acceptStr string) language.Tag {
	acceptLangs, _, err := language.ParseAcceptLanguage(acceptStr)
	if err != nil {
		return l.defaultLang
	}

	for _, acceptLang := range acceptLangs {
		for _, supportLang := range l.supportLangs {
			if strings.HasPrefix(supportLang.String(), acceptLang.String()) {
				return supportLang
			}
		}
	}

	return l.defaultLang
}

func (l *localize) userLocalize(acceptStr string) *userLocalize {
	userLang := l.matchUserLang(acceptStr)

	return &userLocalize{
		l,
		userLang,
	}
}

func newLocalize(defaultStr, supportStr, filePath string) *localize {
	defaultLang := language.Make(defaultStr)
	var supportLangs []language.Tag
	for _, s := range strings.Split(supportStr, ",") {
		supportLangs = append(supportLangs, language.Make(strings.TrimSpace(s)))
	}
	if !tagContains(supportLangs, defaultLang) {
		panic("supportLangs must contains defaultLang")
	}

	messages := make(map[language.Tag]message, len(supportLangs))
	for _, lang := range supportLangs {
		var message message
		f, err := ioutil.ReadFile(path.Join(filePath, fmt.Sprintf("%s.yml", lang)))
		if err != nil {
			panic(err)
		}
		if err := yaml.Unmarshal(f, &message); err != nil {
			panic(err)
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

type userLocalize struct {
	l        *localize
	userLang language.Tag
}

func (u *userLocalize) getMsg(tag string, args ...interface{}) (msg string) {
	if u.l == nil {
		return
	}

	if s, ok := u.l.messages[u.userLang][tag]; ok {
		msg = s
	} else {
		msg = u.l.messages[u.l.defaultLang][tag]
	}
	if len(msg) > 0 && len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}

	return
}
