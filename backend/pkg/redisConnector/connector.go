package redisConnector

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host               string `validate:"required"`
	Port               string `validate:"required"`
	MinIdleConns       int    `validate:"required"`
	PoolSize           int    `validate:"required"`
	PoolTimeout        int    `validate:"required"`
	Password           string
	UseCertificates    bool
	InsecureSkipVerify bool
	CertificatesPaths  struct {
		Cert string
		Key  string
		Ca   string
	}
	DB int
}

func GetConnection(cfg Config) (conn *redis.Client, err error) {
	opts := &redis.Options{}
	if cfg.UseCertificates {
		certs := make([]tls.Certificate, 0)
		if cfg.CertificatesPaths.Cert != "" && cfg.CertificatesPaths.Key != "" {
			cert, err := tls.LoadX509KeyPair(cfg.CertificatesPaths.Cert, cfg.CertificatesPaths.Key)
			if err != nil {
				return nil, errors.Wrapf(
					err,
					"certPath: %v, keyPath: %v",
					cfg.CertificatesPaths.Cert,
					cfg.CertificatesPaths.Key,
				)
			}
			certs = append(certs, cert)
		}
		caCert, err := os.ReadFile(cfg.CertificatesPaths.Ca)
		if err != nil {
			return nil, errors.Wrapf(err, "ca load path: %v", cfg.CertificatesPaths.Ca)
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		opts = &redis.Options{
			Addr:         fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			MinIdleConns: cfg.MinIdleConns,
			PoolSize:     cfg.PoolSize,
			PoolTimeout:  time.Duration(cfg.PoolTimeout) * time.Second,
			Password:     cfg.Password,
			DB:           cfg.DB,
			TLSConfig: &tls.Config{
				InsecureSkipVerify: cfg.InsecureSkipVerify,
				Certificates:       certs,
				RootCAs:            caCertPool,
			},
		}
	} else {
		opts = &redis.Options{
			Addr:         fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			MinIdleConns: cfg.MinIdleConns,
			PoolSize:     cfg.PoolSize,
			PoolTimeout:  time.Duration(cfg.PoolTimeout) * time.Second,
			Password:     cfg.Password,
			DB:           cfg.DB,
		}
	}

	client := redis.NewClient(opts)
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, errors.Wrapf(err, "ping")
	}

	return client, nil
}
