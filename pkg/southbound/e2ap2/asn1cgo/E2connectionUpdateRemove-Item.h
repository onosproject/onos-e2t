/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2AP-PDU-Contents"
 * 	found in "../../../../api/e2ap/v2beta1/e2ap_v2.0.1_Feb3_21.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_E2connectionUpdateRemove_Item_H_
#define	_E2connectionUpdateRemove_Item_H_


#include "asn_application.h"

/* Including external dependencies */
#include "TNLinformation.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* E2connectionUpdateRemove-Item */
typedef struct E2connectionUpdateRemove_Item {
	TNLinformation_t	 tnlInformation;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} E2connectionUpdateRemove_Item_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_E2connectionUpdateRemove_Item;

#ifdef __cplusplus
}
#endif

#endif	/* _E2connectionUpdateRemove_Item_H_ */
#include "asn_internal.h"
