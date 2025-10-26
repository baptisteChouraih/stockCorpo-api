package models

import "time"

const (
	ActionCreate = "CREATE"
	ActionUpdate = "UPDATE"
	ActionDelete = "DELETE"
	ActionLogin  = "LOGIN"
	ActionLogout = "LOGOUT"
)

type Logs struct {
	IdLogs    int       `json:"idLogs"`
	IdUser    int       `json:"idUser"`
	Action    string    `json:"action"`
	Details   *string   `json:"details,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}
