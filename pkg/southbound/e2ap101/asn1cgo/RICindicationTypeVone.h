/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2AP-IEs"
 * 	found in "../../../../api/e2ap/v1beta2/e2ap-v01.01.00.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_RICindicationTypeVone_H_
#define	_RICindicationTypeVone_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum RICindicationTypeVone {
	RICindicationTypeVone_report	= 0,
	RICindicationTypeVone_insert	= 1
	/*
	 * Enumeration is extensible
	 */
} e_RICindicationTypeVone;

/* RICindicationTypeVone */
typedef long	 RICindicationTypeVone_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_RICindicationTypeVone;
asn_struct_free_f RICindicationTypeVone_free;
asn_struct_print_f RICindicationTypeVone_print;
asn_constr_check_f RICindicationTypeVone_constraint;
ber_type_decoder_f RICindicationTypeVone_decode_ber;
der_type_encoder_f RICindicationTypeVone_encode_der;
xer_type_decoder_f RICindicationTypeVone_decode_xer;
xer_type_encoder_f RICindicationTypeVone_encode_xer;
per_type_decoder_f RICindicationTypeVone_decode_uper;
per_type_encoder_f RICindicationTypeVone_encode_uper;
per_type_decoder_f RICindicationTypeVone_decode_aper;
per_type_encoder_f RICindicationTypeVone_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _RICindicationTypeVone_H_ */
#include "asn_internal.h"
