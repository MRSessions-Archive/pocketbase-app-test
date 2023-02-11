package migrations

import (
	"github.com/pocketbase/dbx"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		// add up queries...
		// set Global Time Allotment Settings defaults
		query := db.NewQuery("INSERT INTO globalTimeAllotmentSettings (id, dayOfWeek, hour, minute) VALUES (1, 'Monday', 0, 0), (2, 'Tuesday', 0, 0), (3, 'Wednesday', 0, 0), (4, 'Thursday', 0, 0), (5, 'Friday', 0, 0), (6, 'Saturday', 0, 0), (7, 'Sunday', 0, 0)")
		if _, err := query.Execute(); err != nil {
			return err
		}

		return nil
	}, func(db dbx.Builder) error {
		// add down queries...
		// delete Global Time Allotment Settings defaults
		query := db.NewQuery("DELETE FROM globalTimeAllotmentSettings WHERE id IN (1, 2, 3, 4, 5, 6, 7)")
		if _, err := query.Execute(); err != nil {
			return err
		}

		return nil
	})
}
