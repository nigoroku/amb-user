// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// ShareToken is an object representing the database table.
type ShareToken struct {
	ShareTokenID int       `boil:"share_token_id" json:"share_token_id" toml:"share_token_id" yaml:"share_token_id"`
	UserID       int       `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Token        string    `boil:"token" json:"token" toml:"token" yaml:"token"`
	CreatedBy    int       `boil:"created_by" json:"created_by" toml:"created_by" yaml:"created_by"`
	CreatedAt    time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	ModifiedBy   null.Int  `boil:"modified_by" json:"modified_by,omitempty" toml:"modified_by" yaml:"modified_by,omitempty"`
	ModifiedAt   null.Time `boil:"modified_at" json:"modified_at,omitempty" toml:"modified_at" yaml:"modified_at,omitempty"`

	R *shareTokenR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L shareTokenL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ShareTokenColumns = struct {
	ShareTokenID string
	UserID       string
	Token        string
	CreatedBy    string
	CreatedAt    string
	ModifiedBy   string
	ModifiedAt   string
}{
	ShareTokenID: "share_token_id",
	UserID:       "user_id",
	Token:        "token",
	CreatedBy:    "created_by",
	CreatedAt:    "created_at",
	ModifiedBy:   "modified_by",
	ModifiedAt:   "modified_at",
}

// Generated where

var ShareTokenWhere = struct {
	ShareTokenID whereHelperint
	UserID       whereHelperint
	Token        whereHelperstring
	CreatedBy    whereHelperint
	CreatedAt    whereHelpertime_Time
	ModifiedBy   whereHelpernull_Int
	ModifiedAt   whereHelpernull_Time
}{
	ShareTokenID: whereHelperint{field: "`share_tokens`.`share_token_id`"},
	UserID:       whereHelperint{field: "`share_tokens`.`user_id`"},
	Token:        whereHelperstring{field: "`share_tokens`.`token`"},
	CreatedBy:    whereHelperint{field: "`share_tokens`.`created_by`"},
	CreatedAt:    whereHelpertime_Time{field: "`share_tokens`.`created_at`"},
	ModifiedBy:   whereHelpernull_Int{field: "`share_tokens`.`modified_by`"},
	ModifiedAt:   whereHelpernull_Time{field: "`share_tokens`.`modified_at`"},
}

// ShareTokenRels is where relationship names are stored.
var ShareTokenRels = struct {
	User string
}{
	User: "User",
}

// shareTokenR is where relationships are stored.
type shareTokenR struct {
	User *User
}

// NewStruct creates a new relationship struct
func (*shareTokenR) NewStruct() *shareTokenR {
	return &shareTokenR{}
}

// shareTokenL is where Load methods for each relationship are stored.
type shareTokenL struct{}

var (
	shareTokenAllColumns            = []string{"share_token_id", "user_id", "token", "created_by", "created_at", "modified_by", "modified_at"}
	shareTokenColumnsWithoutDefault = []string{"user_id", "token", "created_by", "created_at", "modified_by", "modified_at"}
	shareTokenColumnsWithDefault    = []string{"share_token_id"}
	shareTokenPrimaryKeyColumns     = []string{"share_token_id"}
)

type (
	// ShareTokenSlice is an alias for a slice of pointers to ShareToken.
	// This should generally be used opposed to []ShareToken.
	ShareTokenSlice []*ShareToken
	// ShareTokenHook is the signature for custom ShareToken hook methods
	ShareTokenHook func(context.Context, boil.ContextExecutor, *ShareToken) error

	shareTokenQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	shareTokenType                 = reflect.TypeOf(&ShareToken{})
	shareTokenMapping              = queries.MakeStructMapping(shareTokenType)
	shareTokenPrimaryKeyMapping, _ = queries.BindMapping(shareTokenType, shareTokenMapping, shareTokenPrimaryKeyColumns)
	shareTokenInsertCacheMut       sync.RWMutex
	shareTokenInsertCache          = make(map[string]insertCache)
	shareTokenUpdateCacheMut       sync.RWMutex
	shareTokenUpdateCache          = make(map[string]updateCache)
	shareTokenUpsertCacheMut       sync.RWMutex
	shareTokenUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var shareTokenBeforeInsertHooks []ShareTokenHook
var shareTokenBeforeUpdateHooks []ShareTokenHook
var shareTokenBeforeDeleteHooks []ShareTokenHook
var shareTokenBeforeUpsertHooks []ShareTokenHook

var shareTokenAfterInsertHooks []ShareTokenHook
var shareTokenAfterSelectHooks []ShareTokenHook
var shareTokenAfterUpdateHooks []ShareTokenHook
var shareTokenAfterDeleteHooks []ShareTokenHook
var shareTokenAfterUpsertHooks []ShareTokenHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ShareToken) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shareTokenBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ShareToken) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shareTokenBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ShareToken) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shareTokenBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ShareToken) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shareTokenBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ShareToken) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shareTokenAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ShareToken) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shareTokenAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ShareToken) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shareTokenAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ShareToken) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shareTokenAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ShareToken) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shareTokenAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddShareTokenHook registers your hook function for all future operations.
func AddShareTokenHook(hookPoint boil.HookPoint, shareTokenHook ShareTokenHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		shareTokenBeforeInsertHooks = append(shareTokenBeforeInsertHooks, shareTokenHook)
	case boil.BeforeUpdateHook:
		shareTokenBeforeUpdateHooks = append(shareTokenBeforeUpdateHooks, shareTokenHook)
	case boil.BeforeDeleteHook:
		shareTokenBeforeDeleteHooks = append(shareTokenBeforeDeleteHooks, shareTokenHook)
	case boil.BeforeUpsertHook:
		shareTokenBeforeUpsertHooks = append(shareTokenBeforeUpsertHooks, shareTokenHook)
	case boil.AfterInsertHook:
		shareTokenAfterInsertHooks = append(shareTokenAfterInsertHooks, shareTokenHook)
	case boil.AfterSelectHook:
		shareTokenAfterSelectHooks = append(shareTokenAfterSelectHooks, shareTokenHook)
	case boil.AfterUpdateHook:
		shareTokenAfterUpdateHooks = append(shareTokenAfterUpdateHooks, shareTokenHook)
	case boil.AfterDeleteHook:
		shareTokenAfterDeleteHooks = append(shareTokenAfterDeleteHooks, shareTokenHook)
	case boil.AfterUpsertHook:
		shareTokenAfterUpsertHooks = append(shareTokenAfterUpsertHooks, shareTokenHook)
	}
}

// One returns a single shareToken record from the query.
func (q shareTokenQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ShareToken, error) {
	o := &ShareToken{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for share_tokens")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ShareToken records from the query.
func (q shareTokenQuery) All(ctx context.Context, exec boil.ContextExecutor) (ShareTokenSlice, error) {
	var o []*ShareToken

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ShareToken slice")
	}

	if len(shareTokenAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ShareToken records in the query.
func (q shareTokenQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count share_tokens rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q shareTokenQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if share_tokens exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *ShareToken) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`user_id` = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "`users`")

	return query
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (shareTokenL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeShareToken interface{}, mods queries.Applicator) error {
	var slice []*ShareToken
	var object *ShareToken

	if singular {
		object = maybeShareToken.(*ShareToken)
	} else {
		slice = *maybeShareToken.(*[]*ShareToken)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &shareTokenR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &shareTokenR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`users`), qm.WhereIn(`users.user_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(shareTokenAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.ShareTokens = append(foreign.R.ShareTokens, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.UserID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.ShareTokens = append(foreign.R.ShareTokens, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the shareToken to the related item.
// Sets o.R.User to related.
// Adds o to related.R.ShareTokens.
func (o *ShareToken) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `share_tokens` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, shareTokenPrimaryKeyColumns),
	)
	values := []interface{}{related.UserID, o.ShareTokenID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.UserID
	if o.R == nil {
		o.R = &shareTokenR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			ShareTokens: ShareTokenSlice{o},
		}
	} else {
		related.R.ShareTokens = append(related.R.ShareTokens, o)
	}

	return nil
}

// ShareTokens retrieves all the records using an executor.
func ShareTokens(mods ...qm.QueryMod) shareTokenQuery {
	mods = append(mods, qm.From("`share_tokens`"))
	return shareTokenQuery{NewQuery(mods...)}
}

// FindShareToken retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindShareToken(ctx context.Context, exec boil.ContextExecutor, shareTokenID int, selectCols ...string) (*ShareToken, error) {
	shareTokenObj := &ShareToken{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `share_tokens` where `share_token_id`=?", sel,
	)

	q := queries.Raw(query, shareTokenID)

	err := q.Bind(ctx, exec, shareTokenObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from share_tokens")
	}

	return shareTokenObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ShareToken) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no share_tokens provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(shareTokenColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	shareTokenInsertCacheMut.RLock()
	cache, cached := shareTokenInsertCache[key]
	shareTokenInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			shareTokenAllColumns,
			shareTokenColumnsWithDefault,
			shareTokenColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(shareTokenType, shareTokenMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(shareTokenType, shareTokenMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `share_tokens` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `share_tokens` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `share_tokens` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, shareTokenPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into share_tokens")
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

	o.ShareTokenID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == shareTokenMapping["share_token_id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ShareTokenID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for share_tokens")
	}

CacheNoHooks:
	if !cached {
		shareTokenInsertCacheMut.Lock()
		shareTokenInsertCache[key] = cache
		shareTokenInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ShareToken.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ShareToken) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	shareTokenUpdateCacheMut.RLock()
	cache, cached := shareTokenUpdateCache[key]
	shareTokenUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			shareTokenAllColumns,
			shareTokenPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update share_tokens, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `share_tokens` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, shareTokenPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(shareTokenType, shareTokenMapping, append(wl, shareTokenPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update share_tokens row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for share_tokens")
	}

	if !cached {
		shareTokenUpdateCacheMut.Lock()
		shareTokenUpdateCache[key] = cache
		shareTokenUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q shareTokenQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for share_tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for share_tokens")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ShareTokenSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), shareTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `share_tokens` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, shareTokenPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in shareToken slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all shareToken")
	}
	return rowsAff, nil
}

var mySQLShareTokenUniqueColumns = []string{
	"share_token_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ShareToken) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no share_tokens provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(shareTokenColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLShareTokenUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	shareTokenUpsertCacheMut.RLock()
	cache, cached := shareTokenUpsertCache[key]
	shareTokenUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			shareTokenAllColumns,
			shareTokenColumnsWithDefault,
			shareTokenColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			shareTokenAllColumns,
			shareTokenPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert share_tokens, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "share_tokens", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `share_tokens` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(shareTokenType, shareTokenMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(shareTokenType, shareTokenMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for share_tokens")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ShareTokenID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == shareTokenMapping["share_token_id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(shareTokenType, shareTokenMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for share_tokens")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for share_tokens")
	}

CacheNoHooks:
	if !cached {
		shareTokenUpsertCacheMut.Lock()
		shareTokenUpsertCache[key] = cache
		shareTokenUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ShareToken record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ShareToken) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ShareToken provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), shareTokenPrimaryKeyMapping)
	sql := "DELETE FROM `share_tokens` WHERE `share_token_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from share_tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for share_tokens")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q shareTokenQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no shareTokenQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from share_tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for share_tokens")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ShareTokenSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(shareTokenBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), shareTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `share_tokens` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, shareTokenPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from shareToken slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for share_tokens")
	}

	if len(shareTokenAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ShareToken) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindShareToken(ctx, exec, o.ShareTokenID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ShareTokenSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ShareTokenSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), shareTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `share_tokens`.* FROM `share_tokens` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, shareTokenPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ShareTokenSlice")
	}

	*o = slice

	return nil
}

// ShareTokenExists checks if the ShareToken row exists.
func ShareTokenExists(ctx context.Context, exec boil.ContextExecutor, shareTokenID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `share_tokens` where `share_token_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, shareTokenID)
	}
	row := exec.QueryRowContext(ctx, sql, shareTokenID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if share_tokens exists")
	}

	return exists, nil
}
