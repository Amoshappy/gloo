// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"

	"github.com/mitchellh/hashstructure"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"go.uber.org/zap"
)

type ApiSnapshot struct {
	Secrets   gloo_solo_io.SecretsByNamespace
	Upstreams gloo_solo_io.UpstreamsByNamespace
	Ingresses IngressesByNamespace
}

func (s ApiSnapshot) Clone() ApiSnapshot {
	return ApiSnapshot{
		Secrets:   s.Secrets.Clone(),
		Upstreams: s.Upstreams.Clone(),
		Ingresses: s.Ingresses.Clone(),
	}
}

func (s ApiSnapshot) snapshotToHash() ApiSnapshot {
	snapshotForHashing := s.Clone()
	for _, secret := range snapshotForHashing.Secrets.List() {
		resources.UpdateMetadata(secret, func(meta *core.Metadata) {
			meta.ResourceVersion = ""
		})
	}
	for _, upstream := range snapshotForHashing.Upstreams.List() {
		resources.UpdateMetadata(upstream, func(meta *core.Metadata) {
			meta.ResourceVersion = ""
		})
		upstream.SetStatus(core.Status{})
	}
	for _, ingress := range snapshotForHashing.Ingresses.List() {
		resources.UpdateMetadata(ingress, func(meta *core.Metadata) {
			meta.ResourceVersion = ""
		})
		ingress.SetStatus(core.Status{})
	}

	return snapshotForHashing
}

func (s ApiSnapshot) Hash() uint64 {
	return s.hashStruct(s.snapshotToHash())
}

func (s ApiSnapshot) HashFields() []zap.Field {
	snapshotForHashing := s.snapshotToHash()
	var fields []zap.Field
	secrets := s.hashStruct(snapshotForHashing.Secrets.List())
	fields = append(fields, zap.Uint64("secrets", secrets))
	upstreams := s.hashStruct(snapshotForHashing.Upstreams.List())
	fields = append(fields, zap.Uint64("upstreams", upstreams))
	ingresses := s.hashStruct(snapshotForHashing.Ingresses.List())
	fields = append(fields, zap.Uint64("ingresses", ingresses))

	return append(fields, zap.Uint64("snapshotHash", s.hashStruct(snapshotForHashing)))
}

func (s ApiSnapshot) hashStruct(v interface{}) uint64 {
	h, err := hashstructure.Hash(v, nil)
	if err != nil {
		panic(err)
	}
	return h
}
