package testutils

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
)

type BinaryCache struct {
	dir   string
	cache sync.Map
}

func NewBinaryCache() *BinaryCache {
	dir, err := os.MkdirTemp(os.TempDir(), "binarycache-*")
	if err != nil {
		log.Fatalf("Failed to create temporary directory: %v", err)
	}

	return &BinaryCache{dir: dir}
}

func (c *BinaryCache) LoadBinary(importPath string) string {
	p, ok := c.cache.Load(importPath)
	if ok {
		return p.(string)
	}

	binPath := filepath.Join(c.dir, RandomBinaryName())
	cmd := exec.Command("go", "build", "-o", binPath, importPath)
	cmd.Env = append(os.Environ(), "GOFLAGS=")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	c.cache.Store(importPath, binPath)
	return binPath
}

func (c *BinaryCache) Clear() {
	_ = os.RemoveAll(c.dir)
}

func RandomBinaryName() string {
	name := RandomName()
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	return name
}

func RandomName() string {
	var raw [8]byte
	_, _ = rand.Read(raw[:])
	name := hex.EncodeToString(raw[:])
	return name
}
