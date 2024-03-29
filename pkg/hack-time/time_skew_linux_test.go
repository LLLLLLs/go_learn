package main

import (
	"os"
	"testing"

	"golearn/pkg/hack-time/log"
	"golearn/pkg/hack-time/test/timer"
)

// These test cases required bin/test/timer as its workload.
// You could use make test-utils to build it.

func TestModifyTime(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Time Suit")
}

var _ = BeforeSuite(func(done Done) {
	By("change working directory")

	err := os.Chdir(".")

	Expect(err).NotTo(HaveOccurred())

	close(done)
})

// These tests are written in BDD-style using Ginkgo framework. Refer to
// http://onsi.github.io/ginkgo to learn more.

var _ = Describe("ModifyTime", func() {
	var t *timer.Timer
	logger, err := log.NewDefaultZapLogger()
	Expect(err).ShouldNot(HaveOccurred())
	BeforeEach(func() {
		var err error

		t, err = timer.StartTimer()
		Expect(err).ShouldNot(HaveOccurred())
	})

	AfterEach(func() {
		err := t.Stop()
		Expect(err).ShouldNot(HaveOccurred())
	})

	Context("Modify Time", func() {
		It("should move forward successfully", func() {

			Expect(t).NotTo(BeNil())
			s, err := GetSkew(logger, NewConfig(10000, 0, 1))
			Expect(err).ShouldNot(HaveOccurred(), "error: %+v", err)

			now, err := t.GetTime()
			Expect(err).ShouldNot(HaveOccurred(), "error: %+v", err)

			sec := now.Unix()

			err = s.Inject(t.Pid())
			Expect(err).ShouldNot(HaveOccurred(), "error: %+v", err)

			newTime, err := t.GetTime()
			Expect(err).ShouldNot(HaveOccurred(), "error: %+v", err)

			newSec := newTime.Unix()

			Expect(newSec-sec).Should(BeNumerically(">=", 10000), "sec %d newSec %d", sec, newSec)
			Expect(newSec-sec).Should(BeNumerically("<=", 10010), "sec %d newSec %d", sec, newSec)
		})

		It("should move backward successfully", func() {
			Expect(t).NotTo(BeNil())
			s, err := GetSkew(logger, NewConfig(10000, 0, 1))
			Expect(err).ShouldNot(HaveOccurred(), "error: %+v", err)

			err = s.Inject(t.Pid())
			Expect(err).ShouldNot(HaveOccurred(), "error: %+v", err)

			now, err := t.GetTime()
			Expect(err).ShouldNot(HaveOccurred(), "error: %+v", err)

			sec := now.Unix()

			err = s.Recover(t.Pid())
			Expect(err).ShouldNot(HaveOccurred(), "error: %+v", err)

			newTime, err := t.GetTime()
			Expect(err).ShouldNot(HaveOccurred(), "error: %+v", err)

			newSec := newTime.Unix()
			Expect(10000-(sec-newSec)).Should(BeNumerically("<=", 1), "sec %d newSec %d", sec, newSec)
		})

		It("should handle nsec overflow", func() {
			Expect(t).NotTo(BeNil())

			s, err := GetSkew(logger, NewConfig(0, 1000000000, 1))
			Expect(err).ShouldNot(HaveOccurred(), "error: %+v", err)

			now, err := t.GetTime()
			Expect(err).ShouldNot(HaveOccurred(), "error: %+v", err)

			sec := now.Unix()

			err = s.Inject(t.Pid())
			Expect(err).ShouldNot(HaveOccurred(), "error: %+v", err)

			newTime, err := t.GetTime()
			Expect(err).ShouldNot(HaveOccurred(), "error: %+v", err)

			newSec := newTime.Unix()
			Expect(newSec-sec).Should(BeNumerically(">=", 1), "sec %d newSec %d", sec, newSec)
		})
	})
})
