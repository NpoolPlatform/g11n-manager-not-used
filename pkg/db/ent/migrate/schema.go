// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AppCountriesColumns holds the columns for the "app_countries" table.
	AppCountriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "country_id", Type: field.TypeUUID, Nullable: true},
	}
	// AppCountriesTable holds the schema information for the "app_countries" table.
	AppCountriesTable = &schema.Table{
		Name:       "app_countries",
		Columns:    AppCountriesColumns,
		PrimaryKey: []*schema.Column{AppCountriesColumns[0]},
	}
	// AppLangsColumns holds the columns for the "app_langs" table.
	AppLangsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "lang_id", Type: field.TypeUUID, Nullable: true},
		{Name: "main", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// AppLangsTable holds the schema information for the "app_langs" table.
	AppLangsTable = &schema.Table{
		Name:       "app_langs",
		Columns:    AppLangsColumns,
		PrimaryKey: []*schema.Column{AppLangsColumns[0]},
	}
	// CountriesColumns holds the columns for the "countries" table.
	CountriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "country", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "flag", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "code", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "short", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// CountriesTable holds the schema information for the "countries" table.
	CountriesTable = &schema.Table{
		Name:       "countries",
		Columns:    CountriesColumns,
		PrimaryKey: []*schema.Column{CountriesColumns[0]},
	}
	// LangsColumns holds the columns for the "langs" table.
	LangsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "lang", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "logo", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "short", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// LangsTable holds the schema information for the "langs" table.
	LangsTable = &schema.Table{
		Name:       "langs",
		Columns:    LangsColumns,
		PrimaryKey: []*schema.Column{LangsColumns[0]},
	}
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "lang_id", Type: field.TypeUUID, Nullable: true},
		{Name: "message_id", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "message", Type: field.TypeString, Nullable: true, Size: 16016, Default: ""},
		{Name: "get_index", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "disabled", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "short", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppCountriesTable,
		AppLangsTable,
		CountriesTable,
		LangsTable,
		MessagesTable,
	}
)

func init() {
}
