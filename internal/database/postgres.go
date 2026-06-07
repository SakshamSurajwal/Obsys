package database

import (
	"Obsys/internal/metrices"
	"log"
	"strconv"
	"time"

	"context"

	"github.com/jackc/pgx/v5"
)

type PostgresDB struct {
	Connection *pgx.Conn
}

func NewPostgresDB(connectionUrl string) (*PostgresDB, error) {
	connection, err := pgx.Connect(
		context.Background(),
		connectionUrl,
	)

	if err != nil {
		return nil, err
	}

	return &PostgresDB{Connection: connection}, nil
}

func (pdb *PostgresDB) SaveMetric(metrics []metrices.Metric) {
	for _, metric := range metrics {
		query := `INSERT INTO metrics(service, name, value, timestamp) VALUES ($1, $2, $3, $4)`
		_, err := pdb.Connection.Exec(context.Background(), query, metric.Service, metric.Name, metric.Value, metric.Timestamp)

		if err != nil {
			log.Println(err)
		}
	}
}

func AddServiceParam(query *string, serviceName string, argPos *int, args *[]interface{}) {
	if serviceName != "" {
		*query += (" AND service = $" + strconv.Itoa(*argPos))

		*args = append(*args, serviceName)
		*argPos++
	}
}

func AddNameParam(query *string, name string, argPos *int, args *[]interface{}) {
	if name != "" {
		*query += (" AND name = $" + strconv.Itoa(*argPos))

		*args = append(*args, name)
		*argPos++
	}
}

func AddFromTimeStamp(query *string, from time.Time, argPos *int, args *[]interface{}) {
	if !from.IsZero() {
		*query += (" AND timestamp >= $" + strconv.Itoa(*argPos))

		*args = append(*args, from)
		*argPos++
	}
}

func AddToimeStamp(query *string, to time.Time, argPos *int, args *[]interface{}) {
	if !to.IsZero() {
		*query += (" AND timestamp <= $" + strconv.Itoa(*argPos))

		*args = append(*args, to)
		*argPos++
	}
}

func AddLimit(query *string, limit string, argPos *int, args *[]interface{}) {
	if limit != "" {
		*query += (" LIMIT $" + strconv.Itoa(*argPos))

		num, _ := strconv.Atoi(limit)
		*args = append(*args, num)
		*argPos++
	}
}

func AddOffset(query *string, offset string, argPos *int, args *[]interface{}) {
	if offset != "" {
		*query += (" OFFSET $" + strconv.Itoa(*argPos))

		num, _ := strconv.Atoi(offset)
		*args = append(*args, num)
		*argPos++
	}
}

func (pdb *PostgresDB) GetMetric(serviceName string, name string, from time.Time, to time.Time, limit string, offset string) []metrices.Metric {
	query := `
		SELECT service,name,value,timestamp
		FROM metrics
		WHERE 1=1
	`

	args := []interface{}{}
	argPos := 1

	AddServiceParam(&query, serviceName, &argPos, &args)
	AddNameParam(&query, name, &argPos, &args)
	AddFromTimeStamp(&query, from, &argPos, &args)
	AddToimeStamp(&query, to, &argPos, &args)
	AddLimit(&query, limit, &argPos, &args)
	AddOffset(&query, offset, &argPos, &args)

	rows, err := pdb.Connection.Query(context.Background(), query, args...)

	if err != nil {
		log.Println(err)
		return []metrices.Metric{}
	}

	defer rows.Close()

	var metrics []metrices.Metric

	for rows.Next() {
		var metric metrices.Metric

		err := rows.Scan(&metric.Service, &metric.Name, &metric.Value, &metric.Timestamp)

		if err != nil {
			log.Println(err)
			continue
		}

		metrics = append(metrics, metric)
	}

	return metrics
}

func (pdb *PostgresDB) GetMetricCount(serviceName string, name string, from time.Time, to time.Time) int64 {

	query := `
		SELECT COUNT(*)
		FROM metrics
		WHERE 1=1
	`

	args := []interface{}{}
	argPos := 1

	AddServiceParam(&query, serviceName, &argPos, &args)
	AddNameParam(&query, name, &argPos, &args)
	AddFromTimeStamp(&query, from, &argPos, &args)
	AddToimeStamp(&query, to, &argPos, &args)

	var count int64

	err := pdb.Connection.QueryRow(context.Background(), query, args...).Scan(&count)

	if err != nil {
		log.Println(err)
		return 0
	}

	return count
}

func (pdb *PostgresDB) GetMetricSum(serviceName string, name string, from time.Time, to time.Time) int64 {

	query := `
		SELECT SUM(value)
		FROM metrics
		WHERE 1=1
	`

	args := []interface{}{}
	argPos := 1

	AddServiceParam(&query, serviceName, &argPos, &args)
	AddNameParam(&query, name, &argPos, &args)
	AddFromTimeStamp(&query, from, &argPos, &args)
	AddToimeStamp(&query, to, &argPos, &args)

	var sum int64

	err := pdb.Connection.QueryRow(context.Background(), query, args...).Scan(&sum)

	if err != nil {
		log.Println(err)
		return 0
	}

	return sum
}

func (pdb *PostgresDB) GetMetricAvg(serviceName string, name string, from time.Time, to time.Time) float64 {

	query := `
		SELECT AVG(value)
		FROM metrics
		WHERE 1=1
	`

	args := []interface{}{}
	argPos := 1

	AddServiceParam(&query, serviceName, &argPos, &args)
	AddNameParam(&query, name, &argPos, &args)
	AddFromTimeStamp(&query, from, &argPos, &args)
	AddToimeStamp(&query, to, &argPos, &args)

	var avg float64

	err := pdb.Connection.QueryRow(context.Background(), query, args...).Scan(&avg)

	if err != nil {
		log.Println(err)
		return 0
	}

	return avg
}

func (pdb *PostgresDB) GetMetricMin(serviceName string, name string, from time.Time, to time.Time) int64 {

	query := `
		SELECT MIN(value)
		FROM metrics
		WHERE 1=1
	`

	args := []interface{}{}
	argPos := 1

	AddServiceParam(&query, serviceName, &argPos, &args)
	AddNameParam(&query, name, &argPos, &args)
	AddFromTimeStamp(&query, from, &argPos, &args)
	AddToimeStamp(&query, to, &argPos, &args)

	var min int64

	err := pdb.Connection.QueryRow(context.Background(), query, args...).Scan(&min)

	if err != nil {
		log.Println(err)
		return 0
	}

	return min
}

func (pdb *PostgresDB) GetMetricMax(serviceName string, name string, from time.Time, to time.Time) int64 {

	query := `
		SELECT MAX(value)
		FROM metrics
		WHERE 1=1
	`

	args := []interface{}{}
	argPos := 1

	AddServiceParam(&query, serviceName, &argPos, &args)
	AddNameParam(&query, name, &argPos, &args)
	AddFromTimeStamp(&query, from, &argPos, &args)
	AddToimeStamp(&query, to, &argPos, &args)

	var max int64

	err := pdb.Connection.QueryRow(context.Background(), query, args...).Scan(&max)

	if err != nil {
		log.Println(err)
		return 0
	}

	return max
}
