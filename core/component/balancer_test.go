// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"testing"

	"github.com/franela/goblin"
)

// TestUnitBalancer
func TestUnitBalancer(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#RandomBalancer", func() {
		g.It("It should satisfy test cases", func() {
			b := NewRandomBalancer([]*Target{
				{
					URL: "A",
				},
				{
					URL: "B",
				},
				{
					URL: "C",
				},
				{
					URL: "D",
				},
				{
					URL: "E",
				},
				{
					URL: "F",
				},
			})

			g.Assert(b.Next().URL != "").Equal(true)
			g.Assert(b.Next().URL != "").Equal(true)
			g.Assert(b.Next().URL != "").Equal(true)
			g.Assert(b.Next().URL != "").Equal(true)
			g.Assert(b.Next().URL != "").Equal(true)
			g.Assert(b.Next().URL != "").Equal(true)
		})
	})

	g.Describe("#RoundRobinBalancer", func() {
		g.It("It should satisfy test cases", func() {
			r := NewRoundRobinBalancer([]*Target{
				{
					URL: "A",
				},
				{
					URL: "B",
				},
				{
					URL: "C",
				},
				{
					URL: "D",
				},
				{
					URL: "E",
				},
				{
					URL: "F",
				},
			})

			g.Assert(r.Next().URL).Equal("A")
			g.Assert(r.Next().URL).Equal("B")
			g.Assert(r.Next().URL).Equal("C")
			g.Assert(r.Next().URL).Equal("D")
			g.Assert(r.Next().URL).Equal("E")
			g.Assert(r.Next().URL).Equal("F")
			g.Assert(r.Next().URL).Equal("A")
			g.Assert(r.Next().URL).Equal("B")
			g.Assert(r.Next().URL).Equal("C")
			g.Assert(r.Next().URL).Equal("D")
			g.Assert(r.Next().URL).Equal("E")
			g.Assert(r.Next().URL).Equal("F")
		})
	})
}
