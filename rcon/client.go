package rcon

import (
	"context"
	"fmt"
	"sync"

	gorcon "github.com/gorcon/rcon"
)

// The rcon client
type rconClient struct {
	conn *gorcon.Conn
	mu   sync.Mutex
}

// An rcon client
type Client interface {
	Close() error
	Execute(ctx context.Context, command string) (string, error)
}

// Creates a new rcon client
func New(address string, port string, password string) (Client, error) {
	conn, err := gorcon.Dial(fmt.Sprintf("%s:%s", address, port), password)
	if err != nil {
		return nil, fmt.Errorf("rcon dial: %w", err)
	}

	return &rconClient{
		conn: conn,
	}, nil
}

// Attempts to aquire the mutex. Gives up if the context is cancelled while waiting.
func (c *rconClient) lockWithContext(ctx context.Context) error {
	acquired := make(chan struct{})

	// Attempts to aquire mutex lock on a new thread.
	go func() {
		c.mu.Lock()
		close(acquired)
	}()

	// In the mean time if the context finishes or cancles then the mutex is opened to prevent deadlock.
	select {
	case <-acquired:
		return nil
	case <-ctx.Done():
		go func() {
			<-acquired
			c.mu.Unlock()
		}()
		return ctx.Err()
	}
}

// Execute an rcon command
func (c *rconClient) Execute(ctx context.Context, command string) (string, error) {

	if err := ctx.Err(); err != nil {
		return "", fmt.Errorf("rcon execute: %w", err)
	}

	if err := c.lockWithContext(ctx); err != nil {
		return "", fmt.Errorf("rcon execute: %w", err)
	}

	defer c.mu.Unlock()

	// Command execution could have been waiting for a while so we check again for any context updates
	if err := ctx.Err(); err != nil {
		return "", fmt.Errorf("rcon execute: %w", err)
	}

	res, err := c.conn.Execute(command)
	if err != nil {
		return res, fmt.Errorf("error executing command '%s': %w", command, err)
	}
	return res, nil
}

// Closes the rcon connection.
func (c *rconClient) Close() error {
	return c.conn.Close()
}
