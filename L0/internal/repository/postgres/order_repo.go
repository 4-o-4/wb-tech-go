package postgres

import (
    "L0/internal/domain"
    "database/sql"
    "encoding/json"
    "errors"
    _ "github.com/lib/pq"
    "time"
)

type OrderRepository struct {
    db *sql.DB
}

func NewPostgresConnection(databaseURL string) (*OrderRepository, error) {
    db, err := sql.Open("postgres", databaseURL)
    if err != nil {
        return nil, err
    }
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)
    db.SetConnMaxLifetime(5 * 60 * time.Second)
    if err := db.Ping(); err != nil {
        return nil, err
    }
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS orders (id SERIAL PRIMARY KEY, data JSONB NOT NULL)`)
    if err != nil {
        return nil, err
    }
    repo := &OrderRepository{db: db}
    return repo, nil
}

func (r *OrderRepository) GetOrderById(id string) (*domain.Order, error) {
    var jsonData string
    query := `SELECT data FROM orders WHERE id = $1`
    err := r.db.QueryRow(query, id).Scan(&jsonData)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, errors.New("order not found")
        }
        return nil, err
    }
    var order domain.Order
    if err := json.Unmarshal([]byte(jsonData), &order); err != nil {
        return nil, err
    }
    return &order, nil
}

func (r *OrderRepository) SaveOrder(order *domain.Order) error {
    query := `INSERT INTO orders (data) VALUES ($1)`
    orderJSON, err := json.Marshal(order)
    if err != nil {
        return err
    }
    _, err = r.db.Exec(query, orderJSON)
    return err
}
