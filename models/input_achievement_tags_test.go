// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testInputAchievementTags(t *testing.T) {
	t.Parallel()

	query := InputAchievementTags()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testInputAchievementTagsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InputAchievementTag{}
	if err = randomize.Struct(seed, o, inputAchievementTagDBTypes, true, inputAchievementTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InputAchievementTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := InputAchievementTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testInputAchievementTagsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InputAchievementTag{}
	if err = randomize.Struct(seed, o, inputAchievementTagDBTypes, true, inputAchievementTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InputAchievementTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := InputAchievementTags().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := InputAchievementTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testInputAchievementTagsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InputAchievementTag{}
	if err = randomize.Struct(seed, o, inputAchievementTagDBTypes, true, inputAchievementTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InputAchievementTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := InputAchievementTagSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := InputAchievementTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testInputAchievementTagsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InputAchievementTag{}
	if err = randomize.Struct(seed, o, inputAchievementTagDBTypes, true, inputAchievementTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InputAchievementTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := InputAchievementTagExists(ctx, tx, o.InputAchievementTagID)
	if err != nil {
		t.Errorf("Unable to check if InputAchievementTag exists: %s", err)
	}
	if !e {
		t.Errorf("Expected InputAchievementTagExists to return true, but got false.")
	}
}

func testInputAchievementTagsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InputAchievementTag{}
	if err = randomize.Struct(seed, o, inputAchievementTagDBTypes, true, inputAchievementTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InputAchievementTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	inputAchievementTagFound, err := FindInputAchievementTag(ctx, tx, o.InputAchievementTagID)
	if err != nil {
		t.Error(err)
	}

	if inputAchievementTagFound == nil {
		t.Error("want a record, got nil")
	}
}

func testInputAchievementTagsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InputAchievementTag{}
	if err = randomize.Struct(seed, o, inputAchievementTagDBTypes, true, inputAchievementTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InputAchievementTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = InputAchievementTags().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testInputAchievementTagsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InputAchievementTag{}
	if err = randomize.Struct(seed, o, inputAchievementTagDBTypes, true, inputAchievementTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InputAchievementTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := InputAchievementTags().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testInputAchievementTagsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	inputAchievementTagOne := &InputAchievementTag{}
	inputAchievementTagTwo := &InputAchievementTag{}
	if err = randomize.Struct(seed, inputAchievementTagOne, inputAchievementTagDBTypes, false, inputAchievementTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InputAchievementTag struct: %s", err)
	}
	if err = randomize.Struct(seed, inputAchievementTagTwo, inputAchievementTagDBTypes, false, inputAchievementTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InputAchievementTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = inputAchievementTagOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = inputAchievementTagTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := InputAchievementTags().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testInputAchievementTagsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	inputAchievementTagOne := &InputAchievementTag{}
	inputAchievementTagTwo := &InputAchievementTag{}
	if err = randomize.Struct(seed, inputAchievementTagOne, inputAchievementTagDBTypes, false, inputAchievementTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InputAchievementTag struct: %s", err)
	}
	if err = randomize.Struct(seed, inputAchievementTagTwo, inputAchievementTagDBTypes, false, inputAchievementTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InputAchievementTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = inputAchievementTagOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = inputAchievementTagTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := InputAchievementTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func inputAchievementTagBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *InputAchievementTag) error {
	*o = InputAchievementTag{}
	return nil
}

func inputAchievementTagAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *InputAchievementTag) error {
	*o = InputAchievementTag{}
	return nil
}

func inputAchievementTagAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *InputAchievementTag) error {
	*o = InputAchievementTag{}
	return nil
}

func inputAchievementTagBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *InputAchievementTag) error {
	*o = InputAchievementTag{}
	return nil
}

func inputAchievementTagAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *InputAchievementTag) error {
	*o = InputAchievementTag{}
	return nil
}

func inputAchievementTagBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *InputAchievementTag) error {
	*o = InputAchievementTag{}
	return nil
}

func inputAchievementTagAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *InputAchievementTag) error {
	*o = InputAchievementTag{}
	return nil
}

func inputAchievementTagBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *InputAchievementTag) error {
	*o = InputAchievementTag{}
	return nil
}

func inputAchievementTagAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *InputAchievementTag) error {
	*o = InputAchievementTag{}
	return nil
}

func testInputAchievementTagsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &InputAchievementTag{}
	o := &InputAchievementTag{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, inputAchievementTagDBTypes, false); err != nil {
		t.Errorf("Unable to randomize InputAchievementTag object: %s", err)
	}

	AddInputAchievementTagHook(boil.BeforeInsertHook, inputAchievementTagBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	inputAchievementTagBeforeInsertHooks = []InputAchievementTagHook{}

	AddInputAchievementTagHook(boil.AfterInsertHook, inputAchievementTagAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	inputAchievementTagAfterInsertHooks = []InputAchievementTagHook{}

	AddInputAchievementTagHook(boil.AfterSelectHook, inputAchievementTagAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	inputAchievementTagAfterSelectHooks = []InputAchievementTagHook{}

	AddInputAchievementTagHook(boil.BeforeUpdateHook, inputAchievementTagBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	inputAchievementTagBeforeUpdateHooks = []InputAchievementTagHook{}

	AddInputAchievementTagHook(boil.AfterUpdateHook, inputAchievementTagAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	inputAchievementTagAfterUpdateHooks = []InputAchievementTagHook{}

	AddInputAchievementTagHook(boil.BeforeDeleteHook, inputAchievementTagBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	inputAchievementTagBeforeDeleteHooks = []InputAchievementTagHook{}

	AddInputAchievementTagHook(boil.AfterDeleteHook, inputAchievementTagAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	inputAchievementTagAfterDeleteHooks = []InputAchievementTagHook{}

	AddInputAchievementTagHook(boil.BeforeUpsertHook, inputAchievementTagBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	inputAchievementTagBeforeUpsertHooks = []InputAchievementTagHook{}

	AddInputAchievementTagHook(boil.AfterUpsertHook, inputAchievementTagAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	inputAchievementTagAfterUpsertHooks = []InputAchievementTagHook{}
}

func testInputAchievementTagsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InputAchievementTag{}
	if err = randomize.Struct(seed, o, inputAchievementTagDBTypes, true, inputAchievementTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InputAchievementTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := InputAchievementTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testInputAchievementTagsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InputAchievementTag{}
	if err = randomize.Struct(seed, o, inputAchievementTagDBTypes, true); err != nil {
		t.Errorf("Unable to randomize InputAchievementTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(inputAchievementTagColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := InputAchievementTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testInputAchievementTagToOneInputAchievementUsingInputAchievement(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local InputAchievementTag
	var foreign InputAchievement

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, inputAchievementTagDBTypes, false, inputAchievementTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InputAchievementTag struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, inputAchievementDBTypes, false, inputAchievementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InputAchievement struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.InputAchievementID = foreign.InputAchievementID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.InputAchievement().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.InputAchievementID != foreign.InputAchievementID {
		t.Errorf("want: %v, got %v", foreign.InputAchievementID, check.InputAchievementID)
	}

	slice := InputAchievementTagSlice{&local}
	if err = local.L.LoadInputAchievement(ctx, tx, false, (*[]*InputAchievementTag)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.InputAchievement == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.InputAchievement = nil
	if err = local.L.LoadInputAchievement(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.InputAchievement == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testInputAchievementTagToOneMCategoryUsingCategory(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local InputAchievementTag
	var foreign MCategory

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, inputAchievementTagDBTypes, false, inputAchievementTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InputAchievementTag struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, mCategoryDBTypes, false, mCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MCategory struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.CategoryID = foreign.CategoryID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Category().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.CategoryID != foreign.CategoryID {
		t.Errorf("want: %v, got %v", foreign.CategoryID, check.CategoryID)
	}

	slice := InputAchievementTagSlice{&local}
	if err = local.L.LoadCategory(ctx, tx, false, (*[]*InputAchievementTag)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Category == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Category = nil
	if err = local.L.LoadCategory(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Category == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testInputAchievementTagToOneSetOpInputAchievementUsingInputAchievement(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a InputAchievementTag
	var b, c InputAchievement

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, inputAchievementTagDBTypes, false, strmangle.SetComplement(inputAchievementTagPrimaryKeyColumns, inputAchievementTagColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, inputAchievementDBTypes, false, strmangle.SetComplement(inputAchievementPrimaryKeyColumns, inputAchievementColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, inputAchievementDBTypes, false, strmangle.SetComplement(inputAchievementPrimaryKeyColumns, inputAchievementColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*InputAchievement{&b, &c} {
		err = a.SetInputAchievement(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.InputAchievement != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.InputAchievementTags[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.InputAchievementID != x.InputAchievementID {
			t.Error("foreign key was wrong value", a.InputAchievementID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.InputAchievementID))
		reflect.Indirect(reflect.ValueOf(&a.InputAchievementID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.InputAchievementID != x.InputAchievementID {
			t.Error("foreign key was wrong value", a.InputAchievementID, x.InputAchievementID)
		}
	}
}
func testInputAchievementTagToOneSetOpMCategoryUsingCategory(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a InputAchievementTag
	var b, c MCategory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, inputAchievementTagDBTypes, false, strmangle.SetComplement(inputAchievementTagPrimaryKeyColumns, inputAchievementTagColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, mCategoryDBTypes, false, strmangle.SetComplement(mCategoryPrimaryKeyColumns, mCategoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, mCategoryDBTypes, false, strmangle.SetComplement(mCategoryPrimaryKeyColumns, mCategoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*MCategory{&b, &c} {
		err = a.SetCategory(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Category != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.CategoryInputAchievementTags[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.CategoryID != x.CategoryID {
			t.Error("foreign key was wrong value", a.CategoryID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CategoryID))
		reflect.Indirect(reflect.ValueOf(&a.CategoryID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CategoryID != x.CategoryID {
			t.Error("foreign key was wrong value", a.CategoryID, x.CategoryID)
		}
	}
}

func testInputAchievementTagsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InputAchievementTag{}
	if err = randomize.Struct(seed, o, inputAchievementTagDBTypes, true, inputAchievementTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InputAchievementTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testInputAchievementTagsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InputAchievementTag{}
	if err = randomize.Struct(seed, o, inputAchievementTagDBTypes, true, inputAchievementTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InputAchievementTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := InputAchievementTagSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testInputAchievementTagsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &InputAchievementTag{}
	if err = randomize.Struct(seed, o, inputAchievementTagDBTypes, true, inputAchievementTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InputAchievementTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := InputAchievementTags().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	inputAchievementTagDBTypes = map[string]string{`InputAchievementTagID`: `int`, `InputAchievementID`: `int`, `CategoryID`: `int`, `CreatedBy`: `int`, `CreatedAt`: `timestamp`, `ModifiedBy`: `int`, `ModifiedAt`: `timestamp`}
	_                          = bytes.MinRead
)

func testInputAchievementTagsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(inputAchievementTagPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(inputAchievementTagAllColumns) == len(inputAchievementTagPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &InputAchievementTag{}
	if err = randomize.Struct(seed, o, inputAchievementTagDBTypes, true, inputAchievementTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InputAchievementTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := InputAchievementTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, inputAchievementTagDBTypes, true, inputAchievementTagPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize InputAchievementTag struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testInputAchievementTagsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(inputAchievementTagAllColumns) == len(inputAchievementTagPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &InputAchievementTag{}
	if err = randomize.Struct(seed, o, inputAchievementTagDBTypes, true, inputAchievementTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize InputAchievementTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := InputAchievementTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, inputAchievementTagDBTypes, true, inputAchievementTagPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize InputAchievementTag struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(inputAchievementTagAllColumns, inputAchievementTagPrimaryKeyColumns) {
		fields = inputAchievementTagAllColumns
	} else {
		fields = strmangle.SetComplement(
			inputAchievementTagAllColumns,
			inputAchievementTagPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := InputAchievementTagSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testInputAchievementTagsUpsert(t *testing.T) {
	t.Parallel()

	if len(inputAchievementTagAllColumns) == len(inputAchievementTagPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLInputAchievementTagUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := InputAchievementTag{}
	if err = randomize.Struct(seed, &o, inputAchievementTagDBTypes, false); err != nil {
		t.Errorf("Unable to randomize InputAchievementTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert InputAchievementTag: %s", err)
	}

	count, err := InputAchievementTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, inputAchievementTagDBTypes, false, inputAchievementTagPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize InputAchievementTag struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert InputAchievementTag: %s", err)
	}

	count, err = InputAchievementTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
