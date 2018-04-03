// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
	"gopkg.in/volatiletech/null.v6"
)

// Support is an object representing the database table.
type Support struct {
	ID               uint64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	SupportedClaimID string      `boil:"supported_claim_id" json:"supported_claim_id" toml:"supported_claim_id" yaml:"supported_claim_id"`
	SupportAmount    float64     `boil:"support_amount" json:"support_amount" toml:"support_amount" yaml:"support_amount"`
	BidState         string      `boil:"bid_state" json:"bid_state" toml:"bid_state" yaml:"bid_state"`
	TransactionHash  null.String `boil:"transaction_hash" json:"transaction_hash,omitempty" toml:"transaction_hash" yaml:"transaction_hash,omitempty"`
	Vout             uint        `boil:"vout" json:"vout" toml:"vout" yaml:"vout"`

	R *supportR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L supportL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SupportColumns = struct {
	ID               string
	SupportedClaimID string
	SupportAmount    string
	BidState         string
	TransactionHash  string
	Vout             string
}{
	ID:               "id",
	SupportedClaimID: "supported_claim_id",
	SupportAmount:    "support_amount",
	BidState:         "bid_state",
	TransactionHash:  "transaction_hash",
	Vout:             "vout",
}

// supportR is where relationships are stored.
type supportR struct {
	SupportedClaim *Claim
}

// supportL is where Load methods for each relationship are stored.
type supportL struct{}

var (
	supportColumns               = []string{"id", "supported_claim_id", "support_amount", "bid_state", "transaction_hash", "vout"}
	supportColumnsWithoutDefault = []string{"supported_claim_id", "transaction_hash", "vout"}
	supportColumnsWithDefault    = []string{"id", "support_amount", "bid_state"}
	supportPrimaryKeyColumns     = []string{"id"}
)

type (
	// SupportSlice is an alias for a slice of pointers to Support.
	// This should generally be used opposed to []Support.
	SupportSlice []*Support

	supportQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	supportType                 = reflect.TypeOf(&Support{})
	supportMapping              = queries.MakeStructMapping(supportType)
	supportPrimaryKeyMapping, _ = queries.BindMapping(supportType, supportMapping, supportPrimaryKeyColumns)
	supportInsertCacheMut       sync.RWMutex
	supportInsertCache          = make(map[string]insertCache)
	supportUpdateCacheMut       sync.RWMutex
	supportUpdateCache          = make(map[string]updateCache)
	supportUpsertCacheMut       sync.RWMutex
	supportUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single support record from the query, and panics on error.
func (q supportQuery) OneP() *Support {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single support record from the query.
func (q supportQuery) One() (*Support, error) {
	o := &Support{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for support")
	}

	return o, nil
}

// AllP returns all Support records from the query, and panics on error.
func (q supportQuery) AllP() SupportSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Support records from the query.
func (q supportQuery) All() (SupportSlice, error) {
	var o []*Support

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to Support slice")
	}

	return o, nil
}

// CountP returns the count of all Support records in the query, and panics on error.
func (q supportQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Support records in the query.
func (q supportQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count support rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q supportQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q supportQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if support exists")
	}

	return count > 0, nil
}

// SupportedClaimG pointed to by the foreign key.
func (o *Support) SupportedClaimG(mods ...qm.QueryMod) claimQuery {
	return o.SupportedClaim(boil.GetDB(), mods...)
}

// SupportedClaim pointed to by the foreign key.
func (o *Support) SupportedClaim(exec boil.Executor, mods ...qm.QueryMod) claimQuery {
	queryMods := []qm.QueryMod{
		qm.Where("claim_id=?", o.SupportedClaimID),
	}

	queryMods = append(queryMods, mods...)

	query := Claims(exec, queryMods...)
	queries.SetFrom(query.Query, "`claim`")

	return query
} // LoadSupportedClaim allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (supportL) LoadSupportedClaim(e boil.Executor, singular bool, maybeSupport interface{}) error {
	var slice []*Support
	var object *Support

	count := 1
	if singular {
		object = maybeSupport.(*Support)
	} else {
		slice = *maybeSupport.(*[]*Support)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &supportR{}
		}
		args[0] = object.SupportedClaimID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &supportR{}
			}
			args[i] = obj.SupportedClaimID
		}
	}

	query := fmt.Sprintf(
		"select * from `claim` where `claim_id` in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Claim")
	}
	defer results.Close()

	var resultSlice []*Claim
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Claim")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		object.R.SupportedClaim = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.SupportedClaimID == foreign.ClaimID {
				local.R.SupportedClaim = foreign
				break
			}
		}
	}

	return nil
}

// SetSupportedClaimG of the support to the related item.
// Sets o.R.SupportedClaim to related.
// Adds o to related.R.SupportedClaimSupports.
// Uses the global database handle.
func (o *Support) SetSupportedClaimG(insert bool, related *Claim) error {
	return o.SetSupportedClaim(boil.GetDB(), insert, related)
}

// SetSupportedClaimP of the support to the related item.
// Sets o.R.SupportedClaim to related.
// Adds o to related.R.SupportedClaimSupports.
// Panics on error.
func (o *Support) SetSupportedClaimP(exec boil.Executor, insert bool, related *Claim) {
	if err := o.SetSupportedClaim(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetSupportedClaimGP of the support to the related item.
// Sets o.R.SupportedClaim to related.
// Adds o to related.R.SupportedClaimSupports.
// Uses the global database handle and panics on error.
func (o *Support) SetSupportedClaimGP(insert bool, related *Claim) {
	if err := o.SetSupportedClaim(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetSupportedClaim of the support to the related item.
// Sets o.R.SupportedClaim to related.
// Adds o to related.R.SupportedClaimSupports.
func (o *Support) SetSupportedClaim(exec boil.Executor, insert bool, related *Claim) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `support` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"supported_claim_id"}),
		strmangle.WhereClause("`", "`", 0, supportPrimaryKeyColumns),
	)
	values := []interface{}{related.ClaimID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.SupportedClaimID = related.ClaimID

	if o.R == nil {
		o.R = &supportR{
			SupportedClaim: related,
		}
	} else {
		o.R.SupportedClaim = related
	}

	if related.R == nil {
		related.R = &claimR{
			SupportedClaimSupports: SupportSlice{o},
		}
	} else {
		related.R.SupportedClaimSupports = append(related.R.SupportedClaimSupports, o)
	}

	return nil
}

// SupportsG retrieves all records.
func SupportsG(mods ...qm.QueryMod) supportQuery {
	return Supports(boil.GetDB(), mods...)
}

// Supports retrieves all the records using an executor.
func Supports(exec boil.Executor, mods ...qm.QueryMod) supportQuery {
	mods = append(mods, qm.From("`support`"))
	return supportQuery{NewQuery(exec, mods...)}
}

// FindSupportG retrieves a single record by ID.
func FindSupportG(id uint64, selectCols ...string) (*Support, error) {
	return FindSupport(boil.GetDB(), id, selectCols...)
}

// FindSupportGP retrieves a single record by ID, and panics on error.
func FindSupportGP(id uint64, selectCols ...string) *Support {
	retobj, err := FindSupport(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindSupport retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSupport(exec boil.Executor, id uint64, selectCols ...string) (*Support, error) {
	supportObj := &Support{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `support` where `id`=?", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(supportObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from support")
	}

	return supportObj, nil
}

// FindSupportP retrieves a single record by ID with an executor, and panics on error.
func FindSupportP(exec boil.Executor, id uint64, selectCols ...string) *Support {
	retobj, err := FindSupport(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Support) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Support) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Support) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Support) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("model: no support provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(supportColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	supportInsertCacheMut.RLock()
	cache, cached := supportInsertCache[key]
	supportInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			supportColumns,
			supportColumnsWithDefault,
			supportColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(supportType, supportMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(supportType, supportMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `support` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `support` () VALUES ()"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `support` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, supportPrimaryKeyColumns))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to insert into support")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == supportMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for support")
	}

CacheNoHooks:
	if !cached {
		supportInsertCacheMut.Lock()
		supportInsertCache[key] = cache
		supportInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single Support record. See Update for
// whitelist behavior description.
func (o *Support) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Support record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Support) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Support, and panics on error.
// See Update for whitelist behavior description.
func (o *Support) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Support.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Support) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	supportUpdateCacheMut.RLock()
	cache, cached := supportUpdateCache[key]
	supportUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			supportColumns,
			supportPrimaryKeyColumns,
			whitelist,
		)

		if len(wl) == 0 {
			return errors.New("model: unable to update support, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `support` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, supportPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(supportType, supportMapping, append(wl, supportPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "model: unable to update support row")
	}

	if !cached {
		supportUpdateCacheMut.Lock()
		supportUpdateCache[key] = cache
		supportUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q supportQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q supportQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "model: unable to update all for support")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o SupportSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o SupportSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o SupportSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SupportSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("model: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), supportPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `support` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, supportPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to update all in support slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Support) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Support) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Support) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Support) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("model: no support provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(supportColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	supportUpsertCacheMut.RLock()
	cache, cached := supportUpsertCache[key]
	supportUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			supportColumns,
			supportColumnsWithDefault,
			supportColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			supportColumns,
			supportPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("model: unable to upsert support, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "support", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `support` WHERE `id`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(supportType, supportMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(supportType, supportMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to upsert for support")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == supportMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for support")
	}

CacheNoHooks:
	if !cached {
		supportUpsertCacheMut.Lock()
		supportUpsertCache[key] = cache
		supportUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single Support record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Support) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Support record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Support) DeleteG() error {
	if o == nil {
		return errors.New("model: no Support provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Support record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Support) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Support record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Support) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("model: no Support provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), supportPrimaryKeyMapping)
	sql := "DELETE FROM `support` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to delete from support")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q supportQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q supportQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("model: no supportQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "model: unable to delete all from support")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o SupportSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o SupportSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("model: no Support slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o SupportSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SupportSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("model: no Support slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), supportPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `support` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, supportPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to delete all from support slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Support) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Support) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Support) ReloadG() error {
	if o == nil {
		return errors.New("model: no Support provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Support) Reload(exec boil.Executor) error {
	ret, err := FindSupport(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *SupportSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *SupportSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SupportSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("model: empty SupportSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SupportSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	supports := SupportSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), supportPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `support`.* FROM `support` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, supportPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&supports)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in SupportSlice")
	}

	*o = supports

	return nil
}

// SupportExists checks if the Support row exists.
func SupportExists(exec boil.Executor, id uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `support` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if support exists")
	}

	return exists, nil
}

// SupportExistsG checks if the Support row exists.
func SupportExistsG(id uint64) (bool, error) {
	return SupportExists(boil.GetDB(), id)
}

// SupportExistsGP checks if the Support row exists. Panics on error.
func SupportExistsGP(id uint64) bool {
	e, err := SupportExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// SupportExistsP checks if the Support row exists. Panics on error.
func SupportExistsP(exec boil.Executor, id uint64) bool {
	e, err := SupportExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}