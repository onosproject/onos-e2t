/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2AP-IEs"
 * 	found in "../../../../api/e2ap/v1beta2/e2ap-v01.01.00.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#include "GlobalE2node-ng-eNB-ID-Vone.h"

asn_TYPE_member_t asn_MBR_GlobalE2node_ng_eNB_ID_Vone_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct GlobalE2node_ng_eNB_ID_Vone, global_ng_eNB_ID),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_GlobalngeNB_ID_Vone,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"global-ng-eNB-ID"
		},
};
static const ber_tlv_tag_t asn_DEF_GlobalE2node_ng_eNB_ID_Vone_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_GlobalE2node_ng_eNB_ID_Vone_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 } /* global-ng-eNB-ID */
};
asn_SEQUENCE_specifics_t asn_SPC_GlobalE2node_ng_eNB_ID_Vone_specs_1 = {
	sizeof(struct GlobalE2node_ng_eNB_ID_Vone),
	offsetof(struct GlobalE2node_ng_eNB_ID_Vone, _asn_ctx),
	asn_MAP_GlobalE2node_ng_eNB_ID_Vone_tag2el_1,
	1,	/* Count of tags in the map */
	0, 0, 0,	/* Optional elements (not needed) */
	1,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_GlobalE2node_ng_eNB_ID_Vone = {
	"GlobalE2node-ng-eNB-ID-Vone",
	"GlobalE2node-ng-eNB-ID-Vone",
	&asn_OP_SEQUENCE,
	asn_DEF_GlobalE2node_ng_eNB_ID_Vone_tags_1,
	sizeof(asn_DEF_GlobalE2node_ng_eNB_ID_Vone_tags_1)
		/sizeof(asn_DEF_GlobalE2node_ng_eNB_ID_Vone_tags_1[0]), /* 1 */
	asn_DEF_GlobalE2node_ng_eNB_ID_Vone_tags_1,	/* Same as above */
	sizeof(asn_DEF_GlobalE2node_ng_eNB_ID_Vone_tags_1)
		/sizeof(asn_DEF_GlobalE2node_ng_eNB_ID_Vone_tags_1[0]), /* 1 */
	{ 0, 0, SEQUENCE_constraint },
	asn_MBR_GlobalE2node_ng_eNB_ID_Vone_1,
	1,	/* Elements count */
	&asn_SPC_GlobalE2node_ng_eNB_ID_Vone_specs_1	/* Additional specs */
};

