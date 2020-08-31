/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2AP-PDU-Contents"
 * 	found in "e2ap-v01.00.00.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#include "RANfunctionsID-List.h"

#include "ProtocolIE-SingleContainer.h"
static asn_per_constraints_t asn_PER_type_RANfunctionsID_List_constr_1 CC_NOTUSED = {
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	{ APC_CONSTRAINED,	 9,  9,  0,  256 }	/* (SIZE(0..256)) */,
	0, 0	/* No PER value map */
};
static asn_TYPE_member_t asn_MBR_RANfunctionsID_List_1[] = {
	{ ATF_POINTER, 0, 0,
		(ASN_TAG_CLASS_UNIVERSAL | (16 << 2)),
		0,
		&asn_DEF_ProtocolIE_SingleContainer_1547P4,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		""
		},
};
static const ber_tlv_tag_t asn_DEF_RANfunctionsID_List_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static asn_SET_OF_specifics_t asn_SPC_RANfunctionsID_List_specs_1 = {
	sizeof(struct RANfunctionsID_List),
	offsetof(struct RANfunctionsID_List, _asn_ctx),
	0,	/* XER encoding is XMLDelimitedItemList */
};
asn_TYPE_descriptor_t asn_DEF_RANfunctionsID_List = {
	"RANfunctionsID-List",
	"RANfunctionsID-List",
	&asn_OP_SEQUENCE_OF,
	asn_DEF_RANfunctionsID_List_tags_1,
	sizeof(asn_DEF_RANfunctionsID_List_tags_1)
		/sizeof(asn_DEF_RANfunctionsID_List_tags_1[0]), /* 1 */
	asn_DEF_RANfunctionsID_List_tags_1,	/* Same as above */
	sizeof(asn_DEF_RANfunctionsID_List_tags_1)
		/sizeof(asn_DEF_RANfunctionsID_List_tags_1[0]), /* 1 */
	{ 0, &asn_PER_type_RANfunctionsID_List_constr_1, SEQUENCE_OF_constraint },
	asn_MBR_RANfunctionsID_List_1,
	1,	/* Single element */
	&asn_SPC_RANfunctionsID_List_specs_1	/* Additional specs */
};

