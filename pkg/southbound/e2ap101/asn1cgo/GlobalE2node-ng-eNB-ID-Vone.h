/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2AP-IEs"
 * 	found in "../../../../api/e2ap/v1beta2/e2ap-v01.01.00.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_GlobalE2node_ng_eNB_ID_Vone_H_
#define	_GlobalE2node_ng_eNB_ID_Vone_H_


#include "asn_application.h"

/* Including external dependencies */
#include "GlobalngeNB-ID-Vone.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* GlobalE2node-ng-eNB-ID-Vone */
typedef struct GlobalE2node_ng_eNB_ID_Vone {
	GlobalngeNB_ID_Vone_t	 global_ng_eNB_ID;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} GlobalE2node_ng_eNB_ID_Vone_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_GlobalE2node_ng_eNB_ID_Vone;
extern asn_SEQUENCE_specifics_t asn_SPC_GlobalE2node_ng_eNB_ID_Vone_specs_1;
extern asn_TYPE_member_t asn_MBR_GlobalE2node_ng_eNB_ID_Vone_1[1];

#ifdef __cplusplus
}
#endif

#endif	/* _GlobalE2node_ng_eNB_ID_Vone_H_ */
#include "asn_internal.h"
