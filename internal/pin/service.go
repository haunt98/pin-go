package pin

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/spf13/cast"
)

type Service interface {
	GetPinBySHA256(ctx context.Context, sha256 string) (Pin, error)
	AddPin(ctx context.Context, pin Pin) error
	InitPin(ctx context.Context, length int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetPinBySHA256(ctx context.Context, sha256 string) (Pin, error) {
	sha256 = strings.ToLower(sha256)

	pin, err := s.repo.SelectPinBySHA256(ctx, sha256)
	if err != nil {
		return Pin{}, fmt.Errorf("failed to select pin by sha256: %w", err)
	}

	return pin, nil
}

func (s *service) AddPin(ctx context.Context, pin Pin) error {
	if pin.Pin == "" {
		return errors.New("pin invalid")
	}

	// SHA256
	hashSHA256 := sha256.New()
	if _, err := hashSHA256.Write([]byte(pin.Pin)); err != nil {
		return fmt.Errorf("failed to write hash sha256: %w", err)
	}
	pin.SHA256Hex = strings.ToLower(hex.EncodeToString(hashSHA256.Sum(nil)))

	if err := s.repo.InsertPin(ctx, pin); err != nil {
		return fmt.Errorf("failed to insert pin: %w", err)
	}

	return nil
}

// Number 4 -> Pin 0000
// Number 5 -> Pin 00000
// Number 6 -> Pin 000000
func (s *service) InitPin(ctx context.Context, length int) error {
	for number := int64(0); number < cast.ToInt64(math.Pow10(length)); number++ {
		numberStr := pinToString(number, length)
		if err := s.AddPin(ctx, Pin{
			Pin: numberStr,
		}); err != nil {
			return err
		}
	}

	return nil
}
