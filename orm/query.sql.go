// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: query.sql

package orm

import (
	"context"
)

const addExOrder = `-- name: AddExOrder :one
insert into exorder ("task_id", "inout_id", "symbol", "enter", "order_type", "order_id", "side",
                     "create_at", "price", "average", "amount", "filled", "status", "fee", "fee_type", "update_at")
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
    RETURNING id
`

type AddExOrderParams struct {
	TaskID    int32
	InoutID   int32
	Symbol    string
	Enter     bool
	OrderType string
	OrderID   string
	Side      string
	CreateAt  int64
	Price     float64
	Average   float64
	Amount    float64
	Filled    float64
	Status    int16
	Fee       float64
	FeeType   string
	UpdateAt  int64
}

func (q *Queries) AddExOrder(ctx context.Context, arg AddExOrderParams) (int64, error) {
	row := q.db.QueryRow(ctx, addExOrder,
		arg.TaskID,
		arg.InoutID,
		arg.Symbol,
		arg.Enter,
		arg.OrderType,
		arg.OrderID,
		arg.Side,
		arg.CreateAt,
		arg.Price,
		arg.Average,
		arg.Amount,
		arg.Filled,
		arg.Status,
		arg.Fee,
		arg.FeeType,
		arg.UpdateAt,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const addIOrder = `-- name: AddIOrder :one
insert into iorder ("task_id", "symbol", "sid", "timeframe", "short", "status",
                    "enter_tag", "init_price", "quote_cost", "exit_tag", "leverage",
                    "enter_at", "exit_at", "strategy", "stg_ver", "max_draw_down",
                    "profit_rate", "profit", "info")
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
    RETURNING id
`

type AddIOrderParams struct {
	TaskID      int32
	Symbol      string
	Sid         int32
	Timeframe   string
	Short       bool
	Status      int16
	EnterTag    string
	InitPrice   float64
	QuoteCost   float64
	ExitTag     string
	Leverage    int32
	EnterAt     int64
	ExitAt      int64
	Strategy    string
	StgVer      int32
	MaxDrawDown float64
	ProfitRate  float64
	Profit      float64
	Info        string
}

func (q *Queries) AddIOrder(ctx context.Context, arg AddIOrderParams) (int64, error) {
	row := q.db.QueryRow(ctx, addIOrder,
		arg.TaskID,
		arg.Symbol,
		arg.Sid,
		arg.Timeframe,
		arg.Short,
		arg.Status,
		arg.EnterTag,
		arg.InitPrice,
		arg.QuoteCost,
		arg.ExitTag,
		arg.Leverage,
		arg.EnterAt,
		arg.ExitAt,
		arg.Strategy,
		arg.StgVer,
		arg.MaxDrawDown,
		arg.ProfitRate,
		arg.Profit,
		arg.Info,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

type AddKHolesParams struct {
	Sid       int32
	Timeframe string
	Start     int64
	Stop      int64
}

const addKInfo = `-- name: AddKInfo :one
insert into kinfo
("sid", "timeframe", "start", "stop")
values ($1, $2, $3, $4)
    returning sid, timeframe, start, stop
`

type AddKInfoParams struct {
	Sid       int32
	Timeframe string
	Start     int64
	Stop      int64
}

func (q *Queries) AddKInfo(ctx context.Context, arg AddKInfoParams) (*KInfo, error) {
	row := q.db.QueryRow(ctx, addKInfo,
		arg.Sid,
		arg.Timeframe,
		arg.Start,
		arg.Stop,
	)
	var i KInfo
	err := row.Scan(
		&i.Sid,
		&i.Timeframe,
		&i.Start,
		&i.Stop,
	)
	return &i, err
}

type AddSymbolsParams struct {
	Exchange string
	Market   string
	Symbol   string
}

const addTask = `-- name: AddTask :one
insert into bottask
("mode", "name", "create_at", "start_at", "stop_at", "info")
values ($1, $2, $3, $4, $5, $6)
returning id, mode, name, create_at, start_at, stop_at, info
`

type AddTaskParams struct {
	Mode     string
	Name     string
	CreateAt int64
	StartAt  int64
	StopAt   int64
	Info     string
}

func (q *Queries) AddTask(ctx context.Context, arg AddTaskParams) (*BotTask, error) {
	row := q.db.QueryRow(ctx, addTask,
		arg.Mode,
		arg.Name,
		arg.CreateAt,
		arg.StartAt,
		arg.StopAt,
		arg.Info,
	)
	var i BotTask
	err := row.Scan(
		&i.ID,
		&i.Mode,
		&i.Name,
		&i.CreateAt,
		&i.StartAt,
		&i.StopAt,
		&i.Info,
	)
	return &i, err
}

const delKHoleRange = `-- name: DelKHoleRange :exec
delete from khole
where sid = $1 and timeframe=$2 and start >= $3 and stop <= $4
`

type DelKHoleRangeParams struct {
	Sid       int32
	Timeframe string
	Start     int64
	Stop      int64
}

func (q *Queries) DelKHoleRange(ctx context.Context, arg DelKHoleRangeParams) error {
	_, err := q.db.Exec(ctx, delKHoleRange,
		arg.Sid,
		arg.Timeframe,
		arg.Start,
		arg.Stop,
	)
	return err
}

const findTask = `-- name: FindTask :one
select id, mode, name, create_at, start_at, stop_at, info from bottask
where mode = $1 and name = $2
order by create_at desc
limit 1
`

type FindTaskParams struct {
	Mode string
	Name string
}

func (q *Queries) FindTask(ctx context.Context, arg FindTaskParams) (*BotTask, error) {
	row := q.db.QueryRow(ctx, findTask, arg.Mode, arg.Name)
	var i BotTask
	err := row.Scan(
		&i.ID,
		&i.Mode,
		&i.Name,
		&i.CreateAt,
		&i.StartAt,
		&i.StopAt,
		&i.Info,
	)
	return &i, err
}

const getExOrders = `-- name: GetExOrders :many
select id, task_id, inout_id, symbol, enter, order_type, order_id, side, create_at, price, average, amount, filled, status, fee, fee_type, update_at from exorder
where inout_id=$1
`

func (q *Queries) GetExOrders(ctx context.Context, inoutID int32) ([]*ExOrder, error) {
	rows, err := q.db.Query(ctx, getExOrders, inoutID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*ExOrder{}
	for rows.Next() {
		var i ExOrder
		if err := rows.Scan(
			&i.ID,
			&i.TaskID,
			&i.InoutID,
			&i.Symbol,
			&i.Enter,
			&i.OrderType,
			&i.OrderID,
			&i.Side,
			&i.CreateAt,
			&i.Price,
			&i.Average,
			&i.Amount,
			&i.Filled,
			&i.Status,
			&i.Fee,
			&i.FeeType,
			&i.UpdateAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getIOrder = `-- name: GetIOrder :one
select id, task_id, symbol, sid, timeframe, short, status, enter_tag, init_price, quote_cost, exit_tag, leverage, enter_at, exit_at, strategy, stg_ver, max_draw_down, profit_rate, profit, info from iorder
where id = $1
`

func (q *Queries) GetIOrder(ctx context.Context, id int64) (*IOrder, error) {
	row := q.db.QueryRow(ctx, getIOrder, id)
	var i IOrder
	err := row.Scan(
		&i.ID,
		&i.TaskID,
		&i.Symbol,
		&i.Sid,
		&i.Timeframe,
		&i.Short,
		&i.Status,
		&i.EnterTag,
		&i.InitPrice,
		&i.QuoteCost,
		&i.ExitTag,
		&i.Leverage,
		&i.EnterAt,
		&i.ExitAt,
		&i.Strategy,
		&i.StgVer,
		&i.MaxDrawDown,
		&i.ProfitRate,
		&i.Profit,
		&i.Info,
	)
	return &i, err
}

const getKHoles = `-- name: GetKHoles :many
select id, sid, timeframe, start, stop from khole
where sid = $1 and timeframe = $2
`

type GetKHolesParams struct {
	Sid       int32
	Timeframe string
}

func (q *Queries) GetKHoles(ctx context.Context, arg GetKHolesParams) ([]*KHole, error) {
	rows, err := q.db.Query(ctx, getKHoles, arg.Sid, arg.Timeframe)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*KHole{}
	for rows.Next() {
		var i KHole
		if err := rows.Scan(
			&i.ID,
			&i.Sid,
			&i.Timeframe,
			&i.Start,
			&i.Stop,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTask = `-- name: GetTask :one
select id, mode, name, create_at, start_at, stop_at, info from bottask
where id = $1
`

func (q *Queries) GetTask(ctx context.Context, id int64) (*BotTask, error) {
	row := q.db.QueryRow(ctx, getTask, id)
	var i BotTask
	err := row.Scan(
		&i.ID,
		&i.Mode,
		&i.Name,
		&i.CreateAt,
		&i.StartAt,
		&i.StopAt,
		&i.Info,
	)
	return &i, err
}

const listExchanges = `-- name: ListExchanges :many
select distinct exchange from exsymbol
`

func (q *Queries) ListExchanges(ctx context.Context) ([]string, error) {
	rows, err := q.db.Query(ctx, listExchanges)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var exchange string
		if err := rows.Scan(&exchange); err != nil {
			return nil, err
		}
		items = append(items, exchange)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listKHoles = `-- name: ListKHoles :many
select id, sid, timeframe, start, stop from khole
order by sid, start
`

func (q *Queries) ListKHoles(ctx context.Context) ([]*KHole, error) {
	rows, err := q.db.Query(ctx, listKHoles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*KHole{}
	for rows.Next() {
		var i KHole
		if err := rows.Scan(
			&i.ID,
			&i.Sid,
			&i.Timeframe,
			&i.Start,
			&i.Stop,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listKInfos = `-- name: ListKInfos :many
select sid, timeframe, start, stop from kinfo
`

func (q *Queries) ListKInfos(ctx context.Context) ([]*KInfo, error) {
	rows, err := q.db.Query(ctx, listKInfos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*KInfo{}
	for rows.Next() {
		var i KInfo
		if err := rows.Scan(
			&i.Sid,
			&i.Timeframe,
			&i.Start,
			&i.Stop,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSymbols = `-- name: ListSymbols :many
select id, exchange, market, symbol, list_ms, delist_ms from exsymbol
where exchange = $1
order by id
`

func (q *Queries) ListSymbols(ctx context.Context, exchange string) ([]*ExSymbol, error) {
	rows, err := q.db.Query(ctx, listSymbols, exchange)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*ExSymbol{}
	for rows.Next() {
		var i ExSymbol
		if err := rows.Scan(
			&i.ID,
			&i.Exchange,
			&i.Market,
			&i.Symbol,
			&i.ListMs,
			&i.DelistMs,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTaskPairs = `-- name: ListTaskPairs :many
select symbol from iorder
where task_id = $1
and enter_at >= $2
and enter_at <= $3
`

type ListTaskPairsParams struct {
	TaskID    int32
	EnterAt   int64
	EnterAt_2 int64
}

func (q *Queries) ListTaskPairs(ctx context.Context, arg ListTaskPairsParams) ([]string, error) {
	rows, err := q.db.Query(ctx, listTaskPairs, arg.TaskID, arg.EnterAt, arg.EnterAt_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var symbol string
		if err := rows.Scan(&symbol); err != nil {
			return nil, err
		}
		items = append(items, symbol)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTasks = `-- name: ListTasks :many
select id, mode, name, create_at, start_at, stop_at, info from bottask
order by id
`

func (q *Queries) ListTasks(ctx context.Context) ([]*BotTask, error) {
	rows, err := q.db.Query(ctx, listTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*BotTask{}
	for rows.Next() {
		var i BotTask
		if err := rows.Scan(
			&i.ID,
			&i.Mode,
			&i.Name,
			&i.CreateAt,
			&i.StartAt,
			&i.StopAt,
			&i.Info,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setExOrder = `-- name: SetExOrder :exec
update exorder set
       "task_id" = $1,
       "inout_id" = $2,
       "symbol" = $3,
       "enter" = $4,
       "order_type" = $5,
       "order_id" = $6,
       "side" = $7,
       "create_at" = $8,
       "price" = $9,
       "average" = $10,
       "amount" = $11,
       "filled" = $12,
       "status" = $13,
       "fee" = $14,
       "fee_type" = $15,
       "update_at" = $16
    where id = $17
`

type SetExOrderParams struct {
	TaskID    int32
	InoutID   int32
	Symbol    string
	Enter     bool
	OrderType string
	OrderID   string
	Side      string
	CreateAt  int64
	Price     float64
	Average   float64
	Amount    float64
	Filled    float64
	Status    int16
	Fee       float64
	FeeType   string
	UpdateAt  int64
	ID        int64
}

func (q *Queries) SetExOrder(ctx context.Context, arg SetExOrderParams) error {
	_, err := q.db.Exec(ctx, setExOrder,
		arg.TaskID,
		arg.InoutID,
		arg.Symbol,
		arg.Enter,
		arg.OrderType,
		arg.OrderID,
		arg.Side,
		arg.CreateAt,
		arg.Price,
		arg.Average,
		arg.Amount,
		arg.Filled,
		arg.Status,
		arg.Fee,
		arg.FeeType,
		arg.UpdateAt,
		arg.ID,
	)
	return err
}

const setIOrder = `-- name: SetIOrder :exec
update iorder set
      "task_id" = $1,
      "symbol" = $2,
      "sid" = $3,
      "timeframe" = $4,
      "short" = $5,
      "status" = $6,
      "enter_tag" = $7,
      "init_price" = $8,
      "quote_cost" = $9,
      "exit_tag" = $10,
      "leverage" = $11,
      "enter_at" = $12,
      "exit_at" = $13,
      "strategy" = $14,
      "stg_ver" = $15,
      "max_draw_down" = $16,
      "profit_rate" = $17,
      "profit" = $18,
      "info" = $19
    WHERE id = $20
`

type SetIOrderParams struct {
	TaskID      int32
	Symbol      string
	Sid         int32
	Timeframe   string
	Short       bool
	Status      int16
	EnterTag    string
	InitPrice   float64
	QuoteCost   float64
	ExitTag     string
	Leverage    int32
	EnterAt     int64
	ExitAt      int64
	Strategy    string
	StgVer      int32
	MaxDrawDown float64
	ProfitRate  float64
	Profit      float64
	Info        string
	ID          int64
}

func (q *Queries) SetIOrder(ctx context.Context, arg SetIOrderParams) error {
	_, err := q.db.Exec(ctx, setIOrder,
		arg.TaskID,
		arg.Symbol,
		arg.Sid,
		arg.Timeframe,
		arg.Short,
		arg.Status,
		arg.EnterTag,
		arg.InitPrice,
		arg.QuoteCost,
		arg.ExitTag,
		arg.Leverage,
		arg.EnterAt,
		arg.ExitAt,
		arg.Strategy,
		arg.StgVer,
		arg.MaxDrawDown,
		arg.ProfitRate,
		arg.Profit,
		arg.Info,
		arg.ID,
	)
	return err
}

const setKHole = `-- name: SetKHole :exec
update khole set start = $2, stop = $3
where id = $1
`

type SetKHoleParams struct {
	ID    int64
	Start int64
	Stop  int64
}

func (q *Queries) SetKHole(ctx context.Context, arg SetKHoleParams) error {
	_, err := q.db.Exec(ctx, setKHole, arg.ID, arg.Start, arg.Stop)
	return err
}

const setKInfo = `-- name: SetKInfo :exec
update kinfo set start = $3, stop = $4
where sid = $1 and timeframe = $2
`

type SetKInfoParams struct {
	Sid       int32
	Timeframe string
	Start     int64
	Stop      int64
}

func (q *Queries) SetKInfo(ctx context.Context, arg SetKInfoParams) error {
	_, err := q.db.Exec(ctx, setKInfo,
		arg.Sid,
		arg.Timeframe,
		arg.Start,
		arg.Stop,
	)
	return err
}

const setListMS = `-- name: SetListMS :exec
update exsymbol set list_ms = $2, delist_ms = $3
where id = $1
`

type SetListMSParams struct {
	ID       int32
	ListMs   int64
	DelistMs int64
}

func (q *Queries) SetListMS(ctx context.Context, arg SetListMSParams) error {
	_, err := q.db.Exec(ctx, setListMS, arg.ID, arg.ListMs, arg.DelistMs)
	return err
}
