/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2AP-PDU-VONE-Contents"
 * 	found in "../../../../api/e2ap/v1beta2/e2ap-v01.01.00.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#include "E2connectionUpdateVone-Item.h"

static asn_TYPE_member_t asn_MBR_E2connectionUpdateVone_Item_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct E2connectionUpdateVone_Item, tnlInformation),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_TNLinformationVone,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"tnlInformation"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct E2connectionUpdateVone_Item, tnlUsage),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_TNLusageVone,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"tnlUsage"
		},
};
static const ber_tlv_tag_t asn_DEF_E2connectionUpdateVone_Item_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_E2connectionUpdateVone_Item_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* tnlInformation */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 } /* tnlUsage */
};
static asn_SEQUENCE_specifics_t asn_SPC_E2connectionUpdateVone_Item_specs_1 = {
	sizeof(struct E2connectionUpdateVone_Item),
	offsetof(struct E2connectionUpdateVone_Item, _asn_ctx),
	asn_MAP_E2connectionUpdateVone_Item_tag2el_1,
	2,	/* Count of tags in the map */
	0, 0, 0,	/* Optional elements (not needed) */
	2,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_E2connectionUpdateVone_Item = {
	"E2connectionUpdateVone-Item",
	"E2connectionUpdateVone-Item",
	&asn_OP_SEQUENCE,
	asn_DEF_E2connectionUpdateVone_Item_tags_1,
	sizeof(asn_DEF_E2connectionUpdateVone_Item_tags_1)
		/sizeof(asn_DEF_E2connectionUpdateVone_Item_tags_1[0]), /* 1 */
	asn_DEF_E2connectionUpdateVone_Item_tags_1,	/* Same as above */
	sizeof(asn_DEF_E2connectionUpdateVone_Item_tags_1)
		/sizeof(asn_DEF_E2connectionUpdateVone_Item_tags_1[0]), /* 1 */
	{ 0, 0, SEQUENCE_constraint },
	asn_MBR_E2connectionUpdateVone_Item_1,
	2,	/* Elements count */
	&asn_SPC_E2connectionUpdateVone_Item_specs_1	/* Additional specs */
};

