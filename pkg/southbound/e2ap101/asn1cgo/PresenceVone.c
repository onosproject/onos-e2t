/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2AP-CommonDataTypes"
 * 	found in "../../../../api/e2ap/v1beta2/e2ap-v01.01.00.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#include "PresenceVone.h"

/*
 * This type is implemented using NativeEnumerated,
 * so here we adjust the DEF accordingly.
 */
asn_per_constraints_t asn_PER_type_PresenceVone_constr_1 CC_NOTUSED = {
	{ APC_CONSTRAINED,	 2,  2,  0,  2 }	/* (0..2) */,
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	0, 0	/* No PER value map */
};
static const asn_INTEGER_enum_map_t asn_MAP_PresenceVone_value2enum_1[] = {
	{ 0,	8,	"optional" },
	{ 1,	11,	"conditional" },
	{ 2,	9,	"mandatory" }
};
static const unsigned int asn_MAP_PresenceVone_enum2value_1[] = {
	1,	/* conditional(1) */
	2,	/* mandatory(2) */
	0	/* optional(0) */
};
const asn_INTEGER_specifics_t asn_SPC_PresenceVone_specs_1 = {
	asn_MAP_PresenceVone_value2enum_1,	/* "tag" => N; sorted by tag */
	asn_MAP_PresenceVone_enum2value_1,	/* N => "tag"; sorted by N */
	3,	/* Number of elements in the maps */
	0,	/* Enumeration is not extensible */
	1,	/* Strict enumeration */
	0,	/* Native long size */
	0
};
static const ber_tlv_tag_t asn_DEF_PresenceVone_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (10 << 2))
};
asn_TYPE_descriptor_t asn_DEF_PresenceVone = {
	"PresenceVone",
	"PresenceVone",
	&asn_OP_NativeEnumerated,
	asn_DEF_PresenceVone_tags_1,
	sizeof(asn_DEF_PresenceVone_tags_1)
		/sizeof(asn_DEF_PresenceVone_tags_1[0]), /* 1 */
	asn_DEF_PresenceVone_tags_1,	/* Same as above */
	sizeof(asn_DEF_PresenceVone_tags_1)
		/sizeof(asn_DEF_PresenceVone_tags_1[0]), /* 1 */
	{ 0, &asn_PER_type_PresenceVone_constr_1, NativeEnumerated_constraint },
	0, 0,	/* Defined elsewhere */
	&asn_SPC_PresenceVone_specs_1	/* Additional specs */
};

