// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "last_login", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UserAccountsColumns holds the columns for the "user_accounts" table.
	UserAccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "passwd", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "owner_id", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// UserAccountsTable holds the schema information for the "user_accounts" table.
	UserAccountsTable = &schema.Table{
		Name:       "user_accounts",
		Columns:    UserAccountsColumns,
		PrimaryKey: []*schema.Column{UserAccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "user_accounts_users_account",
				Columns: []*schema.Column{UserAccountsColumns[5]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserFriendsColumns holds the columns for the "user_friends" table.
	UserFriendsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "friend_id", Type: field.TypeInt},
	}
	// UserFriendsTable holds the schema information for the "user_friends" table.
	UserFriendsTable = &schema.Table{
		Name:       "user_friends",
		Columns:    UserFriendsColumns,
		PrimaryKey: []*schema.Column{UserFriendsColumns[0], UserFriendsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "user_friends_user_id",
				Columns: []*schema.Column{UserFriendsColumns[0]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "user_friends_friend_id",
				Columns: []*schema.Column{UserFriendsColumns[1]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UsersTable,
		UserAccountsTable,
		UserFriendsTable,
	}
)

func init() {
	UserAccountsTable.ForeignKeys[0].RefTable = UsersTable
	UserFriendsTable.ForeignKeys[0].RefTable = UsersTable
	UserFriendsTable.ForeignKeys[1].RefTable = UsersTable
}
