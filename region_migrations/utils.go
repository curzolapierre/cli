package region_migrations

import (
	scalingo "github.com/Scalingo/go-scalingo"
	"github.com/fatih/color"
)

func formatMigrationStatus(status scalingo.RegionMigrationStatus) string {
	strStatus := string(status)
	switch status {
	case scalingo.RegionMigrationStatusScheduled:
		return color.BlueString(strStatus)
	case scalingo.RegionMigrationStatusRunning:
		return color.YellowString(strStatus)
	case scalingo.RegionMigrationStatusDone:
		return color.GreenString(strStatus)
	case scalingo.RegionMigrationStatusPreflightError:
		fallthrough
	case scalingo.RegionMigrationStatusError:
		return color.RedString(strStatus)
	}

	return color.BlueString(strStatus)
}
