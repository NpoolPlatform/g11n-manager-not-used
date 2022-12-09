// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/g11n-manager/pkg/db/ent/message"
	"github.com/NpoolPlatform/g11n-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// MessageUpdate is the builder for updating Message entities.
type MessageUpdate struct {
	config
	hooks     []Hook
	mutation  *MessageMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the MessageUpdate builder.
func (mu *MessageUpdate) Where(ps ...predicate.Message) *MessageUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetCreatedAt sets the "created_at" field.
func (mu *MessageUpdate) SetCreatedAt(u uint32) *MessageUpdate {
	mu.mutation.ResetCreatedAt()
	mu.mutation.SetCreatedAt(u)
	return mu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableCreatedAt(u *uint32) *MessageUpdate {
	if u != nil {
		mu.SetCreatedAt(*u)
	}
	return mu
}

// AddCreatedAt adds u to the "created_at" field.
func (mu *MessageUpdate) AddCreatedAt(u int32) *MessageUpdate {
	mu.mutation.AddCreatedAt(u)
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MessageUpdate) SetUpdatedAt(u uint32) *MessageUpdate {
	mu.mutation.ResetUpdatedAt()
	mu.mutation.SetUpdatedAt(u)
	return mu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (mu *MessageUpdate) AddUpdatedAt(u int32) *MessageUpdate {
	mu.mutation.AddUpdatedAt(u)
	return mu
}

// SetDeletedAt sets the "deleted_at" field.
func (mu *MessageUpdate) SetDeletedAt(u uint32) *MessageUpdate {
	mu.mutation.ResetDeletedAt()
	mu.mutation.SetDeletedAt(u)
	return mu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableDeletedAt(u *uint32) *MessageUpdate {
	if u != nil {
		mu.SetDeletedAt(*u)
	}
	return mu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (mu *MessageUpdate) AddDeletedAt(u int32) *MessageUpdate {
	mu.mutation.AddDeletedAt(u)
	return mu
}

// SetAppID sets the "app_id" field.
func (mu *MessageUpdate) SetAppID(u uuid.UUID) *MessageUpdate {
	mu.mutation.SetAppID(u)
	return mu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableAppID(u *uuid.UUID) *MessageUpdate {
	if u != nil {
		mu.SetAppID(*u)
	}
	return mu
}

// ClearAppID clears the value of the "app_id" field.
func (mu *MessageUpdate) ClearAppID() *MessageUpdate {
	mu.mutation.ClearAppID()
	return mu
}

// SetLangID sets the "lang_id" field.
func (mu *MessageUpdate) SetLangID(u uuid.UUID) *MessageUpdate {
	mu.mutation.SetLangID(u)
	return mu
}

// SetNillableLangID sets the "lang_id" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableLangID(u *uuid.UUID) *MessageUpdate {
	if u != nil {
		mu.SetLangID(*u)
	}
	return mu
}

// ClearLangID clears the value of the "lang_id" field.
func (mu *MessageUpdate) ClearLangID() *MessageUpdate {
	mu.mutation.ClearLangID()
	return mu
}

// SetMessageID sets the "message_id" field.
func (mu *MessageUpdate) SetMessageID(s string) *MessageUpdate {
	mu.mutation.SetMessageID(s)
	return mu
}

// SetNillableMessageID sets the "message_id" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableMessageID(s *string) *MessageUpdate {
	if s != nil {
		mu.SetMessageID(*s)
	}
	return mu
}

// ClearMessageID clears the value of the "message_id" field.
func (mu *MessageUpdate) ClearMessageID() *MessageUpdate {
	mu.mutation.ClearMessageID()
	return mu
}

// SetMessage sets the "message" field.
func (mu *MessageUpdate) SetMessage(s string) *MessageUpdate {
	mu.mutation.SetMessage(s)
	return mu
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableMessage(s *string) *MessageUpdate {
	if s != nil {
		mu.SetMessage(*s)
	}
	return mu
}

// ClearMessage clears the value of the "message" field.
func (mu *MessageUpdate) ClearMessage() *MessageUpdate {
	mu.mutation.ClearMessage()
	return mu
}

// SetGetIndex sets the "get_index" field.
func (mu *MessageUpdate) SetGetIndex(u uint32) *MessageUpdate {
	mu.mutation.ResetGetIndex()
	mu.mutation.SetGetIndex(u)
	return mu
}

// SetNillableGetIndex sets the "get_index" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableGetIndex(u *uint32) *MessageUpdate {
	if u != nil {
		mu.SetGetIndex(*u)
	}
	return mu
}

// AddGetIndex adds u to the "get_index" field.
func (mu *MessageUpdate) AddGetIndex(u int32) *MessageUpdate {
	mu.mutation.AddGetIndex(u)
	return mu
}

// ClearGetIndex clears the value of the "get_index" field.
func (mu *MessageUpdate) ClearGetIndex() *MessageUpdate {
	mu.mutation.ClearGetIndex()
	return mu
}

// SetDisabled sets the "disabled" field.
func (mu *MessageUpdate) SetDisabled(b bool) *MessageUpdate {
	mu.mutation.SetDisabled(b)
	return mu
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableDisabled(b *bool) *MessageUpdate {
	if b != nil {
		mu.SetDisabled(*b)
	}
	return mu
}

// ClearDisabled clears the value of the "disabled" field.
func (mu *MessageUpdate) ClearDisabled() *MessageUpdate {
	mu.mutation.ClearDisabled()
	return mu
}

// SetShort sets the "short" field.
func (mu *MessageUpdate) SetShort(s string) *MessageUpdate {
	mu.mutation.SetShort(s)
	return mu
}

// SetNillableShort sets the "short" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableShort(s *string) *MessageUpdate {
	if s != nil {
		mu.SetShort(*s)
	}
	return mu
}

// ClearShort clears the value of the "short" field.
func (mu *MessageUpdate) ClearShort() *MessageUpdate {
	mu.mutation.ClearShort()
	return mu
}

// Mutation returns the MessageMutation object of the builder.
func (mu *MessageUpdate) Mutation() *MessageMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MessageUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := mu.defaults(); err != nil {
		return 0, err
	}
	if len(mu.hooks) == 0 {
		if err = mu.check(); err != nil {
			return 0, err
		}
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mu.check(); err != nil {
				return 0, err
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MessageUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MessageUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MessageUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MessageUpdate) defaults() error {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		if message.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized message.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := message.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (mu *MessageUpdate) check() error {
	if v, ok := mu.mutation.Message(); ok {
		if err := message.MessageValidator(v); err != nil {
			return &ValidationError{Name: "message", err: fmt.Errorf(`ent: validator failed for field "Message.message": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mu *MessageUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MessageUpdate {
	mu.modifiers = append(mu.modifiers, modifiers...)
	return mu
}

func (mu *MessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   message.Table,
			Columns: message.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: message.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: message.FieldCreatedAt,
		})
	}
	if value, ok := mu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: message.FieldCreatedAt,
		})
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: message.FieldUpdatedAt,
		})
	}
	if value, ok := mu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: message.FieldUpdatedAt,
		})
	}
	if value, ok := mu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: message.FieldDeletedAt,
		})
	}
	if value, ok := mu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: message.FieldDeletedAt,
		})
	}
	if value, ok := mu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: message.FieldAppID,
		})
	}
	if mu.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: message.FieldAppID,
		})
	}
	if value, ok := mu.mutation.LangID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: message.FieldLangID,
		})
	}
	if mu.mutation.LangIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: message.FieldLangID,
		})
	}
	if value, ok := mu.mutation.MessageID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldMessageID,
		})
	}
	if mu.mutation.MessageIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: message.FieldMessageID,
		})
	}
	if value, ok := mu.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldMessage,
		})
	}
	if mu.mutation.MessageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: message.FieldMessage,
		})
	}
	if value, ok := mu.mutation.GetIndex(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: message.FieldGetIndex,
		})
	}
	if value, ok := mu.mutation.AddedGetIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: message.FieldGetIndex,
		})
	}
	if mu.mutation.GetIndexCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: message.FieldGetIndex,
		})
	}
	if value, ok := mu.mutation.Disabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: message.FieldDisabled,
		})
	}
	if mu.mutation.DisabledCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: message.FieldDisabled,
		})
	}
	if value, ok := mu.mutation.Short(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldShort,
		})
	}
	if mu.mutation.ShortCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: message.FieldShort,
		})
	}
	_spec.Modifiers = mu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// MessageUpdateOne is the builder for updating a single Message entity.
type MessageUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *MessageMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (muo *MessageUpdateOne) SetCreatedAt(u uint32) *MessageUpdateOne {
	muo.mutation.ResetCreatedAt()
	muo.mutation.SetCreatedAt(u)
	return muo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableCreatedAt(u *uint32) *MessageUpdateOne {
	if u != nil {
		muo.SetCreatedAt(*u)
	}
	return muo
}

// AddCreatedAt adds u to the "created_at" field.
func (muo *MessageUpdateOne) AddCreatedAt(u int32) *MessageUpdateOne {
	muo.mutation.AddCreatedAt(u)
	return muo
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MessageUpdateOne) SetUpdatedAt(u uint32) *MessageUpdateOne {
	muo.mutation.ResetUpdatedAt()
	muo.mutation.SetUpdatedAt(u)
	return muo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (muo *MessageUpdateOne) AddUpdatedAt(u int32) *MessageUpdateOne {
	muo.mutation.AddUpdatedAt(u)
	return muo
}

// SetDeletedAt sets the "deleted_at" field.
func (muo *MessageUpdateOne) SetDeletedAt(u uint32) *MessageUpdateOne {
	muo.mutation.ResetDeletedAt()
	muo.mutation.SetDeletedAt(u)
	return muo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableDeletedAt(u *uint32) *MessageUpdateOne {
	if u != nil {
		muo.SetDeletedAt(*u)
	}
	return muo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (muo *MessageUpdateOne) AddDeletedAt(u int32) *MessageUpdateOne {
	muo.mutation.AddDeletedAt(u)
	return muo
}

// SetAppID sets the "app_id" field.
func (muo *MessageUpdateOne) SetAppID(u uuid.UUID) *MessageUpdateOne {
	muo.mutation.SetAppID(u)
	return muo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableAppID(u *uuid.UUID) *MessageUpdateOne {
	if u != nil {
		muo.SetAppID(*u)
	}
	return muo
}

// ClearAppID clears the value of the "app_id" field.
func (muo *MessageUpdateOne) ClearAppID() *MessageUpdateOne {
	muo.mutation.ClearAppID()
	return muo
}

// SetLangID sets the "lang_id" field.
func (muo *MessageUpdateOne) SetLangID(u uuid.UUID) *MessageUpdateOne {
	muo.mutation.SetLangID(u)
	return muo
}

// SetNillableLangID sets the "lang_id" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableLangID(u *uuid.UUID) *MessageUpdateOne {
	if u != nil {
		muo.SetLangID(*u)
	}
	return muo
}

// ClearLangID clears the value of the "lang_id" field.
func (muo *MessageUpdateOne) ClearLangID() *MessageUpdateOne {
	muo.mutation.ClearLangID()
	return muo
}

// SetMessageID sets the "message_id" field.
func (muo *MessageUpdateOne) SetMessageID(s string) *MessageUpdateOne {
	muo.mutation.SetMessageID(s)
	return muo
}

// SetNillableMessageID sets the "message_id" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableMessageID(s *string) *MessageUpdateOne {
	if s != nil {
		muo.SetMessageID(*s)
	}
	return muo
}

// ClearMessageID clears the value of the "message_id" field.
func (muo *MessageUpdateOne) ClearMessageID() *MessageUpdateOne {
	muo.mutation.ClearMessageID()
	return muo
}

// SetMessage sets the "message" field.
func (muo *MessageUpdateOne) SetMessage(s string) *MessageUpdateOne {
	muo.mutation.SetMessage(s)
	return muo
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableMessage(s *string) *MessageUpdateOne {
	if s != nil {
		muo.SetMessage(*s)
	}
	return muo
}

// ClearMessage clears the value of the "message" field.
func (muo *MessageUpdateOne) ClearMessage() *MessageUpdateOne {
	muo.mutation.ClearMessage()
	return muo
}

// SetGetIndex sets the "get_index" field.
func (muo *MessageUpdateOne) SetGetIndex(u uint32) *MessageUpdateOne {
	muo.mutation.ResetGetIndex()
	muo.mutation.SetGetIndex(u)
	return muo
}

// SetNillableGetIndex sets the "get_index" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableGetIndex(u *uint32) *MessageUpdateOne {
	if u != nil {
		muo.SetGetIndex(*u)
	}
	return muo
}

// AddGetIndex adds u to the "get_index" field.
func (muo *MessageUpdateOne) AddGetIndex(u int32) *MessageUpdateOne {
	muo.mutation.AddGetIndex(u)
	return muo
}

// ClearGetIndex clears the value of the "get_index" field.
func (muo *MessageUpdateOne) ClearGetIndex() *MessageUpdateOne {
	muo.mutation.ClearGetIndex()
	return muo
}

// SetDisabled sets the "disabled" field.
func (muo *MessageUpdateOne) SetDisabled(b bool) *MessageUpdateOne {
	muo.mutation.SetDisabled(b)
	return muo
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableDisabled(b *bool) *MessageUpdateOne {
	if b != nil {
		muo.SetDisabled(*b)
	}
	return muo
}

// ClearDisabled clears the value of the "disabled" field.
func (muo *MessageUpdateOne) ClearDisabled() *MessageUpdateOne {
	muo.mutation.ClearDisabled()
	return muo
}

// SetShort sets the "short" field.
func (muo *MessageUpdateOne) SetShort(s string) *MessageUpdateOne {
	muo.mutation.SetShort(s)
	return muo
}

// SetNillableShort sets the "short" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableShort(s *string) *MessageUpdateOne {
	if s != nil {
		muo.SetShort(*s)
	}
	return muo
}

// ClearShort clears the value of the "short" field.
func (muo *MessageUpdateOne) ClearShort() *MessageUpdateOne {
	muo.mutation.ClearShort()
	return muo
}

// Mutation returns the MessageMutation object of the builder.
func (muo *MessageUpdateOne) Mutation() *MessageMutation {
	return muo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MessageUpdateOne) Select(field string, fields ...string) *MessageUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Message entity.
func (muo *MessageUpdateOne) Save(ctx context.Context) (*Message, error) {
	var (
		err  error
		node *Message
	)
	if err := muo.defaults(); err != nil {
		return nil, err
	}
	if len(muo.hooks) == 0 {
		if err = muo.check(); err != nil {
			return nil, err
		}
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = muo.check(); err != nil {
				return nil, err
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, muo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Message)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MessageMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MessageUpdateOne) SaveX(ctx context.Context) *Message {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MessageUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MessageUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MessageUpdateOne) defaults() error {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		if message.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized message.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := message.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (muo *MessageUpdateOne) check() error {
	if v, ok := muo.mutation.Message(); ok {
		if err := message.MessageValidator(v); err != nil {
			return &ValidationError{Name: "message", err: fmt.Errorf(`ent: validator failed for field "Message.message": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (muo *MessageUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MessageUpdateOne {
	muo.modifiers = append(muo.modifiers, modifiers...)
	return muo
}

func (muo *MessageUpdateOne) sqlSave(ctx context.Context) (_node *Message, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   message.Table,
			Columns: message.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: message.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Message.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, message.FieldID)
		for _, f := range fields {
			if !message.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != message.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: message.FieldCreatedAt,
		})
	}
	if value, ok := muo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: message.FieldCreatedAt,
		})
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: message.FieldUpdatedAt,
		})
	}
	if value, ok := muo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: message.FieldUpdatedAt,
		})
	}
	if value, ok := muo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: message.FieldDeletedAt,
		})
	}
	if value, ok := muo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: message.FieldDeletedAt,
		})
	}
	if value, ok := muo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: message.FieldAppID,
		})
	}
	if muo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: message.FieldAppID,
		})
	}
	if value, ok := muo.mutation.LangID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: message.FieldLangID,
		})
	}
	if muo.mutation.LangIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: message.FieldLangID,
		})
	}
	if value, ok := muo.mutation.MessageID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldMessageID,
		})
	}
	if muo.mutation.MessageIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: message.FieldMessageID,
		})
	}
	if value, ok := muo.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldMessage,
		})
	}
	if muo.mutation.MessageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: message.FieldMessage,
		})
	}
	if value, ok := muo.mutation.GetIndex(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: message.FieldGetIndex,
		})
	}
	if value, ok := muo.mutation.AddedGetIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: message.FieldGetIndex,
		})
	}
	if muo.mutation.GetIndexCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: message.FieldGetIndex,
		})
	}
	if value, ok := muo.mutation.Disabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: message.FieldDisabled,
		})
	}
	if muo.mutation.DisabledCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: message.FieldDisabled,
		})
	}
	if value, ok := muo.mutation.Short(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldShort,
		})
	}
	if muo.mutation.ShortCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: message.FieldShort,
		})
	}
	_spec.Modifiers = muo.modifiers
	_node = &Message{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
