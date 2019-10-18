// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bytes"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// UserAccount is the model entity for the UserAccount schema.
type UserAccount struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Passwd holds the value of the "passwd" field.
	Passwd string `json:"passwd,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// FromRows scans the sql response data into UserAccount.
func (ua *UserAccount) FromRows(rows *sql.Rows) error {
	var vua struct {
		ID        int
		Name      sql.NullString
		Passwd    sql.NullString
		Email     sql.NullString
		CreatedAt sql.NullTime
	}
	// the order here should be the same as in the `useraccount.Columns`.
	if err := rows.Scan(
		&vua.ID,
		&vua.Name,
		&vua.Passwd,
		&vua.Email,
		&vua.CreatedAt,
	); err != nil {
		return err
	}
	ua.ID = vua.ID
	ua.Name = vua.Name.String
	ua.Passwd = vua.Passwd.String
	ua.Email = vua.Email.String
	ua.CreatedAt = vua.CreatedAt.Time
	return nil
}

// QueryOwner queries the owner edge of the UserAccount.
func (ua *UserAccount) QueryOwner() *UserQuery {
	return (&UserAccountClient{ua.config}).QueryOwner(ua)
}

// Update returns a builder for updating this UserAccount.
// Note that, you need to call UserAccount.Unwrap() before calling this method, if this UserAccount
// was returned from a transaction, and the transaction was committed or rolled back.
func (ua *UserAccount) Update() *UserAccountUpdateOne {
	return (&UserAccountClient{ua.config}).UpdateOne(ua)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (ua *UserAccount) Unwrap() *UserAccount {
	tx, ok := ua.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserAccount is not a transactional entity")
	}
	ua.config.driver = tx.drv
	return ua
}

// String implements the fmt.Stringer.
func (ua *UserAccount) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("UserAccount(")
	buf.WriteString(fmt.Sprintf("id=%v", ua.ID))
	buf.WriteString(fmt.Sprintf(", name=%v", ua.Name))
	buf.WriteString(fmt.Sprintf(", passwd=%v", ua.Passwd))
	buf.WriteString(fmt.Sprintf(", email=%v", ua.Email))
	buf.WriteString(fmt.Sprintf(", createdAt=%v", ua.CreatedAt))
	buf.WriteString(")")
	return buf.String()
}

// UserAccounts is a parsable slice of UserAccount.
type UserAccounts []*UserAccount

// FromRows scans the sql response data into UserAccounts.
func (ua *UserAccounts) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		vua := &UserAccount{}
		if err := vua.FromRows(rows); err != nil {
			return err
		}
		*ua = append(*ua, vua)
	}
	return nil
}

func (ua UserAccounts) config(cfg config) {
	for _i := range ua {
		ua[_i].config = cfg
	}
}
