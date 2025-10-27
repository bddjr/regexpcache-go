package regexpcache_test

import (
	"regexp"
	"sync"
	"testing"

	"github.com/bddjr/regexpcache-go"
)

func wgGo(wg *sync.WaitGroup, f func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		f()
	}()
}

func TestCompile(t *testing.T) {
	println(`TestCompile`)
	var a_re, b_re *regexp.Regexp
	var a_err, b_err error
	var wg sync.WaitGroup
	for _, v := range []string{`\d`, ``} {
		wgGo(&wg, func() { a_re, a_err = regexpcache.Compile(v) })
		wgGo(&wg, func() { b_re, b_err = regexpcache.Compile(v) })
		wg.Wait()
		if a_re != b_re {
			panic(`a_re != b_re`)
		}
		if a_err != b_err {
			panic(`a_err != b_err`)
		}
		if a_re == nil && a_err == nil {
			panic(`a_re == nil && a_err == nil`)
		}
	}
}

func TestCompilePOSIX(t *testing.T) {
	println(`TestCompilePOSIX`)
	var a_re, b_re *regexp.Regexp
	var a_err, b_err error
	var wg sync.WaitGroup
	for _, v := range []string{`\d`, ``} {
		wgGo(&wg, func() { a_re, a_err = regexpcache.CompilePOSIX(v) })
		wgGo(&wg, func() { b_re, b_err = regexpcache.CompilePOSIX(v) })
		wg.Wait()
		if a_re != b_re {
			panic(`a_re != b_re`)
		}
		if a_err != b_err {
			panic(`a_err != b_err`)
		}
		if a_re == nil && a_err == nil {
			panic(`a_re == nil && a_err == nil`)
		}
	}
}
