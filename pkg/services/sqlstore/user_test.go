package sqlstore

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/grafana/grafana/pkg/models"
)

func TestUserDataAccess(t *testing.T) {

	Convey("Testing DB", t, func() {
		InitTestDB(t)

		var err error
		for i := 0; i < 8; i++ {
			err = CreateUser(&models.CreateUserCommand{
				Email: fmt.Sprint("user", i, "@test.com"),
				Name:  fmt.Sprint("user", i),
				Login: fmt.Sprint("user", i),
			})
		}

		So(err, ShouldBeNil)

		Convey("Can return the first page of users and a total count", func() {
			query := models.SearchUsersQuery{Query: "", Page: 0, Limit: 3}
			err = SearchUsers(&query)

			So(err, ShouldBeNil)
			So(query.Result[0].Email, ShouldEqual, "user0@test.com")
			So(query.Result[2].Email, ShouldEqual, "user7@test.com")
			So(query.TotalCount, ShouldEqual, 8)
		})
	})
}
