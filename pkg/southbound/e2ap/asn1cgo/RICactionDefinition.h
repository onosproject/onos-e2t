/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2AP-IEs"
 * 	found in "../../../../api/e2ap/v2/e2ap_v2.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_RICactionDefinition_H_
#define	_RICactionDefinition_H_


#include "asn_application.h"

/* Including external dependencies */
#include "OCTET_STRING.h"

#ifdef __cplusplus
extern "C" {
#endif

/* RICactionDefinition */
typedef OCTET_STRING_t	 RICactionDefinition_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_RICactionDefinition;
asn_struct_free_f RICactionDefinition_free;
asn_struct_print_f RICactionDefinition_print;
asn_constr_check_f RICactionDefinition_constraint;
ber_type_decoder_f RICactionDefinition_decode_ber;
der_type_encoder_f RICactionDefinition_encode_der;
xer_type_decoder_f RICactionDefinition_decode_xer;
xer_type_encoder_f RICactionDefinition_encode_xer;
per_type_decoder_f RICactionDefinition_decode_uper;
per_type_encoder_f RICactionDefinition_encode_uper;
per_type_decoder_f RICactionDefinition_decode_aper;
per_type_encoder_f RICactionDefinition_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _RICactionDefinition_H_ */
#include "asn_internal.h"
