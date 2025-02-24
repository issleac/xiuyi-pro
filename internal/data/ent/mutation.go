// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
	"xiuyiPro/internal/data/ent/predicate"
	"xiuyiPro/internal/data/ent/turtle"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTurtle = "Turtle"
	TypeUser   = "User"
)

// TurtleMutation represents an operation that mutates the Turtle nodes in the graph.
type TurtleMutation struct {
	config
	op            Op
	typ           string
	id            *int64
	qid           *string
	title         *string
	content       *string
	answer        *string
	category      *int32
	addcategory   *int32
	creator       *string
	state         *int32
	addstate      *int32
	ctime         *time.Time
	mtime         *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Turtle, error)
	predicates    []predicate.Turtle
}

var _ ent.Mutation = (*TurtleMutation)(nil)

// turtleOption allows management of the mutation configuration using functional options.
type turtleOption func(*TurtleMutation)

// newTurtleMutation creates new mutation for the Turtle entity.
func newTurtleMutation(c config, op Op, opts ...turtleOption) *TurtleMutation {
	m := &TurtleMutation{
		config:        c,
		op:            op,
		typ:           TypeTurtle,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTurtleID sets the ID field of the mutation.
func withTurtleID(id int64) turtleOption {
	return func(m *TurtleMutation) {
		var (
			err   error
			once  sync.Once
			value *Turtle
		)
		m.oldValue = func(ctx context.Context) (*Turtle, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Turtle.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTurtle sets the old Turtle of the mutation.
func withTurtle(node *Turtle) turtleOption {
	return func(m *TurtleMutation) {
		m.oldValue = func(context.Context) (*Turtle, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TurtleMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TurtleMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Turtle entities.
func (m *TurtleMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TurtleMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TurtleMutation) IDs(ctx context.Context) ([]int64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Turtle.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetQid sets the "qid" field.
func (m *TurtleMutation) SetQid(s string) {
	m.qid = &s
}

// Qid returns the value of the "qid" field in the mutation.
func (m *TurtleMutation) Qid() (r string, exists bool) {
	v := m.qid
	if v == nil {
		return
	}
	return *v, true
}

// OldQid returns the old "qid" field's value of the Turtle entity.
// If the Turtle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TurtleMutation) OldQid(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldQid is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldQid requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldQid: %w", err)
	}
	return oldValue.Qid, nil
}

// ResetQid resets all changes to the "qid" field.
func (m *TurtleMutation) ResetQid() {
	m.qid = nil
}

// SetTitle sets the "title" field.
func (m *TurtleMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *TurtleMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Turtle entity.
// If the Turtle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TurtleMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *TurtleMutation) ResetTitle() {
	m.title = nil
}

// SetContent sets the "content" field.
func (m *TurtleMutation) SetContent(s string) {
	m.content = &s
}

// Content returns the value of the "content" field in the mutation.
func (m *TurtleMutation) Content() (r string, exists bool) {
	v := m.content
	if v == nil {
		return
	}
	return *v, true
}

// OldContent returns the old "content" field's value of the Turtle entity.
// If the Turtle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TurtleMutation) OldContent(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldContent is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldContent requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldContent: %w", err)
	}
	return oldValue.Content, nil
}

// ResetContent resets all changes to the "content" field.
func (m *TurtleMutation) ResetContent() {
	m.content = nil
}

// SetAnswer sets the "answer" field.
func (m *TurtleMutation) SetAnswer(s string) {
	m.answer = &s
}

// Answer returns the value of the "answer" field in the mutation.
func (m *TurtleMutation) Answer() (r string, exists bool) {
	v := m.answer
	if v == nil {
		return
	}
	return *v, true
}

// OldAnswer returns the old "answer" field's value of the Turtle entity.
// If the Turtle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TurtleMutation) OldAnswer(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAnswer is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAnswer requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAnswer: %w", err)
	}
	return oldValue.Answer, nil
}

// ResetAnswer resets all changes to the "answer" field.
func (m *TurtleMutation) ResetAnswer() {
	m.answer = nil
}

// SetCategory sets the "category" field.
func (m *TurtleMutation) SetCategory(i int32) {
	m.category = &i
	m.addcategory = nil
}

// Category returns the value of the "category" field in the mutation.
func (m *TurtleMutation) Category() (r int32, exists bool) {
	v := m.category
	if v == nil {
		return
	}
	return *v, true
}

// OldCategory returns the old "category" field's value of the Turtle entity.
// If the Turtle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TurtleMutation) OldCategory(ctx context.Context) (v int32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCategory is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCategory requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCategory: %w", err)
	}
	return oldValue.Category, nil
}

// AddCategory adds i to the "category" field.
func (m *TurtleMutation) AddCategory(i int32) {
	if m.addcategory != nil {
		*m.addcategory += i
	} else {
		m.addcategory = &i
	}
}

// AddedCategory returns the value that was added to the "category" field in this mutation.
func (m *TurtleMutation) AddedCategory() (r int32, exists bool) {
	v := m.addcategory
	if v == nil {
		return
	}
	return *v, true
}

// ResetCategory resets all changes to the "category" field.
func (m *TurtleMutation) ResetCategory() {
	m.category = nil
	m.addcategory = nil
}

// SetCreator sets the "creator" field.
func (m *TurtleMutation) SetCreator(s string) {
	m.creator = &s
}

// Creator returns the value of the "creator" field in the mutation.
func (m *TurtleMutation) Creator() (r string, exists bool) {
	v := m.creator
	if v == nil {
		return
	}
	return *v, true
}

// OldCreator returns the old "creator" field's value of the Turtle entity.
// If the Turtle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TurtleMutation) OldCreator(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreator is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreator requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreator: %w", err)
	}
	return oldValue.Creator, nil
}

// ResetCreator resets all changes to the "creator" field.
func (m *TurtleMutation) ResetCreator() {
	m.creator = nil
}

// SetState sets the "state" field.
func (m *TurtleMutation) SetState(i int32) {
	m.state = &i
	m.addstate = nil
}

// State returns the value of the "state" field in the mutation.
func (m *TurtleMutation) State() (r int32, exists bool) {
	v := m.state
	if v == nil {
		return
	}
	return *v, true
}

// OldState returns the old "state" field's value of the Turtle entity.
// If the Turtle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TurtleMutation) OldState(ctx context.Context) (v int32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldState is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldState requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldState: %w", err)
	}
	return oldValue.State, nil
}

// AddState adds i to the "state" field.
func (m *TurtleMutation) AddState(i int32) {
	if m.addstate != nil {
		*m.addstate += i
	} else {
		m.addstate = &i
	}
}

// AddedState returns the value that was added to the "state" field in this mutation.
func (m *TurtleMutation) AddedState() (r int32, exists bool) {
	v := m.addstate
	if v == nil {
		return
	}
	return *v, true
}

// ResetState resets all changes to the "state" field.
func (m *TurtleMutation) ResetState() {
	m.state = nil
	m.addstate = nil
}

// SetCtime sets the "ctime" field.
func (m *TurtleMutation) SetCtime(t time.Time) {
	m.ctime = &t
}

// Ctime returns the value of the "ctime" field in the mutation.
func (m *TurtleMutation) Ctime() (r time.Time, exists bool) {
	v := m.ctime
	if v == nil {
		return
	}
	return *v, true
}

// OldCtime returns the old "ctime" field's value of the Turtle entity.
// If the Turtle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TurtleMutation) OldCtime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCtime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCtime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCtime: %w", err)
	}
	return oldValue.Ctime, nil
}

// ResetCtime resets all changes to the "ctime" field.
func (m *TurtleMutation) ResetCtime() {
	m.ctime = nil
}

// SetMtime sets the "mtime" field.
func (m *TurtleMutation) SetMtime(t time.Time) {
	m.mtime = &t
}

// Mtime returns the value of the "mtime" field in the mutation.
func (m *TurtleMutation) Mtime() (r time.Time, exists bool) {
	v := m.mtime
	if v == nil {
		return
	}
	return *v, true
}

// OldMtime returns the old "mtime" field's value of the Turtle entity.
// If the Turtle object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TurtleMutation) OldMtime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMtime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMtime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMtime: %w", err)
	}
	return oldValue.Mtime, nil
}

// ResetMtime resets all changes to the "mtime" field.
func (m *TurtleMutation) ResetMtime() {
	m.mtime = nil
}

// Where appends a list predicates to the TurtleMutation builder.
func (m *TurtleMutation) Where(ps ...predicate.Turtle) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the TurtleMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *TurtleMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Turtle, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *TurtleMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *TurtleMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Turtle).
func (m *TurtleMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TurtleMutation) Fields() []string {
	fields := make([]string, 0, 9)
	if m.qid != nil {
		fields = append(fields, turtle.FieldQid)
	}
	if m.title != nil {
		fields = append(fields, turtle.FieldTitle)
	}
	if m.content != nil {
		fields = append(fields, turtle.FieldContent)
	}
	if m.answer != nil {
		fields = append(fields, turtle.FieldAnswer)
	}
	if m.category != nil {
		fields = append(fields, turtle.FieldCategory)
	}
	if m.creator != nil {
		fields = append(fields, turtle.FieldCreator)
	}
	if m.state != nil {
		fields = append(fields, turtle.FieldState)
	}
	if m.ctime != nil {
		fields = append(fields, turtle.FieldCtime)
	}
	if m.mtime != nil {
		fields = append(fields, turtle.FieldMtime)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TurtleMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case turtle.FieldQid:
		return m.Qid()
	case turtle.FieldTitle:
		return m.Title()
	case turtle.FieldContent:
		return m.Content()
	case turtle.FieldAnswer:
		return m.Answer()
	case turtle.FieldCategory:
		return m.Category()
	case turtle.FieldCreator:
		return m.Creator()
	case turtle.FieldState:
		return m.State()
	case turtle.FieldCtime:
		return m.Ctime()
	case turtle.FieldMtime:
		return m.Mtime()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TurtleMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case turtle.FieldQid:
		return m.OldQid(ctx)
	case turtle.FieldTitle:
		return m.OldTitle(ctx)
	case turtle.FieldContent:
		return m.OldContent(ctx)
	case turtle.FieldAnswer:
		return m.OldAnswer(ctx)
	case turtle.FieldCategory:
		return m.OldCategory(ctx)
	case turtle.FieldCreator:
		return m.OldCreator(ctx)
	case turtle.FieldState:
		return m.OldState(ctx)
	case turtle.FieldCtime:
		return m.OldCtime(ctx)
	case turtle.FieldMtime:
		return m.OldMtime(ctx)
	}
	return nil, fmt.Errorf("unknown Turtle field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TurtleMutation) SetField(name string, value ent.Value) error {
	switch name {
	case turtle.FieldQid:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetQid(v)
		return nil
	case turtle.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case turtle.FieldContent:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetContent(v)
		return nil
	case turtle.FieldAnswer:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAnswer(v)
		return nil
	case turtle.FieldCategory:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCategory(v)
		return nil
	case turtle.FieldCreator:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreator(v)
		return nil
	case turtle.FieldState:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetState(v)
		return nil
	case turtle.FieldCtime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCtime(v)
		return nil
	case turtle.FieldMtime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMtime(v)
		return nil
	}
	return fmt.Errorf("unknown Turtle field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TurtleMutation) AddedFields() []string {
	var fields []string
	if m.addcategory != nil {
		fields = append(fields, turtle.FieldCategory)
	}
	if m.addstate != nil {
		fields = append(fields, turtle.FieldState)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TurtleMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case turtle.FieldCategory:
		return m.AddedCategory()
	case turtle.FieldState:
		return m.AddedState()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TurtleMutation) AddField(name string, value ent.Value) error {
	switch name {
	case turtle.FieldCategory:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddCategory(v)
		return nil
	case turtle.FieldState:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddState(v)
		return nil
	}
	return fmt.Errorf("unknown Turtle numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TurtleMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TurtleMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TurtleMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Turtle nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TurtleMutation) ResetField(name string) error {
	switch name {
	case turtle.FieldQid:
		m.ResetQid()
		return nil
	case turtle.FieldTitle:
		m.ResetTitle()
		return nil
	case turtle.FieldContent:
		m.ResetContent()
		return nil
	case turtle.FieldAnswer:
		m.ResetAnswer()
		return nil
	case turtle.FieldCategory:
		m.ResetCategory()
		return nil
	case turtle.FieldCreator:
		m.ResetCreator()
		return nil
	case turtle.FieldState:
		m.ResetState()
		return nil
	case turtle.FieldCtime:
		m.ResetCtime()
		return nil
	case turtle.FieldMtime:
		m.ResetMtime()
		return nil
	}
	return fmt.Errorf("unknown Turtle field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TurtleMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TurtleMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TurtleMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TurtleMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TurtleMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TurtleMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TurtleMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Turtle unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TurtleMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Turtle edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 0)
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}
