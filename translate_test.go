package yandex_translate

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

const API_KEY = "trnsl.1.1.20130906T224742Z.773de87520381874.34176f81082c2758819298377d038a2b94a0c8d0"

func TestTranslateAPI(t *testing.T) {
	Convey("#GetLangs", t, func() {
		tr := New(API_KEY)

		Convey("On success it returns available codes and languages", func() {
			response, err := tr.GetLangs("en")
			So(err, ShouldBeNil)
			So(response.Dirs, ShouldNotBeEmpty)
			pairs := []string{"ru-en", "ru-de", "ru-it", "ru-fr", "en-de", "en-it", "en-ru"}
			for _, pair := range pairs {
				So(response.Dirs, ShouldContain, pair)
			}
			So(response.Langs, ShouldNotBeNil)
			So(response.Langs["en"], ShouldEqual, "English")
			So(response.Langs["ru"], ShouldEqual, "Russian")
			So(response.Langs["mumbayumba"], ShouldBeBlank)
		})

		Convey("On failure it returns error code and message", func() {
			tr := New(API_KEY + "a")
			response, err := tr.GetLangs("en")
			So(response, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "(401) API key is invalid")
		})
	})

	Convey("#Translate", t, func() {
		tr := New(API_KEY)

		Convey("On success it returns translation of the word", func() {
			response, err := tr.Translate("ru", "A lazy dog")

			So(err, ShouldBeNil)
			So(response, ShouldNotBeNil)

			So(response.Code, ShouldEqual, 200)
			So(response.Message, ShouldBeBlank)

			So(response.Lang, ShouldEqual, "en-ru")
			So(response.Detected["lang"], ShouldEqual, "en")
			So(response.Text, ShouldContain, "Ленивая собака")
		})

		Convey("On fail it returns an error with error code and message", func() {
			response, err := tr.Translate("mumba-yumba", "A lazy dog")

			So(err, ShouldNotBeNil)
			So(response, ShouldBeNil)
			So(err.Error(), ShouldEqual, "(501) The specified translation direction is not supported")
		})
	})

	Convey("#Result", t, func() {
		tr := New(API_KEY)
		response, _ := tr.Translate("ru", "A lazy dog")
		So(response.Result(), ShouldEqual, "Ленивая собака")
	})

}
