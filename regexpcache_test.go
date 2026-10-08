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

func TestConcurrentStress(t *testing.T) {
	const goroutines = 100
	var wg sync.WaitGroup
	patterns := []string{`\d+`, `[a-z]+`, `(foo|bar)`, `\s*`, `[`, `invalid(`}

	for _, p := range patterns {
		p := p
		for i := 0; i < goroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				re, err := regexpcache.Compile(p)
				re2, err2 := regexpcache.Compile(p)
				if (err == nil) != (err2 == nil) {
					t.Errorf("mismatched error presence for %s", p)
				}
				if re != re2 {
					t.Errorf("mismatched regexp pointer for %s", p)
				}
			}()
		}
	}
	wg.Wait()
}

func TestConcurrentPOSIXStress(t *testing.T) {
	const goroutines = 100
	var wg sync.WaitGroup
	patterns := []string{`[0-9]+`, `[a-z]+`, `(foo|bar)`, `[`, `invalid(`}

	for _, p := range patterns {
		p := p
		for i := 0; i < goroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				re, err := regexpcache.CompilePOSIX(p)
				re2, err2 := regexpcache.CompilePOSIX(p)
				if (err == nil) != (err2 == nil) {
					t.Errorf("mismatched error presence for %s", p)
				}
				if re != re2 {
					t.Errorf("mismatched regexp pointer for %s", p)
				}
			}()
		}
	}
	wg.Wait()
}

func BenchmarkCompileHit(b *testing.B) {
	regexpcache.Compile(`\d+`)
	b.ResetTimer()
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = regexpcache.Compile(`\d+`)
		}
	})
}
