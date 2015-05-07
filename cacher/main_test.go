package cacher_test

import (
	"github.com/crit/critical-go/cacher"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type CacherSuite struct{}

func (t *CacherSuite) SetUpSuite(c *C) {

}

func (t *CacherSuite) TearDownSuite(c *C) {

}

var _ = Suite(&CacherSuite{})

func (t *CacherSuite) TestInit(ch *C) {
	opts := cacher.Options{}
	opts.Engine = cacher.LOCALCACHE

	cacher.InitCache(opts)
}

func (t *CacherSuite) TestSet(ch *C) {
	opts := cacher.Options{}
	opts.Engine = cacher.LOCALCACHE

	cacher.InitCache(opts)

	cacher.Set("a", []byte("b"))
}

func (t *CacherSuite) TestGet(ch *C) {
	opts := cacher.Options{}
	opts.Engine = cacher.LOCALCACHE

	cacher.InitCache(opts)

	cacher.Set("c", []byte("d"))

	a := cacher.Get("c")

	ch.Assert(string(a), Equals, "d")
}

func (t *CacherSuite) TestDelete(ch *C) {
	opts := cacher.Options{}
	opts.Engine = cacher.LOCALCACHE

	cacher.InitCache(opts)

	cacher.Set("e", []byte("f"))

	cacher.Delete("e")

	a := cacher.Get("e")

	ch.Assert(len(a), Equals, 0)
}
