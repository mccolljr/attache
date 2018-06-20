package attache

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
)

type Storable interface {
	Table() string

	Insert() (columns []string, values []interface{})
	Update() (columns []string, values []interface{})

	Select() (columns []string, into []interface{})

	KeyColumns() []string
	KeyValues() []interface{}
}

type (
	BeforeInserter interface{ BeforeInsert() error }
	AfterInserter  interface{ AfterInsert(sql.Result) }
)

type (
	BeforeUpdater interface{ BeforeUpdate() (err error) }
	AfterUpdater  interface{ AfterUpdate(sql.Result) }
)

type (
	BeforeDeleter interface{ BeforeDelete() (err error) }
	AfterDeleter  interface{ AfterDelete(sql.Result) }
)

var (
	tStorable = reflect.TypeOf((*Storable)(nil)).Elem()
)

type DB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)

	Insert(s Storable) error
	Update(s Storable) error
	Delete(s Storable) error
	All(func() Storable) ([]Storable, error)
	Find(into Storable, args ...interface{}) error
	FindBy(into Storable, field string, val interface{}) error

	private()
}

type db struct {
	conn *sql.DB
}

func (d db) private() {}

func (d db) Exec(query string, args ...interface{}) (sql.Result, error) {
	sqlLog(query, args)
	return d.conn.Exec(query, args...)
}

func (d db) Query(query string, args ...interface{}) (*sql.Rows, error) {
	sqlLog(query, args)
	return d.conn.Query(query, args...)
}

func (d db) Insert(s Storable) error {
	if impl, ok := s.(BeforeInserter); ok {
		if err := impl.BeforeInsert(); err != nil {
			return err
		}
	}

	cols, vals := s.Insert()

	query := new(bytes.Buffer)
	fmt.Fprintf(query, "INSERT INTO %s (", s.Table())
	for i, name := range cols {
		if i != 0 {
			query.WriteString(", ")
		}

		fmt.Fprintf(query, "%s", name)
	}

	query.WriteString(") VALUES (")
	for i := range vals {
		if i != 0 {
			query.WriteString(", ")
		}
		query.WriteString("?")
	}
	query.WriteString(")")

	result, err := d.Exec(query.String(), vals...)
	if err != nil {
		return err
	}

	if impl, ok := s.(AfterInserter); ok {
		impl.AfterInsert(result)
	}

	return nil
}

func (d db) Update(s Storable) error {
	if impl, ok := s.(BeforeUpdater); ok {
		if err := impl.BeforeUpdate(); err != nil {
			return err
		}
	}

	cols, vals := s.Update()
	query := new(bytes.Buffer)
	fmt.Fprintf(query, "UPDATE %s SET ", s.Table())
	for i, name := range cols {
		if i != 0 {
			query.WriteString(", ")
		}

		fmt.Fprintf(query, "%s=?", name)
	}
	query.WriteString(" WHERE ")
	for i, name := range s.KeyColumns() {
		if i != 0 {
			query.WriteString(" AND ")
		}

		fmt.Fprintf(query, "%s=?", name)
	}

	result, err := d.Exec(query.String(), append(vals, s.KeyValues()...)...)
	if err != nil {
		return err
	}

	if impl, ok := s.(AfterUpdater); ok {
		impl.AfterUpdate(result)
	}

	return nil
}

func (d db) Delete(s Storable) error {
	if impl, ok := s.(BeforeDeleter); ok {
		if err := impl.BeforeDelete(); err != nil {
			return err
		}
	}

	query := new(bytes.Buffer)
	fmt.Fprintf(query, "DELETE FROM %s WHERE ", s.Table())

	for i, name := range s.KeyColumns() {
		if i != 0 {
			query.WriteString(" AND ")
		}

		fmt.Fprintf(query, "%s=?", name)
	}

	result, err := d.Exec(query.String(), s.KeyValues()...)
	if err != nil {
		return err
	}

	if impl, ok := s.(AfterDeleter); ok {
		impl.AfterDelete(result)
	}

	return nil
}

func (d db) All(newFn func() Storable) ([]Storable, error) {
	storable := newFn()
	cols, _ := storable.Select()
	result := make([]Storable, 0, 64)
	query := selectQuery(cols, storable.Table(), nil, false, 0)
	rows, err := d.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		into := newFn()
		_, targs := into.Select()
		if err := rows.Scan(targs...); err != nil {
			return nil, err
		}
		result = append(result, into)
	}

	if len(result) == 0 {
		return nil, sql.ErrNoRows
	}

	return result, nil
}

func (d db) Find(into Storable, args ...interface{}) error {
	cols, targets := into.Select()
	query := selectQuery(cols, into.Table(), into.KeyColumns(), false, 1)
	rows, err := d.Query(query, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	if !rows.Next() {
		return sql.ErrNoRows
	}

	return rows.Scan(targets...)
}

func (d db) FindBy(into Storable, field string, val interface{}) error {
	cols, targets := into.Select()
	query := selectQuery(cols, into.Table(), []string{field}, false, 1)
	rows, err := d.Query(query, val)
	if err != nil {
		return err
	}

	defer rows.Close()

	if !rows.Next() {
		return sql.ErrNoRows
	}

	return rows.Scan(targets...)
}

func selectQuery(cols []string, table string, searchFields []string, or bool, limit int) string {
	query := new(strings.Builder)
	query.WriteString("SELECT ")
	for i, name := range cols {
		if i != 0 {
			query.WriteString(", ")
		}

		query.WriteString(name)
	}

	query.WriteString(" FROM ")
	query.WriteString(table)
	if len(searchFields) > 0 {
		query.WriteString(" WHERE ")
		for i, field := range searchFields {
			if i != 0 {
				if or {
					query.WriteString(" OR ")
				} else {
					query.WriteString(" AND ")
				}
			}

			fmt.Fprintf(query, "%s=?", field)
		}
	}

	if limit > 0 {
		fmt.Fprintf(query, " LIMIT %d", limit)
	}

	query.WriteByte(';')
	return query.String()
}

func sqlLog(query string, args []interface{}) {
	log.Println("SQL:", query, args)
}
