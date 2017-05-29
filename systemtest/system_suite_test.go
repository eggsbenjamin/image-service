// +build system

package system_test

import (
	"net"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

const (
	address   = "localhost:3030"
	pathToSrc = "github.com/eggsbenjamin/image-service/cmd"
)

var (
	pathToBinary string
	session      *gexec.Session
)

//	build the server binary
var _ = BeforeSuite(func() {
	var err error
	pathToBinary, err = gexec.Build(pathToSrc)
	Expect(err).NotTo(HaveOccurred())
})

//	clean up server binary
var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

//	spin up the server
var _ = BeforeEach(func() {
	var err error
	session, err = gexec.Start(exec.Command(pathToBinary), GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())
	Eventually(verifyServerIsListening).Should(Succeed())
})

//	tear down the server
var _ = AfterEach(func() {
	session.Interrupt()
	Eventually(session).Should(gexec.Exit())
})

func verifyServerIsListening() error {
	_, err := net.Dial("tcp", address)
	return err
}

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "System Test Suite")
}
