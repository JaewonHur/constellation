//go:build enterprise

/*
Copyright (c) Edgeless Systems GmbH

SPDX-License-Identifier: AGPL-3.0-only
*/

package measurements

// Regenerate the measurements by running go generate.
// The directive can be found in the file measurements.go since it does not have
// a build tag.
// The enterprise build tag is required to validate the measurements using production
// sigstore certificates.

// revive:disable:var-naming
var (
	aws_AWSNitroTPM          = M{0: {Expected: []byte{0x73, 0x7f, 0x76, 0x7a, 0x12, 0xf5, 0x4e, 0x70, 0xee, 0xcb, 0xc8, 0x68, 0x40, 0x11, 0x32, 0x3a, 0xe2, 0xfe, 0x2d, 0xd9, 0xf9, 0x07, 0x85, 0x57, 0x79, 0x69, 0xd7, 0xa2, 0x01, 0x3e, 0x8c, 0x12}, ValidationOpt: WarnOnly}, 2: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 3: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 4: {Expected: []byte{0x7c, 0x4e, 0xca, 0x99, 0x21, 0x28, 0x69, 0xb6, 0xfa, 0xa6, 0x9a, 0x8a, 0xa9, 0x16, 0xa5, 0x04, 0x61, 0xd7, 0xa2, 0xfa, 0xe8, 0xac, 0xc0, 0xcf, 0xac, 0x44, 0x22, 0x80, 0x32, 0x67, 0x24, 0x5a}, ValidationOpt: Enforce}, 6: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 8: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 9: {Expected: []byte{0x69, 0xd8, 0x35, 0xe7, 0x06, 0x94, 0x9a, 0xb1, 0x19, 0xb7, 0x08, 0xfb, 0x79, 0xa0, 0xc2, 0x09, 0x59, 0xc1, 0x9c, 0xfd, 0x3c, 0xe6, 0x92, 0xf3, 0xc8, 0x75, 0x33, 0xa4, 0x7c, 0x2f, 0xdd, 0x1c}, ValidationOpt: Enforce}, 11: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 12: {Expected: []byte{0xbe, 0x29, 0x7a, 0x98, 0x53, 0x10, 0x54, 0xa6, 0x32, 0x6c, 0x22, 0xe2, 0x55, 0x36, 0xb9, 0xca, 0x9c, 0x99, 0xfa, 0xa6, 0x07, 0x75, 0x76, 0x14, 0x79, 0x44, 0x1d, 0x41, 0xd9, 0x67, 0x1d, 0xd0}, ValidationOpt: Enforce}, 13: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 14: {Expected: []byte{0xd7, 0xc4, 0xcc, 0x7f, 0xf7, 0x93, 0x30, 0x22, 0xf0, 0x13, 0xe0, 0x3b, 0xde, 0xe8, 0x75, 0xb9, 0x17, 0x20, 0xb5, 0xb8, 0x6c, 0xf1, 0x75, 0x3c, 0xad, 0x83, 0x0f, 0x95, 0xe7, 0x91, 0x92, 0x6f}, ValidationOpt: WarnOnly}, 15: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}}
	aws_AWSSEVSNP            = M{0: {Expected: []byte{0x7b, 0x06, 0x8c, 0x0c, 0x3a, 0xc2, 0x9a, 0xfe, 0x26, 0x41, 0x34, 0x53, 0x6b, 0x9b, 0xe2, 0x6f, 0x1d, 0x4c, 0xcd, 0x57, 0x5b, 0x88, 0xd3, 0xc3, 0xce, 0xab, 0xf3, 0x6a, 0xc9, 0x9c, 0x02, 0x78}, ValidationOpt: WarnOnly}, 2: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 3: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 4: {Expected: []byte{0x41, 0x9d, 0xc5, 0x27, 0xea, 0x79, 0x91, 0x23, 0xd3, 0xef, 0x62, 0x92, 0xbb, 0xba, 0x3e, 0xdb, 0x1c, 0xb5, 0x36, 0x11, 0xd1, 0xa4, 0x94, 0x77, 0x70, 0x6f, 0x32, 0x6c, 0x25, 0xe4, 0x0c, 0x8f}, ValidationOpt: Enforce}, 6: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 8: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 9: {Expected: []byte{0x69, 0xd8, 0x35, 0xe7, 0x06, 0x94, 0x9a, 0xb1, 0x19, 0xb7, 0x08, 0xfb, 0x79, 0xa0, 0xc2, 0x09, 0x59, 0xc1, 0x9c, 0xfd, 0x3c, 0xe6, 0x92, 0xf3, 0xc8, 0x75, 0x33, 0xa4, 0x7c, 0x2f, 0xdd, 0x1c}, ValidationOpt: Enforce}, 11: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 12: {Expected: []byte{0x10, 0x50, 0x62, 0x2d, 0xa2, 0x95, 0xc8, 0x6c, 0x90, 0xdd, 0xa8, 0xe3, 0xe2, 0x93, 0x86, 0xda, 0xa8, 0x56, 0x8f, 0xe5, 0xbf, 0x19, 0xe8, 0x4b, 0x91, 0x7c, 0xff, 0x4d, 0x37, 0xe4, 0xc6, 0xef}, ValidationOpt: Enforce}, 13: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 14: {Expected: []byte{0xd7, 0xc4, 0xcc, 0x7f, 0xf7, 0x93, 0x30, 0x22, 0xf0, 0x13, 0xe0, 0x3b, 0xde, 0xe8, 0x75, 0xb9, 0x17, 0x20, 0xb5, 0xb8, 0x6c, 0xf1, 0x75, 0x3c, 0xad, 0x83, 0x0f, 0x95, 0xe7, 0x91, 0x92, 0x6f}, ValidationOpt: WarnOnly}, 15: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}}
	azure_AzureSEVSNP        = M{1: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 2: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 3: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 4: {Expected: []byte{0x2c, 0xa4, 0xb0, 0x11, 0xc9, 0x91, 0x71, 0xd8, 0xa0, 0xf9, 0x1d, 0x67, 0xf2, 0xbd, 0x30, 0x4b, 0x70, 0xf1, 0x83, 0x7f, 0xa0, 0xd6, 0xd8, 0x14, 0x54, 0x10, 0xa1, 0x26, 0x52, 0x0d, 0xcc, 0x09}, ValidationOpt: Enforce}, 8: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 9: {Expected: []byte{0x69, 0xd8, 0x35, 0xe7, 0x06, 0x94, 0x9a, 0xb1, 0x19, 0xb7, 0x08, 0xfb, 0x79, 0xa0, 0xc2, 0x09, 0x59, 0xc1, 0x9c, 0xfd, 0x3c, 0xe6, 0x92, 0xf3, 0xc8, 0x75, 0x33, 0xa4, 0x7c, 0x2f, 0xdd, 0x1c}, ValidationOpt: Enforce}, 11: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 12: {Expected: []byte{0x94, 0xc3, 0x40, 0x59, 0x90, 0x10, 0x10, 0x71, 0x94, 0xd2, 0xce, 0x36, 0x54, 0x78, 0x58, 0x22, 0xba, 0x65, 0x99, 0xf2, 0x6e, 0xcd, 0x0d, 0xe3, 0x84, 0x42, 0x89, 0xf6, 0x0c, 0xf7, 0x61, 0xd5}, ValidationOpt: Enforce}, 13: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 14: {Expected: []byte{0xd7, 0xc4, 0xcc, 0x7f, 0xf7, 0x93, 0x30, 0x22, 0xf0, 0x13, 0xe0, 0x3b, 0xde, 0xe8, 0x75, 0xb9, 0x17, 0x20, 0xb5, 0xb8, 0x6c, 0xf1, 0x75, 0x3c, 0xad, 0x83, 0x0f, 0x95, 0xe7, 0x91, 0x92, 0x6f}, ValidationOpt: WarnOnly}, 15: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}}
	azure_AzureTrustedLaunch M
	gcp_GCPSEVES             = M{1: {Expected: []byte{0x74, 0x5f, 0x2f, 0xb4, 0x23, 0x5e, 0x46, 0x47, 0xaa, 0x0a, 0xd5, 0xac, 0xe7, 0x81, 0xcd, 0x92, 0x9e, 0xb6, 0x8c, 0x28, 0x87, 0x0e, 0x7d, 0xd5, 0xd1, 0xa1, 0x53, 0x58, 0x54, 0x32, 0x5e, 0x56}, ValidationOpt: WarnOnly}, 2: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 3: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 4: {Expected: []byte{0xbb, 0xf5, 0xb6, 0x8e, 0x0e, 0x56, 0x8f, 0x8f, 0x9b, 0x26, 0x7f, 0xb3, 0xf6, 0x49, 0x57, 0xf9, 0x69, 0x0c, 0x04, 0xe6, 0x69, 0x57, 0xff, 0x20, 0xcf, 0x1a, 0xb4, 0x70, 0xf7, 0x9b, 0xa5, 0xcd}, ValidationOpt: Enforce}, 6: {Expected: []byte{0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea, 0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d, 0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a, 0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69}, ValidationOpt: WarnOnly}, 8: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 9: {Expected: []byte{0x69, 0xd8, 0x35, 0xe7, 0x06, 0x94, 0x9a, 0xb1, 0x19, 0xb7, 0x08, 0xfb, 0x79, 0xa0, 0xc2, 0x09, 0x59, 0xc1, 0x9c, 0xfd, 0x3c, 0xe6, 0x92, 0xf3, 0xc8, 0x75, 0x33, 0xa4, 0x7c, 0x2f, 0xdd, 0x1c}, ValidationOpt: Enforce}, 11: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 12: {Expected: []byte{0xdf, 0xfd, 0xf3, 0x44, 0x05, 0x1e, 0x7c, 0xc8, 0x1d, 0x78, 0x32, 0x54, 0x13, 0x3a, 0xa1, 0x68, 0x6e, 0xa8, 0xd8, 0x12, 0x50, 0x97, 0x8c, 0x5b, 0x7c, 0x57, 0x0c, 0xd0, 0x46, 0x23, 0xd1, 0xff}, ValidationOpt: Enforce}, 13: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 14: {Expected: []byte{0xd7, 0xc4, 0xcc, 0x7f, 0xf7, 0x93, 0x30, 0x22, 0xf0, 0x13, 0xe0, 0x3b, 0xde, 0xe8, 0x75, 0xb9, 0x17, 0x20, 0xb5, 0xb8, 0x6c, 0xf1, 0x75, 0x3c, 0xad, 0x83, 0x0f, 0x95, 0xe7, 0x91, 0x92, 0x6f}, ValidationOpt: WarnOnly}, 15: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}}
	qemu_QEMUTDX             M
	qemu_QEMUVTPM            = M{4: {Expected: []byte{0x8c, 0x5d, 0x53, 0x59, 0x4b, 0xe8, 0x60, 0x4f, 0x08, 0x94, 0xac, 0x12, 0x23, 0x5b, 0xb3, 0x34, 0xa1, 0x11, 0xe0, 0xd2, 0xa1, 0x66, 0x22, 0x2f, 0xed, 0xdd, 0x92, 0xe4, 0x52, 0x4a, 0x5f, 0xf3}, ValidationOpt: Enforce}, 8: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 9: {Expected: []byte{0x69, 0xd8, 0x35, 0xe7, 0x06, 0x94, 0x9a, 0xb1, 0x19, 0xb7, 0x08, 0xfb, 0x79, 0xa0, 0xc2, 0x09, 0x59, 0xc1, 0x9c, 0xfd, 0x3c, 0xe6, 0x92, 0xf3, 0xc8, 0x75, 0x33, 0xa4, 0x7c, 0x2f, 0xdd, 0x1c}, ValidationOpt: Enforce}, 11: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 12: {Expected: []byte{0xdb, 0x25, 0x88, 0xe9, 0xe4, 0x60, 0x9d, 0x51, 0x38, 0x64, 0xe3, 0x7b, 0x3c, 0x62, 0xcb, 0x0f, 0x03, 0xe0, 0xae, 0x3a, 0x62, 0x23, 0x43, 0x79, 0xa0, 0x85, 0xb0, 0x99, 0xcd, 0xbe, 0x27, 0x5a}, ValidationOpt: Enforce}, 13: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}, 15: {Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, ValidationOpt: Enforce}}
)
