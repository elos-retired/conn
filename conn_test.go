package conn_test

import (
	. "github.com/elos/conn"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Conn", func() {
	It("Exposes a ConnectionClosedError", func() {
		Expect(ConnectionClosedError).NotTo(BeNil())
	})
})
