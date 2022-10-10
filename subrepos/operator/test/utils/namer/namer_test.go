package namer_test

import (
	"testing"

	"fmt"

	"github.com/artemiscloud/activemq-artemis-operator/pkg/utils/namer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/types"
)

func TestConfigUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Namer Utils Suite")
}

var _ = BeforeSuite(func() {
	fmt.Println("=======Before Config Suite========")
})

var _ = AfterSuite(func() {
	fmt.Println("=======After Config Suite========")
})

var _ = Describe("Namer Util Test", func() {
	Context("TestMatchPodNameAndStatefulSetName: match", func() {
		podName := "ex-aao-ss-0"
		podNamespace := "default"
		ssName := "ex-aao-ss"
		ssNamespace := "default"

		It("Testing pod belonging to statefulset", func() {
			podNamespacedName := types.NamespacedName{
				Name:      podName,
				Namespace: podNamespace,
			}
			ssNamespacedName := types.NamespacedName{
				Name:      ssName,
				Namespace: ssNamespace,
			}
			err, ok, podSerial := namer.PodBelongsToStatefulset(&podNamespacedName, &ssNamespacedName)
			Expect(err).To(BeNil())
			Expect(ok).To(BeTrue())
			Expect(podSerial).To(Equal(0))
		})
	})
	Context("TestMatchPodNameAndStatefulSetName: namespace mismatch", func() {
		podName := "ex-aao-ss-0"
		podNamespace := "default"
		ssName := "ex-aao-ss"
		ssNamespace := "another"

		It("Testing different namespaces", func() {
			podNamespacedName := types.NamespacedName{
				Name:      podName,
				Namespace: podNamespace,
			}
			ssNamespacedName := types.NamespacedName{
				Name:      ssName,
				Namespace: ssNamespace,
			}
			err, ok, podSerial := namer.PodBelongsToStatefulset(&podNamespacedName, &ssNamespacedName)
			Expect(err).NotTo(BeNil())
			Expect(ok).To(BeFalse())
			Expect(podSerial).To(Equal(-1))
		})
	})
	Context("TestMatchPodNameAndStatefulSetName: mismatch-different ss names", func() {
		podName := "ex-aao-ss-0"
		podNamespace := "default"
		ssName := "ex-aao1-ss"
		ssNamespace := "default"

		It("Testing different statefulset name", func() {
			podNamespacedName := types.NamespacedName{
				Name:      podName,
				Namespace: podNamespace,
			}
			ssNamespacedName := types.NamespacedName{
				Name:      ssName,
				Namespace: ssNamespace,
			}
			err, ok, podSerial := namer.PodBelongsToStatefulset(&podNamespacedName, &ssNamespacedName)
			Expect(err).NotTo(BeNil())
			Expect(ok).To(BeFalse())
			Expect(podSerial).To(Equal(-1))
		})
	})
})
