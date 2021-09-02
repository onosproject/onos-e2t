/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2AP-IEs"
 * 	found in "../../../../api/e2ap/v1beta2/e2ap-v01.01.00.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_RANfunctionDefinitionVone_H_
#define	_RANfunctionDefinitionVone_H_


#include "asn_application.h"

/* Including external dependencies */
#include "OCTET_STRING.h"

#ifdef __cplusplus
extern "C" {
#endif

/* RANfunctionDefinitionVone */
typedef OCTET_STRING_t	 RANfunctionDefinitionVone_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_RANfunctionDefinitionVone;
asn_struct_free_f RANfunctionDefinitionVone_free;
asn_struct_print_f RANfunctionDefinitionVone_print;
asn_constr_check_f RANfunctionDefinitionVone_constraint;
ber_type_decoder_f RANfunctionDefinitionVone_decode_ber;
der_type_encoder_f RANfunctionDefinitionVone_encode_der;
xer_type_decoder_f RANfunctionDefinitionVone_decode_xer;
xer_type_encoder_f RANfunctionDefinitionVone_encode_xer;
per_type_decoder_f RANfunctionDefinitionVone_decode_uper;
per_type_encoder_f RANfunctionDefinitionVone_encode_uper;
per_type_decoder_f RANfunctionDefinitionVone_decode_aper;
per_type_encoder_f RANfunctionDefinitionVone_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _RANfunctionDefinitionVone_H_ */
#include "asn_internal.h"
