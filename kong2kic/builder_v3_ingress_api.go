package kong2kic

import (
	"log"

	"github.com/kong/go-database-reconciler/pkg/file"
)

type KICv3IngressAPIBuilder struct {
	kicContent *KICContent
}

func newKICv3IngressAPIBuilder() *KICv3IngressAPIBuilder {
	return &KICv3IngressAPIBuilder{
		kicContent: &KICContent{},
	}
}

func (b *KICv3IngressAPIBuilder) buildServices(content *file.Content) {
	err := populateKICServicesWithAnnotations(content, b.kicContent)
	if err != nil {
		log.Fatal(err)
	}
}

func (b *KICv3IngressAPIBuilder) buildRoutes(content *file.Content) {
	err := populateKICIngressesWithAnnotations(content, b.kicContent)
	if err != nil {
		log.Fatal(err)
	}
}

func (b *KICv3IngressAPIBuilder) buildGlobalPlugins(content *file.Content) {
	err := populateKICKongClusterPlugins(content, b.kicContent)
	if err != nil {
		log.Fatal(err)
	}
}

func (b *KICv3IngressAPIBuilder) buildConsumers(content *file.Content) {
	err := populateKICConsumers(content, b.kicContent)
	if err != nil {
		log.Fatal(err)
	}
}

func (b *KICv3IngressAPIBuilder) buildConsumerGroups(content *file.Content) {
	err := populateKICConsumerGroups(content, b.kicContent)
	if err != nil {
		log.Fatal(err)
	}
}

func (b *KICv3IngressAPIBuilder) buildCACertificates(content *file.Content) {
	populateKICCACertificate(content, b.kicContent)
}

func (b *KICv3IngressAPIBuilder) buildCertificates(content *file.Content) {
	populateKICCertificates(content, b.kicContent)
}

func (b *KICv3IngressAPIBuilder) getContent() *KICContent {
	return b.kicContent
}
