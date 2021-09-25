/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2AP-IEs"
 * 	found in "../../../../api/e2ap/v2/e2ap_v2.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#include "NGENB-DU-ID.h"

int
NGENB_DU_ID_constraint(const asn_TYPE_descriptor_t *td, const void *sptr,
			asn_app_constraint_failed_f *ctfailcb, void *app_key) {
	const INTEGER_t *st = (const INTEGER_t *)sptr;
	long value;
	
	if(!sptr) {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: value not given (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
	
	if(asn_INTEGER2long(st, &value)) {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: value too large (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
	
	if((value >= 0 && value <= 68719476735)) {
		/* Constraint check succeeded */
		return 0;
	} else {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: constraint failed (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
}

/*
 * This type is implemented using INTEGER,
 * so here we adjust the DEF accordingly.
 */
asn_per_constraints_t asn_PER_type_NGENB_DU_ID_constr_1 CC_NOTUSED = {
	{ APC_CONSTRAINED,	 36, -1,  0,  68719476735 }	/* (0..68719476735) */,
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	0, 0	/* No PER value map */
};
static const ber_tlv_tag_t asn_DEF_NGENB_DU_ID_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (2 << 2))
};
asn_TYPE_descriptor_t asn_DEF_NGENB_DU_ID = {
	"NGENB-DU-ID",
	"NGENB-DU-ID",
	&asn_OP_INTEGER,
	asn_DEF_NGENB_DU_ID_tags_1,
	sizeof(asn_DEF_NGENB_DU_ID_tags_1)
		/sizeof(asn_DEF_NGENB_DU_ID_tags_1[0]), /* 1 */
	asn_DEF_NGENB_DU_ID_tags_1,	/* Same as above */
	sizeof(asn_DEF_NGENB_DU_ID_tags_1)
		/sizeof(asn_DEF_NGENB_DU_ID_tags_1[0]), /* 1 */
	{ 0, &asn_PER_type_NGENB_DU_ID_constr_1, NGENB_DU_ID_constraint },
	0, 0,	/* No members */
	0	/* No specifics */
};

