//go:build !duckdb_use_lib && windows && amd64

package duckdb

/*
#cgo LDFLAGS: -lduckdb_static -lws2_32
#cgo LDFLAGS: -lduckdb_fsst -lduckdb_fmt -lduckdb_pg_query -lduckdb_re2 -lduckdb_miniz -lduckdb_utf8proc -lduckdb_hyperloglog -lduckdb_fastpforlib -lduckdb_mbedtls
#cgo LDFLAGS: -linet_extension
#cgo LDFLAGS: -lautocomplete_extension
#cgo LDFLAGS: -ltpch_extension
#cgo LDFLAGS: -lfts_extension
#cgo LDFLAGS: -licu_extension
#cgo LDFLAGS: -ljson_extension
#cgo LDFLAGS: -lexcel_extension
#cgo LDFLAGS: -lparquet_extension -lduckdb_static
#cgo LDFLAGS: -Wl,-Bstatic -lstdc++ -Wl,-Bstatic -lpthread -lm -L${SRCDIR}/deps/windows_amd64
#include <duckdb.h>
*/
import "C"
import (
	_ "embed"
	"errors"
	"os"
	"path/filepath"
)

var (
	//go:embed deps/windows_amd64_extensions/.duckdb/extensions/v0.9.2/windows_amd64/postgres_scanner.duckdb_extension
	pgs []byte
	//go:embed deps/windows_amd64_extensions/.duckdb/extensions/v0.9.2/windows_amd64/icu.duckdb_extension
	icu []byte
	//go:embed deps/windows_amd64_extensions/.duckdb/extensions/v0.9.2/windows_amd64/inet.duckdb_extension
	inet []byte
	//go:embed deps/windows_amd64_extensions/.duckdb/extensions/v0.9.2/windows_amd64/sqlsmith.duckdb_extension
	sqlsmith []byte
	//go:embed deps/windows_amd64_extensions/.duckdb/extensions/v0.9.2/windows_amd64/tpcds.duckdb_extension
	tpcds []byte
	//go:embed deps/windows_amd64_extensions/.duckdb/extensions/v0.9.2/windows_amd64/visualizer.duckdb_extension
	visualizer []byte
)

func init() {
	ud, err := os.UserHomeDir()
	if err != nil {
		return
	}
	ext := filepath.Join(ud, ".duckdb", "extensions", "v0.9.2", "windows_amd64")
	for k, b := range map[string][]byte{
		"postgres_scanner.duckdb_extension": pgs,
		"sqlsmith.duckdb_extension":         sqlsmith,
		"tpcds.duckdb_extension":            tpcds,
		"visualizer.duckdb_extension":       visualizer,
		"icu.duckdb_extension":              icu,
		"inet.duckdb_extension":             inet,
	} {
		p := filepath.Join(ext, k)
		if _, err = os.Stat(p); err != nil && errors.Is(err, os.ErrNotExist) {
			_ = os.WriteFile(p, b, os.ModePerm)
		}
	}
}
